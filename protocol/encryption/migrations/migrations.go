// Code generated by go-bindata. DO NOT EDIT.
// sources:
// 1536754952_initial_schema.down.sql (83B)
// 1536754952_initial_schema.up.sql (962B)
// 1539249977_update_ratchet_info.down.sql (311B)
// 1539249977_update_ratchet_info.up.sql (368B)
// 1540715431_add_version.down.sql (127B)
// 1540715431_add_version.up.sql (265B)
// 1541164797_add_installations.down.sql (26B)
// 1541164797_add_installations.up.sql (216B)
// 1558084410_add_secret.down.sql (56B)
// 1558084410_add_secret.up.sql (301B)
// 1558588866_add_version.down.sql (47B)
// 1558588866_add_version.up.sql (57B)
// 1559627659_add_contact_code.down.sql (32B)
// 1559627659_add_contact_code.up.sql (198B)
// 1561368210_add_installation_metadata.down.sql (35B)
// 1561368210_add_installation_metadata.up.sql (267B)
// doc.go (377B)

package migrations

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var __1536754952_initial_schemaDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\x28\x4e\x2d\x2e\xce\xcc\xcf\x2b\xb6\xe6\x42\x12\x4c\x2a\xcd\x4b\xc9\x49\x45\x15\xcb\x4e\xad\x44\x15\x28\x4a\x2c\x49\xce\x48\x2d\x89\xcf\xcc\x4b\xcb\xb7\xe6\x02\x04\x00\x00\xff\xff\x72\x61\x3f\x92\x53\x00\x00\x00")

func _1536754952_initial_schemaDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__1536754952_initial_schemaDownSql,
		"1536754952_initial_schema.down.sql",
	)
}

func _1536754952_initial_schemaDownSql() (*asset, error) {
	bytes, err := _1536754952_initial_schemaDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1536754952_initial_schema.down.sql", size: 83, mode: os.FileMode(0644), modTime: time.Unix(1599559876, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x44, 0xcf, 0x76, 0x71, 0x1f, 0x5e, 0x9a, 0x43, 0xd8, 0xcd, 0xb8, 0xc3, 0x70, 0xc3, 0x7f, 0xfc, 0x90, 0xb4, 0x25, 0x1e, 0xf4, 0x66, 0x20, 0xb8, 0x33, 0x7e, 0xb0, 0x76, 0x1f, 0xc, 0xc0, 0x75}}
	return a, nil
}

var __1536754952_initial_schemaUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x52\xc1\x8e\x9b\x30\x10\xbd\xe7\x2b\xe6\x98\x48\x39\xf4\xde\x13\xb0\x13\x84\x4a\xcd\xd6\x0b\x52\xf7\x64\x79\xe3\x69\xb0\x16\x1b\x64\x3b\xab\xe6\xef\x2b\x20\xa5\xb8\x65\xdb\xde\x98\xc7\x9b\x99\x37\xef\x39\xe3\x98\xd4\x08\x75\x92\x96\x08\x9e\xbc\xd7\xbd\xf5\xb0\xdf\x01\xa8\xd6\x41\x5a\x56\xe9\x71\xfa\xf6\x62\xb8\xbe\x74\xfa\x1c\x43\x4e\xbf\xc9\x40\x0b\xe6\xfa\x3e\x88\x73\x2b\xb5\x15\xaf\x74\x5b\x60\x4f\x56\xfd\x1d\xb6\x50\xb0\x1a\x73\xe4\xd3\x14\x3a\xbf\x6d\xd0\x57\x70\x44\xf7\x81\x86\x75\x3d\x58\x58\x97\x5a\x4d\x13\x80\x55\x35\xb0\xa6\x2c\xe1\x91\x17\x9f\x13\xfe\x0c\x9f\xf0\x79\xfc\xdf\xb0\xe2\x4b\x83\x7b\xad\x0e\x50\x31\xc8\x2a\x76\x2a\x8b\xac\x06\x8e\x8f\x65\x92\xe1\xee\xf0\x71\xb7\x8b\x3c\x7a\xa5\xdb\xec\xcf\xec\xc7\x22\x71\x59\x30\x0e\x35\xfe\x22\xec\xd5\xac\x75\x18\xf2\x5e\x5e\x68\x9b\x3f\x8b\x80\xfd\xbd\xef\xb8\x66\xff\xa7\xae\x97\xab\x55\x1d\xcd\xd2\xb4\x22\x1b\x74\xd8\x58\xa4\xad\x0f\xb2\xeb\x64\xd0\xbd\x15\x5a\x41\x8d\x5f\xeb\x88\x70\x8f\x34\x0e\x4a\x5f\x2c\x29\x31\xb8\x0d\xf5\x6b\x3b\x23\xa1\x45\xce\x2a\x8e\x63\x7b\xd0\x86\x7c\x90\x66\x80\x86\x3d\x15\x39\xc3\x07\x48\x8b\x7c\xf4\x26\xda\x4c\xdf\x07\xed\x48\x41\x5a\x55\x25\x26\x0c\x1e\xf0\x94\x34\x65\x0d\x1f\xfe\xbc\xd5\xc9\x70\x6e\x29\x08\x6d\xbf\xf5\xd3\xc1\xf3\xf1\xe2\xf7\xac\xa7\xb1\x43\x4b\x86\x9c\xec\xa2\x93\xde\x77\xc8\xdf\x8c\xa1\xe0\xde\x4b\xf6\x9f\x06\xde\xdf\xd3\xa2\xe8\xb8\xec\xda\x0c\x72\x6c\x39\x55\x1c\x8b\x9c\x4d\x16\xfe\x6a\x3c\x00\xc7\x13\x72\x64\x19\x3e\xfd\x4c\x77\x1f\x47\x71\x18\xad\xf9\x11\x00\x00\xff\xff\xa9\x50\xa8\xb2\xc2\x03\x00\x00")

func _1536754952_initial_schemaUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1536754952_initial_schemaUpSql,
		"1536754952_initial_schema.up.sql",
	)
}

func _1536754952_initial_schemaUpSql() (*asset, error) {
	bytes, err := _1536754952_initial_schemaUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1536754952_initial_schema.up.sql", size: 962, mode: os.FileMode(0644), modTime: time.Unix(1599559876, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xea, 0x90, 0x5a, 0x59, 0x3e, 0x3, 0xe2, 0x3c, 0x81, 0x42, 0xcd, 0x4c, 0x9a, 0xe8, 0xda, 0x93, 0x2b, 0x70, 0xa4, 0xd5, 0x29, 0x3e, 0xd5, 0xc9, 0x27, 0xb6, 0xb7, 0x65, 0xff, 0x0, 0xcb, 0xde}}
	return a, nil
}

var __1539249977_update_ratchet_infoDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8f\x41\x4b\xc4\x30\x10\x85\xef\xf9\x15\xef\xd8\xc2\x9e\xbc\xee\xa9\x8d\x53\x29\x86\x64\x8d\x29\xe8\x29\xd4\xed\xe8\x06\xdb\xec\xd2\x46\xa1\xff\x5e\x22\x52\x59\xf4\x3a\xdf\xf7\x78\x6f\x6e\xad\x39\xc0\x55\xb5\x22\xcc\x7d\x3a\x9e\x38\xf9\x10\x5f\xcf\xfe\xf3\x66\x2f\x84\xb4\x54\x39\xfa\x07\xa3\x10\xc0\xcb\x47\x1c\x46\xf6\x61\x40\xad\x4c\x0d\x6d\x1c\x74\xa7\xd4\x4e\x00\x7c\x39\xf1\xc4\x73\x3f\xfa\x77\x5e\xbf\x71\xbe\x86\x81\x63\x0a\x69\xfd\xeb\x2f\xeb\x34\x71\x9a\xc3\x71\xf3\xaf\x70\x88\x4b\xea\xc7\xb1\x4f\xe1\x1c\x73\x9f\xa3\x27\x77\x25\x74\xba\x7d\xe8\xa8\xd8\x16\xed\xb6\xae\x12\x46\x43\x1a\xdd\xa8\x56\x3a\x58\x3a\xa8\x4a\x52\x8e\x34\xc6\x52\x7b\xa7\x71\x4f\xcf\xf8\x0d\x96\xb0\xd4\x90\x25\x2d\xe9\xf1\xe7\xc1\xa5\x58\xc2\x5b\xe4\xc1\x5f\x66\xce\xf3\x4a\x51\xee\xc5\x57\x00\x00\x00\xff\xff\x69\x51\x9b\xb4\x37\x01\x00\x00")

func _1539249977_update_ratchet_infoDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__1539249977_update_ratchet_infoDownSql,
		"1539249977_update_ratchet_info.down.sql",
	)
}

func _1539249977_update_ratchet_infoDownSql() (*asset, error) {
	bytes, err := _1539249977_update_ratchet_infoDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1539249977_update_ratchet_info.down.sql", size: 311, mode: os.FileMode(0644), modTime: time.Unix(1599559876, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x1, 0xa4, 0xeb, 0xa0, 0xe6, 0xa0, 0xd4, 0x48, 0xbb, 0xad, 0x6f, 0x7d, 0x67, 0x8c, 0xbd, 0x25, 0xde, 0x1f, 0x73, 0x9a, 0xbb, 0xa8, 0xc9, 0x30, 0xb7, 0xa9, 0x7c, 0xaf, 0xb5, 0x1, 0x61, 0xdd}}
	return a, nil
}

var __1539249977_update_ratchet_infoUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8f\x41\x4f\x84\x30\x10\x85\xef\xfd\x15\x73\x84\x84\x93\x57\x4e\x50\x07\x43\xac\xed\x5a\x4b\xa2\xa7\x06\x97\xd1\x6d\x16\xca\x86\x56\x13\xfe\xbd\xa9\x31\x28\xea\xf5\xbd\x6f\xde\x7b\x73\x8d\x02\x0d\x42\xa3\xd5\x1d\x04\x0a\xc1\xcd\x3e\x94\xec\xa7\x7a\xa6\x35\x29\x5a\x1d\xc0\x54\xb5\x40\x58\xfa\x78\x3c\x51\xb4\xce\xbf\xcc\x25\x63\x5c\x63\x65\xf0\x1f\xcf\xbe\x5f\x41\xc6\x00\x9e\xdf\xfc\x30\x92\x75\x03\xd4\x42\xd5\x20\x95\x01\xd9\x09\x51\x30\x00\xba\x9c\x68\xa2\xa5\x1f\xed\x99\xd6\x4f\x3b\xa9\x6e\x20\x1f\x5d\x5c\xff\xf2\x61\x9d\x26\x8a\x8b\x3b\x6e\xfc\xce\x76\x3e\xc4\x7e\x1c\xfb\xe8\x66\x9f\xfa\x0c\x3e\x9a\x1d\xd0\xc9\xf6\xbe\xc3\x6c\x5b\x54\x6c\x5d\xc5\xef\xe3\x1c\x94\x04\xae\x64\x23\x5a\x6e\x40\xe3\x41\x54\x1c\x53\x46\xa3\x34\xb6\x37\x12\x6e\xf1\x09\xbe\x93\x72\xd0\xd8\xa0\x46\xc9\xf1\xe1\xeb\xe3\x90\x05\xf7\xea\x69\xb0\x97\x85\xd2\xde\x9c\xe5\x25\xfb\x08\x00\x00\xff\xff\xb6\x31\x2b\x32\x70\x01\x00\x00")

func _1539249977_update_ratchet_infoUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1539249977_update_ratchet_infoUpSql,
		"1539249977_update_ratchet_info.up.sql",
	)
}

func _1539249977_update_ratchet_infoUpSql() (*asset, error) {
	bytes, err := _1539249977_update_ratchet_infoUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1539249977_update_ratchet_info.up.sql", size: 368, mode: os.FileMode(0644), modTime: time.Unix(1599559876, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xc, 0x8e, 0xbf, 0x6f, 0xa, 0xc0, 0xe1, 0x3c, 0x42, 0x28, 0x88, 0x1d, 0xdb, 0xba, 0x1c, 0x83, 0xec, 0xba, 0xd3, 0x5f, 0x5c, 0x77, 0x5e, 0xa7, 0x46, 0x36, 0xec, 0x69, 0xa, 0x4b, 0x17, 0x79}}
	return a, nil
}

var __1540715431_add_versionDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\xc8\x4e\xad\x2c\x56\x70\x09\xf2\x0f\x50\x70\xf6\xf7\x09\xf5\xf5\x53\x28\x4e\x2d\x2e\xce\xcc\xcf\x8b\xcf\x4c\xb1\xe6\x42\x56\x08\x15\x47\x55\x0c\xd2\x1d\x9f\x9c\x5f\x9a\x57\x82\xaa\x38\xa9\x34\x2f\x25\x27\x15\x55\x6d\x59\x6a\x11\xc8\x00\x6b\x2e\x40\x00\x00\x00\xff\xff\xda\x5d\x80\x2d\x7f\x00\x00\x00")

func _1540715431_add_versionDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__1540715431_add_versionDownSql,
		"1540715431_add_version.down.sql",
	)
}

func _1540715431_add_versionDownSql() (*asset, error) {
	bytes, err := _1540715431_add_versionDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1540715431_add_version.down.sql", size: 127, mode: os.FileMode(0644), modTime: time.Unix(1599559876, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xf5, 0x9, 0x4, 0xe3, 0x76, 0x2e, 0xb8, 0x9, 0x23, 0xf0, 0x70, 0x93, 0xc4, 0x50, 0xe, 0x9d, 0x84, 0x22, 0x8c, 0x94, 0xd3, 0x24, 0x9, 0x9a, 0xc1, 0xa1, 0x48, 0x45, 0xfd, 0x40, 0x6e, 0xe6}}
	return a, nil
}

var __1540715431_add_versionUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\xcd\xb1\x0e\x02\x21\x0c\xc6\xf1\xdd\xa7\xf8\x1e\xc1\xdd\x09\xa4\x67\x4c\x7a\x90\x90\x32\x93\xe8\x31\x5c\x54\x2e\x8a\x98\xf8\xf6\x06\xe3\xc2\xa2\xae\x6d\xff\xbf\x1a\x62\x12\xc2\xe0\xdd\x88\x53\x7a\x96\xcd\x4a\xb1\x90\x87\x28\xcd\xf4\x9e\x40\x19\x83\xad\xe3\x30\x5a\x94\x74\x8d\xb9\x5e\xb0\xb7\x42\x3b\xf2\xb0\x4e\x60\x03\x33\x0c\x0d\x2a\xb0\x60\xfd\xab\x2f\x65\x5e\x72\x9c\x27\x68\x76\xba\x3f\xfe\x2c\xbb\xa0\x01\xf1\xb8\xd4\x7c\xff\xfb\xe7\xa1\xe6\xe9\x9c\x3a\xe5\x91\x6e\x4d\xfe\x4a\xbc\x02\x00\x00\xff\xff\x0e\x27\x2c\x52\x09\x01\x00\x00")

func _1540715431_add_versionUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1540715431_add_versionUpSql,
		"1540715431_add_version.up.sql",
	)
}

func _1540715431_add_versionUpSql() (*asset, error) {
	bytes, err := _1540715431_add_versionUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1540715431_add_version.up.sql", size: 265, mode: os.FileMode(0644), modTime: time.Unix(1599559876, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xc7, 0x4c, 0x36, 0x96, 0xdf, 0x16, 0x10, 0xa6, 0x27, 0x1a, 0x79, 0x8b, 0x42, 0x83, 0x23, 0xc, 0x7e, 0xb6, 0x3d, 0x2, 0xda, 0xa4, 0xb4, 0xd, 0x27, 0x55, 0xba, 0xdc, 0xb2, 0x88, 0x8f, 0xa6}}
	return a, nil
}

var __1541164797_add_installationsDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xc8\xcc\x2b\x2e\x49\xcc\xc9\x49\x2c\xc9\xcc\xcf\x2b\xb6\xe6\x02\x04\x00\x00\xff\xff\xd8\xbf\x14\x75\x1a\x00\x00\x00")

func _1541164797_add_installationsDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__1541164797_add_installationsDownSql,
		"1541164797_add_installations.down.sql",
	)
}

func _1541164797_add_installationsDownSql() (*asset, error) {
	bytes, err := _1541164797_add_installationsDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1541164797_add_installations.down.sql", size: 26, mode: os.FileMode(0644), modTime: time.Unix(1599559876, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xf5, 0xfd, 0xe6, 0xd8, 0xca, 0x3b, 0x38, 0x18, 0xee, 0x0, 0x5f, 0x36, 0x9e, 0x1e, 0xd, 0x19, 0x3e, 0xb4, 0x73, 0x53, 0xe9, 0xa5, 0xac, 0xdd, 0xa1, 0x2f, 0xc7, 0x6c, 0xa8, 0xd9, 0xa, 0x88}}
	return a, nil
}

var __1541164797_add_installationsUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\xce\xb1\x6a\xc3\x30\x14\x85\xe1\xdd\x4f\x71\x46\x1b\xbc\x74\xee\x24\xc9\xd7\x46\x70\xb9\x6a\x5d\x09\xba\x05\x05\x6b\x10\xd8\x4a\xc0\x5a\xf2\xf6\xc1\x43\x20\xce\xfc\x7f\x70\x8e\x99\x49\x79\x82\x57\x9a\x09\xb9\xec\x35\xae\x6b\xac\xf9\x56\x76\xa0\x6d\x80\xbc\xa4\x52\x73\x7d\x40\xb3\xd3\x10\xe7\x21\x81\xb9\x3f\xca\x1b\xbe\xe4\x05\x9e\xfe\xfd\x09\xd4\xbc\xa5\xbd\xc6\xed\x8e\x20\x7f\x76\x12\x1a\xa0\xed\x04\x2b\x67\x96\x4a\xbc\xae\x69\x81\x76\x8e\x49\x09\x06\x1a\x55\x60\x8f\xaf\x23\x06\xb1\xbf\x81\xda\xd7\x8b\xfe\x73\xb5\x83\x13\x18\x27\x23\x5b\xe3\x31\xd3\x0f\x2b\x43\x4d\xf7\xdd\x3c\x03\x00\x00\xff\xff\x28\x14\xac\x9d\xd8\x00\x00\x00")

func _1541164797_add_installationsUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1541164797_add_installationsUpSql,
		"1541164797_add_installations.up.sql",
	)
}

func _1541164797_add_installationsUpSql() (*asset, error) {
	bytes, err := _1541164797_add_installationsUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1541164797_add_installations.up.sql", size: 216, mode: os.FileMode(0644), modTime: time.Unix(1599559876, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x2d, 0x18, 0x26, 0xb8, 0x88, 0x47, 0xdb, 0x83, 0xcc, 0xb6, 0x9d, 0x1c, 0x1, 0xae, 0x2f, 0xde, 0x97, 0x82, 0x3, 0x30, 0xa8, 0x63, 0xa1, 0x78, 0x4b, 0xa5, 0x9, 0x8, 0x75, 0xa2, 0x57, 0x81}}
	return a, nil
}

var __1558084410_add_secretDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\x28\x4e\x4d\x2e\x4a\x2d\x89\xcf\xcc\x2b\x2e\x49\xcc\xc9\x49\x2c\xc9\xcc\xcf\x8b\xcf\x4c\x29\xb6\xe6\xc2\x50\x53\x6c\xcd\x05\x08\x00\x00\xff\xff\xd3\xcd\x41\x83\x38\x00\x00\x00")

func _1558084410_add_secretDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__1558084410_add_secretDownSql,
		"1558084410_add_secret.down.sql",
	)
}

func _1558084410_add_secretDownSql() (*asset, error) {
	bytes, err := _1558084410_add_secretDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1558084410_add_secret.down.sql", size: 56, mode: os.FileMode(0644), modTime: time.Unix(1599559876, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x49, 0xb, 0x65, 0xdf, 0x59, 0xbf, 0xe9, 0x5, 0x5b, 0x6f, 0xd5, 0x3a, 0xb7, 0x57, 0xe8, 0x78, 0x38, 0x73, 0x53, 0x57, 0xf7, 0x24, 0x4, 0xe4, 0xa2, 0x49, 0x22, 0xa2, 0xc6, 0xfd, 0x80, 0xa4}}
	return a, nil
}

var __1558084410_add_secretUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x50\xcf\x0a\x82\x30\x1c\xbe\xef\x29\xbe\xa3\x82\x6f\xd0\x49\xc7\x4f\x19\xad\xdf\x6a\x4d\xc8\x93\x48\xf3\x30\x10\x83\xdc\xa5\xb7\x0f\x23\x45\xa1\xce\xdf\xff\x4f\x5a\xca\x1d\xc1\xe5\x85\x26\x4c\xfd\xfd\xd9\xc7\x09\x89\x00\x82\xef\xc7\x18\xe2\x0b\x85\x36\x05\xd8\x38\x70\xad\x35\xce\x56\x9d\x72\xdb\xe0\x48\x0d\x0c\x43\x1a\x2e\xb5\x92\x0e\xaa\x62\x63\x29\x13\xf8\x9a\xec\x65\x22\x3d\x08\xf1\x23\xaa\x0d\xe3\x14\xbb\x61\xe8\x62\x78\x8c\x6d\xf0\x4b\x34\x1c\xdd\xdc\xaa\xce\x36\x75\xda\xe0\xf7\xd6\x33\x58\xb3\xba\xd4\x94\x04\x9f\x6d\x79\xe9\x9f\x82\xa5\xb1\xa4\x2a\xfe\x4c\x48\x76\x7c\x4b\x25\x59\x62\x49\xd7\xe5\x8a\x15\x4f\xe7\x09\xef\x00\x00\x00\xff\xff\xa6\xbb\x2c\x23\x2d\x01\x00\x00")

func _1558084410_add_secretUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1558084410_add_secretUpSql,
		"1558084410_add_secret.up.sql",
	)
}

func _1558084410_add_secretUpSql() (*asset, error) {
	bytes, err := _1558084410_add_secretUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1558084410_add_secret.up.sql", size: 301, mode: os.FileMode(0644), modTime: time.Unix(1599559876, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xf5, 0x32, 0x36, 0x8e, 0x47, 0xb0, 0x8f, 0xc1, 0xc6, 0xf7, 0xc6, 0x9f, 0x2d, 0x44, 0x75, 0x2b, 0x26, 0xec, 0x6, 0xa0, 0x7b, 0xa5, 0xbd, 0xc8, 0x76, 0x8a, 0x82, 0x68, 0x2, 0x42, 0xb5, 0xf4}}
	return a, nil
}

var __1558588866_add_versionDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\xc8\xcc\x2b\x2e\x49\xcc\xc9\x49\x2c\xc9\xcc\xcf\x2b\x56\x70\x09\xf2\x0f\x50\x70\xf6\xf7\x09\xf5\xf5\x53\x28\x4b\x2d\x2a\xce\xcc\xcf\xb3\xe6\x02\x04\x00\x00\xff\xff\xdf\x6b\x9f\xbb\x2f\x00\x00\x00")

func _1558588866_add_versionDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__1558588866_add_versionDownSql,
		"1558588866_add_version.down.sql",
	)
}

func _1558588866_add_versionDownSql() (*asset, error) {
	bytes, err := _1558588866_add_versionDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1558588866_add_version.down.sql", size: 47, mode: os.FileMode(0644), modTime: time.Unix(1599559876, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xde, 0x52, 0x34, 0x3c, 0x46, 0x4a, 0xf0, 0x72, 0x47, 0x6f, 0x49, 0x5c, 0xc7, 0xf9, 0x32, 0xce, 0xc4, 0x3d, 0xfd, 0x61, 0xa1, 0x8b, 0x8f, 0xf2, 0x31, 0x34, 0xde, 0x15, 0x49, 0xa6, 0xde, 0xb9}}
	return a, nil
}

var __1558588866_add_versionUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\xc8\xcc\x2b\x2e\x49\xcc\xc9\x49\x2c\xc9\xcc\xcf\x2b\x56\x70\x74\x71\x51\x28\x4b\x2d\x2a\xce\xcc\xcf\x53\xf0\xf4\x0b\x71\x75\x77\x0d\x52\x70\x71\x75\x73\x0c\xf5\x09\x51\x30\xb0\xe6\x02\x04\x00\x00\xff\xff\x14\x7b\x07\xb5\x39\x00\x00\x00")

func _1558588866_add_versionUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1558588866_add_versionUpSql,
		"1558588866_add_version.up.sql",
	)
}

func _1558588866_add_versionUpSql() (*asset, error) {
	bytes, err := _1558588866_add_versionUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1558588866_add_version.up.sql", size: 57, mode: os.FileMode(0644), modTime: time.Unix(1599559876, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x2a, 0xea, 0x64, 0x39, 0x61, 0x20, 0x83, 0x83, 0xb, 0x2e, 0x79, 0x64, 0xb, 0x53, 0xfa, 0xfe, 0xc6, 0xf7, 0x67, 0x42, 0xd3, 0x4f, 0xdc, 0x7e, 0x30, 0x32, 0xe8, 0x14, 0x41, 0xe9, 0xe7, 0x3b}}
	return a, nil
}

var __1559627659_add_contact_codeDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\x48\xce\xcf\x2b\x49\x4c\x2e\x89\x4f\xce\x4f\x49\x8d\x4f\xce\xcf\x4b\xcb\x4c\xb7\xe6\x02\x04\x00\x00\xff\xff\x73\x7b\x50\x80\x20\x00\x00\x00")

func _1559627659_add_contact_codeDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__1559627659_add_contact_codeDownSql,
		"1559627659_add_contact_code.down.sql",
	)
}

func _1559627659_add_contact_codeDownSql() (*asset, error) {
	bytes, err := _1559627659_add_contact_codeDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1559627659_add_contact_code.down.sql", size: 32, mode: os.FileMode(0644), modTime: time.Unix(1599559876, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x5d, 0x64, 0x6d, 0xce, 0x24, 0x42, 0x20, 0x8d, 0x4f, 0x37, 0xaa, 0x9d, 0xc, 0x57, 0x98, 0xc1, 0xd1, 0x1a, 0x34, 0xcd, 0x9f, 0x8f, 0x34, 0x86, 0xb3, 0xd3, 0xdc, 0xf1, 0x7d, 0xe5, 0x1b, 0x6e}}
	return a, nil
}

var __1559627659_add_contact_codeUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\xce\xc1\x8e\x82\x30\x18\x04\xe0\x7b\x9f\x62\x6e\x40\xb2\x07\xf6\xcc\xa9\xbb\xfb\xaf\x21\xd6\x62\x4a\x31\x72\x22\xb5\xa0\x34\x21\x45\xa1\xf8\xfc\x06\x13\xe3\xc5\xeb\xe4\x9b\xc9\xfc\x2a\xe2\x9a\xa0\xf9\x8f\x20\xd8\xd1\x07\x63\x43\x63\xc7\xb6\x6b\xec\xe8\xcf\xee\x82\x98\x01\x58\xbc\xbb\x2d\xcf\x68\x0e\x93\x71\x3e\xe0\x6e\x26\xdb\x9b\x29\xfe\x4e\x20\x0b\x0d\x59\x09\x81\xbd\xca\x77\x5c\xd5\xd8\x52\x8d\x3f\xfa\xe7\x95\xd0\x88\x8e\xd1\x17\x03\x06\x33\x87\xe6\xba\x9c\x06\x37\xf7\x5d\x8b\x5c\x6a\xda\x90\x7a\x57\x5f\x3c\x65\x49\xc6\x58\x2e\x4b\x52\x7a\x55\xc5\xc7\x4f\x07\x2e\x2a\x2a\x11\xaf\xe3\x48\x93\x8c\x3d\x02\x00\x00\xff\xff\xdc\x7c\x0c\xd3\xc6\x00\x00\x00")

func _1559627659_add_contact_codeUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1559627659_add_contact_codeUpSql,
		"1559627659_add_contact_code.up.sql",
	)
}

func _1559627659_add_contact_codeUpSql() (*asset, error) {
	bytes, err := _1559627659_add_contact_codeUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1559627659_add_contact_code.up.sql", size: 198, mode: os.FileMode(0644), modTime: time.Unix(1599559876, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x16, 0xf6, 0xc2, 0x62, 0x9c, 0xd2, 0xc9, 0x1e, 0xd8, 0xea, 0xaa, 0xea, 0x95, 0x8f, 0x89, 0x6a, 0x85, 0x5d, 0x9d, 0x99, 0x78, 0x3c, 0x90, 0x66, 0x99, 0x3e, 0x4b, 0x19, 0x62, 0xfb, 0x31, 0x4d}}
	return a, nil
}

var __1561368210_add_installation_metadataDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xc8\xcc\x2b\x2e\x49\xcc\xc9\x49\x2c\xc9\xcc\xcf\x2b\x8e\xcf\x4d\x2d\x49\x4c\x49\x2c\x49\xb4\xe6\x02\x04\x00\x00\xff\xff\x03\x72\x7f\x08\x23\x00\x00\x00")

func _1561368210_add_installation_metadataDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__1561368210_add_installation_metadataDownSql,
		"1561368210_add_installation_metadata.down.sql",
	)
}

func _1561368210_add_installation_metadataDownSql() (*asset, error) {
	bytes, err := _1561368210_add_installation_metadataDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1561368210_add_installation_metadata.down.sql", size: 35, mode: os.FileMode(0644), modTime: time.Unix(1599559876, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xa8, 0xde, 0x3f, 0xd2, 0x4a, 0x50, 0x98, 0x56, 0xe3, 0xc0, 0xcd, 0x9d, 0xb0, 0x34, 0x3b, 0xe5, 0x62, 0x18, 0xb5, 0x20, 0xc9, 0x3e, 0xdc, 0x6a, 0x40, 0x36, 0x66, 0xea, 0x51, 0x8c, 0x71, 0xf5}}
	return a, nil
}

var __1561368210_add_installation_metadataUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\xce\xc1\x8a\x83\x30\x10\xc6\xf1\xbb\x4f\xf1\xdd\x54\xf0\x0d\xf6\x14\xb3\x23\x08\x21\xd9\x95\x04\x7a\x93\x60\x52\x08\xd5\x58\xe8\x50\xf0\xed\x8b\x87\x42\xed\xc1\xeb\xcc\xef\x83\xbf\x1c\x48\x58\x82\x15\xad\x22\xa4\xfc\x60\x3f\xcf\x9e\xd3\x9a\xc7\x25\xb2\x0f\x9e\x3d\x50\x15\x40\x0a\x31\x73\xe2\x0d\xad\x32\x2d\xb4\xb1\xd0\x4e\xa9\x66\xff\x7c\x8e\x52\x80\xa5\x8b\x3d\x80\xec\x97\x78\xbc\xe2\x97\x3a\xe1\x94\x45\x59\xee\x20\xc4\x67\x9a\xe2\xc8\xdb\xfd\xdc\x5d\xa7\x65\xe4\xf5\x16\xf3\xa9\x72\xba\xff\x77\x54\xbd\x83\x9b\xef\xc0\x1a\x46\x43\x1a\xdd\xa9\x5e\x5a\x0c\xf4\xa7\x84\xa4\xa2\xfe\x29\x5e\x01\x00\x00\xff\xff\x5d\x6f\xe6\xd3\x0b\x01\x00\x00")

func _1561368210_add_installation_metadataUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1561368210_add_installation_metadataUpSql,
		"1561368210_add_installation_metadata.up.sql",
	)
}

func _1561368210_add_installation_metadataUpSql() (*asset, error) {
	bytes, err := _1561368210_add_installation_metadataUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1561368210_add_installation_metadata.up.sql", size: 267, mode: os.FileMode(0644), modTime: time.Unix(1599559876, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xb4, 0x71, 0x8f, 0x29, 0xb1, 0xaa, 0xd6, 0xd1, 0x8c, 0x17, 0xef, 0x6c, 0xd5, 0x80, 0xb8, 0x2c, 0xc3, 0xfe, 0xec, 0x24, 0x4d, 0xc8, 0x25, 0xd3, 0xb4, 0xcd, 0xa9, 0xac, 0x63, 0x61, 0xb2, 0x9c}}
	return a, nil
}

var _docGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x8f\xbb\x6e\xc3\x30\x0c\x45\x77\x7f\xc5\x45\x96\x2c\xb5\xb4\x74\xea\xd6\xb1\x7b\x7f\x80\x91\x68\x89\x88\x1e\xae\x48\xe7\xf1\xf7\x85\xd3\x02\xcd\xd6\xf5\x00\xe7\xf0\xd2\x7b\x7c\x66\x51\x2c\x52\x18\xa2\x68\x1c\x58\x95\xc6\x1d\x27\x0e\xb4\x29\xe3\x90\xc4\xf2\x76\x72\xa1\x57\xaf\x46\xb6\xe9\x2c\xd5\x57\x49\x83\x8c\xfd\xe5\xf5\x30\x79\x8f\x40\xed\x68\xc8\xd4\x62\xe1\x47\x4b\xa1\x46\xc3\xa4\x25\x5c\xc5\x32\x08\xeb\xe0\x45\x6e\x0e\xef\x86\xc2\xa4\x06\xcb\x64\x47\x85\x65\x46\x20\xe5\x3d\xb3\xf4\x81\xd4\xe7\x93\xb4\x48\x46\x6e\x47\x1f\xcb\x13\xd9\x17\x06\x2a\x85\x23\x96\xd1\xeb\xc3\x55\xaa\x8c\x28\x83\x83\xf5\x71\x7f\x01\xa9\xb2\xa1\x51\x65\xdd\xfd\x4c\x17\x46\xeb\xbf\xe7\x41\x2d\xfe\xff\x11\xae\x7d\x9c\x15\xa4\xe0\xdb\xca\xc1\x38\xba\x69\x5a\x29\x9c\x29\x31\xf4\xab\x88\xf1\x34\x79\x9f\xfa\x5b\xe2\xc6\xbb\xf5\xbc\x71\x5e\xcf\x09\x3f\x35\xe9\x4d\x31\x77\x38\xe7\xff\x80\x4b\x1d\x6e\xfa\x0e\x00\x00\xff\xff\x9d\x60\x3d\x88\x79\x01\x00\x00")

func docGoBytes() ([]byte, error) {
	return bindataRead(
		_docGo,
		"doc.go",
	)
}

func docGo() (*asset, error) {
	bytes, err := docGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "doc.go", size: 377, mode: os.FileMode(0644), modTime: time.Unix(1603694100, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xef, 0xaf, 0xdf, 0xcf, 0x65, 0xae, 0x19, 0xfc, 0x9d, 0x29, 0xc1, 0x91, 0xaf, 0xb5, 0xd5, 0xb1, 0x56, 0xf3, 0xee, 0xa8, 0xba, 0x13, 0x65, 0xdb, 0xab, 0xcf, 0x4e, 0xac, 0x92, 0xe9, 0x60, 0xf1}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"1536754952_initial_schema.down.sql": _1536754952_initial_schemaDownSql,

	"1536754952_initial_schema.up.sql": _1536754952_initial_schemaUpSql,

	"1539249977_update_ratchet_info.down.sql": _1539249977_update_ratchet_infoDownSql,

	"1539249977_update_ratchet_info.up.sql": _1539249977_update_ratchet_infoUpSql,

	"1540715431_add_version.down.sql": _1540715431_add_versionDownSql,

	"1540715431_add_version.up.sql": _1540715431_add_versionUpSql,

	"1541164797_add_installations.down.sql": _1541164797_add_installationsDownSql,

	"1541164797_add_installations.up.sql": _1541164797_add_installationsUpSql,

	"1558084410_add_secret.down.sql": _1558084410_add_secretDownSql,

	"1558084410_add_secret.up.sql": _1558084410_add_secretUpSql,

	"1558588866_add_version.down.sql": _1558588866_add_versionDownSql,

	"1558588866_add_version.up.sql": _1558588866_add_versionUpSql,

	"1559627659_add_contact_code.down.sql": _1559627659_add_contact_codeDownSql,

	"1559627659_add_contact_code.up.sql": _1559627659_add_contact_codeUpSql,

	"1561368210_add_installation_metadata.down.sql": _1561368210_add_installation_metadataDownSql,

	"1561368210_add_installation_metadata.up.sql": _1561368210_add_installation_metadataUpSql,

	"doc.go": docGo,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"1536754952_initial_schema.down.sql":            &bintree{_1536754952_initial_schemaDownSql, map[string]*bintree{}},
	"1536754952_initial_schema.up.sql":              &bintree{_1536754952_initial_schemaUpSql, map[string]*bintree{}},
	"1539249977_update_ratchet_info.down.sql":       &bintree{_1539249977_update_ratchet_infoDownSql, map[string]*bintree{}},
	"1539249977_update_ratchet_info.up.sql":         &bintree{_1539249977_update_ratchet_infoUpSql, map[string]*bintree{}},
	"1540715431_add_version.down.sql":               &bintree{_1540715431_add_versionDownSql, map[string]*bintree{}},
	"1540715431_add_version.up.sql":                 &bintree{_1540715431_add_versionUpSql, map[string]*bintree{}},
	"1541164797_add_installations.down.sql":         &bintree{_1541164797_add_installationsDownSql, map[string]*bintree{}},
	"1541164797_add_installations.up.sql":           &bintree{_1541164797_add_installationsUpSql, map[string]*bintree{}},
	"1558084410_add_secret.down.sql":                &bintree{_1558084410_add_secretDownSql, map[string]*bintree{}},
	"1558084410_add_secret.up.sql":                  &bintree{_1558084410_add_secretUpSql, map[string]*bintree{}},
	"1558588866_add_version.down.sql":               &bintree{_1558588866_add_versionDownSql, map[string]*bintree{}},
	"1558588866_add_version.up.sql":                 &bintree{_1558588866_add_versionUpSql, map[string]*bintree{}},
	"1559627659_add_contact_code.down.sql":          &bintree{_1559627659_add_contact_codeDownSql, map[string]*bintree{}},
	"1559627659_add_contact_code.up.sql":            &bintree{_1559627659_add_contact_codeUpSql, map[string]*bintree{}},
	"1561368210_add_installation_metadata.down.sql": &bintree{_1561368210_add_installation_metadataDownSql, map[string]*bintree{}},
	"1561368210_add_installation_metadata.up.sql":   &bintree{_1561368210_add_installation_metadataUpSql, map[string]*bintree{}},
	"doc.go": &bintree{docGo, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory.
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}