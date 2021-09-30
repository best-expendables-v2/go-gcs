package gogcs

import (
	"cloud.google.com/go/storage"
	"io"
)

var (
	MapExtensionWithContentType = map[string]string{
		".csv": "text/csv",
	}
)

type File struct {
	Path     string
	Name     string
	Body     io.Reader
	IsPublic bool
}

type UploadedFile struct {
	Name        string
	MD5         string
	IsPublic    bool
	Url         string
	Size        int64
	ObjectAttrs *storage.ObjectAttrs
}

type DownloadedFile struct {
	Object   string
	Location *FileLocation
	Data     []byte
}

type FileLocation struct {
	Name string
	Path string
}

type ListFile struct {
	Name string
	Url  string
	Size int64
}
