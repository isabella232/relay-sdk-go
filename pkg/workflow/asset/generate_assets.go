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
		"/schemas/v1/Workflow.json": &vfsgen۰CompressedFileInfo{
			name:             "Workflow.json",
			modTime:          time.Time{},
			uncompressedSize: 2856,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x55\x4b\x4f\x1b\x31\x10\xbe\xe7\x57\x8c\x0c\x47\xe8\x96\x53\xa5\xdc\x10\x5c\x90\xaa\x82\x84\xda\x1e\xaa\x1c\xcc\xee\x6c\x30\xec\xda\xc6\xf6\x86\xa2\x68\xff\x7b\xe5\x47\xbc\xeb\x7d\xa4\x49\x53\x2e\xc4\xb3\x33\xdf\x3c\xbe\x6f\xec\xed\x02\x80\x9c\xeb\xfc\x19\x6b\x4a\x96\x40\x9e\x8d\x91\xcb\x2c\x7b\xd1\x82\x5f\x7a\xeb\x27\xa1\xd6\x59\xa1\x68\x69\x2e\x3f\x7f\xc9\xbc\xed\x8c\x5c\xd8\x38\xc3\x4c\x85\x36\xea\xa7\x50\xaf\x65\x25\xde\xbd\xb9\x40\x9d\x2b\x26\x0d\x13\xdc\x7e\xbc\x06\xa9\xc4\x0b\xe6\x06\x38\x3e\x35\x15\x85\xf7\xe0\x0d\x05\x96\x8c\x33\xe7\xe7\xf1\x3e\xa4\x83\x13\x4f\xd6\xdd\xdb\xa4\x12\x12\x95\x61\xa8\xc9\x12\x6c\xb5\x00\x84\x4a\xf6\x03\x95\xf6\xf8\xde\xd6\x8b\xd6\x46\x31\xbe\x76\xd1\xce\x3e\x28\xe7\xdb\xa0\x08\xdf\x11\x6c\x02\x60\x0c\x43\xde\xd4\x64\x09\xbf\xc2\x19\x80\x6c\xae\x48\x38\xac\xdc\xff\xd6\xfb\x92\x57\xc6\x8b\x7f\x2f\x24\xe4\x77\x41\x7b\x92\xc7\x11\x4f\x96\x90\x42\x1f\x5b\xc9\x35\x07\xe1\x0e\xb4\x82\x46\xa3\xba\x94\x4a\x6c\x58\x81\x45\x9f\xaa\x2e\x22\xc9\x2c\xa9\xa2\x35\x1a\x54\x7a\x2a\x71\x8f\x49\xcf\x5c\x51\x30\x9f\xe8\x61\xcc\xab\xf3\x38\x57\x58\xda\xc8\xb3\xac\x53\x87\xce\x1e\x76\x59\x76\xed\xb7\x49\x11\xda\xa0\x9c\xcc\x4f\x95\xa2\x1f\xb3\x7d\x7f\x65\xda\x80\x28\x7b\x5a\x70\x38\xd1\xbd\x66\xfc\xce\x60\x6d\x91\xaf\xa2\x91\x05\xcb\x5f\x6b\x7e\x34\x28\x07\xe5\x2e\x42\xc9\x44\xe1\x5b\xc3\x14\x16\x91\xe3\xd0\xc2\x02\x60\x15\x76\x28\x02\x75\xba\xef\xa6\x70\xc0\xa8\x47\x5b\x18\xbb\x8c\x94\x0d\x17\xd0\x13\x3a\x43\x4c\x81\x25\x6d\x2a\x93\x18\xc7\x69\x6e\xbd\x57\x2f\xc7\x86\x56\x0d\x92\x18\xd2\x5e\x2c\x66\x62\x13\xd8\x19\xe5\x4e\xa5\xbc\xdf\x49\xb7\xdf\xd7\x40\xad\x1d\x07\x43\xe9\x38\x96\x4e\x1a\xa7\x65\xee\xa8\x49\x72\x5a\xe3\x09\xfd\x7e\xe7\xec\xad\xc1\x41\x7e\x87\x39\x33\x64\x89\xbc\xd0\xf7\xa3\x11\x0f\x60\x1f\x7d\x1b\xd6\x19\x79\x6e\x8b\x4e\x8a\x10\x1c\xef\xcb\xe4\x4a\xb2\x7f\xdb\xe4\x34\xee\x23\xf9\xdc\x5e\x1c\x16\x9b\x2e\x6d\xfc\x3c\x5e\xbc\x83\xb2\x76\xc4\x8f\x4f\xab\xb1\x38\x22\x7b\xa3\x15\x85\xc8\xdd\xee\x0a\x8e\xbe\xe3\xe1\x6c\xe7\x6e\x85\x1b\xc1\x0d\x65\x1c\x95\x13\x5e\x7f\x26\xb3\x21\xd7\xd2\xde\xc7\xb4\x0a\x11\x93\x0f\x40\x0a\xdb\xd3\xf3\x9c\x08\xc3\xc8\x12\x45\xe4\x82\x6b\xbb\xe0\xf6\x87\x47\x9b\x96\x14\xab\xe9\xfa\x14\x05\xdf\x8a\xfc\x15\x15\x38\x18\x28\x85\xf2\x12\xc6\xdf\x98\x37\x83\x95\xed\x25\xcd\x45\x5d\xd3\xe4\xa1\x3d\x36\xed\x8d\x47\x00\x23\x80\x69\x3d\x77\x29\x51\xb5\x1e\xaa\x6c\x8f\x32\xe7\x72\x50\xb5\x6e\x6a\xe4\x66\xb0\x45\xd3\x1a\xde\xa3\xdf\x76\x9a\x00\x2e\x9b\xd1\x4d\x7c\x78\x91\x77\x36\x1c\xbc\xc9\x4e\xc3\x4f\x1e\xff\x5f\xa9\x07\x2e\x93\xd7\xd1\xa4\x9e\x13\xcd\x9f\x2a\x67\x1a\xc0\x26\xde\x82\xc9\x0a\x3d\xda\xaa\x7b\xb3\x17\xed\x9f\x00\x00\x00\xff\xff\x4b\xf6\x4c\x6d\x28\x0b\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/schemas"].(os.FileInfo),
	}
	fs["/schemas"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/schemas/v1"].(os.FileInfo),
	}
	fs["/schemas/v1"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/schemas/v1/Workflow.json"].(os.FileInfo),
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
