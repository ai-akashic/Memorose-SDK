import { MemoroseClient } from '../src/client';
import { IngestRequest, RetrieveRequest } from '../src/types';

global.fetch = jest.fn() as jest.Mock;

describe('MemoroseClient', () => {
    let client: MemoroseClient;

    beforeEach(() => {
        client = new MemoroseClient('http://localhost:8000', 'test_key');
        (global.fetch as jest.Mock).mockClear();
    });

    it('uses current stream route and x-api-key header for ingest', async () => {
        (global.fetch as jest.Mock).mockResolvedValue({
            ok: true,
            text: async () => JSON.stringify({ status: 'accepted', event_id: 'evt_1' }),
        });

        const payload: IngestRequest = { content: 'hello' };
        const result = await client.ingestEvent('user-1', 'stream-1', payload);

        expect(result.event_id).toBe('evt_1');
        expect(global.fetch).toHaveBeenCalledWith(
            'http://localhost:8000/v1/users/user-1/streams/stream-1/events',
            expect.objectContaining({
                method: 'POST',
                headers: expect.objectContaining({
                    'x-api-key': 'test_key',
                }),
            }),
        );
    });

    it('uses current retrieve payload shape', async () => {
        (global.fetch as jest.Mock).mockResolvedValue({
            ok: true,
            text: async () =>
                JSON.stringify({
                    stream_id: 'stream-1',
                    query: 'hello',
                    results: [],
                    query_time_ms: 14,
                }),
        });

        const payload: RetrieveRequest = { query: 'hello', limit: 3, org_id: 'org-1' };
        const result = await client.retrieveMemory('user-1', 'stream-1', payload);

        expect(result.query_time_ms).toBe(14);
        expect(global.fetch).toHaveBeenCalledWith(
            'http://localhost:8000/v1/users/user-1/streams/stream-1/retrieve',
            expect.objectContaining({
                method: 'POST',
                body: JSON.stringify(payload),
            }),
        );
    });
});
