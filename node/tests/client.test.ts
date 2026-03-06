import { MemoroseClient } from '../src/client';

// Simple mock for fetch
global.fetch = jest.fn() as jest.Mock;

describe('MemoroseClient', () => {
    let client: MemoroseClient;

    beforeEach(() => {
        client = new MemoroseClient('http://localhost:8000', 'test_key');
        (global.fetch as jest.Mock).mockClear();
    });

    it('should add a memory', async () => {
        const mockResponse = { id: '1', content: 'test memory', metadata: {} };
        (global.fetch as jest.Mock).mockResolvedValue({
            ok: true,
            json: async () => mockResponse,
        });

        const result = await client.addMemory('test memory', { key: 'value' });
        expect(result.id).toBe('1');
        expect(result.content).toBe('test memory');
        
        expect(global.fetch).toHaveBeenCalledWith(
            'http://localhost:8000/v1/memories',
            expect.objectContaining({
                method: 'POST',
                headers: expect.objectContaining({
                    'Authorization': 'Bearer test_key'
                })
            })
        );
    });

    it('should search memories', async () => {
        const mockResponse = [{ id: '1', content: 'test memory', metadata: {} }];
        (global.fetch as jest.Mock).mockResolvedValue({
            ok: true,
            json: async () => mockResponse,
        });

        const result = await client.searchMemories('test query');
        expect(result.length).toBe(1);
        expect(result[0].id).toBe('1');
    });

    it('should delete a memory', async () => {
        (global.fetch as jest.Mock).mockResolvedValue({
            ok: true,
            json: async () => ({}),
        });

        const result = await client.deleteMemory('1');
        expect(result).toBe(true);
    });
});
