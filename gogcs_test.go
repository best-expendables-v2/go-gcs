package gogcs

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"os"
	"testing"
)

var (
	_   = os.Setenv("GCS_BUCKET", "gank-staging")
	_   = os.Setenv("GCS_PROJECT_ID", "gank-staging-276406")
	_   = os.Setenv("GCS_BASE_URL", "https://cdn-staging.ganknow.com")
	_   = os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "./staging.json")
	ctx = context.Background()
)

func TestGoGSCClient_GoGSCClient(t *testing.T) {
	gcsClient, err := NewGCSClient(ctx)
	if err != nil {
		fmt.Printf("[Error] Init gcs client %v \n", err)
		return
	}
	err = gcsClient.RemoveFiles([]string{"branded/import-inventories/new/Sample-Stockhero.csv"})
	if err != nil {
		fmt.Printf("[Error] delete object %v \n", err)
		return
	}
}

func TestGoGSCClient_GSCClient(t *testing.T) {
	client, err := storage.NewClient(ctx)
	if err != nil {
		fmt.Printf("[Error] Init gcs client %v \n", err)
		return
	}
	obj := client.Bucket("gank-staging").Object("branded/import-inventories/new/Sample-Stockhero.csv")
	if obj != nil {
		attr, err := obj.Attrs(ctx)
		if err != nil {
			fmt.Printf("[Error] Object Attribute %v \n", err)
			return
		}
		fmt.Printf("Object Attribute %#v\n", attr)
		err = obj.Delete(ctx)
		if err != nil {
			fmt.Printf("[Error] Object Delete %v \n", err)
			return
		}
	}
}

func TestGoGSCClient_ListFiles(t *testing.T) {
	gcsClient, err := NewGCSClient(ctx)
	if err != nil {
		fmt.Printf("[Error] Init gcs client %v \n", err)
		return
	}
	downloadFile := []DownloadedFile{{
		Object: "branded/import-inventories/new/Sample-Stockhero.csv",
	}}
	err = gcsClient.DownloadFiles(downloadFile)
	if err != nil {
		fmt.Printf("[Error] download file %v \n", err)
		return
	}
	fmt.Println(string(downloadFile[0].Data))
	_ = gcsClient.CloneFile("branded/import-inventories/pending/Sample-Stockhero.csv", "branded/import-inventories/new/Sample-Stockhero.csv", true)

}
