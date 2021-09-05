# go-gcs
Google Cloud Storage Wrapper in GoLang

## Installation

```
go get github.com/hawow/go-gcs
```

## Preparation

You will need to setup your environment variables see `.env.example`

```dotenv
GCS_BUCKET=
GCS_PROJECT_ID=
GCS_JSON_PATH=
```

## Usage

```go
package main

import (
 "context"
 "log"
 "github.com/hawow/go-gcs/"
)

ctx := context.Background()
client := NewGCSClient(ctx)
f, err := os.Open("sample.txt")
if err != nil {
    log.Fatal(err)
    return
}
file := File{
    Name:     "test.txt",
    Path:     "new/test/file",
    Body:     f,
    IsPublic: true,
}
result, err := client.UploadSingleFile(ctx, file)
```