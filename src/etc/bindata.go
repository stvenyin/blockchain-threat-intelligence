// Code generated by go-bindata. DO NOT EDIT.
// sources:
// etc/.DS_Store
// etc/bindata.go
// etc/rule.xml
// etc/server.crt
// etc/server.key
package etc

import (
	"bytes"
	"compress/gzip"
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
	bytes []byte
	info  os.FileInfo
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

var _etcDs_store = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x98\xb1\x4e\xc3\x30\x14\x45\xef\x73\x32\x44\xb0\x78\x64\xcc\x17\x54\xe2\x0f\xac\xaa\x0c\xcc\xfc\x00\x6a\xc9\x84\x51\xa5\x40\x11\x6c\xf9\x73\x90\xed\x5b\xb5\xa8\x41\x88\x21\x22\x54\xf7\x48\xed\xa9\xe4\x5e\xc7\xf1\xe0\xbc\x17\x00\xb6\xdc\x3d\x5c\x03\x1e\x40\x83\x62\x4b\x3f\x46\x68\xf8\x39\xc1\xd1\x55\x0a\xe7\x39\x16\x78\xc3\x13\xe2\x6d\xdc\x6e\xc6\xe7\x9a\x1d\x55\xbe\xb9\x1e\x3b\x44\x74\xc7\xeb\x5f\xc7\xed\x3a\x6d\xcc\xc7\x17\x72\xe6\x02\xcf\xe8\xd0\xe3\x35\x7f\x2f\xb0\x41\x8f\x97\x5f\x66\x1e\xd1\xe1\xfd\x87\x8c\x10\x42\x08\x31\x05\x56\xd4\x5c\xfe\xf5\x42\x84\x10\xb3\x23\x9d\x0f\x2d\x1d\xe8\xa1\xd8\x38\xee\xe8\xfa\x28\xe3\xe9\x96\x0e\xf4\x50\x6c\xfc\x9f\xa3\x6b\xba\xa1\x3d\xdd\xd2\x81\x1e\x8a\x79\x68\x19\x9b\x0f\xe3\x95\xf7\xcd\x8b\x79\xba\xa5\xc3\x34\x7b\x23\xc4\x7f\xa7\x2a\xf2\xe9\xf9\x7f\xf3\x7d\xff\x2f\x84\x38\x63\xac\x5e\xdd\xad\x96\x87\x86\xe0\x04\xc7\x42\xe0\x7e\x1f\x60\x21\x80\x91\x22\xc0\x95\x97\x85\x57\x38\x8c\xab\x10\x10\x62\x66\x7c\x06\x00\x00\xff\xff\x7a\x87\x14\x61\x04\x18\x00\x00")

func etcDs_storeBytes() ([]byte, error) {
	return bindataRead(
		_etcDs_store,
		"etc/.DS_Store",
	)
}

func etcDs_store() (*asset, error) {
	bytes, err := etcDs_storeBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "etc/.DS_Store", size: 6148, mode: os.FileMode(420), modTime: time.Unix(1529907444, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _etcBindataGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func etcBindataGoBytes() ([]byte, error) {
	return bindataRead(
		_etcBindataGo,
		"etc/bindata.go",
	)
}

func etcBindataGo() (*asset, error) {
	bytes, err := etcBindataGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "etc/bindata.go", size: 0, mode: os.FileMode(420), modTime: time.Unix(1530026561, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _etcRuleXml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x57\xcd\x6e\x1b\x47\x0c\xbe\xe7\x29\x84\x01\x7a\x53\xe2\xf9\xff\x11\x14\x07\x30\x82\xa2\x87\xa2\x2d\x52\xfb\x64\x18\x06\x87\xe4\x48\xdb\xae\x76\x85\xdd\x55\xe3\xd4\xc8\xbb\x17\xab\x95\x65\xb7\xa8\x0b\x24\xb6\x4f\xcd\x9c\x86\xe4\xec\xcc\x47\xf2\x1b\x72\x76\xf9\xee\x66\x53\xcf\xfe\xe0\xae\xaf\xda\xe6\xad\x50\x6f\xa4\x98\x71\x83\x2d\x55\xcd\xea\xad\xb8\x38\xff\xfe\x75\x14\xef\x4e\x5f\x2d\x3b\xee\xdb\x5d\x87\xdc\x9f\xbe\x9a\xcd\x66\xb3\x65\xb7\xab\x79\x9a\xee\xc5\x0d\x0f\xeb\x96\x4e\x7f\xf9\xf9\xd7\xf3\xe5\xc9\x41\x78\x60\x85\x01\xd7\xa7\x3c\xac\xaf\x73\xdd\xe2\xef\x3f\xed\x36\x99\xbb\xe5\xc9\xa4\xbe\x5f\xd6\xf1\xb6\x06\xe4\xd3\x5b\xf1\x5b\xdf\x36\xdd\x16\xc5\x42\xe8\x37\x52\xcc\x45\x45\x62\xf1\x5d\x3f\x17\x1d\xf7\xbb\x7a\x10\x0b\x21\x6f\x9c\x85\x12\x82\xf8\xbc\x3c\xb9\xfb\x6c\x02\x76\x72\x8f\xec\x6b\x41\x02\x62\xbb\x6b\x86\xfe\x29\x08\x2f\x85\xbc\x09\x05\x2c\xe6\x0c\x09\x6c\x51\x56\x5a\x49\xa0\x62\x29\xe8\x43\x88\xa8\x1d\x5b\xc4\xa0\x8a\x49\x62\x3e\xfa\xe3\x23\x78\x13\xb4\xf5\x41\xe6\x04\x36\xb8\xe2\xbc\x73\x96\x12\x69\x36\x1e\x95\xcb\xd9\xa0\xa3\xfd\xe2\x82\xb2\x64\x4a\x20\x5d\x76\x39\x6b\x2a\xac\x99\xbc\x2b\xa0\x8c\x0a\x94\x6d\x91\x32\x79\x62\xb3\x5f\x6c\x18\x20\xe9\xa0\xa3\x77\xd2\x27\x6d\x0c\x27\xe0\x64\x9c\xa3\xa0\xa5\xe1\xe2\x15\x2b\x4d\xd6\xed\x17\xa3\x24\xc9\x24\x7d\xcc\xc1\xb2\x94\x25\x95\xa2\xc0\x78\xeb\x50\xa3\x4e\x92\x8d\x35\xba\xb0\x9b\x76\xf6\xd1\x30\x26\x1b\x48\x83\x95\x26\xf9\xa4\x90\x43\x26\x50\x3a\x05\xe9\x51\xea\x62\x46\x9f\xac\xb8\x7a\x89\x34\xad\x78\x38\x83\x1a\x1a\xe4\xff\x4a\xd4\x5d\x62\xfe\x99\xb0\x07\x54\x92\x46\x39\x63\x31\x82\x31\x29\x00\x64\x17\x5f\x84\x56\x3d\x37\x74\xde\x41\xd3\x03\x0e\x55\xdb\x7c\x0d\xbb\xb8\xeb\xda\x4e\x2c\x6e\x05\xb6\xc4\x62\xf1\xda\x68\x29\xe5\x5c\x6c\xb8\xef\x61\xc5\x62\x21\x0e\xd4\x9d\x55\xfd\x6c\xbc\x69\x4c\xe2\xf3\x4b\xb9\xf2\x01\x3e\x3e\xd1\x9b\x07\x29\x28\xab\x71\xa0\xb2\x26\xbf\x48\xec\x47\xae\x8c\x01\x39\xfb\xf4\xf4\xe2\x73\x2b\xa8\x2a\xa5\xc2\x5d\x3d\x7c\xda\x83\xb7\xf2\x30\xc4\x5c\xf0\xcd\xd0\xc1\x7b\x18\x60\x6f\x51\x2a\x67\x8e\x94\x2d\x1b\x1b\xb2\xe5\x88\xc9\x04\x54\x18\x4d\x90\x6c\xb3\x63\x32\x06\x28\x1b\xca\x3e\x61\xa6\x1c\xc0\x44\x56\xec\x64\x56\x39\xea\x02\x62\x2e\x56\xd0\xff\x58\x6d\xaa\x29\x4c\xca\xc4\x38\xe9\x2e\x7a\xa6\x89\xbc\x62\x2e\xd6\xd0\xaf\xf7\x02\x59\x76\x3e\x58\x59\x62\xf0\xc0\x25\xa2\x54\x32\x47\x0f\x56\xd2\x58\x4e\x82\x75\xa0\x54\x24\x99\xa4\x07\x63\xd9\x27\x60\x8c\x28\x29\x2b\xcc\xb1\xc0\x78\xab\xeb\x76\xd5\x9f\xd5\x6d\xbb\x99\x76\xff\x36\xfe\xd7\x43\xcc\xc5\xa6\x6a\xb8\xfb\x22\x32\x88\xb9\x68\xda\x06\xf9\x5f\x3e\xb2\x7a\x34\xee\xaf\xe0\x91\xbd\x5b\xe8\xb8\x19\x7e\xb8\xe3\xf0\x33\x40\xee\x18\xb9\xda\x0e\xfd\x87\xb6\x3d\x3c\x16\x3c\x47\x55\x54\x50\x19\xd1\x39\xf0\xa5\x44\x63\x1d\xfb\xa4\x51\x96\xe8\xd9\x65\x1b\x59\xaa\x9c\x92\x47\x20\x94\x52\x79\xad\x4b\x76\x6c\xbc\xc9\x56\x2b\x31\x17\xfd\x1a\xcc\x45\x83\x35\xf7\xd3\x4d\x24\x44\x4b\x1c\x89\x31\x38\x1a\xbb\x46\x74\xd9\xf9\x90\x3d\x22\x59\x05\x64\x94\xb6\x4e\xe5\x64\x23\x04\xab\x4c\x91\xa0\xac\x2e\x64\x25\xd9\x64\x6c\x18\x77\xac\xfe\x9c\x42\xa4\x15\x8e\xe2\x00\x03\x1f\x11\x53\x28\x31\x05\x5b\xb2\x03\x0c\x91\x12\xa0\x4c\x29\x27\x20\x27\x55\xcc\x4c\xa8\x91\x25\x04\x4d\x40\x2a\xea\x00\x2a\xc8\x44\x60\xa4\x8b\xb2\x48\x67\xad\x98\x8b\xa1\xda\x70\x3f\xc0\x66\x7b\x8c\xf4\xd0\x0e\x50\xbf\x7f\xbc\x78\x0d\xf7\xf5\xbc\x17\x8b\xcb\xab\xbf\x6b\x9e\x33\x9a\xbb\x43\x24\x2f\xaf\x9e\xbd\x3f\x35\x3c\x5c\x1f\xde\xb3\x4f\x6a\x4a\xea\xd9\x1b\xd1\xc8\xed\xc7\x21\x7d\xc9\x61\xfb\x2d\xba\x2d\x5e\x6f\x5a\xda\xd5\xfc\xa4\x97\xea\xad\x00\xda\x54\xcd\xe8\xf1\xde\x4e\x9c\x77\xab\xa3\xc4\xc3\xfa\x38\xbf\xab\x04\x93\xd4\xf0\x70\x9c\x6f\xb9\xeb\xdb\x06\xea\xa3\x62\x3a\x70\x9a\x0f\x37\xdb\xb6\xbd\x37\x7d\xe4\x6c\x0e\xc2\xa3\xb9\x1f\xd5\x77\xff\x1b\x7f\x05\x00\x00\xff\xff\x74\xc7\x58\x4a\xa2\x0c\x00\x00")

func etcRuleXmlBytes() ([]byte, error) {
	return bindataRead(
		_etcRuleXml,
		"etc/rule.xml",
	)
}

func etcRuleXml() (*asset, error) {
	bytes, err := etcRuleXmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "etc/rule.xml", size: 3234, mode: os.FileMode(511), modTime: time.Unix(1529921188, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _etcServerCrt = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x94\xb9\xae\xab\x58\x13\x85\x73\x9e\xe2\xcf\xad\x5f\xc6\x36\xf6\x81\xe0\x06\x7b\x60\x66\x83\x37\x33\x64\x06\x33\x18\xcc\x74\xc1\x80\x79\xfa\x96\xcf\x95\xba\x8f\xd4\x5d\xd9\xb7\x54\x41\xad\xa5\xd2\xfa\xff\x67\xa0\x28\xab\xe6\xff\x90\x68\xbb\xaa\xa4\x22\xe0\x8a\xdf\x2a\x43\x54\x15\x07\x15\x42\xa0\x56\x11\xa2\xb8\x2e\x52\x3d\xcf\xdd\x8e\x6a\x15\x30\x61\x51\x0f\x65\xfd\x90\x85\x85\x85\x80\x7a\x12\xc0\x70\x26\x74\x5c\x10\x8d\xb0\x4f\xa9\x2c\x2e\x5a\xc5\x24\x95\x78\x25\x80\x95\xc1\xc1\x13\x11\x58\xe4\xb8\xf5\x87\x5b\x20\xbd\x08\xe5\x17\xfc\x67\x51\xc1\x20\xe6\x6e\x81\x34\xc5\x01\xb7\xa2\x0d\x68\xb0\x30\x7d\x08\x3a\x02\x5a\x73\x62\x08\xed\x16\x04\xbe\x17\x0d\x0c\x24\xee\x07\xe3\x6f\xb6\xeb\x45\x5a\x22\x4d\xef\x62\xb5\x9c\x53\x13\xd0\x1a\x4a\x45\xcf\x51\xd9\x7e\x26\x01\xff\x62\xa2\xa3\x30\x11\xc8\x85\xd8\x15\x39\x82\xa3\x37\xa9\xc0\x81\xb8\x60\x33\x0f\x5d\x88\x5d\xf5\xa3\xad\x96\xfb\xb7\xb6\x24\x9b\x68\x10\x50\x7f\x5f\x0c\x4b\x82\x98\xe8\xc8\xad\x78\xfb\xf8\xfd\x9c\x55\x10\xd8\xc4\x87\x5b\xf3\x2c\x7f\x5a\x83\xcb\x22\x67\xf2\xb3\x4c\x02\xff\xf5\x23\x03\x1d\x03\x6d\x63\x12\x57\xd4\x09\x28\xfe\x64\xb0\x2c\x30\xc3\xff\x30\xf8\xc3\x31\x81\xa9\x8c\x9c\x41\x76\xd4\xe4\x84\xa9\xa8\x01\x3b\xd2\x33\x11\xd6\x71\x70\x98\x19\xa3\x31\xe7\xc4\x45\x08\x38\xea\x27\xb4\x1f\x5e\x45\x08\x29\xc0\x45\x21\x5e\x01\x46\x08\xd0\x0e\x15\x85\x08\x81\x31\x0a\xbe\x70\x3f\xed\xc2\xc5\x34\x05\xf5\xca\x4c\x6b\xc3\xee\x5f\xa4\x6c\xb5\xa8\xab\x77\x2c\xd4\xee\x1b\x6b\x56\xd3\x3d\xef\x40\x59\x38\xa9\x1c\x8b\xce\xfd\x78\x0d\xce\xec\x83\x2f\x76\xf4\x15\xa4\x8a\xe9\x66\xab\x6c\x3f\xb6\x42\xc6\x9c\xe3\x33\xb7\xc7\x0e\x5b\x9a\x3e\x2e\xec\x08\x2b\x3b\xa3\xe2\x65\x88\xda\xbe\x36\x0f\x64\x26\xce\x75\x90\x4a\xae\x5b\x5e\x88\x6e\x9b\xd0\x83\xad\x77\xd4\xf4\x64\x59\x46\x2e\xc7\x6f\xdc\x0d\x47\x12\xd7\x4c\x07\xce\x82\xc6\x47\x3e\xae\x76\x0f\x7a\x62\x5b\xda\xd2\xb3\x67\x87\x2a\xa8\x57\x22\xcf\xba\xcc\x77\xdc\x85\x77\x85\x51\xab\x0f\x79\x30\xf2\x6c\x7d\x72\x5b\xb0\xb3\x9b\xaf\xd6\x32\x3a\xd5\xa9\x99\xe1\xf6\xa2\xf7\xe4\xcd\x3f\x42\xeb\xe2\xf4\x2a\xbd\xb1\x70\xf7\xee\xd9\x18\xd6\x6b\xc0\x55\xb8\xba\x64\xa5\xcb\xfe\x76\x58\x25\x78\x88\x67\x0f\x36\x13\x34\xfb\x31\x6f\xed\x21\x05\x8d\x50\x5c\x5f\xcc\x7c\x1a\x94\xde\xf1\x95\x93\xfc\xbe\x8c\x9c\x75\x38\x50\xba\x28\x89\x26\xab\x46\x68\xc6\xfd\x73\xef\xbd\x95\x6d\xdf\x4e\x30\xa1\x79\x44\x2f\x92\xf4\x74\x81\xab\x3f\xb3\xd0\x3e\x25\x08\x00\x7d\x63\x38\x4d\x79\x7a\xdc\xe7\x8b\x45\x00\xdc\xff\xf8\x7f\x60\x21\x40\x45\xa0\xdd\xe8\x4a\x2a\xb7\x6a\xc9\x4c\xb4\x8d\xbd\x9b\x71\x9c\xa9\x23\xff\x64\xde\x6f\xfe\x98\xb4\x81\xf4\xfe\xaa\x37\xbd\x87\x4d\x97\xdb\xa6\x34\xaf\x9d\x7f\x31\x5e\x81\xb1\x09\x8d\xae\xa8\x24\x0b\xe4\xc1\x75\x1b\x35\x6c\x48\xb0\x57\x58\x9b\xda\x48\x08\xf1\xcd\xef\xce\x80\xd9\xd9\xb2\xb3\x2f\x9e\x5f\xe9\xcc\x1e\x5b\xfe\xe1\x9e\x9a\xf2\x1a\xe9\xc6\x43\xf9\x82\xfd\x70\x50\xe6\xe9\x34\xad\x01\x5f\x1d\x9c\xfe\xa8\xfb\x7b\xb7\x1e\x68\xa6\x3c\x47\xe3\xa2\xa7\x77\x16\x7e\x79\x3a\x33\xe5\xbf\x67\x8d\xbd\xcd\x47\x2f\xe9\xef\xfc\xe8\x08\x56\x31\x84\x37\x1d\x41\xcd\x7b\x6d\xb0\xd0\xfd\x17\xef\x88\x15\xbc\x5a\xe1\xac\x0d\xb4\x4a\xc3\x05\x82\x43\x70\x0c\xdf\x0d\xa6\xe7\x52\xf3\x98\x73\x49\xe6\xfe\xce\x1a\xe6\xf3\x62\x5e\x0c\xcf\x95\x92\xd8\x67\xb3\xc5\xf1\x9c\xb6\x95\xee\x95\xb2\xdb\xe1\x26\x2b\xfa\x98\x4f\x80\x4e\x4d\x9a\xed\x6c\x25\x47\x01\x4f\xce\xe3\x64\x9b\x79\x2e\x31\xf2\x5e\x2a\x2d\x93\x9e\x2c\xcd\x08\xda\xf2\xb1\x9e\x9a\xaf\x9b\xe4\x40\xc7\xdb\x9f\xbd\x6b\xd3\x8a\x5a\xfa\xdb\x4e\xd2\xa3\x3a\xe2\x68\x1c\xcb\xbc\x86\x88\x8d\xd7\xd0\xbe\x1a\xd3\x20\x99\xf4\xd7\x2f\xe6\xbb\x97\x44\x13\xff\xbb\xab\xfe\x0a\x00\x00\xff\xff\x9a\xb5\x06\x10\xc8\x04\x00\x00")

func etcServerCrtBytes() ([]byte, error) {
	return bindataRead(
		_etcServerCrt,
		"etc/server.crt",
	)
}

func etcServerCrt() (*asset, error) {
	bytes, err := etcServerCrtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "etc/server.crt", size: 1224, mode: os.FileMode(420), modTime: time.Unix(1529644237, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _etcServerKey = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\xd5\x37\x12\xab\x56\x00\x85\xe1\x9e\x55\xbc\x9e\xf1\x90\x83\xca\x4b\x14\x39\x83\xa0\x23\x5e\x10\x19\x81\x40\xac\xde\xe3\x57\xfb\xb4\xa7\xf9\xbb\xef\x9f\xff\x26\xc8\xaa\x66\xff\xf1\x03\xf0\xc7\xf5\xb5\x18\x84\xf2\x1f\x43\x4e\xff\x3e\x88\xa5\x69\xf2\x7c\x6a\x02\x00\x86\x08\x3c\x19\x1c\x37\xf1\x22\xa8\x86\x69\x00\x4e\xe1\x10\xa5\xd4\x22\x44\xe9\x9f\x5a\x0e\x5d\x17\x72\x9e\x3c\x50\xae\x47\x3a\x04\x81\x42\x51\x15\xf4\x33\xeb\x11\x5f\xa7\xd2\x47\x31\x05\xe6\x4f\x62\x44\x26\x3b\x4b\xc2\xe6\xe4\x4c\x35\x1d\x90\xba\xad\x1e\xaf\x26\x6d\xd1\xfd\x7a\x4b\xc1\xa9\xda\x0a\x03\xc2\x75\xe9\xea\xd1\xa3\xa2\x1f\x7f\x69\xe8\x9c\x3c\x91\xee\x2d\xd2\xba\xe4\xba\x64\x2f\x39\x83\xd6\xde\x25\x4d\xef\x7c\x36\x69\x4e\xb7\xa5\xd7\x18\x40\x69\xc2\xa7\xab\x8d\x1c\x97\xd5\xa5\x26\xa8\xc4\x4a\x1a\x7c\xa5\x84\x62\x28\x9f\x39\x3f\x17\x3f\xe4\xfd\xde\x2e\x97\x3c\xc7\x30\x7e\xe4\xf7\x1d\x84\x95\x53\x4a\x8c\x5a\x1c\x25\x7d\xc0\xd6\x08\x96\x95\x11\x88\xaf\xf9\xd3\x4b\x6e\x31\xfa\x56\xd8\xbc\x27\x67\x4c\x7e\xaf\x86\x4a\xd1\x59\x8e\xbb\x20\xac\xe2\x06\xbb\xe9\x55\xb9\x16\x0e\x9e\x9a\xe3\x32\x39\x5e\x68\xa5\x2e\xa7\x58\x90\x00\x65\xb1\x7a\xae\x07\x3d\x6a\xca\xc2\x5c\x6f\xc8\xbd\x62\x41\x02\xd5\xa7\x4f\xe1\x5e\x12\xe3\xf8\x7a\x84\x88\xd6\xb8\x68\x8d\x2b\xbb\xf0\x68\xa5\x39\x4a\x62\x4b\xb0\xd6\x98\xa9\x9e\xd5\x47\x03\x60\x73\x61\x5f\xc7\x21\xf4\x34\x09\x78\x40\x00\xb3\x26\x00\x4f\x1c\x36\x2e\x39\x5b\xed\xfd\x6d\x5a\xae\x44\xac\x91\x2a\x3c\x03\x85\x0e\xf3\x7e\x9c\xb7\x68\xda\x32\x95\x08\x9e\x68\x15\xa5\xae\x49\x15\xd5\x7f\x58\x8d\x1f\xab\x78\xb6\x45\xd7\xe0\xe7\xd1\x41\x9f\xae\xdb\x65\x9b\x76\x09\x1f\x95\xc4\xe9\x04\x11\xf9\x64\xc2\x18\x7c\xd3\xab\x9a\xa4\x3f\x07\x80\xf8\x7d\x13\x82\x5b\xb0\xbd\x9e\x76\x3e\xa0\x74\xee\x4b\xdc\x40\x4a\xa0\x9c\xd3\x73\xbf\x6c\x29\x5d\x3b\xbe\x45\xb4\xf2\x7a\xf3\x38\x56\x31\x88\x49\x37\x24\xdb\x97\x1f\x5a\xdd\xe8\x74\x88\x34\x3f\x5b\xa6\x5e\x0d\x2e\x67\x7e\x27\x22\x3b\x45\x5f\xfe\xb2\x9e\xc0\xb3\xac\xe6\xcb\xb3\xec\xc8\xb6\xf9\xd4\xed\x2b\xb4\xac\xc5\x3a\x5b\xea\x46\x11\x3b\x7c\xfc\x50\xb0\x3a\x2e\x8a\x42\xc7\xb3\x56\x68\x50\xf4\xb3\x1f\x82\x63\xcc\x64\x3c\xd7\x9c\x94\x65\x5f\x90\x70\x0c\x4b\x22\xd3\x06\x2b\xf9\xd3\xaa\xae\xbd\xd2\x27\x21\xc0\x4e\x4c\x97\x4a\xa4\x30\x1f\x93\x73\xdf\x53\x30\x0a\x8c\x7e\x53\x24\x20\xb1\x2f\x68\x71\x8a\xf5\xe7\x68\xc0\xc7\x70\x9c\x75\xf8\xa1\xa7\xe6\x87\x63\x7a\xa8\xda\xcc\xa1\x92\x8e\x9c\xbd\x95\x8f\xd2\x91\x3e\x88\x50\x64\xd1\x49\xfb\x9c\x30\x01\xcc\xaa\x00\x9c\x5f\xf2\xbe\xdf\xbe\xc7\x63\xd2\xd7\xd7\x29\x67\x7f\x50\xa2\x89\x3f\x5d\x98\x06\xcb\x20\x9a\xe2\x2c\x55\xf3\x6f\x1f\x8f\xb7\x5d\x11\x1a\xfa\x0b\x2d\xd3\x44\xe2\x9f\x80\xe6\xd8\x87\xad\x4f\xde\x75\x45\xaa\x0a\x84\xe4\xfa\x36\x5d\x84\xd3\x2f\xf9\x3b\x17\xf2\xd8\x4f\xbf\x7b\x78\x0f\x96\x94\x24\x36\x8c\xd5\xa2\xac\x4b\xdd\xd9\x1f\xe1\x6e\x0c\x1d\x5c\x51\xe4\xa2\x0a\x76\x77\xea\xe1\x60\xc2\x8c\xc7\xa9\x8e\x97\x1b\xde\x08\x19\x29\x6c\x4b\x08\x1f\xea\xad\xe0\xfe\xe6\x9f\x8b\x2f\x82\x79\x39\xf4\xa2\x39\x47\x77\x15\x94\x27\xf3\xb7\xd8\x5a\x73\xc4\x75\xc0\xa5\x0c\x41\xe3\x7f\xb2\xa1\xee\x96\xac\x11\x6b\x33\x19\xea\xbd\xaf\x95\x5f\x7e\x52\x39\xbb\x0b\xf5\x8e\x2f\x39\x35\x4f\x2e\x35\xba\x4a\xf8\x68\x6d\xb6\x26\x22\x74\xe0\xae\x0e\xa0\x10\xb9\x74\x53\x38\xf2\xfd\xe8\x67\x62\x7b\xf4\x9b\x41\xaf\x7c\xea\x74\xd4\xfc\x45\x37\xbb\xe0\xa9\xad\x3b\xb4\xeb\xba\x4f\xd1\x97\x41\xe7\x96\xc2\xd6\xfd\xd8\x72\xad\x08\x66\xaf\xd5\xc7\x1a\x64\x09\x22\xd7\x9c\xe5\xed\xca\xc8\x45\x64\x18\xc7\x96\x7f\xe6\x76\x48\x5f\x85\xc7\xce\x79\x9f\xf7\x3e\x76\x32\x46\xae\x8f\x98\x0e\x66\x15\x04\x04\x0a\x44\x6a\x9b\xa8\x2d\xad\x4b\x22\x4f\x18\x20\x66\x88\xe7\xd7\xb0\xc7\x44\x4a\xcd\xbf\x5d\x54\x3e\x8d\xcd\x4f\x03\xfb\x2d\x6f\x35\x84\x38\xd9\x33\x26\x69\x54\x8e\x96\x8a\x84\x5a\x9e\xed\xe9\x72\x6d\x7e\x63\xa7\x6e\x87\xb6\xb3\x92\x28\x8c\x4d\x0e\x89\xf4\xc9\x8f\x32\x33\xdf\x41\x23\xa0\x42\x26\x62\x61\xb3\x6e\x77\x2d\x1c\xe2\x15\xa0\xc1\xa3\x6f\xd6\xcb\x98\x00\xcd\x50\x6c\xc9\xcb\x84\xcc\xd7\xa5\xa1\xe8\x61\x73\x28\x50\xab\x29\xfb\x3e\x91\xae\x08\x9a\xcf\x60\xe0\x0d\x10\xfa\xfd\xe5\xd7\xe2\x1e\x7a\x5e\xda\x8b\x30\x15\x18\x0f\x16\x95\x4f\xbb\x36\x5f\xa7\x3f\xfd\x31\x3b\xe6\x26\xa6\xea\x50\x35\xf8\xbc\xc4\x4f\x85\x1a\x63\xb7\xe9\x11\xae\xdf\x63\x37\x24\xdf\x77\x41\x46\x5d\xb0\xa7\xfd\x3c\x3c\xeb\xf1\x87\x29\x23\xf0\x8e\x11\xfb\xf8\x4d\x1a\x3c\x35\x67\x95\xad\x03\xbb\x1e\x74\xf4\x2a\x1c\xbc\x9a\xe4\x91\xfa\x52\xef\xce\xa1\x50\x64\xbe\xaf\xd0\xee\x60\xa8\x8e\xbb\xa9\x5c\x46\xed\xdc\x89\x1b\x19\xb2\x24\x6f\x9d\x68\xc5\x4a\x1e\x7d\x33\xe3\x86\xb5\xee\x4e\x37\x8b\x7f\x28\xc3\xb8\xcc\xc9\x75\x0d\x88\x5f\xf9\xa4\x9a\x22\x8f\xec\xaf\x17\x77\x7a\x86\x00\x25\xbe\xa7\x2a\x2f\x51\xad\xdb\x27\x84\x96\x0e\x85\x90\xe4\x67\x0e\x0a\x01\x39\x1f\xd9\x94\x9d\xa6\xdd\xd2\x79\x32\xf2\xf9\xf8\x1e\x56\x7e\xe8\x92\xeb\x65\x78\xa8\x87\x74\xbc\x6b\x08\x6f\xc6\x1e\x5f\xed\x2b\x80\xb2\x79\x71\x1d\x69\xb7\x59\x7e\x64\xf2\x0b\x3c\x89\xa5\xef\xd7\x25\x5a\xbd\xe8\xfb\xba\x2e\x69\x72\x8e\x8f\xb8\x61\xf5\x12\x3b\xeb\x1d\x32\x23\xc7\x62\x48\xe1\x58\x9e\x49\x39\xd5\xc6\x74\xee\xa7\xe6\x9a\x5d\x65\x72\x35\xbb\x59\xeb\x3e\x25\x3d\x67\xde\x8d\x64\xc4\xfc\xa1\x3a\xf9\x8b\x1f\x24\x45\xf4\xc5\x65\xf7\x91\xbf\xa4\xc8\xb6\xf4\xff\xd4\xfc\x1b\x00\x00\xff\xff\x07\xe9\xfd\xb3\x8b\x06\x00\x00")

func etcServerKeyBytes() ([]byte, error) {
	return bindataRead(
		_etcServerKey,
		"etc/server.key",
	)
}

func etcServerKey() (*asset, error) {
	bytes, err := etcServerKeyBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "etc/server.key", size: 1675, mode: os.FileMode(420), modTime: time.Unix(1529644229, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
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

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
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
	"etc/.DS_Store": etcDs_store,
	"etc/bindata.go": etcBindataGo,
	"etc/rule.xml": etcRuleXml,
	"etc/server.crt": etcServerCrt,
	"etc/server.key": etcServerKey,
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
	"etc": &bintree{nil, map[string]*bintree{
		".DS_Store": &bintree{etcDs_store, map[string]*bintree{}},
		"bindata.go": &bintree{etcBindataGo, map[string]*bintree{}},
		"rule.xml": &bintree{etcRuleXml, map[string]*bintree{}},
		"server.crt": &bintree{etcServerCrt, map[string]*bintree{}},
		"server.key": &bintree{etcServerKey, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
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
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
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
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

