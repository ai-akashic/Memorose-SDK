import {
    IngestRequest,
    IngestResponse,
    RetrieveRequest,
    RetrieveResponse,
    UpdateTaskStatusRequest,
    AddEdgeRequest,
    AddEdgeResponse,
    JoinClusterRequest,
    GoalTree,
    L3Task,
    PendingCountResponse,
    ClusterResponse
} from './types';

export class MemoroseAPIError extends Error {
    constructor(public statusCode: number, public message: string, public rawResponse: string) {
        super(`API Error ${statusCode}: ${message}`);
        this.name = 'MemoroseAPIError';
    }
}

export class MemoroseClient {
    private endpoint: string;
    private apiKey: string;
    private timeoutMs: number;

    constructor(endpoint: string, apiKey: string, timeoutMs: number = 10000) {
        this.endpoint = endpoint.replace(/\/$/, '');
        this.apiKey = apiKey;
        this.timeoutMs = timeoutMs;
    }

    private async request<T>(path: string, options: RequestInit = {}): Promise<T> {
        const url = `${this.endpoint}${path}`;
        const headers = {
            'Authorization': `Bearer ${this.apiKey}`,
            'Content-Type': 'application/json',
            ...(options.headers || {})
        };

        const controller = new AbortController();
        const id = setTimeout(() => controller.abort(), this.timeoutMs);

        try {
            const response = await fetch(url, { 
                ...options, 
                headers,
                signal: controller.signal 
            });

            const text = await response.text();

            if (!response.ok) {
                let errorMsg = text;
                try {
                    const parsed = JSON.parse(text);
                    if (parsed.error) errorMsg = parsed.error;
                } catch {
                    // Ignore parse error for fallback to text
                }
                throw new MemoroseAPIError(response.status, errorMsg, text);
            }
            
            return text ? JSON.parse(text) : {};
        } catch (error) {
            if (error instanceof Error && error.name === 'AbortError') {
                throw new Error(`Request timeout after ${this.timeoutMs}ms`);
            }
            throw error;
        } finally {
            clearTimeout(id);
        }
    }

    async ingestEvent(userId: string, appId: string, streamId: string, eventData: IngestRequest): Promise<IngestResponse> {
        return this.request<IngestResponse>(`/v1/users/${userId}/apps/${appId}/streams/${streamId}/events`, {
            method: 'POST',
            body: JSON.stringify(eventData)
        });
    }

    async retrieveMemory(userId: string, appId: string, streamId: string, queryData: RetrieveRequest): Promise<RetrieveResponse> {
        return this.request<RetrieveResponse>(`/v1/users/${userId}/apps/${appId}/streams/${streamId}/retrieve`, {
            method: 'POST',
            body: JSON.stringify(queryData)
        });
    }

    async getTaskTree(userId: string, appId: string, streamId: string): Promise<GoalTree[]> {
        return this.request<GoalTree[]>(`/v1/users/${userId}/apps/${appId}/streams/${streamId}/tasks/tree`);
    }

    async getAllTaskTrees(userId: string): Promise<GoalTree[]> {
        return this.request<GoalTree[]>(`/v1/users/${userId}/tasks/tree`);
    }

    async getReadyTasks(userId: string): Promise<L3Task[]> {
        return this.request<L3Task[]>(`/v1/users/${userId}/tasks/ready`);
    }

    async updateTaskStatus(userId: string, taskId: string, statusData: UpdateTaskStatusRequest): Promise<L3Task> {
        return this.request<L3Task>(`/v1/users/${userId}/tasks/${taskId}/status`, {
            method: 'PUT',
            body: JSON.stringify(statusData)
        });
    }

    async addEdge(userId: string, edgeData: AddEdgeRequest): Promise<AddEdgeResponse> {
        return this.request<AddEdgeResponse>(`/v1/users/${userId}/graph/edges`, {
            method: 'POST',
            body: JSON.stringify(edgeData)
        });
    }

    async getPendingCount(): Promise<PendingCountResponse> {
        return this.request<PendingCountResponse>('/v1/status/pending');
    }

    async initializeCluster(): Promise<ClusterResponse> {
        return this.request<ClusterResponse>('/v1/cluster/initialize', {
            method: 'POST'
        });
    }

    async joinCluster(joinData: JoinClusterRequest): Promise<ClusterResponse> {
        return this.request<ClusterResponse>('/v1/cluster/join', {
            method: 'POST',
            body: JSON.stringify(joinData)
        });
    }

    async leaveCluster(nodeId: string): Promise<boolean> {
        await this.request<void>(`/v1/cluster/nodes/${nodeId}`, { method: 'DELETE' });
        return true;
    }
}
