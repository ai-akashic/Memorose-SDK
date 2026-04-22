import {
    AddEdgeRequest,
    AddEdgeResponse,
    BatchIngestRequest,
    BatchIngestResponse,
    ClusterResponse,
    GoalTree,
    IngestRequest,
    IngestResponse,
    JoinClusterRequest,
    L3Task,
    MemoryContextRequest,
    MemoryContextResponse,
    PendingCountResponse,
    RetrieveRequest,
    RetrieveResponse,
    SemanticMemoryExecuteRequest,
    SemanticMemoryPreviewRequest,
    UpdateTaskStatusRequest,
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
            'x-api-key': this.apiKey,
            'Content-Type': 'application/json',
            ...(options.headers || {}),
        };

        const controller = new AbortController();
        const id = setTimeout(() => controller.abort(), this.timeoutMs);

        try {
            const response = await fetch(url, {
                ...options,
                headers,
                signal: controller.signal,
            });

            const text = await response.text();

            if (!response.ok) {
                let errorMsg = text;
                try {
                    const parsed = JSON.parse(text);
                    if (parsed.error) errorMsg = parsed.error;
                } catch {
                    // ignore
                }
                throw new MemoroseAPIError(response.status, errorMsg, text);
            }

            return text ? JSON.parse(text) : ({} as T);
        } catch (error) {
            if (error instanceof Error && error.name === 'AbortError') {
                throw new Error(`Request timeout after ${this.timeoutMs}ms`);
            }
            throw error;
        } finally {
            clearTimeout(id);
        }
    }

    async ingestEvent(userId: string, streamId: string, eventData: IngestRequest): Promise<IngestResponse> {
        return this.request<IngestResponse>(`/v1/users/${userId}/streams/${streamId}/events`, {
            method: 'POST',
            body: JSON.stringify(eventData),
        });
    }

    async ingestEventsBatch(
        userId: string,
        streamId: string,
        payload: BatchIngestRequest,
    ): Promise<BatchIngestResponse> {
        return this.request<BatchIngestResponse>(`/v1/users/${userId}/streams/${streamId}/events/batch`, {
            method: 'POST',
            body: JSON.stringify(payload),
        });
    }

    async retrieveMemory(userId: string, streamId: string, queryData: RetrieveRequest): Promise<RetrieveResponse> {
        return this.request<RetrieveResponse>(`/v1/users/${userId}/streams/${streamId}/retrieve`, {
            method: 'POST',
            body: JSON.stringify(queryData),
        });
    }

    async buildMemoryContext(requestData: MemoryContextRequest): Promise<MemoryContextResponse> {
        return this.request<MemoryContextResponse>('/v1/memory/context', {
            method: 'POST',
            body: JSON.stringify(requestData),
        });
    }

    async deleteMemory(userId: string, memoryId: string): Promise<boolean> {
        await this.request<void>(`/v1/users/${userId}/memories/${memoryId}`, {
            method: 'DELETE',
        });
        return true;
    }

    async previewSemanticMemory(
        userId: string,
        requestData: SemanticMemoryPreviewRequest,
    ): Promise<Record<string, unknown>> {
        return this.request<Record<string, unknown>>(`/v1/users/${userId}/memories/semantic/preview`, {
            method: 'POST',
            body: JSON.stringify(requestData),
        });
    }

    async executeSemanticMemory(
        userId: string,
        requestData: SemanticMemoryExecuteRequest,
    ): Promise<Record<string, unknown>> {
        return this.request<Record<string, unknown>>(`/v1/users/${userId}/memories/semantic/execute`, {
            method: 'POST',
            body: JSON.stringify(requestData),
        });
    }

    async getTaskTree(userId: string, streamId: string): Promise<GoalTree[]> {
        return this.request<GoalTree[]>(`/v1/users/${userId}/streams/${streamId}/tasks/tree`);
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
            body: JSON.stringify(statusData),
        });
    }

    async addEdge(userId: string, edgeData: AddEdgeRequest): Promise<AddEdgeResponse> {
        return this.request<AddEdgeResponse>(`/v1/users/${userId}/graph/edges`, {
            method: 'POST',
            body: JSON.stringify(edgeData),
        });
    }

    async getPendingCount(): Promise<PendingCountResponse> {
        return this.request<PendingCountResponse>('/v1/status/pending');
    }

    async listOrganizationKnowledge(
        orgId: string,
        query: Record<string, string> = {},
    ): Promise<Record<string, unknown>> {
        const params = new URLSearchParams(query);
        const suffix = params.toString() ? `?${params.toString()}` : '';
        return this.request<Record<string, unknown>>(`/v1/organizations/${orgId}/knowledge${suffix}`);
    }

    async getOrganizationKnowledge(orgId: string, knowledgeId: string): Promise<Record<string, unknown>> {
        return this.request<Record<string, unknown>>(`/v1/organizations/${orgId}/knowledge/${knowledgeId}`);
    }

    async getOrganizationKnowledgeMetrics(orgId: string): Promise<Record<string, unknown>> {
        return this.request<Record<string, unknown>>(`/v1/organizations/${orgId}/knowledge/metrics`);
    }

    async initializeCluster(): Promise<ClusterResponse> {
        return this.request<ClusterResponse>('/v1/cluster/initialize', {
            method: 'POST',
        });
    }

    async joinCluster(joinData: JoinClusterRequest): Promise<ClusterResponse> {
        return this.request<ClusterResponse>('/v1/cluster/join', {
            method: 'POST',
            body: JSON.stringify(joinData),
        });
    }

    async leaveCluster(nodeId: string): Promise<boolean> {
        await this.request<void>(`/v1/cluster/nodes/${nodeId}`, { method: 'DELETE' });
        return true;
    }
}
