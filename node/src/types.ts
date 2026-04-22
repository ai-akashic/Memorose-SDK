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
    description?: string | null;
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
    stream_id: string;
    memory_type: 'Factual' | 'Procedural' | string;
    domain?: string | null;
    namespace_key?: string | null;
    share_policy?: Record<string, unknown>;
    content: string;
    embedding?: number[] | null;
    visible?: boolean;
    materialization_state?: string | null;
    materialized_at?: string | null;
    keywords: string[];
    importance: number;
    level: number;
    transaction_time: string;
    valid_time?: string | null;
    last_accessed_at: string;
    access_count: number;
    references: string[];
    assets: Asset[];
    extracted_facts?: Record<string, unknown>[];
    task_metadata?: TaskMetadata | null;
}

export interface L3Task {
    task_id: string;
    org_id?: string | null;
    user_id: string;
    agent_id?: string | null;
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
    goal: Record<string, unknown>;
    tasks: L3TaskTree[];
}

export interface IngestRequest {
    content: string;
    content_type?: string;
    org_id?: string;
    level?: number;
    parent_id?: string;
    task_status?: string;
    task_progress?: number;
}

export interface BatchIngestRequest {
    events: IngestRequest[];
}

export interface IngestResponse {
    status: string;
    event_id: string;
}

export interface BatchIngestResponse {
    status: string;
    event_ids: string[];
    count: number;
}

export interface RetrieveRequest {
    query: string;
    limit?: number;
    enable_arbitration?: boolean;
    min_score?: number;
    token_budget?: number;
    graph_depth?: number;
    start_time?: string;
    end_time?: string;
    as_of?: string;
    org_id?: string;
    agent_id?: string;
    image?: string;
    audio?: string;
    video?: string;
}

export interface RetrieveResponse {
    stream_id: string;
    query: string;
    results: Array<Record<string, unknown>>;
    query_time_ms: number;
}

export interface MemoryContextRequest {
    user_id: string;
    query: string;
    limit?: number;
    enable_arbitration?: boolean;
    min_score?: number;
    token_budget?: number;
    graph_depth?: number;
    start_time?: string;
    end_time?: string;
    as_of?: string;
    org_id?: string;
    agent_id?: string;
    format?: string;
    image?: string;
    audio?: string;
    video?: string;
}

export interface MemoryContextResponse {
    query: string;
    format: string;
    strategy: string;
    token_budget: number;
    used_token_estimate: number;
    matched_count: number;
    included_count: number;
    truncated: boolean;
    context: string;
    hits: Array<Record<string, unknown>>;
    query_time_ms: number;
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

export interface SemanticMemoryPreviewRequest {
    instruction: string;
    org_id?: string;
    mode?: string;
    forget_mode?: string;
    limit?: number;
}

export interface SemanticMemoryExecuteRequest {
    plan_id: string;
    org_id?: string;
    confirm: boolean;
    reviewer?: string;
    note?: string;
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
    shards?: Array<Record<string, unknown>>;
}
