// Code generated by vfsgen; DO NOT EDIT.

package asset

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// assets statically implements the virtual filesystem provided to vfsgen.
var assets = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Time{},
		},
		"/schemas": &vfsgen۰DirInfo{
			name:    "schemas",
			modTime: time.Time{},
		},
		"/schemas/v1": &vfsgen۰DirInfo{
			name:    "v1",
			modTime: time.Time{},
		},
		"/schemas/v1/StepContainer-common.json": &vfsgen۰CompressedFileInfo{
			name:             "StepContainer-common.json",
			modTime:          time.Time{},
			uncompressedSize: 1718,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x54\x4f\x8f\xda\x3e\x10\xbd\xf3\x29\x2c\x2f\xa7\x9f\x08\x81\xd3\x4f\x70\xab\xba\xaa\xb4\xa7\xad\x8a\xd4\x43\xa3\x6c\xe5\x8d\x27\xe0\x6d\xfc\xa7\xf6\x80\x44\x57\xf9\xee\x95\xed\x24\x24\x40\x28\xea\xa5\x27\xd0\x78\xe6\xbd\x37\x33\x6f\xf2\x3e\x21\x84\x4e\x5d\xb1\x03\xc9\xe8\x9a\xd0\x1d\xa2\x59\xa7\xe9\x9b\xd3\x2a\x89\xd1\xb9\xb6\xdb\x94\x5b\x56\x62\xb2\xf8\x3f\x8d\xb1\x07\x3a\x0b\x75\x82\xf7\x6a\x14\xbc\xee\x2b\x36\x37\x7b\x63\x00\xe7\x85\x96\x4d\xb2\x4b\x0b\xad\x90\x09\x05\x36\x3d\x2c\xd3\x0d\x82\xf9\xd8\x06\x92\x42\x4b\xa9\xd5\xdc\xf3\x45\x4c\x0e\xae\xb0\xc2\xa0\xd0\xca\x63\x6f\x76\xcc\x02\x27\xc6\x6a\x03\x16\x05\x38\xa2\x4b\xc2\x88\x43\x30\xa4\x83\x25\x4c\xf1\xf3\x10\x82\x34\x15\x43\x88\xa8\x78\x34\xe0\xe1\xf4\xeb\x1b\x14\x18\x63\x27\x4c\xba\x26\x7e\x0e\x84\x50\xa1\x76\x60\x05\xfa\x00\x9d\x5a\x28\x7d\xcd\x43\xca\xa1\x14\x4a\x78\x49\x2e\xfd\x24\x2a\xf8\x02\x25\xad\x67\xb1\xc2\xf1\x1f\x5f\xc1\xba\x28\x37\xa2\xf4\xf8\x1c\x5a\xa1\xb6\x74\xd6\xc6\x0d\x43\x04\x1b\x3a\x7b\x39\x64\xcb\x64\x95\x67\x8b\x64\x95\xff\x37\xa5\x21\xa3\x05\x15\x92\x6d\x7b\xb2\xae\x37\xd0\x07\xfc\x7c\xd9\x4b\x78\x7e\xc9\x3e\x24\xdf\x58\xf2\x2b\x6f\x7e\x17\xc9\xea\x7b\xe2\xe9\xc6\x1b\x7c\xf2\xdc\xb4\x6e\x30\xea\x8e\x89\x71\x1e\x32\x58\x35\x20\x2b\x59\xe5\x60\x20\xde\x01\xa2\x50\xdb\x7f\x24\x7f\x13\xd9\xff\xaa\x81\x49\x93\x4e\x7b\x80\x27\x6f\xb4\x9b\xef\xb5\xc5\xd4\xf1\xd9\x07\xb2\x4e\xf2\xfb\xf9\xea\x3b\x7e\xd2\x6b\x6c\x7c\x22\xe1\xcd\xc2\xcf\xbd\xb0\xe0\x8f\x2b\xa3\xa5\xd5\x92\xce\x08\x55\x4c\x02\xcd\x07\x79\x57\x1c\xdc\xbd\x85\xb2\xf5\x4d\x3d\x21\x2f\xc0\x5e\xcb\xeb\xa5\x9d\xfe\xb7\xff\xf2\xc1\xc6\xa3\x65\xee\x58\xf7\xa0\xb1\xee\x42\xf3\x93\x1d\x46\x7c\xd0\xa5\xde\x7d\x97\xa1\x8c\x83\x01\xc5\xdd\xb3\x3a\x1b\x4f\xa7\x8f\x59\xcb\x8e\xc3\xd9\x4b\xa1\x9e\x10\xa4\x17\xb0\x1c\x3c\x88\x26\x7a\xf3\x6e\x1e\x3b\xc6\xcb\x99\xd5\x97\x33\x7b\xbc\x22\xf0\xae\x6f\xc7\xd8\x5d\x0c\x28\xda\x53\xb8\x6d\xd8\x7b\x4d\x79\xcb\x6c\xc3\x8f\xf6\x1f\x3d\x77\x60\xd5\x3e\xec\x72\xc4\x64\xbd\x93\x39\x33\x9c\xbf\xd0\x49\x3d\xf9\x1d\x00\x00\xff\xff\x64\xec\x3c\x37\xb6\x06\x00\x00"),
		},
		"/schemas/v1/StepContainer.json": &vfsgen۰CompressedFileInfo{
			name:             "StepContainer.json",
			modTime:          time.Time{},
			uncompressedSize: 822,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x92\xcd\x4e\xc3\x30\x10\x84\xef\x79\x8a\xd5\xb6\x27\x54\xc7\xe5\x84\xc8\xad\xe2\x01\x38\x70\xa3\x0a\xc8\x4d\xb6\xad\x4b\x62\x1b\x7b\x53\x09\xaa\xbe\x3b\xca\x4f\x23\x17\x15\xf1\x77\xb1\x95\xd1\xce\x37\xb3\x49\x0e\x09\x00\x4e\x43\xb1\xa5\x5a\x61\x06\xb8\x65\x76\x99\x94\xbb\x60\x8d\xe8\xd5\xd4\xfa\x8d\x2c\xbd\x5a\xb3\x98\xdf\xc8\x5e\x9b\xe0\xac\xf3\xe9\x32\xf2\x18\x5a\x35\x95\x4a\x5d\xe3\x1c\x71\x5a\xd8\x7a\x18\x0e\xb2\xb0\x86\x95\x36\xe4\xe5\xfe\x5a\x3e\x30\xb9\xbb\x93\x90\xb6\x41\x3d\xac\xa4\x50\x78\xed\x58\x5b\xd3\x42\x17\x10\x98\x1c\x8c\xd6\x7e\x48\x55\xd5\xfd\x1a\x33\x58\x26\x00\x00\x07\x9c\x7a\x5a\xff\xb3\x82\x28\x6c\x5d\x5b\xd3\x35\x99\xe0\x71\xd6\x93\xbb\x13\x00\xf9\xcd\x51\x1b\x60\x57\x3b\x2a\x18\x67\x27\xdd\xd3\x6b\xa3\x3d\x95\x63\x97\x4e\xdd\x93\x0f\x7a\x58\x68\x90\x5e\xb4\x29\x71\x78\xcc\x47\xbb\xf3\xd6\x91\x67\x4d\x01\xb3\x31\x2c\x02\xc4\x62\xd4\x22\xb0\xd7\x66\x13\xe1\x01\x90\x4c\x53\xb7\x2d\x30\x5e\x11\xf3\x71\xe4\xf8\xb9\xcc\x5f\xd8\x67\x6f\xec\x0b\xb8\x51\x35\xfd\x06\xee\x14\x33\xf9\xee\x63\x3f\x2d\x17\xe2\x51\x89\xf7\x7c\xb8\xe7\xe2\xf6\x59\xe4\x57\x53\xbc\x18\xc4\x9a\xab\xef\x92\x2e\x1a\xcf\x7f\xb1\x9f\xd9\x93\xf8\x6e\xcf\x3c\x39\x26\x1f\x01\x00\x00\xff\xff\xf3\x35\x17\x2b\x36\x03\x00\x00"),
		},
		"/schemas/v1/StepContainerTemplate.json": &vfsgen۰CompressedFileInfo{
			name:             "StepContainerTemplate.json",
			modTime:          time.Time{},
			uncompressedSize: 622,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x91\xbb\x6a\xf3\x40\x10\x85\x7b\x3d\xc5\x30\x76\x69\x69\xff\xbf\x0a\xa8\x0b\x79\x80\x14\x49\x17\x54\xac\x57\x23\x7b\x1d\xed\x25\xb3\x23\x43\x30\x7a\xf7\xa0\x8b\xc5\x12\x52\x25\xcd\x08\x1d\xe6\x7c\xe7\x0c\x7b\x2b\x00\x70\x9f\xcc\x99\x9c\xc6\x1a\xf0\x2c\x12\x6b\xa5\x2e\x29\xf8\x72\x51\xab\xc0\x27\xd5\xb2\xee\xa4\xfc\xf7\xa0\x16\x6d\x87\x87\xd9\x67\xdb\xcc\xe3\xe9\x38\xf4\xba\x8a\x43\x8c\x24\x95\x09\x6e\x5d\x4e\xca\x04\x2f\xda\x7a\x62\x75\xfd\xaf\x5e\x84\xe2\xd3\x5d\x78\x25\x17\x7b\x2d\x54\x4d\x81\x0b\xb4\xa5\x64\xd8\x46\xb1\xc1\x4f\xf0\x47\x48\x42\x11\x36\x04\xc8\x6a\x59\xb6\x75\xdf\x3f\x77\x58\xc3\x5b\x01\x00\x70\xc3\x3d\x53\xf7\xc7\x4e\xa5\x09\xce\x05\x3f\x57\xda\xe1\x78\x58\xc8\xf3\x04\x40\xf9\x8c\x34\x05\x84\xe3\x85\x8c\xe0\xe1\xae\x33\x7d\x0c\x96\xa9\xdd\xba\xcc\xea\x95\x38\xd9\xf5\xb2\x55\x7a\xb7\xbe\xc5\xf5\xb7\xd9\xec\x91\x43\x24\x16\x4b\x09\xeb\x2d\x2c\x03\xe4\x62\xd6\x22\x09\x5b\x7f\xca\xf0\x00\x48\x7e\x70\x53\x0b\xcc\x4f\xc4\x66\x5b\x19\xbf\x97\xf9\x0d\xfb\xc7\x57\xcc\x43\x8a\xfc\x3b\xcd\xa6\x18\x8b\xaf\x00\x00\x00\xff\xff\xbd\x98\xe6\x9a\x6e\x02\x00\x00"),
		},
		"/scripts": &vfsgen۰DirInfo{
			name:    "scripts",
			modTime: time.Time{},
		},
		"/scripts/build.tpl": &vfsgen۰CompressedFileInfo{
			name:             "build.tpl",
			modTime:          time.Time{},
			uncompressedSize: 785,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x52\x5f\x6f\xd3\x4e\x10\x7c\xdf\x4f\x31\xbf\xb3\xfb\x4f\x95\x53\xf5\xf7\xe8\xa8\x40\xa8\x43\x14\x15\x12\x94\xb8\xbc\x10\x84\xce\xf6\xda\x3e\xe1\xdc\x45\xbe\x0b\x50\x2c\x7f\x77\x64\x3b\x34\x51\xe1\x6d\x6f\x76\x66\x76\x3c\xb2\xf7\xdf\xcd\xde\xd6\x37\x89\xd2\x37\xac\xbf\x23\x91\xb6\x24\xf2\x10\x97\xca\x22\x57\x15\x43\x59\xc8\xbd\x33\x5b\xe9\x54\x2a\xab\xea\x09\x05\x6b\xae\xa5\xe3\x0c\xc9\x13\x5c\xc9\x58\x70\xb2\xaf\x24\xd6\xd1\xc3\x08\xd1\x12\x8b\x65\x8c\x69\x34\x8f\x47\x44\x96\x1d\x02\xde\x1b\xec\xd4\x8e\x73\xa9\x2a\x22\x8f\x3c\xdc\x9b\xed\x56\xea\xcc\x92\x47\x14\x2d\xef\x1f\xa6\xab\x3b\xe1\x37\xc3\x14\x06\x99\x49\xbf\x71\xdd\x8a\x81\xfb\x49\xd6\x4a\x26\x15\x9f\x90\xbf\x4e\x56\xb3\xf5\x51\xd1\x3f\xc3\xa0\x15\x7f\xd6\x6f\x1f\xe7\xef\xa3\x97\xa4\x23\xd8\x53\x3b\xef\xce\x31\xe3\xb4\x92\x35\x23\x90\x88\x27\xb3\x35\xd1\x8f\xb2\xfb\xe8\x82\x9d\xd9\x39\x8b\x8b\xd0\x85\x17\x58\x7e\x8c\x17\x93\x0f\xd3\x31\x32\x43\x40\x2a\x2d\x43\xf8\xcd\x01\x6d\x05\x94\x26\xc0\x5d\x11\x80\xde\xe5\xfa\xee\xf2\x40\x98\xac\x66\xad\xc0\xb0\x19\x8f\x09\xd8\xbc\x1e\x1e\x9c\x96\x06\xe2\x51\xd7\x9c\x9a\x42\xab\x5f\x9c\xc1\xec\x9c\x32\x3a\x44\x70\xa2\x7c\x75\xfe\xff\x40\xff\xa9\x1c\x6e\x8f\x36\x6c\x65\x4a\x99\xd1\x4c\x64\x4b\x95\x3b\xf8\x97\x97\x5d\xcc\xf9\x22\x42\x80\x5b\x5c\x5d\x11\x35\x0d\x6a\xa9\x0b\xc6\x68\xbe\x95\x05\x5b\xb4\x2d\xf9\x43\x1b\xf0\x4f\xaa\x43\xb2\x57\x55\xf6\x0c\x1d\x8b\x42\x10\x38\x59\xa0\x69\x30\x5a\x71\x8e\xb6\x45\x10\xf4\xbf\x44\x87\xbc\x53\x15\x6b\xb9\xe5\x0e\x1e\x51\xd3\x04\x60\x9d\x75\x27\x28\x37\x75\xd7\x02\x94\x86\xdf\x74\x75\x7c\x7e\xf3\xe5\x5a\x3c\x8f\xad\x68\xfb\x22\x3b\xc9\x5f\xf9\x80\x7f\x26\x7c\x11\x43\x74\xf3\x62\x38\x1e\xf6\xc6\xad\xc0\xf9\x39\x36\x7d\x41\xbb\x5a\x69\x97\x43\x78\x88\x65\x51\x70\x86\x33\x1b\x9e\xd9\x8d\x16\x38\x91\x41\x1c\x74\xa7\xd1\xfb\x42\x7f\x07\x00\x00\xff\xff\xc9\x6b\x30\xd1\x11\x03\x00\x00"),
		},
		"/templates": &vfsgen۰DirInfo{
			name:    "templates",
			modTime: time.Time{},
		},
		"/templates/bash.v1": &vfsgen۰DirInfo{
			name:    "bash.v1",
			modTime: time.Time{},
		},
		"/templates/bash.v1/Dockerfile.standalone.tpl": &vfsgen۰CompressedFileInfo{
			name:             "Dockerfile.standalone.tpl",
			modTime:          time.Time{},
			uncompressedSize: 704,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x92\x4f\x8f\xda\x30\x10\xc5\xef\xf9\x14\x4f\x11\xa0\xb6\x92\x6d\x09\x15\x0e\x54\x3d\x54\xaa\x50\xff\xd2\x2a\x88\x3d\xed\xc5\x24\x0e\x58\xc4\x8e\x65\x3b\x6c\xb4\x59\x7f\xf7\x15\xb0\x2b\xac\x64\xb9\xec\xd5\x33\xef\x37\xf3\xde\xb8\xeb\x08\x46\xff\x79\x7e\xe0\x3b\xb1\x94\x95\x58\x71\x25\xb0\xf8\x0a\x63\xa5\xf6\x25\x52\x2d\xc9\xd8\x91\x4a\xea\xa6\x25\x5c\x15\xf3\xcf\xd4\x73\x4b\xdb\xc7\x14\x74\xfd\xfd\xf7\x9d\xb0\x4e\xd6\x1a\x24\x84\x24\x26\xad\x7f\x7c\x9b\xce\xe6\x6f\xf1\xc6\x8e\xba\x3d\x9f\xce\xe6\xe9\x70\x6c\x9f\x92\x09\x53\x6f\xb2\x3f\xb1\x7c\xef\xbd\x71\x0b\xc6\xcc\xa5\xc3\x51\x2d\xb6\x4d\xc5\xa9\x69\x8c\x11\x9e\x6a\xe1\x99\x2b\x0e\x4c\x4b\x36\x76\xc3\x15\x97\xd9\xbf\xbf\xe8\x3a\xd0\x9f\xea\x2c\xde\x72\x27\x68\x26\x4a\x84\x90\x64\x9b\x15\x9c\xf0\x20\xa2\x69\xf1\x05\xf7\x09\x00\xa8\x43\x21\x2d\x88\x01\xf3\xca\x30\x2d\x31\x99\xbc\x54\xf2\x62\xf8\xf6\xb0\x13\xfe\x84\xbf\x7a\x3d\x6f\xd1\xb7\x33\xf0\x1d\xc2\xbb\x19\xbd\xa0\x63\x92\xc8\xf7\x35\xd2\xd1\x07\xe4\xfc\x4c\xbc\x2d\xf9\x88\xb8\x1e\x55\x52\x3c\xe1\x72\x2d\xd7\x28\x90\x1c\xe4\x8a\xf7\xdc\x82\xb4\xc7\x5f\xe5\x0d\xed\xb5\x53\x1d\xa1\x25\x39\xc5\x1e\x9d\x23\x84\x4f\xf1\xaf\x02\x6b\x9c\x65\x55\x9d\xf3\x8a\x6d\xa5\xee\x47\x1d\xcd\xb5\x0a\xa4\xb4\xaf\xe1\x27\xcf\x01\x00\x00\xff\xff\xf0\xde\xe3\x2c\xc0\x02\x00\x00"),
		},
		"/templates/bash.v1/Dockerfile.tpl": &vfsgen۰CompressedFileInfo{
			name:             "Dockerfile.tpl",
			modTime:          time.Time{},
			uncompressedSize: 1270,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x93\x51\x4f\xdc\x3a\x10\x85\xdf\xf7\x57\x1c\x45\x80\x00\xe1\x58\x42\x17\xa4\xcb\xd5\x7d\x40\x20\x54\x68\x81\xd5\x52\x2a\x55\x2d\x0f\x5e\x7b\x36\x71\x37\xb1\x5d\xdb\xa1\xdb\x86\xfc\xf7\x2a\xd9\x85\x4d\xb3\x45\x95\xfa\x16\x65\x66\xbe\x19\x9f\x39\x53\xd7\x0c\x7c\x1f\x97\x26\x44\x51\x14\xda\x64\x30\x1a\xda\x20\xe6\x14\x08\xd2\x9a\x28\xb4\x21\x1f\x10\x72\x5b\x15\x0a\x53\x42\xa4\xd2\x59\x2f\xfc\x77\xec\x3a\x32\xaa\xad\xf1\x14\x6c\x51\x45\x6d\x0d\xec\x0c\xe3\x1b\x76\xf4\xef\x5e\x8a\x7d\x0e\xd6\x34\xa3\x51\xdb\x63\x6b\x2c\xe4\x5c\x64\x74\xa1\x0b\xba\x11\x25\xe1\xe4\x7f\x38\xaf\x4d\x9c\x21\x31\x9a\x6d\x07\x56\x68\x53\x2d\x98\x28\xd5\xf1\x3f\x69\x14\x3e\x5d\xfc\x48\x90\xde\x9d\xbf\xfd\x40\x3e\xb4\xe0\x16\xd5\x27\xdd\xbd\x39\x3d\x3c\x3a\xfe\x1d\x6f\x3b\xa4\x21\x17\x87\x47\xc7\xc9\x66\xdb\x21\x65\x42\xce\xde\x4f\xde\xf5\xcb\xf3\x18\x5d\x38\xe1\xdc\x2d\x33\x42\x6a\x68\x5a\x15\x22\x75\x95\x73\x14\x53\x43\x91\x07\x35\xe7\x46\xf3\xed\xb0\x39\xe2\x92\xde\xf6\x1b\x8b\x98\xf7\xb9\x7c\xc9\xe1\x21\x92\x63\xdd\x8c\x09\xd2\x97\xa1\x2e\x26\xb7\xd7\xa8\x6b\xa4\x77\x14\xa3\x36\x59\x48\x2f\x4b\x91\x11\x9a\x66\x34\xb9\xbf\x41\xa0\x08\x46\xd5\x02\xff\xe1\xf3\x08\x00\xca\xb9\xd2\x1e\xcc\x81\xc7\xd2\x71\xa3\xb1\xb3\xb3\x8a\x48\xb5\xf9\xef\x5b\x46\xb1\xa5\xaf\x25\xea\x86\x1f\xaa\xb0\x21\x57\xd3\xfc\x35\x63\xb0\x9f\x3e\x89\x64\x6e\x91\x6c\xed\x42\x8a\x8e\xf8\x7a\xc9\x1e\xfa\xf1\x5e\x24\xc1\x13\x96\x4b\x0e\x55\x09\x26\xc1\xd6\xf8\x28\x3c\xd8\xe2\xf1\x6a\xf6\x4a\xed\x3a\xb3\x7c\x84\xd1\xac\x53\x7d\xbd\xc5\xa6\xd9\xef\x9b\x11\xbc\x0a\x9e\x17\x56\x8a\x82\x4f\xb5\x19\x4a\xdd\xeb\xeb\x4b\xb0\x99\x7f\x16\xbf\x5b\x9b\x70\x73\x30\x66\x2c\x93\x42\xe6\x04\xa1\x14\xa6\x22\xe4\x90\x82\x49\xf2\x51\xcf\xb4\x14\x91\x02\x64\xe5\x0b\x64\x3a\xe2\xcb\x57\x58\x47\x26\x84\xbc\xe5\x56\x4e\x89\x48\x6c\x90\xdd\x59\x4c\xcf\x7a\x56\x39\x55\x4a\xb7\xd7\x27\x8a\xd5\x6b\xc3\xb3\x6f\x36\x07\xa8\x6b\x78\x61\x32\xfa\x53\x79\x67\x46\x34\x4d\x5d\x83\x8c\xc2\xea\x6e\x7a\x9f\xaf\x53\xce\x6c\x59\x0a\xa3\x5e\x86\xf8\x94\x74\xca\xb5\x2f\x4f\x0e\x90\x30\x99\x1c\x2c\xe9\x4f\x28\xab\x10\xdf\xdb\xab\xd0\xe9\xfe\xd0\x6f\x71\x76\x3b\xfe\x88\x24\xe5\xbf\xdc\xc4\x8a\xdc\x9d\x56\x6b\x82\xa4\xdd\xf0\xcb\xb1\x35\x4d\x32\x3a\xbb\x3e\x1f\xf6\x1b\xe6\x3c\x8c\x7e\x06\x00\x00\xff\xff\x79\xb4\xac\xaf\xf6\x04\x00\x00"),
		},
		"/templates/bash.v1/container.yaml": &vfsgen۰CompressedFileInfo{
			name:             "container.yaml",
			modTime:          time.Time{},
			uncompressedSize: 550,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x90\xb1\x6a\xf3\x40\x10\x84\xfb\x7b\x8a\x01\xd7\xf6\x8f\xf9\x3b\x75\x8e\xd3\x84\x04\x62\x88\x49\x63\x5c\x6c\x74\x1b\xeb\xd0\x69\x75\x68\x57\x22\x8f\x1f\x74\x96\x62\x27\xc4\xa4\x12\x2c\xdf\x7c\x9a\x9b\x81\x3b\x0d\xad\x14\x28\x5b\x31\x0a\xc2\xdd\xbf\x61\xed\xea\x20\xbe\xc0\x8b\x71\xda\xce\xe7\x3d\x37\x29\x92\xb1\x53\x5f\xbf\xce\xa1\x61\xed\x42\x43\x27\xd6\xc2\x01\x6f\xa4\x3c\x7e\x01\x9b\xd8\x02\xf7\x6d\x59\x73\xf7\x1e\x22\xaf\x2c\x45\x07\x2c\xa0\x46\xe2\x29\xb6\x92\xe1\xc5\x2d\xfc\x82\x7d\x25\x01\xcf\x89\xc5\xeb\xb3\xcc\xd1\x65\xfe\xab\x53\x36\x0b\x72\xca\x35\xb6\x6d\xd3\x90\xf8\x1d\x59\x75\x6e\xe3\x59\xcb\x2e\x24\xcb\x8d\xf7\x15\x23\x91\x55\xb0\x16\x56\x31\xb4\xe2\x18\x71\x06\xc6\x5b\xd7\x4b\x0e\x0d\x14\x7b\x2e\xa0\xc6\x69\xa5\x95\x03\x1e\xc6\x77\xde\x10\x6e\x62\x0a\xc2\x78\x0a\xd2\x7f\x2c\xc7\x42\x1e\x79\x96\x51\xd8\x2b\x5f\x0b\x29\xa3\xc5\x7f\x07\x6c\xbc\x0f\xa3\x83\xe2\x8e\xca\x7a\x5e\xf1\x87\xfe\x02\x61\xb3\x7b\x44\x9a\xc8\xd1\x1c\x44\x8d\x62\xbc\xb6\x1f\x8e\xdf\xbc\xd3\x14\x7f\x78\xef\x48\x2b\x94\x13\xfa\xcb\x06\x87\xa3\xfb\x0c\x00\x00\xff\xff\x39\xc0\x15\xe3\x26\x02\x00\x00"),
		},
		"/templates/go.v1": &vfsgen۰DirInfo{
			name:    "go.v1",
			modTime: time.Time{},
		},
		"/templates/go.v1/Dockerfile.tpl": &vfsgen۰CompressedFileInfo{
			name:             "Dockerfile.tpl",
			modTime:          time.Time{},
			uncompressedSize: 618,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x91\x5f\x6b\xdb\x30\x14\xc5\xdf\xf5\x29\x0e\x62\xeb\x53\x65\xef\xb9\xd0\x87\x34\xc9\x4a\xb7\x35\x0e\xee\xd6\x31\x4a\x19\x37\x92\xec\x88\xda\x92\x91\xe4\x87\xa1\xe9\xbb\x0f\x3b\xa5\x24\x1d\xa5\x6f\xd2\xfd\x73\xce\xe1\xfe\x52\x12\xf8\x70\x65\x2c\xf9\x3f\xb8\xb8\xc4\xe0\x8d\x8d\x0d\x78\x39\x06\x5f\x76\x4e\x52\x57\xee\x8c\x2d\xad\xde\x8d\x1d\x89\x8f\x81\xa3\xd8\x50\xaf\x21\x72\x66\x9f\xeb\xea\x16\xad\xeb\xc8\xb6\x17\x29\xa1\xb8\xd3\x31\x1a\xdb\x86\xe2\xda\xdd\x6b\x1f\x8c\xb3\xc8\x59\x50\x37\x18\xab\xb1\xb8\xc3\x6e\x34\x9d\xd2\x9e\xad\x37\xf7\x58\x5e\x57\xbf\xd7\x9b\xc5\xd5\xb7\xf5\x0a\x9f\xd8\xcf\xaa\xfe\xba\xba\xa9\x51\xce\x23\x6c\x59\x6d\x7f\xa1\x40\xc1\xea\x1f\x1b\xb4\xee\xb0\x08\x41\x10\x0e\x3c\xa5\x97\xbc\x39\x73\xf0\xa2\x3c\xf1\x5e\xba\xbe\x27\xab\xb6\x14\xf7\x53\x9f\x1d\x52\x9e\x8c\xdc\xf4\xd4\x6a\xe4\x3c\xcb\xd3\xf0\x04\x21\xac\x13\x92\xe4\x5e\x83\x94\x82\x24\x21\xb5\x8f\xa6\x31\x92\xa2\x0e\x38\x3b\xc3\x38\x28\x8a\x5a\xbc\xea\xb0\xe9\x78\xa6\x39\x92\x5e\x28\x65\xa2\x71\x96\xba\x2d\xc9\x27\x6a\x75\x78\xdb\x27\x25\x78\xb2\xad\x7e\x6f\x7d\x0e\x8f\x9c\x53\x82\xb6\x6a\xd2\x9b\x6c\x8f\x9e\x6f\xab\x3c\x1f\xe3\x25\xc4\x03\x9f\x69\x86\x3d\x3f\x07\x17\x92\x9f\x1f\xb4\xff\xa2\x1f\x43\xfc\xee\xbe\x84\x19\xd9\xe3\xb1\xc1\xcc\x42\x88\xc6\xbb\xfe\xf2\x19\xe0\xff\x0c\x4e\xff\x6c\x79\xbb\xc2\xc3\xab\xe2\x23\xfb\x17\x00\x00\xff\xff\x95\xc6\x1f\x70\x6a\x02\x00\x00"),
		},
		"/templates/go.v1/container.yaml": &vfsgen۰CompressedFileInfo{
			name:             "container.yaml",
			modTime:          time.Time{},
			uncompressedSize: 590,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x91\x41\x6b\xe3\x30\x10\x85\xef\xfa\x15\x43\xce\x9b\x04\x93\x9b\x6f\xd9\x2c\x84\x65\xf7\x10\x68\xe8\xa5\x14\x32\xb6\x26\xf6\x60\x59\x12\xd2\xd8\x6d\xfe\x7d\x91\x62\x93\x34\xb4\xf4\x64\xb0\xde\xfb\xde\xcc\x9b\x91\x42\x64\x67\x4b\xa8\x9d\x15\x64\x4b\x61\x3d\x16\xaa\x63\xab\x4b\x78\x12\xf2\xbb\xf9\xf7\x91\x7a\x6f\x50\x48\x45\xdd\x3d\xcf\xa6\xb1\x50\xdc\x63\x43\xb1\x54\x00\x15\x46\x4a\x5f\x00\x99\xb4\x25\xfc\x71\x75\x47\xe1\xcc\x86\x56\xe2\x8d\x8a\x24\xc2\xb6\xc9\xf2\x9d\xeb\x7b\xb4\xfa\x80\xd2\x5e\x5d\x9a\x62\x1d\xd8\x4b\x26\x1f\x5b\x02\x8f\xd2\x82\x38\x90\x96\xa0\xbe\xaa\xa1\x62\x8b\xe1\xf2\x0b\x0c\x77\x04\xa7\xba\xd7\xeb\xfe\xb2\x9c\x1e\x4f\x0a\x60\xef\xe6\xe1\xbe\x66\x4e\xfb\x82\x3b\x67\xec\xde\x81\x41\xdb\x0c\xd8\x50\x4a\xaa\x06\x36\x1a\xde\x58\xda\xec\x1e\xd1\x0c\x54\xc2\xa2\x58\x15\x9b\x85\x02\xf8\x9b\x76\xfd\x06\xbc\x35\x9e\x2d\xc1\x7f\xb6\xc3\xfb\x32\x55\xa1\x21\x57\x93\xb0\x43\x24\x38\xbb\x00\x9a\xa3\x04\xae\x86\xd4\x41\x8e\x8f\x42\xfe\xd6\xfc\x7d\x26\x66\x5c\xb9\x51\x00\x5b\xad\x39\xe5\xa0\x39\x60\xdd\xcd\x6d\x3f\x8c\x70\x13\xc1\xf6\xf0\x0f\xfc\xa4\x4c\xe9\x6c\xa3\xa0\x31\xf7\xf4\x97\xd7\x4f\xdc\xe9\x14\x3f\x70\x7f\x63\x6c\xe7\x3b\x64\x70\x18\xec\x03\xf4\x23\x00\x00\xff\xff\x61\xfa\xc0\x81\x4e\x02\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/schemas"].(os.FileInfo),
		fs["/scripts"].(os.FileInfo),
		fs["/templates"].(os.FileInfo),
	}
	fs["/schemas"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/schemas/v1"].(os.FileInfo),
	}
	fs["/schemas/v1"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/schemas/v1/StepContainer-common.json"].(os.FileInfo),
		fs["/schemas/v1/StepContainer.json"].(os.FileInfo),
		fs["/schemas/v1/StepContainerTemplate.json"].(os.FileInfo),
	}
	fs["/scripts"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/scripts/build.tpl"].(os.FileInfo),
	}
	fs["/templates"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/templates/bash.v1"].(os.FileInfo),
		fs["/templates/go.v1"].(os.FileInfo),
	}
	fs["/templates/bash.v1"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/templates/bash.v1/Dockerfile.standalone.tpl"].(os.FileInfo),
		fs["/templates/bash.v1/Dockerfile.tpl"].(os.FileInfo),
		fs["/templates/bash.v1/container.yaml"].(os.FileInfo),
	}
	fs["/templates/go.v1"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/templates/go.v1/Dockerfile.tpl"].(os.FileInfo),
		fs["/templates/go.v1/container.yaml"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr:                        gr,
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}
