package wallet

import (
	"context"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/log"
)

var numberOfBlocksCheckedPerIteration = 40

type ethHistoricalCommand struct {
	db           *Database
	eth          TransferDownloader
	address      common.Address
	client       *walletClient
	balanceCache *balanceCache
	feed         *event.Feed
	foundHeaders []*DBHeader
	error        error
	noLimit      bool

	from              *LastKnownBlock
	to, resultingFrom *big.Int
}

func (c *ethHistoricalCommand) Command() Command {
	return FiniteCommand{
		Interval: 5 * time.Second,
		Runable:  c.Run,
	}.Run
}

func (c *ethHistoricalCommand) Run(ctx context.Context) (err error) {
	start := time.Now()
	if c.from.Number != nil && c.from.Balance != nil {
		c.balanceCache.addBalanceToCache(c.address, c.from.Number, c.from.Balance)
	}
	if c.from.Number != nil && c.from.Nonce != nil {
		c.balanceCache.addNonceToCache(c.address, c.from.Number, c.from.Nonce)
	}
	from, headers, err := findBlocksWithEthTransfers(ctx, c.client, c.balanceCache, c.eth, c.address, c.from.Number, c.to, c.noLimit)

	if err != nil {
		c.error = err
		return nil
	}

	c.foundHeaders = headers
	c.resultingFrom = from

	log.Info("eth historical downloader finished successfully", "address", c.address, "from", from, "to", c.to, "total blocks", len(headers), "time", time.Since(start))

	//err = c.db.ProcessBlocks(c.address, from, c.to, headers, ethTransfer)
	if err != nil {
		log.Error("failed to save found blocks with transfers", "error", err)
		return err
	}
	log.Debug("eth transfers were persisted. command is closed")
	return nil
}

type erc20HistoricalCommand struct {
	db      *Database
	erc20   BatchDownloader
	address common.Address
	client  *walletClient
	feed    *event.Feed

	iterator     *IterativeDownloader
	to           *big.Int
	from         *big.Int
	foundHeaders []*DBHeader
}

func (c *erc20HistoricalCommand) Command() Command {
	return FiniteCommand{
		Interval: 5 * time.Second,
		Runable:  c.Run,
	}.Run
}

func (c *erc20HistoricalCommand) Run(ctx context.Context) (err error) {
	start := time.Now()
	if c.iterator == nil {
		c.iterator, err = SetupIterativeDownloader(
			c.db, c.client, c.address,
			c.erc20, erc20BatchSize, c.to, c.from)
		if err != nil {
			log.Error("failed to setup historical downloader for erc20")
			return err
		}
	}
	for !c.iterator.Finished() {
		headers, _, _, err := c.iterator.Next(ctx)
		if err != nil {
			log.Error("failed to get next batch", "error", err)
			return err
		}
		c.foundHeaders = append(c.foundHeaders, headers...)

		/*err = c.db.ProcessBlocks(c.address, from, to, headers, erc20Transfer)
		if err != nil {
			c.iterator.Revert()
			log.Error("failed to save downloaded erc20 blocks with transfers", "error", err)
			return err
		}*/
	}
	log.Info("wallet historical downloader for erc20 transfers finished", "in", time.Since(start))
	return nil
}

// controlCommand implements following procedure (following parts are executed sequeantially):
// - verifies that the last header that was synced is still in the canonical chain
// - runs fast indexing for each account separately
// - starts listening to new blocks and watches for reorgs
type controlCommand struct {
	accounts           []common.Address
	db                 *Database
	eth                *ETHTransferDownloader
	erc20              *ERC20TransfersDownloader
	chain              *big.Int
	client             *walletClient
	feed               *event.Feed
	errorsCount        int
	nonArchivalRPCNode bool
}

// run fast indexing for every accont up to canonical chain head minus safety depth.
// every account will run it from last synced header.
func (c *findAndCheckBlockRangeCommand) fastIndex(ctx context.Context, bCache *balanceCache, fromByAddress map[common.Address]*LastKnownBlock, toByAddress map[common.Address]*big.Int) (map[common.Address]*big.Int, map[common.Address][]*DBHeader, error) {
	start := time.Now()
	group := NewGroup(ctx)

	commands := make([]*ethHistoricalCommand, len(c.accounts))
	for i, address := range c.accounts {
		eth := &ethHistoricalCommand{
			db:           c.db,
			client:       c.client,
			balanceCache: bCache,
			address:      address,
			eth: &ETHTransferDownloader{
				chain:    c.chain,
				client:   c.client,
				accounts: []common.Address{address},
				signer:   types.NewEIP155Signer(c.chain),
				db:       c.db,
			},
			feed:    c.feed,
			from:    fromByAddress[address],
			to:      toByAddress[address],
			noLimit: c.noLimit,
		}
		commands[i] = eth
		group.Add(eth.Command())
	}
	select {
	case <-ctx.Done():
		return nil, nil, ctx.Err()
	case <-group.WaitAsync():
		resultingFromByAddress := map[common.Address]*big.Int{}
		headers := map[common.Address][]*DBHeader{}
		for _, command := range commands {
			if command.error != nil {
				return nil, nil, command.error
			}
			resultingFromByAddress[command.address] = command.resultingFrom
			headers[command.address] = command.foundHeaders
		}
		log.Info("fast indexer finished", "in", time.Since(start))
		return resultingFromByAddress, headers, nil
	}
}

// run fast indexing for every accont up to canonical chain head minus safety depth.
// every account will run it from last synced header.
func (c *findAndCheckBlockRangeCommand) fastIndexErc20(ctx context.Context, fromByAddress map[common.Address]*big.Int, toByAddress map[common.Address]*big.Int) (map[common.Address][]*DBHeader, error) {
	start := time.Now()
	group := NewGroup(ctx)

	commands := make([]*erc20HistoricalCommand, len(c.accounts))
	for i, address := range c.accounts {
		erc20 := &erc20HistoricalCommand{
			db:           c.db,
			erc20:        NewERC20TransfersDownloader(c.client, []common.Address{address}, types.NewEIP155Signer(c.chain)),
			client:       c.client,
			feed:         c.feed,
			address:      address,
			from:         fromByAddress[address],
			to:           toByAddress[address],
			foundHeaders: []*DBHeader{},
		}
		commands[i] = erc20
		group.Add(erc20.Command())
	}
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-group.WaitAsync():
		headres := map[common.Address][]*DBHeader{}
		for _, command := range commands {
			headres[command.address] = command.foundHeaders
		}
		log.Info("fast indexer Erc20 finished", "in", time.Since(start))
		return headres, nil
	}
}

func getTransfersByBlocks(ctx context.Context, db *Database, downloader *ETHTransferDownloader, address common.Address, blocks []*big.Int) ([]Transfer, error) {
	allTransfers := []Transfer{}

	for _, block := range blocks {
		transfers, err := downloader.GetTransfersByNumber(ctx, block)
		if err != nil {
			return nil, err
		}
		log.Debug("loadTransfers", "block", block, "new transfers", len(transfers))
		if len(transfers) > 0 {
			allTransfers = append(allTransfers, transfers...)
		}
	}

	return allTransfers, nil
}

func loadTransfers(ctx context.Context, accounts []common.Address, db *Database, client *walletClient, chain *big.Int, limit int, blocksByAddress map[common.Address][]*big.Int) (map[common.Address][]Transfer, error) {
	start := time.Now()
	group := NewGroup(ctx)

	commands := []*transfersCommand{}
	for _, address := range accounts {
		blocks, ok := blocksByAddress[address]

		if !ok {
			blocks, _ = db.GetBlocksByAddress(address, numberOfBlocksCheckedPerIteration)
		}
		for _, block := range blocks {
			transfers := &transfersCommand{
				db:      db,
				client:  client,
				address: address,
				eth: &ETHTransferDownloader{
					chain:    chain,
					client:   client,
					accounts: []common.Address{address},
					signer:   types.NewEIP155Signer(chain),
					db:       db,
				},
				block: block,
			}
			commands = append(commands, transfers)
			group.Add(transfers.Command())
		}
	}
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-group.WaitAsync():
		transfersByAddress := map[common.Address][]Transfer{}
		for _, command := range commands {
			if len(command.fetchedTransfers) == 0 {
				continue
			}

			transfers, ok := transfersByAddress[command.address]
			if !ok {
				transfers = []Transfer{}
			}

			for _, transfer := range command.fetchedTransfers {
				transfersByAddress[command.address] = append(transfers, transfer)
			}
		}
		log.Info("loadTransfers finished", "in", time.Since(start))
		return transfersByAddress, nil
	}
}

func (c *controlCommand) LoadTransfers(ctx context.Context, downloader *ETHTransferDownloader, limit int) (map[common.Address][]Transfer, error) {
	return loadTransfers(ctx, c.accounts, c.db, c.client, c.chain, limit, make(map[common.Address][]*big.Int))
}

func findFirstRange(c context.Context, account common.Address, initialTo *big.Int, client *walletClient) (*big.Int, error) {
	from := big.NewInt(0)
	to := initialTo
	goal := uint64(20)

	if from.Cmp(to) == 0 {
		return to, nil
	}

	firstNonce, err := client.NonceAt(c, account, to)
	log.Info("find range with 20 <= len(tx) <= 25", "account", account, "firstNonce", firstNonce, "from", from, "to", to)

	if err != nil {
		return nil, err
	}

	if firstNonce <= goal {
		return zero, nil
	}

	nonceDiff := firstNonce
	iterations := 0
	for iterations < 50 {
		iterations = iterations + 1

		if nonceDiff > goal {
			// from = (from + to) / 2
			from = from.Add(from, to)
			from = from.Div(from, big.NewInt(2))
		} else {
			// from = from - (from + to) / 2
			// to = from
			diff := big.NewInt(0).Sub(to, from)
			diff.Div(diff, big.NewInt(2))
			to = big.NewInt(from.Int64())
			from.Sub(from, diff)
		}
		fromNonce, err := client.NonceAt(c, account, from)
		if err != nil {
			return nil, err
		}
		nonceDiff = firstNonce - fromNonce

		log.Info("next nonce", "from", from, "n", fromNonce, "diff", firstNonce-fromNonce)

		if goal <= nonceDiff && nonceDiff <= (goal+5) {
			log.Info("range found", "account", account, "from", from, "to", to)
			return from, nil
		}
	}

	log.Info("range found", "account", account, "from", from, "to", to)

	return from, nil
}

func findFirstRanges(c context.Context, accounts []common.Address, initialTo *big.Int, client *walletClient) (map[common.Address]*big.Int, error) {
	res := map[common.Address]*big.Int{}

	for _, address := range accounts {
		from, err := findFirstRange(c, address, initialTo, client)
		if err != nil {
			return nil, err
		}

		res[address] = from
	}

	return res, nil
}

func (c *controlCommand) Run(parent context.Context) error {
	log.Info("start control command")
	ctx, cancel := context.WithTimeout(parent, 3*time.Second)
	head, err := c.client.HeaderByNumber(ctx, nil)
	cancel()
	if err != nil {
		if c.NewError(err) {
			return nil
		}
		return err
	}

	c.feed.Send(Event{
		Type:     EventFetchingRecentHistory,
		Accounts: c.accounts,
	})

	log.Info("current head is", "block number", head.Number)
	lastKnownEthBlocks, accountsWithoutHistory, err := c.db.GetLastKnownBlockByAddresses(c.accounts)
	if err != nil {
		log.Error("failed to load last head from database", "error", err)
		if c.NewError(err) {
			return nil
		}
		return err
	}

	fromMap := map[common.Address]*big.Int{}

	if !c.nonArchivalRPCNode {
		fromMap, err = findFirstRanges(parent, accountsWithoutHistory, head.Number, c.client)
		if err != nil {
			if c.NewError(err) {
				return nil
			}
			return err
		}
	}

	target := head.Number
	fromByAddress := map[common.Address]*LastKnownBlock{}
	toByAddress := map[common.Address]*big.Int{}

	for _, address := range c.accounts {
		from, ok := lastKnownEthBlocks[address]
		if !ok {
			from = &LastKnownBlock{Number: fromMap[address]}
		}
		if c.nonArchivalRPCNode {
			from = &LastKnownBlock{Number: big.NewInt(0).Sub(target, big.NewInt(100))}
		}

		fromByAddress[address] = from
		toByAddress[address] = target
	}

	bCache := newBalanceCache()
	cmnd := &findAndCheckBlockRangeCommand{
		accounts:      c.accounts,
		db:            c.db,
		chain:         c.chain,
		client:        c.client,
		balanceCache:  bCache,
		feed:          c.feed,
		fromByAddress: fromByAddress,
		toByAddress:   toByAddress,
	}

	err = cmnd.Command()(parent)
	if err != nil {
		if c.NewError(err) {
			return nil
		}
		return err
	}

	if cmnd.error != nil {
		if c.NewError(cmnd.error) {
			return nil
		}
		return cmnd.error
	}

	downloader := &ETHTransferDownloader{
		chain:    c.chain,
		client:   c.client,
		accounts: c.accounts,
		signer:   types.NewEIP155Signer(c.chain),
		db:       c.db,
	}
	_, err = c.LoadTransfers(parent, downloader, 40)
	if err != nil {
		if c.NewError(err) {
			return nil
		}
		return err
	}

	events := map[common.Address]Event{}
	for _, address := range c.accounts {
		event := Event{
			Type:     EventNewTransfers,
			Accounts: []common.Address{address},
		}
		for _, header := range cmnd.foundHeaders[address] {
			if event.BlockNumber == nil || header.Number.Cmp(event.BlockNumber) == 1 {
				event.BlockNumber = header.Number
			}
		}
		if event.BlockNumber != nil {
			events[address] = event
		}
	}

	for _, event := range events {
		c.feed.Send(event)
	}

	c.feed.Send(Event{
		Type:        EventRecentHistoryReady,
		Accounts:    c.accounts,
		BlockNumber: target,
	})

	log.Info("end control command")
	return err
}

func nonArchivalNodeError(err error) bool {
	return strings.Contains(err.Error(), "missing trie node") ||
		strings.Contains(err.Error(), "project ID does not have access to archive state")
}

func (c *controlCommand) NewError(err error) bool {
	c.errorsCount++
	log.Error("controlCommand error", "error", err, "counter", c.errorsCount)
	if nonArchivalNodeError(err) {
		log.Info("Non archival node detected")
		c.nonArchivalRPCNode = true
		c.feed.Send(Event{
			Type: EventNonArchivalNodeDetected,
		})
	}
	if c.errorsCount >= 3 {
		c.feed.Send(Event{
			Type:    EventFetchingHistoryError,
			Message: err.Error(),
		})
		return true
	}
	return false
}

func (c *controlCommand) Command() Command {
	return FiniteCommand{
		Interval: 5 * time.Second,
		Runable:  c.Run,
	}.Run
}

type transfersCommand struct {
	db               *Database
	eth              *ETHTransferDownloader
	block            *big.Int
	address          common.Address
	client           *walletClient
	fetchedTransfers []Transfer
}

func (c *transfersCommand) Command() Command {
	return FiniteCommand{
		Interval: 5 * time.Second,
		Runable:  c.Run,
	}.Run
}

func (c *transfersCommand) Run(ctx context.Context) (err error) {
	allTransfers, err := getTransfersByBlocks(ctx, c.db, c.eth, c.address, []*big.Int{c.block})
	if err != nil {
		log.Info("getTransfersByBlocks error", "error", err)
		return err
	}

	err = c.db.SaveTranfers(c.address, allTransfers, []*big.Int{c.block})
	if err != nil {
		log.Error("SaveTranfers error", "error", err)
		return err
	}

	c.fetchedTransfers = allTransfers
	log.Debug("transfers loaded", "address", c.address, "len", len(allTransfers))
	return nil
}

type loadTransfersCommand struct {
	accounts                []common.Address
	db                      *Database
	chain                   *big.Int
	client                  *walletClient
	blocksByAddress         map[common.Address][]*big.Int
	foundTransfersByAddress map[common.Address][]Transfer
}

func (c *loadTransfersCommand) Command() Command {
	return FiniteCommand{
		Interval: 5 * time.Second,
		Runable:  c.Run,
	}.Run
}

func (c *loadTransfersCommand) LoadTransfers(ctx context.Context, downloader *ETHTransferDownloader, limit int, blocksByAddress map[common.Address][]*big.Int) (map[common.Address][]Transfer, error) {
	return loadTransfers(ctx, c.accounts, c.db, c.client, c.chain, limit, blocksByAddress)
}

func (c *loadTransfersCommand) Run(parent context.Context) (err error) {
	downloader := &ETHTransferDownloader{
		chain:    c.chain,
		client:   c.client,
		accounts: c.accounts,
		signer:   types.NewEIP155Signer(c.chain),
		db:       c.db,
	}
	transfersByAddress, err := c.LoadTransfers(parent, downloader, 40, c.blocksByAddress)
	if err != nil {
		return err
	}
	c.foundTransfersByAddress = transfersByAddress

	return
}

type findAndCheckBlockRangeCommand struct {
	accounts      []common.Address
	db            *Database
	chain         *big.Int
	client        *walletClient
	balanceCache  *balanceCache
	feed          *event.Feed
	fromByAddress map[common.Address]*LastKnownBlock
	toByAddress   map[common.Address]*big.Int
	foundHeaders  map[common.Address][]*DBHeader
	noLimit       bool
	error         error
}

func (c *findAndCheckBlockRangeCommand) Command() Command {
	return FiniteCommand{
		Interval: 5 * time.Second,
		Runable:  c.Run,
	}.Run
}

func (c *findAndCheckBlockRangeCommand) Run(parent context.Context) (err error) {
	log.Debug("start findAndCHeckBlockRangeCommand")
	newFromByAddress, ethHeadersByAddress, err := c.fastIndex(parent, c.balanceCache, c.fromByAddress, c.toByAddress)
	if err != nil {
		c.error = err
		return nil
	}
	if c.noLimit {
		newFromByAddress = map[common.Address]*big.Int{}
		for _, address := range c.accounts {
			newFromByAddress[address] = c.fromByAddress[address].Number
		}
	}
	erc20HeadersByAddress, err := c.fastIndexErc20(parent, newFromByAddress, c.toByAddress)
	if err != nil {
		return err
	}

	foundHeaders := map[common.Address][]*DBHeader{}
	maxBlockNumber := big.NewInt(0)
	for _, address := range c.accounts {
		ethHeaders := ethHeadersByAddress[address]
		erc20Headers := erc20HeadersByAddress[address]
		allHeaders := append(ethHeaders, erc20Headers...)

		uniqHeadersByHash := map[common.Hash]*DBHeader{}
		for _, header := range allHeaders {
			uniqHeader, ok := uniqHeadersByHash[header.Hash]
			if ok {
				if len(header.Erc20Transfers) > 0 {
					uniqHeader.Erc20Transfers = append(uniqHeader.Erc20Transfers, header.Erc20Transfers...)
				}
				uniqHeadersByHash[header.Hash] = uniqHeader
			} else {
				uniqHeadersByHash[header.Hash] = header
			}
		}

		uniqHeaders := []*DBHeader{}
		for _, header := range uniqHeadersByHash {
			uniqHeaders = append(uniqHeaders, header)
		}

		foundHeaders[address] = uniqHeaders

		for _, header := range allHeaders {
			if header.Number.Cmp(maxBlockNumber) == 1 {
				maxBlockNumber = header.Number
			}
		}

		lastBlockNumber := c.toByAddress[address]
		log.Debug("saving headers", "len", len(uniqHeaders), "lastBlockNumber", lastBlockNumber, "balance", c.balanceCache.ReadCachedBalance(address, lastBlockNumber), "nonce", c.balanceCache.ReadCachedNonce(address, lastBlockNumber))
		to := &LastKnownBlock{
			Number:  lastBlockNumber,
			Balance: c.balanceCache.ReadCachedBalance(address, lastBlockNumber),
			Nonce:   c.balanceCache.ReadCachedNonce(address, lastBlockNumber),
		}
		err = c.db.ProcessBlocks(address, newFromByAddress[address], to, uniqHeaders)
		if err != nil {
			return err
		}
	}

	c.foundHeaders = foundHeaders

	return
}
