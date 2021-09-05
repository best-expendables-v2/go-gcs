package gogcs

import (
	"cloud.google.com/go/storage"
	"io"
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

// DownloadedFile is used when downloading file from Google Could Storage (GCS).
// Object is the object path that provided from GCS.
// Name is the file name that you wanted to save locally.
// Path is where do you want to have your file stored. It should be a full path.
// Data is raw bytes.
type DownloadedFile struct {
	Object string
	Name   string
	Path   string
	Data   *[]byte
}
