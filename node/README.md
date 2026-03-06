# Memorose Node.js SDK

Node.js / TypeScript client for the Memorose Hybrid AI Memory Storage Engine.

## Installation

```bash
npm install memorose-client
```

## Usage

```typescript
import { MemoroseClient } from 'memorose-client';

const client = new MemoroseClient('http://localhost:8000', 'your_api_key');

async function main() {
  // Add a memory
  const memory = await client.addMemory('The capital of Japan is Tokyo.', { source: 'geography' });
  console.log('Added memory:', memory);

  // Search memories
  const results = await client.searchMemories('What is the capital of Japan?');
  console.log('Search results:', results);

  // Get a specific memory
  const fetchedMemory = await client.getMemory(memory.id);

  // Delete a memory
  await client.deleteMemory(memory.id);
}

main();
```
