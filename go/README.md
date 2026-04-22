# Memorose Go SDK

Go client for the current Memorose `/v1` runtime.

## Installation

```bash
go get github.com/memorose/memorose-sdk/go
```

## Usage

```go
package main

import (
	"context"
	"fmt"
	"log"

	memorose "github.com/memorose/memorose-sdk/go"
)

func main() {
	ctx := context.Background()
	client := memorose.NewClient("http://127.0.0.1:3000", "your_api_key")
	streamID := "11111111-1111-1111-1111-111111111111"

	_, err := client.IngestEvent(ctx, "user_123", streamID, memorose.IngestRequest{
		Content: "Dylan prefers concise summaries.",
	})
	if err != nil {
		log.Fatal(err)
	}

	response, err := client.RetrieveMemory(ctx, "user_123", streamID, memorose.RetrieveRequest{
		Query: "What does Dylan prefer?",
		Limit: 5,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(response.Results)
}
```

The SDK sends `x-api-key` and targets stream-scoped routes such as `/v1/users/:user_id/streams/:stream_id/retrieve`.
