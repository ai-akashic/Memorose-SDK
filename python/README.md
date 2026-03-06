# Memorose Python SDK

Python client for the Memorose Hybrid AI Memory Storage Engine.

## Installation

```bash
pip install .
```

## Usage

```python
from memorose import MemoroseClient

# Initialize the client
client = MemoroseClient(endpoint="http://localhost:8000", api_key="your_api_key")

# Add a memory
memory = client.add_memory(
    content="The capital of France is Paris.",
    metadata={"source": "wikipedia", "confidence": 0.99}
)
print("Added memory:", memory)

# Search memories
results = client.search_memories(query="What is the capital of France?", limit=5)
print("Search results:", results)

# Get a specific memory
memory = client.get_memory(memory_id=memory["id"])

# Delete a memory
client.delete_memory(memory_id=memory["id"])
```
