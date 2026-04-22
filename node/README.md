# Memorose Node.js SDK

Node.js / TypeScript client for the current Memorose `/v1` runtime.

## Installation

```bash
npm install memorose-client
```

## Usage

```ts
import { MemoroseClient } from 'memorose-client';

const client = new MemoroseClient('http://127.0.0.1:3000', 'your_api_key');
const streamId = '11111111-1111-1111-1111-111111111111';

await client.ingestEvent('user_123', streamId, {
  content: 'Dylan prefers concise summaries.',
  org_id: 'default',
});

const response = await client.retrieveMemory('user_123', streamId, {
  query: 'What does Dylan prefer?',
  limit: 5,
  org_id: 'default',
});

console.log(response.results);
```

The SDK sends `x-api-key` and mirrors the server `/v1` REST routes.
