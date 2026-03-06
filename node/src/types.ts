export type TaskStatus =
    | 'Pending'
    | 'InProgress'
    | 'Completed'
    | 'Cancelled'
    | { Blocked: string }
    | { Failed: string };

export interface Asset {
    storage_key: string;
    original_name: string;
    asset_type: string;
    metadata: Record<string, string>;
}

export interface TaskMetadata {
    status: TaskStatus;
    progress: number;
}

export interface MemoryUnit {
    id: string;
    org_id?: string | null;
    user_id: string;
    agent_id?: string | null;
    app_id: string;
    stream_id: string;
    memory_type: 'Factual' | 'Procedural';
    content: string;
    embedding?: number[] | null;
    keywords: string[];
    importance: number;
    level: number;
    transaction_time: string;
    valid_time?: string | null;
    last_accessed_at: string;
    access_count: number;
    references: string[];
    assets: Asset[];
    task_metadata?: TaskMetadata | null;
}

export interface L3Task {
    task_id: string;
    org_id?: string | null;
    user_id: string;
    agent_id?: string | null;
    app_id: string;
    parent_id?: string | null;
    title: string;
    description: string;
    status: TaskStatus;
    progress: number;
    dependencies: string[];
    context_refs: string[];
    created_at: string;
    updated_at: string;
    result_summary?: string | null;
}

export interface L3TaskTree {
    task: L3Task;
    children: L3TaskTree[];
}

export interface GoalTree {
    goal: MemoryUnit;
    tasks: L3TaskTree[];
}

export interface IngestRequest {
    content: string;
    content_type?: string;
    level?: number;
    parent_id?: string;
    task_status?: string;
    task_progress?: number;
}

export interface IngestResponse {
    status: string;
    event_id: string;
}

export interface RetrieveRequest {
    query: string;
    include_vector?: boolean;
    enable_arbitration?: boolean;
    min_score?: number;
    graph_depth?: number;
    start_time?: string;
    end_time?: string;
    as_of?: string;
    agent_id?: string;
}

export interface RetrieveResponse {
    stream_id: string;
    query: string;
    results: any[];
}

export interface UpdateTaskStatusRequest {
    status: TaskStatus;
    progress?: number;
    result_summary?: string;
}

export interface AddEdgeRequest {
    source_id: string;
    target_id: string;
    relation: string;
    weight?: number;
}

export interface JoinClusterRequest {
    node_id: number;
    address?: string;
}

export interface AddEdgeResponse {
    status: string;
    edge_id?: string;
}

export interface PendingCountResponse {
    pending: number;
    ready: boolean;
}

export interface ClusterResponse {
    status: string;
    node_id?: number;
    message?: string;
}
