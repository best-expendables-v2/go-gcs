package gogcs

import (
	"cloud.google.com/go/storage"
	"context"
	"google.golang.org/api/iterator"
	"io"
	"io/ioutil"
)

type GoGCSClient interface {
	UploadFiles(file []File) ([]UploadedFile, error)
	DownloadFiles(downloads []DownloadedFile) error
	RemoveFiles(objectNames []string) error
	CloneFile(sourceName, destinationName string, isRemoveSource bool) error
	ListFile(path string) ([]ListFile, error)
	GetBaseUrl() string
}

type GoGSCClient struct {
	Client    *storage.Client
	ProjectID string
	Bucket    string
	BaseUrl   string
	Context   context.Context
}

func NewGCSClient(ctx context.Context) (*GoGSCClient, error) {
	config := LoadGSCConfig()
	client, err := storage.NewClient(ctx)

	if err != nil {
		return nil, err
	}

	return &GoGSCClient{
		Client:    client,
		ProjectID: config.ProjectID,
		Bucket:    config.Bucket,
		BaseUrl:   config.BaseUrl,
		Context:   ctx,
	}, nil
}

func (s GoGSCClient) UploadFiles(files []File) ([]UploadedFile, error) {
	bh := s.Client.Bucket(s.Bucket)
	var results []UploadedFile
	for _, file := range files {
		obj := bh.Object(GetFullPath(file.Path, file.Name))
		w := obj.NewWriter(s.Context)

		if _, err := io.Copy(w, file.Body); err != nil {
			return results, err
		}
		if err := w.Close(); err != nil {
			return results, err
		}
		if file.IsPublic {
			if err := obj.ACL().Set(s.Context, storage.AllUsers, storage.RoleReader); err != nil {
				return results, err
			}
		}
		objAttrs, err := obj.Attrs(s.Context)
		if objAttrs == nil {
			return results, err
		}
		results = append(results, UploadedFile{
			Name:        file.Name,
			Size:        objAttrs.Size,
			IsPublic:    file.IsPublic,
			MD5:         MD5BytesToString(objAttrs.MD5),
			Url:         ObjectToUrl(s.BaseUrl, objAttrs),
			ObjectAttrs: objAttrs,
		})
	}
	return results, nil
}

func (s GoGSCClient) downloadFile(download DownloadedFile) (*DownloadedFile, error) {
	rc, err := s.Client.Bucket(s.Bucket).Object(download.Object).NewReader(s.Context)
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(rc)
	if err != nil {
		return nil, err
	}
	if download.Location != nil {
		err = ioutil.WriteFile(GetFullPath(download.Location.Path, download.Location.Name), data, 0644)
		if err != nil {
			return nil, err
		}
	}
	download.Data = data
	return &download, nil
}

func (s GoGSCClient) removeFile(objectName string) error {
	object := s.Client.Bucket(s.Bucket).Object(objectName)
	if err := object.Delete(s.Context); err != nil {
		return err
	}
	return nil
}

func (s GoGSCClient) DownloadFiles(downloads []DownloadedFile) error {
	for k, download := range downloads {
		result, err := s.downloadFile(download)
		if err != nil {
			return err
		}
		downloads[k].Data = result.Data
	}
	return nil
}

func (s GoGSCClient) RemoveFiles(objectNames []string) error {
	for _, objectName := range objectNames {
		err := s.removeFile(objectName)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s GoGSCClient) CloneFile(sourceName, destinationName string, isRemoveSource bool) error {
	src := s.Client.Bucket(s.Bucket).Object(sourceName)
	dst := s.Client.Bucket(s.Bucket).Object(destinationName)
	_, err := dst.CopierFrom(src).Run(s.Context)
	if err != nil {
		return err
	}
	if isRemoveSource {
		return s.removeFile(sourceName)
	}
	return nil
}

func (s GoGSCClient) ListFile(path string) ([]ListFile, error) {
	q := storage.Query{
		Prefix: path,
	}
	var files []ListFile
	it := s.Client.Bucket(s.Bucket).Objects(s.Context, &q)
	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		files = append(files, ListFile{
			Name: attrs.Name,
			Url:  ObjectToUrl(s.BaseUrl, attrs),
			Size: attrs.Size,
		})
	}
	return files, nil
}

func (s GoGSCClient) GetBaseUrl() string {
	return s.BaseUrl
}