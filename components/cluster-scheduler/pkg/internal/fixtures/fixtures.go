// Code generated by vfsgen; DO NOT EDIT.

package fixtures

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

// Fixtures statically implements the virtual filesystem provided to vfsgen.
var Fixtures = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(2022, 1, 14, 10, 57, 24, 402123249, time.UTC),
		},
		"/download.yaml": &vfsgen۰CompressedFileInfo{
			name:             "download.yaml",
			modTime:          time.Date(2022, 1, 18, 6, 43, 17, 268443520, time.UTC),
			uncompressedSize: 16866,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xec\x3c\x5d\x73\xe2\xb8\xb2\xef\xfc\x0a\x15\x2f\x7b\xce\xa9\x35\x03\x93\xc9\xd6\x1e\xaa\xee\x43\x00\x9b\x84\x24\x7c\x18\x6c\x12\x4e\x9d\x9a\x12\x76\x83\x1d\x64\xc9\x2b\xc9\x24\xcc\xd4\xfc\xf7\x5b\x92\x6d\x30\x60\x32\x49\x76\x37\x77\xf6\xd6\xe6\x65\xa0\xd5\xea\x6e\xb5\xfa\x4b\x2d\x31\x86\x61\x54\x70\x1c\xba\xc0\x45\xc8\x68\x13\xad\x1b\x95\x55\x48\xfd\x26\xea\xe3\x08\x44\x8c\x3d\xa8\x44\x20\xb1\x8f\x25\x6e\x56\x10\x22\x78\x0e\x44\xa8\x4f\x08\x79\x8c\x4a\xce\x88\x11\x13\x4c\xa1\x99\x7f\x25\xc0\x8d\x08\x53\xbc\x04\x5e\x41\x88\xe2\x08\x9a\x88\xe3\x05\xde\x18\x62\x23\x24\x44\x95\x93\x1c\xc7\xc0\xd7\xa1\x07\x17\x9e\xc7\x12\x2a\xf7\xd8\xa6\x64\x52\x02\x86\xc0\x19\x44\x8b\xf7\x1d\xea\x31\x23\xa1\xb7\xf9\xb0\x6e\xcc\x41\xe2\x9c\xd3\x90\xf9\x63\xf0\x12\x1e\xca\xcd\x50\x8f\x97\x30\x4b\xa9\xc6\x3c\x5c\x87\x04\x96\xe0\x1b\xb1\x88\x2b\x22\x06\x4f\xa1\xec\xc0\x4d\x24\x79\x02\x15\x84\x30\x21\xec\x71\x98\xc3\x4d\xe1\x61\x82\xa5\x16\xa1\x88\x00\x7e\x1b\xc7\x78\x1e\x92\x50\x86\x90\xe9\xd1\x40\xd5\x7f\x55\x2b\x08\xad\x19\x49\xa2\x43\x60\xc0\x84\xec\x83\x7c\x64\x7c\xb5\xa5\xa4\x60\x43\xc6\xe5\x16\x35\x0a\x69\x13\xd5\xf5\x17\x84\x22\xfc\xd4\x44\xbf\x9c\x9f\x9f\x9d\x67\xa8\x57\xc3\xf6\xfe\xd4\xab\xce\xf6\x3b\x4f\xe8\x85\x70\x04\xf0\x94\x14\x4f\x08\x34\x51\xd5\x56\xd0\x0b\xba\x51\x02\x08\xb8\x09\x69\xf2\x74\x7a\x3c\x89\x63\x02\x11\x50\x89\x49\x97\xb3\x24\x16\x27\x51\x17\x42\x23\x9c\x18\x57\x3b\xf7\xf5\xab\x81\xc2\x05\xc2\xd4\x47\xb5\x36\x49\x84\x04\x5e\xbb\xcd\xb6\xa6\x76\xa3\x6d\x0f\xfd\x03\x7e\x43\xff\x08\xa9\x0f\x4f\xa7\x71\xaa\x7a\xf7\x6a\x3e\xac\x3f\xac\x92\x39\x70\x0a\x12\xc4\x90\xb3\x75\xe8\x03\xaf\xfe\x13\x55\x07\x43\xb3\x3f\xbe\xbc\xb2\x26\xd5\x7f\xa2\x6f\xdf\xf6\x2c\x46\x64\x96\x51\x63\x31\x50\x11\x84\x0b\x59\x0b\xd9\x87\x82\x95\xa6\xc3\x6d\x46\x25\x3c\xc9\x36\xa3\x42\x72\x1c\x52\x29\x5e\x62\x44\xc2\xf3\x94\x29\x50\xca\xa4\xb6\x8e\x4c\x59\x21\xf5\x48\xe2\x43\x8d\x03\x01\x2c\x60\x9f\x75\x38\x8f\x0c\x8f\xb0\xc4\xcf\xfc\xca\x6f\xa2\xaa\xda\xbe\xea\xf7\xa7\x0a\x20\x8b\x7c\x96\x11\x84\xcb\xc0\xc0\x6b\x1c\x92\xd4\x00\x37\xaf\xa0\x13\xd2\x25\x01\x83\x32\x1f\x0c\x1f\xd6\x40\x58\x0c\x7c\x6f\xfa\x4e\xcd\x0a\xdf\x07\xe1\xf1\x30\x4e\xed\xff\xa7\xdd\xfa\x53\x27\x10\x08\x7b\x1e\x08\x81\x24\x53\x00\x54\x1c\xa7\xbe\xb6\xd1\xcc\x96\x17\x80\x65\xc2\x41\x68\xb8\x0c\x00\x65\x92\xab\x99\x3c\xa1\x08\xab\x91\x0d\x4a\x04\xf0\x9f\xf5\xa7\xa5\x32\xb0\xf4\x63\x66\x6d\x3f\xeb\xb9\x8f\xa1\x0c\x32\x9a\x6a\x6c\x6c\x6a\xa3\xd6\x31\x0b\x9e\x64\x0d\xa1\xe9\x85\xdd\xbf\xea\x77\x9b\x48\x06\xa1\x40\xa1\xd0\xdc\x22\x26\x24\xe2\x40\xf0\x13\xf8\x68\xdc\x6e\x6b\x52\x22\x60\x09\xf1\xd1\x1c\x14\x5b\x3f\x23\xca\x28\xd9\xa0\x05\xe3\xc8\x4b\x2d\x12\x61\x3f\x0a\x69\xa8\x4c\x43\x29\xa1\x86\xba\x1c\x53\xa9\xc5\x40\x1e\x4e\x34\xec\xa7\xd4\x0f\xca\xf4\xed\x71\xc0\x12\x0c\x45\xb5\xa0\xe5\x62\xf0\xe5\x10\x1b\x2a\x28\x10\x86\x95\x3d\x94\x19\x5a\xb5\xa2\xb5\x7d\xc9\x84\xec\x84\xdc\xd5\xd1\x65\x48\x92\x65\x98\xc7\xa4\xed\xf0\x2e\x42\x6c\x41\xfb\x31\x67\x0b\xde\xc5\x8e\x1d\x48\x47\xa2\x02\xf0\x74\x10\xdc\x1f\xf6\x95\x0b\xe1\x90\x2a\x4b\xda\x0d\x1f\x46\x48\x03\xfd\xf4\xaf\x9f\xf2\x21\x87\x0a\xbc\x80\xf1\x46\x78\x92\x6c\xc7\x7c\x58\xe0\x84\xc8\x0b\x7f\x7f\x26\xa2\x09\x21\x95\x42\xcc\x91\x9b\x18\x9a\x28\x0f\x38\x95\x65\x16\xac\x8c\x2c\xaf\x34\xb3\xad\x33\xf4\xd6\x89\x1d\x5c\x59\x7d\xe1\x6b\x84\x15\x96\xa8\xc4\x3c\x64\x5c\x7b\x91\x66\xc4\x01\xfb\x03\x4a\x36\x36\x63\xd2\x0a\x09\x64\xd8\x68\x81\x89\x80\x0a\x87\xdf\x92\x90\x83\xdf\xe1\x2c\x2e\x91\x72\x2f\x0c\x1f\xc8\x99\x05\xe0\x2c\xde\x94\x22\x78\x1e\x8b\xe2\x21\x67\x0b\xc5\x37\x57\x4b\x79\x60\x3e\x98\xab\x3c\xa7\xa8\x03\x91\xe6\x60\x9c\xe6\xe0\x66\x31\xb5\x36\x33\x3d\xbf\x0c\x79\x97\xab\x5f\x84\x1e\xd2\x25\x07\x21\x0c\xba\x0c\xe9\xd3\x1b\xa6\xe8\x4d\x13\x2a\x80\xbf\x6c\xf2\x12\x4b\x58\x01\xc4\xf9\x76\xbf\x7a\x56\x12\xfb\xca\x43\xb7\x65\x88\xa1\x9d\xf3\x3b\x64\x42\xba\xe0\xb8\x99\xfb\x2a\x8b\x40\x06\x90\x08\x03\xfb\x38\x96\xc0\xdf\x38\x99\x00\x97\x79\xc1\xf5\x26\x0a\x01\x90\xc8\x80\xa7\x98\xf1\x37\x0b\xa1\xa2\xbf\x21\xa4\xd2\x48\x04\x92\x87\x9e\x78\x1b\x9d\x6c\xb2\xa1\xe6\xbc\x55\x96\x57\xcc\x7d\x99\x45\xa7\xb8\x2f\x35\xe8\x13\x52\xe9\xd4\xb9\x55\xf2\xb6\xd4\x4b\x5d\x35\x2f\x7d\x80\xfa\x87\xd5\x08\x9f\x63\xaf\x86\x13\x19\x30\x1e\x7e\x49\x73\xc9\xea\x57\xb1\x57\x91\x64\x45\x90\xcd\x08\x9c\x2a\x41\x9a\xb9\x81\xa8\xaa\x4b\xb3\xc5\x71\xb8\x0b\x0a\xa9\x14\x2a\xa9\x08\x96\x70\x0f\xf6\x80\x6b\xe0\xf3\x02\xc0\x40\x94\x51\x3b\x43\x74\xec\x9b\xe7\x71\x0f\xd8\x64\x21\xb6\xb6\x2d\xcf\x8e\x99\x6e\x3d\x4a\xec\xd3\x4b\x53\xa2\xfe\xe8\x03\x81\xec\xe3\x12\xa4\xfe\x97\x84\x22\xfd\x10\x63\xe9\x05\xfa\x53\xea\xa1\xfa\xe3\xa3\x06\xfe\x3e\x69\x3e\x28\xf3\x4e\x0e\x84\xca\xd9\x1f\x71\x7d\x0b\x2f\x89\xc5\x8a\x80\xfc\x21\xd6\x9d\xcb\xf2\x3e\xab\xfe\x61\x96\xfc\xda\xf5\x1e\x1c\x36\x5f\xe7\xac\xad\x90\xfa\x21\x5d\x9e\xf4\xd9\xbd\xd4\x93\x39\xb0\xc1\x19\x81\x79\x36\x51\x7d\xb6\x61\xa1\xe6\xe5\xcb\x7d\x46\x86\x0a\x42\xc7\xf1\xa2\x3c\x4a\x88\x64\xfe\x00\x9e\xd4\x81\xa2\xf4\x6c\xfe\xaa\x13\xf9\xf3\x24\xf2\x18\xfc\xf2\x23\xfd\x1f\x14\x12\x63\xce\x9e\x36\x5a\x9f\x27\xa3\xa2\x62\x02\x54\x86\xde\x81\x1e\x0f\x2d\x87\xad\x80\x72\x58\x87\xf0\x58\x6e\xc9\xa5\x84\x4b\xf6\xe7\x80\x6e\xb6\x0d\xe9\x61\xe9\x59\xfa\xef\x67\x87\x3b\xad\xe5\x56\x78\x7a\xe3\xfe\x70\xfb\x2c\x6c\xd9\x8f\x6d\xa2\xc2\x0b\xc0\x4f\x48\x48\x97\x05\xc5\xef\x1d\x8a\x27\xea\x98\x99\x1f\x21\x90\x47\xb0\x10\x07\x27\x4b\x7d\x9c\xd4\x5c\x50\x56\x68\xa0\x98\xf9\x42\x1f\x35\x6b\x79\x23\x2b\x23\xd0\x56\xf3\x4f\xf6\x1f\xf2\x83\x8d\xc7\x43\x65\xcb\xa4\xb2\xc6\x24\x81\x26\x6a\xd4\xf3\xbf\xb7\x58\xd0\xb3\xce\x45\x00\xfb\xc0\x0d\x20\xe0\xa9\xc9\xe9\x9e\x3d\x63\x29\x27\x1c\xb0\x5a\x3d\xf6\x0a\x8f\xd1\x45\xb8\x8c\x70\x7c\x22\x4a\x6f\xf3\xc3\xe3\x36\x5c\x17\x12\x4a\x21\x55\xec\xc2\x79\x96\x65\x5e\xc9\xfe\xd9\x5c\x71\xc4\xe7\x45\xc4\x61\x0d\x54\xfe\x61\x3e\xfe\x12\xe7\x2e\xdb\xa8\x3f\xcd\xb3\x4b\x5c\xba\xd4\x50\x7e\x6c\xe7\xc6\x71\x08\x4f\x12\xa8\xfa\x26\x8e\x03\x6b\x22\x24\x8b\xf2\x02\xb9\x03\x8b\x90\x86\x6a\x65\x25\x3b\xb0\x2b\x2e\x6b\xc7\xf5\x49\xde\x6a\xf6\x18\x5d\x67\x9c\x75\xe3\x47\xb7\x94\x60\xb9\x69\xa2\x3e\xa3\x4a\x9b\xcb\x54\xfb\x65\x15\x8e\x66\x90\x4e\x3b\x6c\xea\x2b\x98\xf2\x94\xeb\x7d\xf8\x4d\x98\xf5\xde\x62\x92\x70\x4c\x9a\xfb\xe5\x38\x52\x21\x8a\xcb\x3e\x2e\xb4\xa9\x39\xcd\x46\x42\xba\x4c\x08\xe6\x85\x29\x15\x84\x84\xc7\x62\x28\xd0\xf7\x53\xdb\x16\x79\xeb\xd3\xc8\x54\xb1\xfe\x98\x12\x51\x67\xb7\x5d\x4f\x5d\xad\x96\x71\xbc\x84\x22\xc4\x0b\x20\xc2\xcd\xbc\xef\x16\x03\xbd\x18\x5e\xb9\x67\xe3\x3d\x30\x42\x31\x67\x31\xf0\x5d\x93\x3d\xeb\xfd\xed\xb6\xb1\x00\xcd\xbb\x22\x42\xf2\xd4\xec\xf3\x3f\xad\xb4\x97\x20\x16\xf7\xf6\x10\x99\x69\x5b\x2e\x0c\xe4\x3b\xfb\x5d\x44\x84\x9e\x8c\x5d\x5f\xd5\x88\x39\x68\xfd\x18\x09\x5d\x51\xf6\x48\x8d\x45\x08\xc4\x17\x05\xdd\x64\xf4\x75\x58\xfa\xf3\x38\x1c\xd1\x12\xc9\x7c\x2f\x8e\x15\x84\x40\x5f\xbf\x55\x76\xf2\xa8\x4a\x26\x96\xe0\xf7\x0f\xed\xb2\x5a\xdd\x33\x3a\xfd\xd5\x63\xd4\x0f\xd3\x1e\x39\xfa\xcf\x7f\x2b\xa9\x2d\x80\xef\xe6\xd6\xa3\x80\x7f\x9a\x6b\xe6\xe7\x9f\xf7\x72\xcc\x49\xca\xef\xc0\x2d\x33\xe8\x91\x53\x16\x4e\x8a\xa5\x2e\x29\xc9\x81\x4b\xca\x2d\xf9\xbf\x1d\xf2\x6f\x87\xfc\xab\x3a\xe4\xbb\x7a\x63\x89\x2b\x96\xfa\xe1\x69\x27\x2c\xf1\xc1\xbf\x1d\xf0\x6f\x07\x7c\x5f\x07\x3c\x7c\x5d\xb1\xe7\x5a\x47\xf7\xd0\xbb\x86\xb5\xf2\xd3\x98\x71\xd9\x44\xd5\x5f\x3f\x7d\x3a\xab\x96\x0c\x6b\xbb\x83\x26\x0a\xa4\x8c\x45\xe9\x38\xc7\x4a\x2f\xa5\xf7\x97\x2f\x7f\x3c\x72\x3c\xb4\x77\x4d\x10\x7a\xcf\x1d\x2c\xb7\x0f\x36\xf2\xe7\x12\xb9\x7b\x15\x84\xd6\xab\x54\x8b\xd4\x5f\x25\xe6\x4b\xd0\x97\x9a\x3b\x24\xa1\x8f\x27\x8c\xbf\x54\xf4\xe3\x28\x18\x8b\x5d\xd4\xeb\x40\x4c\xd8\x26\x82\xd2\x37\x2e\x27\x15\x51\xb6\xba\xd7\xab\x54\x39\x7c\x5c\xdb\xbf\xac\xcf\xdf\x07\xcc\x37\x4d\x74\x09\x24\x2a\x33\x0c\x25\x68\x2d\x00\x12\xd5\x44\xf0\x21\xbb\xb2\x36\xf2\x40\x95\xde\x79\x3c\x8f\xf7\xdc\xe6\xec\xab\x37\x52\xc7\xe5\x9b\xc2\xba\x5e\xb6\x32\x0e\x31\x09\x3d\x2c\x9a\xa8\x51\x41\x48\x42\x14\x13\x2c\x61\x27\x7d\x31\x18\x91\x3d\xea\x2f\xd5\x5c\x31\x4a\xe1\x85\xce\x55\x9b\x1d\x0d\xca\x7c\xb8\x38\x82\x2a\xb9\xb2\x9b\xdf\x44\x45\xc8\xf1\xb6\x31\x74\xb5\xa4\x6c\x0b\x36\x9f\xc0\xd3\x8f\x02\xf6\x63\x94\xa2\x39\xce\x94\x33\x01\x1e\x1d\x84\x30\xfd\xf4\x47\x69\xcb\x7c\x52\x71\x4a\xec\xb6\xeb\x10\x6b\x05\x9b\xe6\xc1\x13\x0d\x26\x4a\x30\x75\xfa\xe0\x58\xed\x05\xba\xa2\xa5\x08\xba\x77\x54\xca\x26\x65\x45\x42\x9a\x3c\xbd\x54\x08\xcc\xbd\xe0\xcf\x11\x03\x47\xfe\x2f\x9f\x2a\x79\xe6\x2b\x74\xc9\xfa\xcf\x35\xc6\xb2\x48\xbd\xd7\x29\xe8\x1f\x75\x1a\x72\x93\xd1\x6f\x17\xb6\x32\x18\xc8\x63\x51\x84\x8b\xf9\xd1\x40\x1f\x8a\x06\xa4\x0d\x87\x2f\x45\x11\xc1\x30\x80\xe2\x39\x01\xe3\xa0\x11\xb2\x45\x09\x23\x9d\xed\xab\x5f\xbf\xa2\x5a\x87\x3d\x52\xc2\xb0\xdf\xc1\x12\xd7\xda\x5b\x2b\xbd\x52\x28\xe8\xdb\xb7\xea\xce\x18\xb5\xcc\x87\xbc\x8f\x52\x52\x5a\xe2\x44\xa1\x3c\x50\xa4\x17\x27\x4d\xf4\xf1\xbc\x1e\xed\x41\x23\x88\x18\xdf\x34\x51\xe3\xe3\xaf\xb7\xe1\x81\x85\x83\x78\x3b\x0d\x09\x3c\x0a\xa9\x8e\x36\x5d\x8e\x3d\x18\x02\x0f\xf5\x03\x3d\x46\x55\xa6\x6d\xd4\x4f\x27\xb5\xb6\x6e\xc6\xdd\xe2\xb8\x3c\x92\x52\xed\x39\x46\xda\xb2\x7b\x71\x1c\xfd\x81\x22\x64\x25\x5f\x52\x16\x92\x2e\x7c\x9f\x3f\x63\x0b\x6a\x38\xb3\x03\x1c\x87\xa7\x90\x2f\x86\x57\x05\x44\x7d\x79\x91\xa1\xe5\xcf\xe9\xc6\x31\x78\xf9\x97\x94\x7e\xfa\x79\xac\xeb\x96\xda\x44\x4d\xd1\xd3\xdf\xb4\x31\x69\x17\xff\xaf\xbb\x29\xf9\x6a\x54\x75\x30\x54\x6b\xd1\xda\x0b\x17\x07\xfa\xd3\x43\xa9\x1e\xbe\x7d\xfb\xfa\x15\x3d\x87\x50\x33\x75\x10\xf0\x91\xc2\x34\x9e\xc1\xbb\xcc\x79\x22\x43\xa1\x02\xf5\xb7\xff\x54\x33\x91\xc4\xff\x85\x4c\xe2\x19\xa1\x28\x7b\x6f\x89\xfa\xec\x19\x71\xb4\xf9\x5d\x24\x32\x78\x47\x81\xba\x20\x87\x39\xdb\x6d\x77\xbc\x5c\xbc\x39\x63\x52\x1d\x5e\xe3\xf6\xc5\x3b\x0a\xd8\xda\x71\x2d\x17\x4b\x3f\xc0\xbb\xa2\xfa\x71\x2c\x6c\xb1\xff\x50\x09\xd1\xf3\x22\x5e\x94\x4a\x50\x26\xed\xef\x28\xbf\xb3\xba\x20\x4f\x1d\x7f\xc5\xf0\xf4\xe2\xaa\x1a\xc7\x71\xd9\x82\x7f\x4f\x21\x7d\x82\xe4\xdf\xb5\xf3\x5f\xb6\x76\xfe\x7e\x71\xf6\xbe\x35\xf6\x41\x6d\xf7\xa2\x32\x39\x45\x3d\xae\x92\xd3\x57\x78\xb7\x4a\x80\xbd\x9a\xfc\x64\xfd\xb8\xad\x62\xd5\x94\x21\x56\x29\xe4\x03\x48\xef\xc3\x01\x02\xd0\xf5\x31\xb9\xe1\xa0\xf3\xb9\x7f\x71\x6b\x56\x0e\xf6\xc1\xe2\x2c\xda\xdf\x0a\xdd\x62\xca\x6e\x57\x8f\xe0\x29\xd7\xdc\x0f\x6b\x8a\xf6\x49\x56\xe3\xe1\x45\xfb\x4f\xe0\xb7\xbb\x40\x2c\x32\xbd\x9c\x4c\x86\x9f\x87\xf6\xe0\xee\xfe\xfb\x1c\xbd\xbc\x3a\xbc\x86\x4d\x09\xe3\xd2\x32\xb1\xf8\xa7\xbd\x61\x5b\x7e\x55\x0e\xcd\x5f\x19\xab\x6e\xcf\x16\x7b\x74\x45\x39\xc7\xef\x2e\xa8\x78\x83\xa4\xfd\xc1\xbb\x8a\x99\x95\x69\xaf\x93\xb1\x35\x18\x4c\xc6\x13\xfb\x62\xf8\xb9\x7d\xf1\xb9\x6d\xda\x93\x77\x12\xb6\x50\x25\xbd\x4e\xe0\x8b\x9b\x9b\xc1\xf4\xf3\x55\x7f\x6c\xb6\x1d\xdb\xfc\xbc\x95\xff\x9d\xe4\x2e\x2f\xa3\x9e\x5d\x82\x06\x1d\xfb\xb8\x66\xb4\x57\x4f\xbe\xd3\x12\xb6\xf5\xf3\xcb\x14\x4f\xc2\x35\x50\x10\xca\x01\xe6\x50\xe4\xa8\xdc\xa2\x0b\x72\x5f\x88\x38\x0d\xaa\x01\x60\x22\x83\x2f\xfb\x43\x69\x9f\xb6\xfe\x6b\x7d\x0f\x9c\x77\xa1\x95\x5f\x17\x06\xf4\x4d\x12\x26\x1d\x20\x78\xb3\xcd\x56\xff\x2e\x4e\x8d\xf7\x33\xd9\x59\x71\x4c\x86\x11\xb0\x44\x1e\xa7\x39\x1d\x17\x71\x48\x12\x0e\x93\x80\x83\x08\x18\xf1\x9b\xe8\xac\x52\x4c\x28\xa7\xb3\xd5\xa1\x56\xb7\xfb\x51\x54\xc2\x89\x39\xaf\x6a\xec\x17\x8b\x59\x9f\x87\x0b\xf9\xff\xa7\x90\x3d\x51\x62\xc6\xc5\xdf\x43\xc6\x9c\x49\xe6\x31\xd2\x44\x93\x76\x6e\x14\x5b\xf3\x69\xe4\x95\x4d\xa1\xd1\xaf\xc1\x47\xc7\x86\xfc\x27\x25\x1c\x96\xbb\x5f\x72\x1d\xde\x61\xba\x98\x84\x3e\x96\x21\x5d\x4e\x61\x1e\x30\xb6\x4a\x8f\x2b\x49\x8a\x5e\xb2\x2b\xd5\xc2\xb6\x18\xeb\x74\xf6\xd1\xf5\xc8\x0f\xb4\x0d\x8f\xe9\xb2\xf6\x6e\x4f\xf2\xdf\x57\xaa\x35\xd4\x8a\x6b\x48\x7f\xd7\x59\xf0\x80\xdd\xa3\x37\xf4\x9f\xea\xbf\xaa\xff\xad\x1c\x5f\x3c\x1e\x0d\xa5\x05\x6c\xd9\xc8\xae\x83\xa9\x06\x3e\x14\x87\xb2\x7b\xd5\xf4\x57\xb3\x08\x79\x24\x04\x2a\xd3\xbd\xd8\x5e\x23\xe0\x56\x42\x7d\x02\x4d\x74\x33\xae\xcb\x9b\x71\xa3\x6d\x3b\x7e\x6f\x12\xb6\x3a\xb6\xdb\x73\xc6\xce\xac\x37\xaa\x5b\x8e\xad\xc7\xea\xb2\xbd\x6a\xf4\xc6\x4e\xbf\xed\x9a\xfd\xce\x68\x6a\x9d\xcf\xea\xd6\xd9\xd8\xe9\xb5\x66\x75\xe2\xce\x2e\x2f\xea\x03\xcb\xfa\x65\xe2\xba\x37\xa3\xa9\xd5\xf1\xce\x48\xef\xb6\x1e\xcc\x6e\xea\x33\x67\x7e\xd7\xfa\x04\xab\x19\xf7\xeb\x7d\x3a\x75\xc8\xcd\xbc\x11\xf7\x20\x62\x8f\xb6\x63\x9d\x8d\xef\xbc\xeb\xa9\x6b\x5b\x13\xb3\xd1\xc2\x75\xbf\x75\xeb\xba\xd6\x28\x0a\xfa\xa3\xc6\x6c\x3a\xb9\x0b\x3a\xb0\xb2\xae\x47\x91\x3f\x70\x57\xbd\xd6\xac\x61\xb7\x70\xfd\xbc\x3d\x71\x7b\x3d\xbf\xee\xb6\xa6\x8e\x3d\x75\x5c\xeb\xca\xbe\xf3\x85\xdb\xb5\xef\xfa\x77\x2e\x99\x5a\x31\xbb\x5f\xf9\x2e\x5c\xc4\x5d\xbc\xb2\xdc\x7c\xee\xbc\x61\x9b\xee\x2a\x66\xd3\xa8\xd1\xed\x3b\xa4\x3b\x79\x70\x1f\x6e\x1b\x3d\x71\x4f\x82\x3e\x98\xbe\x33\x72\x83\xf6\xac\x7e\x3e\x1d\xad\xac\xfe\xc4\x71\xbb\xe3\x28\xc0\x73\xe7\xbe\x31\xed\x58\xab\xb9\x45\x24\xbe\x9c\x85\xb7\xd4\x8e\xda\xc4\x7c\xb4\xef\xfc\x2b\xdc\xed\x0f\x26\xd3\xb8\x05\x66\xc3\xb1\x27\x6e\xdf\x36\xc9\xa7\x09\xb1\x5b\x7e\xc3\x67\xf7\xf5\xf3\x3e\x26\xee\xa7\x89\x6b\x5b\x7d\xb3\x61\x8e\xef\x82\x81\x6b\x5a\x67\xee\xc7\xb8\x8d\xcd\x86\xed\x9d\xf9\x1d\xc7\x25\xa6\x4b\x2c\xdb\xae\xbb\x67\x63\x72\x7f\xed\xde\xc5\xfb\x3a\xa8\xf7\xfa\xa3\x86\xe2\x15\x58\x78\x65\xd9\xf9\x3a\xee\x1b\x76\xc7\x5d\x7d\x7a\xba\x8f\xea\x8d\xfe\xd4\x96\x36\xed\x63\xd7\x74\xef\x26\x4e\xcf\xb6\xeb\xe6\x93\xeb\xb8\x1d\x6c\x36\x06\xee\xb4\xd1\x9d\x8f\xe2\xd9\xb4\x2b\xa9\x3b\xe9\x0b\x88\xec\x4b\x97\xda\x0f\xb0\x72\xf1\xc4\xe9\x3d\x64\xb8\x2d\xff\xcc\xb7\x35\xae\x4b\xee\xe7\xdd\xd9\xf4\x96\xcc\xe2\x7b\xc7\x3b\xf7\x23\xf3\x69\xe0\xd8\xb6\x6b\xf6\x14\x7e\xdf\xae\xf7\xce\x3d\xc7\x9f\x4c\x26\xa3\xf3\xf6\xca\xa2\xb6\xe3\x77\x46\x67\xd6\xa5\x53\xaf\xd7\x07\x8e\x75\x66\x3b\x41\xeb\xd6\x24\xad\xd1\xaa\x77\x3d\xb3\x5a\x4f\xd7\xf5\x59\x30\x9d\xb4\x7e\xb9\x6e\xb4\x6e\x9d\x46\x20\xe6\x91\x77\xee\x99\xd6\xdd\xf4\xce\xd9\xe0\xce\xec\x1e\xba\x44\x38\x34\x78\xc0\xdd\xd9\x95\x7b\x19\x5f\x3a\x8d\xd1\xf5\xec\xc1\x3e\xef\xaf\x66\x13\x7f\x35\xeb\xdf\x76\xfd\x4f\x33\x72\xb1\x76\xeb\xbf\xd6\xe7\x51\xef\x71\xe0\xf4\xbe\x8c\xcf\xdc\x70\x42\x5b\xb7\xf7\x2b\x42\x9c\xcb\x78\x78\xbf\x22\xc9\xc0\xb5\xa3\xdb\x28\xb6\xf1\xaa\xd7\x9a\x38\x66\x5d\xaf\x6b\x6a\x9b\x7e\xdd\x6d\xdf\x9c\xf9\xb6\x3d\x8a\x5b\x7e\x9d\xb4\x47\x51\xdc\x72\xcc\x1e\x9d\x90\xd9\x95\xb3\x6a\xb4\x47\xd3\xfb\x4f\xb6\xd3\x73\x46\x8e\xdd\x72\x9c\xe5\x7a\xe2\x5c\x3d\xe6\x73\x67\x0d\xeb\x6e\xb4\xea\x8d\xfd\x8f\xe7\xbf\x79\xab\xc6\xf8\xde\xb5\x1f\x26\x97\xb1\xd3\x9f\xb2\x27\x6c\x9e\x37\x5c\xa7\xf7\xd0\xa6\x2d\x36\x9a\x5c\x7d\xc1\x2b\xeb\x66\x14\xf9\xd4\x9b\x06\x1c\xaf\xfe\x3d\x72\x5c\xcb\x1c\x4d\xfd\xc1\xd8\xb4\xcc\xd1\xca\x6d\x61\xc7\x12\xb7\x2e\x61\xf3\xbb\x40\x8c\xc8\xed\xb9\x17\xcd\x5a\x83\x69\xf0\xe8\x5b\x81\x73\xdb\x75\x36\xf3\x8f\x81\x83\x1f\x2c\x3e\xb8\x74\x1f\xe6\x77\xa3\x6b\x7c\x69\x3d\xb8\x5f\x5a\xc2\xe9\x2c\xbf\x38\x0e\xa1\x8e\x1b\x10\xc7\x71\x1a\xb6\xd3\x1b\x4d\x9d\xd5\xa3\x53\x3f\x37\x27\x5f\x66\xc2\x59\x9d\x0f\xc7\x1f\xeb\x8f\xa3\x87\x96\x33\xb7\x2c\x67\x6e\xca\x47\xbb\x61\xc7\xb7\x51\x70\xe7\x47\xfd\xd9\xfc\x4b\xfd\x3a\xf3\x69\x6b\xb2\x1a\x2d\x47\x75\x77\xec\x9a\xa4\x3b\x76\xfa\x2d\xd7\x74\xb4\x3f\xdf\x8c\xd8\xff\xec\x1f\x3e\x9b\x7b\xf7\x25\x59\xec\xab\x16\x83\xdf\xe1\x8d\x4a\x31\x76\xef\xc6\xd2\xa2\xa5\xfa\x41\x83\x03\xac\x42\x0c\x2f\x8c\xee\x27\x9e\x6d\x56\xb1\xf5\x83\xfe\x62\x0c\x5c\x37\xaa\x3f\xa3\x6a\xf6\x5f\x15\x64\xa1\x4d\x84\x3e\x98\x8b\x85\x7e\x0d\x9b\x3f\x65\xd9\x96\x20\xe9\x7f\x5e\xd0\x44\x69\x3f\xa4\x52\x5a\xb9\x54\xfe\x37\x00\x00\xff\xff\x80\xc9\x83\x15\xe2\x41\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/download.yaml"].(os.FileInfo),
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