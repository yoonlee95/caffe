// Code generated by go-bindata.
// sources:
// builtin_models/bvlc_alexnet.yml
// builtin_models/bvlc_googlenet.yml
// builtin_models/bvlc_reference_caffenet.yml
// builtin_models/bvlc_reference_rcnn_ilsvrc13.yml
// DO NOT EDIT!

package caffe

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

func bindataRead(data, name string) ([]byte, error) {
	gz, err := gzip.NewReader(strings.NewReader(data))
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

var _bvlc_alexnetYml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x56\x41\x6f\xe3\x36\x13\xbd\xeb\x57\x0c\x60\x2c\x90\x7c\x9f\x25\x59\xb2\x93\x38\x2a\x50\xa0\x4d\x2f\x05\x8a\x1c\x16\x6d\x2f\x45\x61\x8c\xc8\x91\xc5\x5d\x8a\x14\xc8\x91\x13\xef\xaf\x2f\x48\x49\x96\xd3\xcd\x76\x73\x70\x24\xce\x9b\xe1\xf0\xcd\xcc\x13\x0d\x76\x54\xc1\xcf\x7f\xfe\xf6\x94\xfe\xa4\xe9\xf5\x99\x18\x56\x10\x16\xc1\x36\x70\xb6\x83\x83\xce\x4a\xd2\x49\xe3\xb0\xa3\x17\xeb\x3e\x57\x09\xc0\xe8\xf4\x84\x4d\x43\xb0\x82\x8b\x09\x1a\xeb\x80\x5b\x9a\x5c\x00\x4e\xe4\xbc\xb2\xa6\x82\x22\xdb\xbc\x01\x4e\x06\x10\xd6\xb0\x43\x65\x38\xf9\x17\x74\x06\x28\xd3\x58\xd7\x21\x8f\xcf\xe0\xa9\x43\xc3\x4a\x5c\xec\xa3\x35\x09\x71\x50\x19\x72\x15\xac\xe0\xf2\xe2\x61\xf0\x24\x81\x2d\xf4\xe4\x02\x72\x4c\x0c\x7a\x47\x52\x89\x10\x33\x81\xe5\x6f\x05\xdd\xa0\x59\xf5\x9a\xa0\xd7\xc8\x01\xef\x41\xa0\x81\x9a\xc0\xf7\x24\x54\xa3\x48\x26\x00\xd8\xc9\xfb\x5d\x15\x3d\x8f\xfd\x50\x81\x43\xd5\x3b\xfb\x89\x04\xe7\x02\x5d\xa7\x53\x11\x78\xa9\x22\x2c\x15\xfd\x10\x91\xe2\xbb\xc8\x63\x44\xf6\xbd\xb8\xdf\x69\xaa\xbe\xeb\x34\x01\x27\xb7\xff\x4e\xe5\x1a\x2b\xc9\x0b\xa7\x7a\x8e\x5c\xff\x98\x00\xfc\xde\x2a\x3f\xf1\xa2\x3c\x20\x38\xea\xb5\x12\x23\xe3\xb6\x59\xca\x09\xa3\x67\x4d\x32\x14\x22\x2c\xcf\xfd\xd2\x0f\xf5\xec\x91\x25\x00\xbf\xa8\xa6\x21\x47\x46\x90\xaf\xc0\x58\x86\x58\x61\x65\x8e\xf0\xa2\xb8\x8d\x9e\x8e\xb4\x3a\xb6\x1c\xd6\x24\x32\xa6\x38\x1c\x3b\x32\x1c\x43\xfc\x90\x00\x28\xa3\x58\xa1\x56\x5f\x02\xc2\x58\x93\x7e\x21\x67\xa1\x56\xe8\xc9\x87\x72\x6e\xb2\x02\x94\xf1\x4c\x28\x43\x8e\x05\xdc\x34\x76\x30\x12\x0c\x09\xf2\x1e\xdd\x79\xec\xc4\x69\xdf\x35\xa0\x5f\x42\x8e\x27\x63\x0b\x05\x1c\xf1\x44\xd0\x68\x64\xd0\xd6\xfb\xdb\x2c\xb2\x41\x50\x0f\x46\x6a\x92\x0b\x29\x21\x65\xc5\xe4\x46\xd7\xed\xfd\x66\xbd\xd9\x6c\xc0\x1b\xec\x7d\x6b\x39\x1b\x9d\xc8\x33\x9c\x50\x2b\x89\x53\x63\x4d\x3d\x87\x46\x10\xc8\xc1\x85\xa3\x2c\x4c\x84\x8c\x96\x88\x77\xfb\x18\x31\xf2\xb3\xc4\x00\x14\x62\x70\x28\xce\x09\xc0\xdd\x43\x56\xde\xed\x3f\x00\x1a\x19\x93\x85\x22\xdb\x6f\x1f\x77\xfb\xec\xba\x7e\xb6\x0e\x7d\x1f\x8a\xc8\xb6\x4f\x8b\x8b\x7f\xf0\x2e\x46\x5f\x4c\x20\x1a\xef\x16\xe3\x7e\x93\x95\x1f\xc0\x8e\x35\xbd\xda\xdd\x13\xaf\x61\xf0\x21\xdd\x4f\x83\xe7\x68\x16\x64\x98\x1c\x08\x67\xfb\xc0\xd6\xcd\x1f\xd1\x1c\x2c\x78\x22\x87\xc7\xa8\x1a\xc5\x26\x02\xfc\x1a\x6e\x76\xf0\x7f\x28\x26\xaf\x5b\xf8\x1f\x94\xd0\x29\xe7\xac\x5b\x83\x6f\xed\xa0\xe5\x94\x72\x98\x2a\xa8\x15\x43\xab\x8e\x2d\xb9\x4b\x6e\xd9\x6d\xe2\x68\x69\xa6\x15\x2c\x6f\x71\xaa\xb1\x0f\x43\x9e\xc3\x0b\xd5\x5e\x31\x85\x47\x62\x91\x65\x73\xaf\xce\xc9\xcd\x7a\x94\x42\xcb\xdc\x57\x79\x3e\x7a\x66\x46\xf5\x3e\x13\x62\x7c\xcd\x77\xfb\x72\x97\xaa\x0e\x8f\x64\x88\x53\xa1\xd1\x7b\xd5\x4c\x8d\x9d\x86\xda\xa4\x92\xa8\x4f\x85\x35\x27\xab\x87\xb0\x8a\x3a\x35\x34\xb8\xf8\x8f\x83\xb0\xf9\xac\x97\xcd\x65\x23\x5f\xe5\xf9\x51\x71\x3b\xd4\x99\xb0\x5d\x1e\x44\x36\x8f\x13\x99\xb3\x23\xca\x3b\xf4\x4c\x2e\x8f\xc9\xf9\xbc\x3e\x69\x71\x40\x4d\xaf\x86\x38\x59\x81\x56\x82\x8c\xa7\x37\x23\x98\x4c\x8b\x15\x0c\xc6\x91\x67\xa7\x04\x93\x4c\x56\xa0\x4c\x3f\x70\xa4\x64\xc1\x8e\x6b\x41\x48\x56\xd0\x28\xe7\x79\x44\x01\x9f\x7b\xfa\x4a\xa8\xd3\xb8\x5c\x41\x3c\x7d\x32\xaa\xe1\x95\x52\xcc\x59\x5c\xc5\x89\xa0\x37\x62\x12\xa7\x24\x6e\xb1\x44\xe9\x31\x48\x3e\x93\x8b\xd5\x8b\x5b\x2f\x4b\x93\xf8\x4a\xd5\x91\x09\x62\xee\x2b\xf8\xab\x58\xc3\x76\x0d\x65\xf9\x10\x7f\xfe\x9e\x20\x1d\xa1\xa9\xde\xa3\xf4\xa3\xad\xc9\xf1\x66\x5f\x94\x79\xa8\x8d\x47\x2d\x89\x73\x87\x2f\x33\xb9\x91\xee\xd4\xa3\xce\x83\xd0\xe4\x4a\xfb\x93\x13\x45\x99\xcf\x65\x3e\x84\xd0\x59\xad\x0c\xba\x73\xef\x2c\xdb\xc4\x0e\xdc\x0f\x3c\xf2\x16\x8e\x14\x93\x9e\xce\x3f\xda\xc2\x08\x45\xb6\x1a\x42\x1e\x1c\x45\x28\xbe\xc7\xd7\x88\x5f\x8e\x9c\xbc\x43\xd9\x84\xd1\x58\xc7\x4a\x5c\x31\x36\xd5\xe1\x3d\xd6\xa6\x9d\xfd\x61\x70\xfa\x9a\x18\xcf\xd9\xc8\xce\xe0\xc9\x85\xef\x20\x19\x8e\x44\xa1\xb1\xe6\xdc\xd9\xc1\xe7\xb8\x2f\x77\x35\x6e\xb7\xdb\x1a\x1f\xcb\xf2\x71\x8b\xb4\xdf\x48\xb9\x97\xf7\xf7\x0f\x72\xfb\x58\x14\x91\xbd\xba\x94\x85\x2c\x05\x4a\xdc\x3e\x10\x16\xbb\x7d\xb1\xc7\xfa\x41\x6c\xa9\x2c\x1f\xc5\xee\xb1\xae\xcb\x52\xec\xee\xb6\x77\x71\xcb\x46\x69\x2a\x32\x7e\xe5\x24\xb6\x53\xa8\xf4\xfc\xbd\x9c\x95\xe9\xe8\xb0\x6f\xa3\x00\xbd\x50\x50\x7e\x0f\x8e\xbc\x1d\x9c\xa0\x70\x9e\x68\x3d\xf4\xc8\xed\x72\x16\x87\x2f\xdf\x38\xca\xd5\x18\x7d\x7b\x82\x72\x49\xbd\xb6\xe7\x2c\xd6\x34\xa4\x06\xf3\xce\x57\xfb\x54\x79\x2e\x75\x16\x43\x65\x35\xb9\xcf\xa4\xe9\x7c\x52\xa1\x15\x33\xeb\x8e\x6f\xe2\x8d\xa8\x79\x5c\x94\x3f\xa0\x13\xad\x3a\x85\x26\x40\xed\xc3\x2d\x48\x35\xa3\x62\x72\x4b\xa3\x94\xd6\xe8\x29\xd4\x67\xfc\xac\x86\x07\xb6\x80\x06\x26\xcf\xeb\xab\xc7\xd5\x1d\x24\x78\x2e\x7c\x5c\x53\x36\x2e\xc4\xf0\x92\x8c\x65\x0a\xcf\xdf\x88\x12\x4a\x12\x6f\x69\x7e\x6e\xc5\xaf\x2b\x10\x14\x6d\xfa\x92\xcf\x29\x21\xb3\x53\xf5\xc0\xa3\xd8\xd2\x2b\x3b\x84\x49\xd9\x60\xb1\x25\x00\x9f\x95\x91\x15\x3c\x3d\x3f\x4f\x19\x87\xf7\xb0\xd3\xa8\x86\x17\x9f\x9b\xa7\xe7\xe7\x35\x7c\x0c\x3f\x59\x96\xdd\x86\xc1\x99\xbe\x80\x87\x30\x8e\x9e\xb8\x82\x5f\xc3\x24\x8e\xd7\xce\x69\xed\x72\x6b\xbb\xfe\x88\x27\x00\x1d\x1a\xd5\x90\xe7\x03\x0e\xdc\x5a\x57\x01\xd6\x72\xd0\x32\xf9\x27\x00\x00\xff\xff\x16\x18\xb9\xf4\xc3\x0a\x00\x00"

func bvlc_alexnetYmlBytes() ([]byte, error) {
	return bindataRead(
		_bvlc_alexnetYml,
		"bvlc_alexnet.yml",
	)
}

func bvlc_alexnetYml() (*asset, error) {
	bytes, err := bvlc_alexnetYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bvlc_alexnet.yml", size: 2755, mode: os.FileMode(420), modTime: time.Unix(1502164405, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _bvlc_googlenetYml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x56\x5b\x6f\xe3\xb8\x15\x7e\xd7\xaf\x38\x98\x60\x80\xa4\xb5\x25\x4b\xf6\x64\x1c\x2d\xd0\x87\xa6\x40\x51\x60\x91\x87\xed\xed\xa1\x28\x8c\x23\xf2\x48\xe2\x86\x22\x59\xf2\xc8\x8e\xf7\xd7\x17\xa4\x24\xdb\x99\xcd\xec\xe6\x21\xb0\xc8\xef\xdc\xbe\x73\xa3\xc1\x81\x6a\xf8\xf3\xbf\x7e\x7c\x5e\xff\xd5\xda\xee\x47\x7a\x21\x86\x3b\x88\xc7\x60\x5b\x38\xdb\xd1\xc3\x60\x25\xe9\xac\xf5\x38\xd0\xc9\xfa\xd7\x3a\x03\x98\xc4\x9e\xb1\x6d\x09\xee\xe0\x72\x05\xad\xf5\xc0\x3d\xcd\x22\x00\x47\xf2\x41\x59\x53\x43\x99\x6f\xde\x01\xe7\x0b\x10\xd6\xb0\x47\x65\x38\xfb\x06\xba\x00\x94\x69\xad\x1f\x90\xa7\xdf\x10\x68\x40\xc3\x4a\x5c\xee\xa7\xdb\x2c\xea\x41\x65\xc8\xd7\x70\x07\x97\x8f\x00\x63\x20\x09\x6c\xc1\x91\x8f\xc8\xc9\x31\x70\x9e\xa4\x12\x51\x67\x06\xd7\xbf\x3b\x18\x46\xcd\xca\x69\x02\xa7\x91\x23\x3e\x80\x40\x03\x0d\x41\x70\x24\x54\xab\x48\x66\x00\x38\xc8\xc7\x5d\x9d\x24\x3b\x37\xd6\xe0\x51\x39\x6f\x7f\x26\xc1\x85\x40\x3f\xe8\xb5\x88\xbc\xd4\x09\xb6\x16\x6e\x4c\x48\xf1\xbb\xc8\x2e\x21\x9d\x13\x8f\x3b\x4d\xf5\xef\x0a\xcd\xc0\x59\xec\xb7\x5d\xb9\xc5\x4a\x0a\xc2\x2b\xc7\x89\xeb\x3f\x65\x00\xff\xe8\x55\x98\x79\x51\x01\x10\x3c\x39\xad\xc4\xc4\xb8\x6d\xaf\xe9\x84\x49\xb2\x21\x19\x13\x11\x8f\x63\xc5\xe8\x54\x31\x6e\x6c\x16\x99\x1c\xfe\x4d\x70\xb2\xa3\x96\xa0\xd5\x2b\x45\xf2\xb9\x47\xf3\x0a\xcf\xbd\x57\x81\x15\x1a\xf8\xfb\x2f\xd4\x91\x3c\xa7\x6a\x41\xad\x21\x3a\xd0\x93\x76\x8b\xde\x6f\x3c\xb8\x9a\x49\x7e\xe4\x19\xc0\x5f\x54\xdb\x92\x27\x23\x28\xa4\x72\xb4\x0c\xa9\x8c\x94\xe9\xe0\xa4\xb8\x9f\xd5\x68\xd5\xf5\x1c\xcf\x24\x32\xae\x71\xec\x06\x32\x9c\xf4\xfe\xf0\x5d\xa9\x20\x50\x13\x44\xcf\x62\xce\x79\xed\x23\xfe\x63\x05\x63\xa0\x00\x9f\xde\xf0\xa8\xc8\x7f\x8a\x81\x2a\xa3\x58\xa1\x56\xbf\x50\x52\x75\xa2\x68\x3f\x80\x32\x81\x09\x65\x8c\xe5\x53\x87\x63\x08\x0a\xcd\xa7\xa8\xe0\x7f\xa3\x12\xaf\x87\x60\xf5\x91\x7c\xee\xbc\x65\xcb\x6f\x3c\xa9\x45\x90\x73\x8c\x0c\x9a\xd0\x27\x27\x3d\x32\x81\x24\x81\x67\x70\x56\x2b\x71\x4e\xd4\x26\x5b\xd6\xab\x4e\x19\xd4\xf0\x8d\xb6\x55\x84\x70\xa4\xd9\x9e\xa2\xd6\x61\x14\x3d\xb4\x18\x98\xfc\x35\xf8\xfb\xc7\x0d\x90\xb3\xa2\x0f\x70\x0c\x50\x7d\x59\xbe\x1e\x7e\x48\xf5\x41\xd0\x8c\x46\x6a\x92\xd7\x32\x89\x26\x15\x93\x9f\x92\x54\xad\x76\x9b\xcd\x6a\xb3\xd9\x40\x30\xe8\x42\x6f\xf9\x46\xe5\x03\x8c\x21\x1a\xf9\x30\xd8\xa5\xfe\xde\x1b\xb0\x4d\xec\xdf\xe8\x2e\x5b\xb7\x2e\x01\x85\x18\x3d\x8a\x33\x3c\xee\xf3\xaf\x9f\xe1\x7e\x5b\xe6\xdb\xcf\x40\xde\x5b\xff\x00\x68\xe4\x0c\xfc\x72\x05\xee\xf7\xf9\xd3\x67\xb8\x2f\xcb\xbc\xbc\x00\xed\xc4\xd4\x11\xb5\x92\x93\xdf\x81\x78\x35\x7b\xf7\xf3\x18\x38\x5d\x0b\x32\x91\x1b\xe1\xad\xcb\xe1\xfe\x9f\xe9\x32\x9e\xe3\x91\x3c\x76\x69\x2a\x96\x9b\x74\x1d\x56\x70\xbf\x83\x3f\x42\x39\xcb\x3c\xc0\x1f\xa0\x82\x41\x45\x6b\x2b\x08\x7d\xea\x81\x29\x14\x40\x68\x14\x43\xaf\xba\x9e\xfc\xc5\xcb\xfc\x21\xf3\x74\xa9\x63\xb8\x83\xeb\x57\x9a\x59\xe8\xe2\x08\x2b\xe0\x44\x4d\x50\x4c\xf1\x27\xb1\xc8\xf3\xa5\x13\x17\xd7\x96\x69\xbb\x86\x9e\xd9\x85\xba\x28\xd0\xbf\xa9\x63\x6e\x7d\x57\x38\xd9\x16\xe5\x6e\xf3\x94\xef\xf6\xbb\x2a\x77\xb2\x7d\x87\xeb\x14\xf7\x63\x93\x0b\x3b\x14\x71\x07\x14\x69\x5c\x14\xec\x89\x8a\x21\xd5\x48\x91\x74\x87\xa2\x39\x6a\x71\xe8\x52\x23\x1a\xe2\xec\x0e\xb4\x12\x64\x02\xbd\x9b\x10\xd9\x7c\x58\xc3\x68\x3c\x05\xf6\x4a\x30\xc9\xec\x0e\x94\x71\x23\x87\x69\x14\x2c\xd8\xe9\x2c\x76\xef\x1d\xb4\xca\x07\x9e\x50\xc0\x67\x47\xbf\xda\x23\xeb\x74\x5c\x83\x1a\xb0\xa3\x6c\x1a\xd6\x37\x83\x6c\xf1\xe2\x46\x4f\x02\xbd\x9b\x75\xa9\x64\x93\x89\xab\x16\x87\x71\x23\x31\xf9\x44\x7f\x32\x7d\x3d\x9a\x77\x83\x54\x03\x99\xb8\x6b\x42\x0d\xff\x29\x57\xb0\x5d\x41\x55\xed\xd2\xbf\xff\xce\x90\x81\xd0\xd4\x1f\x91\xfa\x93\x6d\xc8\xf3\x66\x5f\x56\x85\x24\x72\x01\xb5\x24\x2e\x3c\x9e\x16\x7a\x13\xe1\xeb\x80\xba\x88\x13\xa6\x50\x3a\x1c\xbd\x28\xab\x22\xb9\x68\x88\x0f\x51\x75\xde\x28\x83\xfe\x9c\x7a\x26\xb3\x23\xbb\x91\x27\xde\x62\x48\xc9\xe9\x39\xfe\xe9\x2e\x83\x99\xad\x96\x90\x47\x4f\x09\x8a\x1f\xf1\x35\xe1\xaf\x21\x67\x1f\x50\x36\x63\x34\x36\x29\x13\x37\x8c\xcd\x79\xf8\x88\xb5\xd9\x72\x38\x8c\x5e\xdf\x12\x13\x38\x9f\xd8\x19\x03\xf9\xb8\xa6\xc9\x70\x22\x0a\x8d\x35\xe7\xc1\x8e\xa1\xc0\x7d\xb5\x6b\x70\xbb\xdd\x36\xf8\x54\x55\x4f\x5b\xa4\xfd\x46\xca\xbd\x7c\x7c\xfc\x2a\xb7\x4f\x65\x99\xd8\x6b\x2a\x59\xca\x4a\xa0\xc4\xed\x57\xc2\x72\xb7\x2f\xf7\xd8\x7c\x15\x5b\xaa\xaa\x27\xb1\x7b\x6a\x9a\xaa\x12\xbb\x2f\xdb\x2f\xc9\x64\xab\x34\x95\x79\x1c\x36\xa9\x9c\x62\xa6\x97\x75\xbe\x2c\xbe\xce\xa3\xeb\xd3\x24\x59\x66\xb6\xa7\x60\x47\x2f\x28\xc6\x93\x6e\x0f\x0e\xb9\xbf\xc6\xe2\xf1\xf4\x9d\x50\x6e\x1a\xe9\xb7\x7a\xa8\x90\xe4\xb4\x3d\xdf\x4e\xc2\xd9\xf6\x8d\xa5\xba\x28\xa4\xce\x93\xb2\xbc\x21\xff\x4a\x9a\xce\x47\x15\x8b\x31\x35\xf7\x7b\x8d\x13\x6e\x69\x19\x15\x0e\xe8\x45\xaf\x8e\xb1\x10\x50\x87\xf8\x50\x53\xed\x34\xf2\xb8\xa7\x69\x16\x36\x18\x28\xe6\x68\xda\xfc\xf1\x07\x5b\x40\x03\xb3\xe4\xed\xeb\xe8\xe6\x99\x14\x25\xaf\x9c\xdc\xd2\x36\x1d\x24\xf5\x92\x8c\xe5\xb4\x07\xbf\xa3\x25\xa6\x25\x3d\x24\xc3\x52\x8e\xbf\xce\x42\xdc\xca\xf3\xa3\x60\x71\x09\x99\xbd\x6a\x46\x9e\x26\x26\xbd\xb1\x47\x30\xc4\xe9\x55\x79\xbd\xcb\x00\x5e\x95\x91\x35\x3c\xbf\xbc\xcc\x1e\xc7\xef\x68\xc9\xd0\xe8\x51\x5f\x64\xee\x9f\x5f\x5e\x56\xf0\x53\xfc\x97\xe7\xf9\x43\x6c\x9e\x79\x2d\x1e\x62\x4b\x06\xe2\x1a\xfe\x16\xbb\x71\x7a\x19\xcf\x67\x97\x87\x65\x9a\x54\xb3\x40\x06\x30\xa0\x51\x2d\x05\x3e\xe0\xc8\xbd\xf5\x35\x60\x23\x47\x2d\xb3\xff\x07\x00\x00\xff\xff\x6c\xcd\x6e\xa0\x68\x0b\x00\x00"

func bvlc_googlenetYmlBytes() ([]byte, error) {
	return bindataRead(
		_bvlc_googlenetYml,
		"bvlc_googlenet.yml",
	)
}

func bvlc_googlenetYml() (*asset, error) {
	bytes, err := bvlc_googlenetYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bvlc_googlenet.yml", size: 2920, mode: os.FileMode(420), modTime: time.Unix(1502164387, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _bvlc_reference_caffenetYml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x56\x4d\x6f\xe3\x36\x13\xbe\xeb\x57\x0c\x60\xbc\x40\xf6\x6d\x2c\x59\xb2\x93\x38\x2a\x50\xb4\xcd\x5e\xb6\x28\x72\x58\x6c\x7b\x29\x0a\x63\x44\x8e\x2c\xee\x52\xa4\x40\x8e\xec\x75\x7f\x7d\x41\x52\xb6\x9c\x6e\xb6\xdb\x1c\x1c\x91\xf3\xcc\x70\xf8\xcc\x17\x0d\xf6\x54\xc3\xcf\xbf\xff\xfa\xb4\x7c\x4f\x2d\x39\x32\x82\x96\x4f\xd8\xb6\xf4\x4c\x0c\x0b\x08\x72\xb0\x2d\x9c\xec\xe8\xa0\xb7\x92\x74\xd6\x3a\xec\xe9\x68\xdd\xa7\x3a\x03\x48\xfa\x51\x01\x16\x70\x11\x41\x6b\x1d\x70\x47\x93\x0a\xc0\x81\x9c\x57\xd6\xd4\x50\xe6\xab\x17\xc0\x49\x00\xc2\x1a\x76\xa8\x0c\x67\xff\x80\x9e\x01\xca\xb4\xd6\xf5\xc8\xe9\x1b\x3c\xf5\x68\x58\x89\x8b\x3c\x49\xb3\x60\x07\x95\x21\x57\xc3\x02\x2e\x0b\x0f\xa3\x27\x09\x6c\x61\x20\x17\x90\xc9\x31\x18\x1c\x49\x25\x82\xcd\x0c\xe6\xbf\x05\xf4\xa3\x66\x35\x68\x82\x41\x23\x07\xbc\x07\x81\x06\x1a\x02\x3f\x90\x50\xad\x22\x99\x01\x60\x2f\xef\x37\x75\xd4\xdc\x0f\x63\x0d\x0e\xd5\xe0\xec\x47\x12\x5c\x08\x74\xbd\x5e\x8a\xc0\x4b\x1d\x61\x4b\x31\x8c\x11\x29\xbe\x89\xdc\x47\xe4\x30\x88\xfb\x8d\xa6\xfa\x9b\x4a\x13\x70\x52\xfb\x77\x57\xae\xb1\x92\xbc\x70\x6a\xe0\xc8\xf5\x0f\x19\xc0\x87\x4e\xf9\x89\x17\xe5\x63\xf4\x1c\xf9\x51\x73\x88\x7f\x6b\xb5\xb6\x47\x65\xf6\x71\x3f\xc5\xfb\x5d\x8f\xfb\x98\x26\x49\x27\x86\x2f\x20\x94\xf1\xec\xc6\xc8\xaa\xcf\xe1\x1d\x07\x6b\x08\x8e\x06\xad\x44\x8a\x9f\x6d\xe7\xe4\x80\xe4\x47\x43\x32\x84\x35\x6c\xff\xa4\xe9\x73\x30\x3b\x8c\xcd\x45\xe3\xa8\xb8\x03\x6f\x7b\x02\xa9\xda\x29\x4f\x7d\x9e\x01\xbc\x9d\x97\x31\x1f\x2d\xcf\x8e\x44\xa5\x74\x0f\xad\xf6\x1d\x87\x3d\x89\x8c\x4b\x1c\xf7\x3d\x19\x8e\xa6\xbf\xcf\x20\x62\xac\x93\xe4\x82\x67\x83\xb5\x3a\x20\xd1\x48\x30\x21\xa9\xb4\xfa\x2b\x39\xa1\xf1\x14\x52\x49\x79\xf0\x47\xc5\xa2\x23\x09\x37\xca\xc0\xb9\x5c\x6e\x2f\xaa\xca\x83\xb4\x86\xa0\xa1\xd6\x3a\x7a\x69\xe5\x4d\xfe\x05\xd5\xde\xe0\xe0\x3b\x1b\x89\x56\x4c\x2e\x9d\xb6\x2e\x57\xb7\xab\xd5\x2a\x87\x0f\x5d\xb0\xe4\x19\x0e\xa8\x95\x4c\xc2\x29\x8d\xd1\x08\x02\x39\xba\x18\x98\xcb\xb5\xd1\xbf\x30\xb3\x0e\x66\x12\x19\x57\x16\x50\x88\xd1\xa1\x38\xc1\xdd\x43\xbe\x29\xab\xff\xc5\xfb\x6a\xeb\x3d\x94\xf9\xb6\x5a\x57\xdb\xfc\xda\x4b\xdb\x84\x42\x0a\x71\x64\x3b\x2c\xcb\x97\xda\x49\x37\x89\xee\x66\xd1\x76\x15\x44\x36\x05\xf5\xea\x64\x1f\xa8\x1a\x7d\x70\xf5\xe3\xe8\x39\x8a\x05\x19\x26\x07\xc2\xd9\x21\x87\x9b\xdf\xfc\x39\xd3\xf0\x40\x0e\xf7\xb1\x05\x95\xab\x28\xf6\xb7\x70\xb3\x81\xef\xa0\x9c\x74\xde\xc0\xff\xa1\x82\x5e\x39\x67\xdd\x2d\xf8\xce\x8e\x5a\x4e\xee\x02\x42\xa3\x18\x3a\xb5\xef\xc8\xcd\x7e\x79\x56\x5a\xe7\x6f\x5e\x86\x21\x90\x16\x19\x24\x09\xcd\x09\x7e\xa1\xb6\x85\xb7\xd6\x60\x37\x12\xfc\xf8\x91\xda\x56\xa6\x45\xe6\xe8\x92\x6f\xb0\x80\x79\x15\x9b\x0b\x0e\x21\x41\x0a\x38\x52\xe3\x15\x53\xf8\x24\x16\x79\x7e\x4e\xf2\xf3\xb5\xce\x6d\x71\x09\x1d\xf3\x50\x17\x45\xd2\xcc\x8d\x1a\x7c\x2e\x44\x5a\x16\x9b\x6d\xb5\x59\xaa\x50\x65\x86\x78\x29\x34\x7a\xaf\xda\xa9\x22\x96\x21\x9e\x4b\x49\x34\x2c\x85\x35\x07\xab\xc7\xb0\x8b\x7a\x69\x68\x74\xf1\x1f\x87\xfe\xea\xf3\x41\xb6\x97\x83\x7c\x5d\x14\x7b\xc5\xdd\xd8\xe4\xc2\xf6\x45\x68\xfb\x45\x6c\x0c\x05\x3b\xa2\xa2\x47\xcf\xe4\x8a\xe8\x9c\x2f\x9a\x83\x16\xbb\xcb\xfd\x76\x11\x67\x88\xb3\x05\x68\x25\xc8\x78\x7a\x51\xc6\xd9\xb4\x59\xc3\x68\x1c\x79\x76\x4a\x30\xc9\x6c\x01\xca\x0c\x23\x47\x76\x66\x6c\xda\x0b\xf5\xba\x80\x56\x39\xcf\x09\x05\x7c\x1a\xe8\x8b\xd1\xb1\x8c\xdb\x35\x44\x22\xb2\xd4\x9f\xaf\x7a\xd7\xd9\x8b\x2b\x3b\x11\xf4\xa2\xbd\x05\x40\x3a\x62\xb6\x32\x60\x18\x42\x4c\x2e\x06\x32\x1e\x3d\x6f\x4d\xe3\x40\xaa\x9e\x4c\x18\x2f\xbe\x86\x3f\xca\x5b\x58\xdf\x42\x55\x3d\xc4\x9f\x3f\x27\x48\x4f\x68\xea\xd7\xd8\x7d\x6f\x1b\x72\xbc\xda\x96\x55\x11\xc2\xe4\x51\x4b\xe2\xc2\xe1\xf1\xcc\x73\x64\x74\xe9\x51\x17\xa1\x29\x15\x4a\xfb\x83\x13\x65\x55\x9c\x23\xbe\x0b\xa6\xf3\x46\x19\x74\xa7\xc1\x59\xb6\x99\x1d\x79\x18\x39\xf1\x16\xae\x14\x9d\x9e\xee\x9f\x64\xa1\x99\x45\xb6\x5a\x42\x1e\x1d\x45\x28\xbe\xc6\x57\xc2\xcf\x57\xce\x5e\xa1\x6c\xc2\x68\x6c\x62\x24\xae\x18\x9b\xe2\xf0\x1a\x6b\xd3\xc9\x7e\x37\x3a\x7d\x4d\x8c\xe7\x3c\xb1\x33\x7a\x72\x61\x32\x93\xe1\x48\x14\x1a\x6b\x4e\xbd\x1d\x7d\x81\xdb\x6a\xd3\xe0\x7a\xbd\x6e\xf0\xb1\xaa\x1e\xd7\x48\xdb\x95\x94\x5b\x79\x7f\xff\x20\xd7\x8f\x65\x19\xd9\x6b\x2a\x59\xca\x4a\xa0\xc4\xf5\x03\x61\xb9\xd9\x96\x5b\x6c\x1e\xc4\x9a\xaa\xea\x51\x6c\x1e\x9b\xa6\xaa\xc4\xe6\x6e\x7d\x17\x8f\x6c\x95\xa6\x32\xe7\xcf\x9c\xc5\x74\x0a\x91\x3e\x4f\xf0\x73\xe5\xef\x1d\x0e\x5d\xec\x60\x47\x0a\x53\xc2\x87\xb1\x67\x47\x27\x28\xdc\x27\x4a\x77\x03\x72\x37\xdf\xc5\xe1\xf1\x2b\x57\xb9\xaa\xa8\xff\x54\x4c\x85\xa4\x41\xdb\x53\x1e\xc3\x1b\xbc\x84\xb3\x13\x57\x47\xd6\x45\x21\x75\x1e\x55\xf2\x86\xdc\x27\xd2\x74\x3a\xa8\x90\x95\xb9\x75\xfb\xaf\x99\x4e\x0a\xa9\x88\x32\x00\xe5\x77\xe8\x44\xa7\x0e\x21\x37\x50\xfb\xf0\x5c\x53\x6d\xea\xc5\xdc\x51\x6a\xd2\x0d\x7a\x0a\x61\x4b\x13\x3b\x7c\xb0\x05\x34\x30\x69\x5e\xbf\x91\xae\x1e\x4b\x41\x73\xa6\xe9\x9a\xc9\xb4\x11\xcd\x4b\x32\x96\x29\x7c\x7f\xc5\x4a\x88\x54\x7c\x4e\xfa\x73\x86\x7e\x19\x98\xd0\xf3\xa6\x47\xc2\xd9\x25\x64\x76\xaa\x19\x39\xb5\x63\xfa\xcc\x0e\x61\xea\x7d\x30\xcb\x32\x80\x4f\xca\xc8\x1a\x9e\x9e\x9f\x27\x8f\xc3\x3a\x9c\x94\xfa\xe5\x45\xe7\xe6\xe9\xf9\xf9\x16\xde\x87\x9f\x3c\x8f\x43\xe2\x3c\x57\x77\xa1\x4a\x3d\x71\x3d\x3f\x7c\x16\x30\xed\x5d\x9e\x97\xb1\x79\x4d\x0a\x19\x40\x8f\x46\xb5\xe4\x79\x87\x23\x77\xd6\xd5\x80\x8d\x1c\xb5\xcc\xfe\x0e\x00\x00\xff\xff\x21\xde\xfe\x11\x77\x0b\x00\x00"

func bvlc_reference_caffenetYmlBytes() ([]byte, error) {
	return bindataRead(
		_bvlc_reference_caffenetYml,
		"bvlc_reference_caffenet.yml",
	)
}

func bvlc_reference_caffenetYml() (*asset, error) {
	bytes, err := bvlc_reference_caffenetYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bvlc_reference_caffenet.yml", size: 2935, mode: os.FileMode(420), modTime: time.Unix(1502164369, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _bvlc_reference_rcnn_ilsvrc13Yml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x56\x4d\x8f\xdb\x36\x10\xbd\xeb\x57\x0c\xe0\x4b\x0b\xac\x25\x7f\x66\x13\x1d\x8a\x22\x0b\xb4\x08\x90\xfa\xb0\x09\x72\x29\x0a\x63\x44\x8e\x2c\x66\x25\x92\x18\x8e\xec\xf5\xbf\x2f\x48\x49\xb6\xb7\xd9\x24\xf5\xc1\xa0\xc8\xc7\x99\x37\x6f\x3e\x24\x8b\x1d\x95\xf0\xfe\xcb\xc7\x87\xf9\x23\xd5\xc4\x64\x15\xcd\x1f\x1f\x76\xbb\xf9\x87\x8f\x9f\xbe\x3c\x3e\x2c\xd7\x30\x83\x08\x02\x57\xc3\xd9\xf5\x0c\x9d\xd3\xd4\x66\x35\x63\x47\x27\xc7\x4f\x65\x06\x30\x18\x79\xc0\xba\x26\x98\xc1\xe5\x08\x6a\xc7\x20\x0d\x8d\x57\x00\x8e\xc4\xc1\x38\x5b\xc2\x32\x5f\xbc\x00\x8e\x07\xa0\x9c\x15\x46\x63\x25\xfb\x0f\x74\x02\x18\x5b\x3b\xee\x50\x86\x35\x04\xea\xd0\x8a\x51\x97\xf3\xe1\x34\x8b\x76\xd0\x58\xe2\x12\x66\x70\x79\x08\xd0\x07\xd2\x20\x0e\x3c\x71\x44\x0e\xc4\xc0\x33\x69\xa3\xa2\xcd\x0c\xae\xbf\x19\x74\x7d\x2b\xc6\xb7\x04\xbe\x45\x89\xf8\x00\x0a\x2d\x54\x04\xc1\x93\x32\xb5\x21\x9d\x01\x60\xa7\xdf\x6c\xca\x74\xf3\xe0\xfb\x12\x18\x8d\x67\xf7\x95\x94\x14\x0a\xb9\x6b\xe7\x2a\xea\x52\x26\xd8\x5c\xf9\x3e\x21\xd5\x4f\x91\x87\x84\xf4\x5e\xbd\xd9\xb4\x54\xfe\xf4\xd2\x08\x1c\xaf\xfd\x98\xca\x2d\x56\x53\x50\x6c\xbc\x24\xad\x7f\xcb\x00\x3e\x37\x04\xbe\x67\x1a\xd3\x69\x6c\x90\x28\xf1\x20\xb9\xab\x53\x3e\x1f\xe7\x0f\xbb\xdd\x28\x5e\xcc\xf1\xa5\x54\x34\x09\x25\x21\x73\xf8\xdc\x98\x30\x42\x4e\x18\xa0\x43\x4d\x50\x9d\x41\x18\x6d\xf0\x6d\x34\x69\x0f\x37\xc6\x3e\x7d\xf9\x0b\x54\x8b\x21\x44\x59\x39\x80\xb1\xe2\x00\xa1\x56\x73\x56\xd6\x5e\x8e\xd4\xc0\xa3\xc5\x33\xf1\x1d\x78\x76\x47\xa3\x49\x43\x43\x4c\x80\x01\x30\x52\xac\xe7\xd2\xd0\x3c\x34\xd4\xd6\x63\x10\x03\x2d\xc7\x39\x7c\xe6\x73\x72\x7a\x21\x0a\xf4\x8c\x5d\xcc\xb1\x38\x08\x44\x60\x24\xd6\x15\x0e\x41\x64\x00\xbb\xfc\x7d\x0e\x7f\x38\x06\xa6\x40\xc8\xaa\x89\xe2\x78\x17\x28\xdc\x41\x87\x4f\x14\x4b\x6a\x92\xc5\xd5\xb5\x51\x06\xdb\x31\x24\x8f\xea\x09\x0f\x04\x68\x35\x58\x27\x20\x51\x91\xd1\x5d\x9e\x94\x7e\xa1\x50\x2a\x7c\xd2\x51\xa4\x47\x17\x02\xfc\x69\x38\x34\x46\x3d\xc1\xef\x5c\x1d\xc6\x75\xc6\x53\x8f\x86\x58\xd9\xd7\xa7\x54\xd4\xe8\xa3\x72\x05\x9c\xa8\x0a\x46\x28\x2e\x49\x54\x9e\xc3\x90\xe4\x6a\x52\x7c\x6a\xc7\x39\x34\x22\x3e\x94\x45\x81\xfc\x6c\x8e\xb9\xe3\x43\xe1\x75\x5d\x2c\xd7\xcb\x65\xbe\xda\xae\x36\xb9\xd7\xf5\x0b\xdc\xc1\x48\xd3\x57\xb9\x72\x5d\x11\x47\x46\x91\xea\xa9\x10\x26\x2a\x3a\x0c\x42\x5c\x24\xdb\xa1\xa8\x8e\xad\xda\x5f\xe8\xed\x63\x0e\xf7\xa6\x0d\x47\x56\xcb\xf5\xc5\x62\x59\x14\xb6\x3a\x1a\x3a\x11\xe7\x5f\x7b\x7f\x16\xe2\xc4\x61\xf0\x72\xeb\xa1\x6a\x5d\x35\x79\x18\x15\x0c\xc5\xb5\xd8\x8c\x3f\xdb\x2a\x9b\x41\x6b\x14\xd9\x6b\x3e\x86\x38\xc7\xcd\x12\x7a\xcb\x14\x84\x8d\x12\xd2\xd9\x0c\x8c\xf5\xbd\x24\xe1\xae\xd8\x61\x2f\x76\xdb\x0c\x6a\xc3\x41\x06\x14\xc8\xd9\xd3\x37\xd3\x6c\x9e\xb6\x4b\x30\x1d\x1e\x28\x1b\x46\xc6\x4d\x3b\x4d\x2c\x6e\xec\x24\xd0\x8b\x8e\x8b\x80\xc1\xc5\xd5\x8a\xc7\x38\x17\x85\x38\xe5\x38\xb9\xbe\x6e\x8d\x13\x4a\x9b\x8e\x6c\x9c\x78\xa1\x84\xbf\x97\x77\xb0\xbe\x83\xd5\xea\x3e\xfd\xfd\x93\xb9\x5e\x7c\x2f\x43\x14\xd1\x41\x32\x31\xd5\x68\x3a\xcb\x60\xe4\x5e\x13\x4a\xcf\x94\xa0\xf8\x1a\xfb\x01\x7f\x25\x90\xbd\x12\xc0\x88\x69\xb1\x4a\xba\xdc\xf0\x1f\x55\x79\x2d\x86\xd1\x73\xd8\xf7\xdc\x96\x37\x05\x16\x24\x1f\xf2\xdf\x07\xe2\x38\xba\xc9\x4a\x2a\x38\xb4\xce\x9e\x3b\xd7\x87\x42\x6f\xd7\x1b\xbd\x5e\x6d\xb6\x8b\x0d\xd2\x62\xb5\xad\xd5\xfd\xe6\x6d\x55\x61\x85\xf7\xb4\x58\x6e\x57\x05\xe3\xa9\x58\xd3\x1a\xf5\xbb\x7b\xac\xb5\xda\xaa\x6a\xad\xef\x17\xd5\xbb\xc5\xf2\xcd\x62\xbd\xde\xa8\x7a\xb9\x7e\xfb\x56\x2f\xdf\x2d\x17\xdb\x65\x2c\xa3\x7d\x38\xdb\x40\xb2\x3f\x39\xd6\x21\x97\x67\xc9\x52\x8e\xa3\xfc\xd3\xa4\x9f\x3a\xf5\xc0\xe8\x9b\xd4\xd2\x27\x32\x87\x46\x42\x1c\x0c\xae\x67\x45\x31\xac\x74\xba\xf7\x28\xcd\x35\x24\xc6\xd3\x77\x22\xba\x29\xf0\xff\xdf\x3d\x85\x26\xdf\xba\x73\xee\xd9\x89\x8b\x54\x61\x62\x72\xe3\xb7\x2c\x0a\xdd\xe6\xc9\x74\x5e\x11\x3f\x51\x4b\xe7\xa3\x89\xf5\x92\x1a\xec\x47\xf6\x87\x5b\x53\x8d\x9b\xb0\x8f\x43\xcf\x1c\x63\xad\x60\x1b\xe2\xfb\xdd\xd4\x10\x48\xee\x62\xea\x6d\xca\x7f\x85\x81\x62\x1a\xc1\x04\x40\x88\x8b\x38\xbb\x2d\x8c\x37\x6f\x5f\xaa\x37\x6f\xd7\x78\xf3\xaa\xd7\xad\xa4\xc3\x46\x32\xaf\xc9\x3a\xa1\xb8\xfe\x8e\x95\xda\xb4\x94\xbe\x3f\xc2\x54\xb1\xdf\x66\xe8\x64\xa4\x31\x03\xd5\x89\x12\x8a\xb0\xa9\x7a\x19\xe6\x28\x3d\x0b\x23\x58\x92\xf4\x31\x72\x3d\xcb\x00\x9e\x8c\xd5\x25\xc4\x71\x3e\x30\x8e\xcf\xd1\x93\xa5\x9e\xb1\xbd\xdc\xf9\xe5\x61\xb7\xbb\x83\xc7\xf8\x97\xe7\xf9\xaf\xb1\xbf\xe2\x38\x37\xf6\xb0\xd7\x28\x18\x48\x4a\xf8\x10\x3b\x7c\x47\x12\xc7\xc4\xb0\x77\xf9\x1e\x49\xa3\x65\xbc\x90\x01\x74\x68\x4d\x4d\x41\xf6\xd8\x4b\xe3\xb8\x04\xac\x74\xdf\xea\xec\xdf\x00\x00\x00\xff\xff\xa1\x04\x77\xa8\xad\x09\x00\x00"

func bvlc_reference_rcnn_ilsvrc13YmlBytes() ([]byte, error) {
	return bindataRead(
		_bvlc_reference_rcnn_ilsvrc13Yml,
		"bvlc_reference_rcnn_ilsvrc13.yml",
	)
}

func bvlc_reference_rcnn_ilsvrc13Yml() (*asset, error) {
	bytes, err := bvlc_reference_rcnn_ilsvrc13YmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bvlc_reference_rcnn_ilsvrc13.yml", size: 2477, mode: os.FileMode(420), modTime: time.Unix(1502164753, 0)}
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
	"bvlc_alexnet.yml": bvlc_alexnetYml,
	"bvlc_googlenet.yml": bvlc_googlenetYml,
	"bvlc_reference_caffenet.yml": bvlc_reference_caffenetYml,
	"bvlc_reference_rcnn_ilsvrc13.yml": bvlc_reference_rcnn_ilsvrc13Yml,
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
	"bvlc_alexnet.yml": &bintree{bvlc_alexnetYml, map[string]*bintree{}},
	"bvlc_googlenet.yml": &bintree{bvlc_googlenetYml, map[string]*bintree{}},
	"bvlc_reference_caffenet.yml": &bintree{bvlc_reference_caffenetYml, map[string]*bintree{}},
	"bvlc_reference_rcnn_ilsvrc13.yml": &bintree{bvlc_reference_rcnn_ilsvrc13Yml, map[string]*bintree{}},
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

