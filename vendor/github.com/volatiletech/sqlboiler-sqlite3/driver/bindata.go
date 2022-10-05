// Code generated by go-bindata. DO NOT EDIT.
// sources:
// override/templates_test/singleton/sqlite3_main_test.go.tpl (1.665kB)

package driver

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
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
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

var _templates_testSingletonSqlite3_main_testGoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x54\x5b\x4f\xdb\x3e\x14\x7f\x8e\x3f\xc5\xf9\x47\xe2\x2f\x1b\x82\x99\xc4\x1b\x53\x54\x8d\x56\x9b\x76\x81\xb1\x95\x89\x17\xa4\x11\x92\x93\x72\x44\x62\x1b\xdb\x25\x54\xa8\xdf\x7d\x72\xd2\x96\x04\x15\xb6\x87\xad\x2f\x95\x4e\x8e\x7f\x97\x73\xbb\xcf\x2c\xd8\xd9\xc3\xf4\xdb\x17\xf2\x78\x8b\x0b\x48\xc1\xe2\x0c\x1f\x8c\x3c\x99\x3b\x3f\xd6\xb5\xa1\x0a\xf9\x15\x1f\xd5\x24\x38\x4f\x2e\x95\x18\x5d\xba\xbd\x52\x5b\xa4\x99\x82\x5b\x5c\xc8\xdd\xd1\xa5\x12\x7b\x57\x82\x31\xbf\x30\x08\xee\xae\x22\x8f\xe7\xe8\x3c\x5a\x70\xde\xce\x73\x0f\x8f\x2c\x2a\xae\xc7\x5a\x29\xd8\x75\x77\x95\x9c\x1c\xb3\x10\x38\xcd\x6a\x8c\x9c\xb7\xa4\x66\x2c\xf2\xe8\xfc\xe4\x38\x84\x60\x15\x5a\x32\x56\xce\x55\x0e\xa4\xc8\x73\xd1\x61\x9c\x64\xa4\x20\x85\xff\xfb\x24\x8f\xcb\x4d\x2a\x77\x2d\xc1\xe6\x93\x00\x87\x7e\x6e\xb8\x00\xb4\x56\xdb\x80\x11\xfc\xa2\xb5\x5d\x80\x31\x00\x00\x27\x3b\x2d\x90\xc2\x3d\x19\xb4\xf2\x03\xfa\x69\xab\x81\xc7\x1d\xda\xa1\x2c\xae\x55\x56\x63\x2c\xda\x07\x54\x42\x85\x8a\xaf\xdf\x09\x48\x53\x78\x03\x8f\xed\xb7\xf0\xb3\xe8\xe7\x56\x75\x14\x4e\x9e\x62\xc3\x63\xa5\xa1\x83\x00\x67\x30\xa7\x92\xb0\x58\x81\x2d\x19\x8b\x9c\xec\xd9\x4f\xa1\xa4\x0a\x4d\xe6\x6f\xe4\x27\x4d\x8a\x6b\x27\xcf\xb1\x36\x13\xb2\x5c\x24\x50\xd6\x5e\x4e\x8d\x25\xe5\x4b\x1e\x5f\x6b\xaa\xf6\x57\x12\xf7\x77\x0a\xe9\xee\xaa\x38\x01\x9b\xa9\x42\x7e\x54\x9e\x0b\x21\x42\xa1\xe7\xb5\x19\xd7\x05\x1c\xa5\x80\x0f\x98\xcb\xb1\xae\xeb\x4c\x15\x1b\x6f\x71\x02\xf1\x7e\x5e\x17\xe1\x5f\x86\xe4\x38\xd9\x94\x44\xb0\x28\xb7\x98\x79\xfc\x0d\x40\xdf\x40\xe0\xb4\x09\x34\x21\x9f\xb4\x3c\x23\x83\x5c\x6c\x64\xc8\xa9\x2f\xf4\xdc\x43\x0a\x4d\x0f\x3b\x44\xdb\xce\x2a\x6c\xde\x7f\xc6\xc5\x04\x9d\xb7\x7a\x81\x96\xf7\x67\x33\x01\x1b\xc0\xa9\x6c\x3b\x98\xc2\x13\x64\x66\x3d\x17\x6f\xdb\xf0\x7f\x29\x28\xaa\x42\xab\xa3\x61\x1f\x2e\x6c\x66\x38\x5a\x9b\x40\x5c\x66\x54\x61\x01\x5e\x83\x0b\x4f\x57\x33\x7b\xd8\x02\x42\xde\xd9\x8b\x05\x8b\x96\x3d\xb2\xbe\xd6\xbf\x44\xd7\x41\x0e\x09\xb7\xd8\xbb\xc8\x68\x1b\x5d\x98\x84\xb3\x30\x08\x95\x0a\x3c\xe2\xcf\x14\x34\x19\x79\x28\xb5\x7d\xc5\x33\x8b\x1a\x39\xae\xb4\x43\x2e\xe0\xe0\x00\xde\x95\x61\x93\xd7\x53\x44\x0e\x0a\xad\x30\x81\x3c\x64\x80\xbf\x41\x68\x2c\x79\x04\x54\x05\xe8\xb2\x0d\x18\x32\xc8\xb6\x96\xee\xdf\x5a\xd9\x5e\xcf\x15\x92\xa2\xea\x95\x3b\xe1\x31\xb3\x85\x6e\x54\xff\x54\x50\xd9\xee\x41\x7b\xb7\x7a\x62\xd7\xb1\x75\x8d\x06\x2c\xda\xc9\xef\x58\xeb\x7b\xe4\xc3\x95\x78\x99\x39\xd7\x2a\xb0\xf2\xd5\x65\x4c\x3a\x7e\xf1\xa2\x80\x67\x47\x66\x9d\x90\xb4\x06\x5b\x29\xc3\x13\x17\x3d\x65\x74\xdd\x08\x3c\x5f\x0d\xaa\xfe\xfe\x0e\xae\x4a\x38\x3f\x47\x3b\x6e\xf4\xb3\xd2\x79\xfa\xe3\x7c\xfc\x7c\xbf\xc5\x46\xc3\xaa\xc3\x2f\x69\x53\x54\xb5\xa4\x83\x0a\x0d\x05\x2f\xd9\xaf\x00\x00\x00\xff\xff\x89\xba\x4f\x58\x81\x06\x00\x00")

func templates_testSingletonSqlite3_main_testGoTplBytes() ([]byte, error) {
	return bindataRead(
		_templates_testSingletonSqlite3_main_testGoTpl,
		"templates_test/singleton/sqlite3_main_test.go.tpl",
	)
}

func templates_testSingletonSqlite3_main_testGoTpl() (*asset, error) {
	bytes, err := templates_testSingletonSqlite3_main_testGoTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates_test/singleton/sqlite3_main_test.go.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x78, 0x99, 0x6d, 0x40, 0xb2, 0xcc, 0x52, 0x9f, 0x55, 0x7a, 0xaa, 0x4c, 0xf9, 0xbe, 0xe, 0xe, 0xf4, 0x98, 0xbf, 0x1c, 0xdb, 0x1d, 0x83, 0xdf, 0x45, 0xc4, 0xc, 0x62, 0x36, 0xbe, 0xd7, 0x53}}
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
	"templates_test/singleton/sqlite3_main_test.go.tpl": templates_testSingletonSqlite3_main_testGoTpl,
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
	"templates_test": &bintree{nil, map[string]*bintree{
		"singleton": &bintree{nil, map[string]*bintree{
			"sqlite3_main_test.go.tpl": &bintree{templates_testSingletonSqlite3_main_testGoTpl, map[string]*bintree{}},
		}},
	}},
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
