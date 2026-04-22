# Memorose Python SDK

Python client for the current Memorose `/v1` runtime.

## Installation

```bash
pip install memorose
```

## Usage

```python
from memorose import MemoroseClient
from memorose.types import IngestRequest, RetrieveRequest

client = MemoroseClient("http://127.0.0.1:3000", "your_api_key")
stream_id = "11111111-1111-1111-1111-111111111111"

client.ingest_event(
    "user_123",
    stream_id,
    IngestRequest(content="Dylan prefers concise summaries.", org_id="default"),
)

response = client.retrieve_memory(
    "user_123",
    stream_id,
    RetrieveRequest(query="What does Dylan prefer?", limit=5, org_id="default"),
)

print(response.results)
```

The SDK sends `x-api-key` and targets stream-scoped routes such as `/v1/users/:user_id/streams/:stream_id/events`.
