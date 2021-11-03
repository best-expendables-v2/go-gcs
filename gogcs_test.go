package gogcs

import (
	"bytes"
	"cloud.google.com/go/storage"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"testing"
	"time"
)

var (
	_   = os.Setenv("GCS_BUCKET", "supply-chain-3pl-inventory")
	_   = os.Setenv("GCS_PROJECT_ID", "branded-dev-sandbox")
	_   = os.Setenv("GCS_BASE_URL", "https://cdn-staging.ganknow.com")
	_   = os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "./staging.json")
	_   = os.Setenv("GCS_SERVICE_ACCOUNT", "./staging.json")
	ctx = context.Background()
)

func TestGoGSCClient_GoGSCClient(t *testing.T) {
	gcsClient, err := NewGCSClient(ctx)
	if err != nil {
		fmt.Printf("[Error] Init gcs client %v \n", err)
		return
	}
	err = gcsClient.RemoveFiles([]string{"Copy of STOCK 09032021.csv"})
	if err != nil {
		fmt.Printf("[Error] delete object %v \n", err)
		return
	}
}
func TestGoGSCClient_GSCClientListBuckets(t *testing.T) {
	client, err := storage.NewClient(ctx)
	if err != nil {
		fmt.Printf("[Error] Init gcs client %v \n", err)
		return
	}
	butketAttribute, err := client.Bucket("supply-chain-3pl-inventory").Attrs(ctx)
	if err != nil {
		fmt.Printf("[Error] delete object %v \n", err)
		return
	}
	fmt.Printf("Butket Attribute %#v\n", butketAttribute)
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
	now := time.Now()
	fmt.Println(now.Format("01022006"))
	gcsClient, err := NewGCSClient(ctx)
	if err != nil {
		fmt.Printf("[Error] Init gcs client %v \n", err)
		return
	}

	files, err := gcsClient.ListFile("weekly_3pl_reports/STOCK 09072021.csv")
	if err != nil {
		fmt.Printf("[Error] cannot list files %v \n", err)
		return
	}
	b, err := json.Marshal(files)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
	fmt.Printf("Files %#v\n", files)
}

func TestGoGSCClient_DownloadFiles(t *testing.T) {
	gcsClient, err := NewGCSClient(ctx)
	if err != nil {
		fmt.Printf("[Error] Init gcs client %v \n", err)
		return
	}
	downloadFile := []DownloadedFile{{
		Object: "STOCK 09032021.csv",
	}}
	err = gcsClient.DownloadFiles(downloadFile)
	if err != nil {
		fmt.Printf("[Error] download file %v \n", err)
		return
	}
	fmt.Println(string(downloadFile[0].Data))
	_ = gcsClient.CloneFile("branded/import-inventories/pending/Sample-Stockhero.csv", "branded/import-inventories/new/Sample-Stockhero.csv", true)
}

func TestGoGSCClient_CloneFiles(t *testing.T) {
	gcsClient, err := NewGCSClient(ctx)
	if err != nil {
		fmt.Printf("[Error] Init gcs client %v \n", err)
		return
	}
	_ = gcsClient.CloneFile("STOCK 09072021.csv", "new/STOCK 09072021.csv", true)
}

func TestGoGSCClient_UploadFiles(t *testing.T) {
	gcsClient, err := NewGCSClient(ctx)
	if err != nil {
		fmt.Printf("[Error] Init gcs client %v \n", err)
		return
	}
	file := File{
		Path:     "",
		Name:     "data.csv",
		Body:     bytes.NewBufferString("this is a test"),
		IsPublic: false,
	}
	_, err = gcsClient.UploadFiles([]File{file})
	if err != nil {
		fmt.Printf("[Error] can not upload file %v \n", err)
		return
	}
}

func TestGoGSCClient_SignedURL(t *testing.T) {
	object := "tasks/64831e83-423a-48d5-b04a-7f6d47c99ad9/STOCK 09172021.csv"
	//object = "task-templates/inventory/data.csv"
	gcsClient, err := NewGCSClient(ctx)
	if err != nil {
		fmt.Printf("[Error] Init gcs client %v \n", err)
		return
	}
	signedURL, err := gcsClient.GetSignedURL(object, 600*time.Second)
	if err != nil {
		fmt.Printf("[Error] can not get signed url %v \n", err)
		return
	}
	fmt.Printf("signed url %v\n", signedURL)
}
