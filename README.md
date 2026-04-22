# Memorose SDKs

Official multi-language client libraries for the current Memorose `/v1` runtime.

The SDKs follow the live server model in `Memorose/crates/memorose-server/src/main.rs`:

- API-key auth uses `x-api-key: <api-key>`
- memory ingest is stream-scoped: `/v1/users/:user_id/streams/:stream_id/events`
- retrieval is stream-scoped: `/v1/users/:user_id/streams/:stream_id/retrieve`
- task, graph, memory maintenance, organization knowledge, and cluster routes mirror the server REST API

## Supported Languages

| Language | Package | Directory |
| --- | --- | --- |
| Python | `memorose` | `python/` |
| Node.js / TypeScript | `memorose-client` | `node/` |
| Go | `github.com/memorose/memorose-sdk/go` | `go/` |
| Java | `dev.memorose.sdk:memorose-client` | `java/` |

## Quick Start

### Python

```python
from memorose import MemoroseClient
from memorose.types import IngestRequest, RetrieveRequest

client = MemoroseClient("http://127.0.0.1:3000", "your_api_key")
client.ingest_event("user_123", "11111111-1111-1111-1111-111111111111", IngestRequest(content="Dylan prefers concise summaries."))
result = client.retrieve_memory("user_123", "11111111-1111-1111-1111-111111111111", RetrieveRequest(query="What does Dylan prefer?", limit=5))
```

### Node.js / TypeScript

```ts
import { MemoroseClient } from 'memorose-client';

const client = new MemoroseClient('http://127.0.0.1:3000', 'your_api_key');
await client.ingestEvent('user_123', streamId, { content: 'Dylan prefers concise summaries.' });
const result = await client.retrieveMemory('user_123', streamId, { query: 'What does Dylan prefer?', limit: 5 });
```

### Go

```go
client := memorose.NewClient("http://127.0.0.1:3000", "your_api_key")
_, _ = client.IngestEvent(ctx, "user_123", streamID, memorose.IngestRequest{Content: "Dylan prefers concise summaries."})
result, _ := client.RetrieveMemory(ctx, "user_123", streamID, memorose.RetrieveRequest{Query: "What does Dylan prefer?", Limit: 5})
```

### Java

```java
MemoroseClient client = new MemoroseClient("http://127.0.0.1:3000", "your_api_key");
IngestRequest ingest = new IngestRequest("Dylan prefers concise summaries.");
client.ingestEvent("user_123", streamId, ingest);
```

## Development

- Python: `python3 -m unittest python/tests/test_client.py`
- Node.js: `npm test -- --runTestsByPath tests/client.test.ts && npm run build`
- Go: `GOCACHE=/tmp/go-build go test ./...`
- Java: `mvn test`
