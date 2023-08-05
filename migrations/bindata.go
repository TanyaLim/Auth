package migrations

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

var __000001_init_down_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xc8\xc9\x2c\x2e\x29\x8e\xcf\x2c\x49\xcd\x2d\xb6\xe6\xe5\xe2\xe5\x42\x92\x2a\x2d\x4e\x2d\x2a\x8e\x07\x2b\x40\x97\x2a\xc9\x4f\xc9\xc7\x2e\x03\xd6\x84\x55\x39\xc2\x0e\x40\x00\x00\x00\xff\xff\xe1\xdd\xfe\xe4\x7f\x00\x00\x00")

func _000001_init_down_sql() ([]byte, error) {
	return bindata_read(
		__000001_init_down_sql,
		"000001_init.down.sql",
	)
}

var __000001_init_up_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x52\xbd\x6e\x83\x40\x0c\xde\x91\x78\x07\x8f\x20\x75\xaa\x94\xa9\x53\x5a\xb1\x75\xaa\xb2\x23\x97\x33\xc2\xd2\xc5\x97\x9e\x4d\xfb\xfa\x15\x84\x84\x12\x1d\x0a\x43\xa4\x7a\x43\xf8\xfb\xf3\x77\x6f\x1f\xd5\xfe\x50\xc1\x61\xff\xfa\x5e\x41\xaf\x14\x35\xcf\x8a\x3c\x03\x00\x60\x07\x7f\x46\x29\x32\xfa\xe9\x43\x82\x81\xf4\xde\x43\x2f\xfc\xd5\xd3\xd3\x19\x20\x78\xa4\x19\xf0\x8d\xb1\xe9\x30\x16\xcf\xbb\x5d\x79\x05\x4c\x9b\x83\xd0\xbc\x9d\xdc\x5c\x52\x9f\x50\xf5\x27\x44\x57\x77\xa8\x5d\x1a\x90\x67\xe5\x4b\x9e\xe5\xd9\x22\x91\x05\x17\x6a\xcf\x6a\xe9\x58\x1b\x42\x19\x9b\xa7\x0d\x91\x1c\x69\x13\xf9\x64\x1c\x64\xb1\x97\xb4\x35\x1e\x7a\xc5\xd7\xc2\xd3\x96\x49\xfb\x1e\x24\x6a\x76\xc0\x62\x10\xa9\xa5\x48\xd2\x90\x9e\x95\xa1\x60\x57\x42\x10\x70\xe4\xc9\x08\x1a\xd4\x06\x1d\x2d\xe9\x26\x9e\xc1\x64\x82\x67\x3e\xec\x1a\xd9\xfd\x5e\xd8\xe8\xf8\x6f\xbd\x5c\x7e\x07\xb9\xbe\xd9\xcf\x10\x3c\xa1\xdc\xa8\x3a\x6a\xb1\xf7\x06\x2d\x7a\xa5\x4b\x98\x9b\x38\xe3\x21\x56\xf2\x3c\xa8\xcf\x81\x7c\xad\x87\x51\xf8\x5e\x0f\x8f\xed\xf3\x37\x00\x00\xff\xff\x82\x3b\xa5\x42\x38\x04\x00\x00")

func _000001_init_up_sql() ([]byte, error) {
	return bindata_read(
		__000001_init_up_sql,
		"000001_init.up.sql",
	)
}

var _bindata_go = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func bindata_go() ([]byte, error) {
	return bindata_read(
		_bindata_go,
		"bindata.go",
	)
}

var _generate_go = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\xca\xb1\x0d\xc5\x20\x0c\x04\xd0\x9e\x29\x6e\x01\xa0\xff\xdb\xdc\x4f\xac\x13\x42\xb1\x11\xf1\xfe\x4a\x93\x22\xf5\x7b\x8b\xc7\xa4\x0c\xd7\xd0\x66\x8e\xf0\xbb\x94\xde\x15\x3f\x99\xdb\x66\x1a\x14\xf5\x3f\xfc\x64\x12\x75\x4d\x7d\x26\x6a\xe0\xa5\xa6\x00\x5a\x79\x02\x00\x00\xff\xff\xe6\x24\xd8\x86\x4e\x00\x00\x00")

func generate_go() ([]byte, error) {
	return bindata_read(
		_generate_go,
		"generate.go",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
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
var _bindata = map[string]func() ([]byte, error){
	"000001_init.down.sql": _000001_init_down_sql,
	"000001_init.up.sql":   _000001_init_up_sql,
	"bindata.go":           bindata_go,
	"generate.go":          generate_go,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//
//	data/
//	  foo.txt
//	  img/
//	    a.png
//	    b.png
//
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
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
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func     func() ([]byte, error)
	Children map[string]*_bintree_t
}

var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"000001_init.down.sql": &_bintree_t{_000001_init_down_sql, map[string]*_bintree_t{}},
	"000001_init.up.sql":   &_bintree_t{_000001_init_up_sql, map[string]*_bintree_t{}},
	"bindata.go":           &_bintree_t{bindata_go, map[string]*_bintree_t{}},
	"generate.go":          &_bintree_t{generate_go, map[string]*_bintree_t{}},
}}
