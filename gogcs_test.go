package gogcs

import (
	"context"
	"fmt"
	"os"
	"testing"
)

func TestGoGSCClient_ListFiles(t *testing.T) {
	_ = os.Setenv("GCS_BUCKET", "gank-staging")
	_ = os.Setenv("GCS_PROJECT_ID", "gank-staging-276406")
	_ = os.Setenv("GCS_BASE_URL", "https://cdn-staging.ganknow.com")
	_ = os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "./staging.json")
	ctx := context.Background()
	gscClient, err := NewGCSClient(ctx)
	if err != nil {
		fmt.Printf("[Error] Init gcs client %v \n", err)
		return
	}

	files, err := gscClient.ListFile("branded/import-inventories/new")
	if err != nil {
		fmt.Printf("[Error] generate image resizable %v \n", err)
		return
	}
	fmt.Println(files)

}
