package gogcs

import (
	"cloud.google.com/go/storage"
	"fmt"
	"path"
	"strings"
)

func ObjectToUrl(baseUrl string, objAttrs *storage.ObjectAttrs) string {
	return fmt.Sprintf("%s/%s", baseUrl, objAttrs.Name)
}

func GenerateUrlFromPath(baseUrl string, filePath string) string {
	return fmt.Sprintf("%s/%s", baseUrl, filePath)
}

func MD5BytesToString(bytes []byte) string {
	return fmt.Sprintf("%x", bytes)
}

func GetFullPath(path string, name string) string {
	return fmt.Sprintf("%s/%s", path, name)
}

func GetFileExtension(fileName string) string {
	return strings.ToLower(path.Ext(fileName))
}
