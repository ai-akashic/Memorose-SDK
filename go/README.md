# Memorose Go SDK

Go client for the Memorose Hybrid AI Memory Storage Engine.

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

	"github.com/memorose/memorose-sdk/go"
)

func main() {
	client := memorose.NewClient("http://localhost:8000", "your_api_key")
	ctx := context.Background()

	// Add a memory
	metadata := map[string]interface{}{
		"source": "wikipedia",
	}
	mem, err := client.AddMemory(ctx, "The capital of Germany is Berlin.", metadata)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Added memory ID: %s
", mem.ID)

	// Search memories
	results, err := client.SearchMemories(ctx, "What is the capital of Germany?", 5)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Found %d results
", len(results))

	// Delete a memory
	err = client.DeleteMemory(ctx, mem.ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Memory deleted.")
}
```
