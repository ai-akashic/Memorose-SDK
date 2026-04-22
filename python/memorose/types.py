from dataclasses import dataclass, field
from typing import Any, Dict, List, Optional, Union

TaskStatus = Union[str, Dict[str, str]]


@dataclass
class Asset:
    storage_key: str
    original_name: str
    asset_type: str
    description: Optional[str] = None
    metadata: Dict[str, str] = field(default_factory=dict)


@dataclass
class TaskMetadata:
    status: TaskStatus
    progress: float


@dataclass
class MemoryUnit:
    id: str
    user_id: str
    stream_id: str
    memory_type: str
    domain: Optional[str] = None
    namespace_key: Optional[str] = None
    share_policy: Dict[str, Any] = field(default_factory=dict)
    content: str = ""
    keywords: List[str] = field(default_factory=list)
    importance: float = 0.0
    level: int = 0
    transaction_time: str = ""
    last_accessed_at: str = ""
    access_count: int = 0
    references: List[str] = field(default_factory=list)
    assets: List[Asset] = field(default_factory=list)
    org_id: Optional[str] = None
    agent_id: Optional[str] = None
    embedding: Optional[List[float]] = None
    valid_time: Optional[str] = None
    visible: Optional[bool] = None
    materialization_state: Optional[str] = None
    materialized_at: Optional[str] = None
    extracted_facts: List[Dict[str, Any]] = field(default_factory=list)
    task_metadata: Optional[TaskMetadata] = None


@dataclass
class L3Task:
    task_id: str
    user_id: str
    title: str
    description: str
    status: TaskStatus
    progress: float
    dependencies: List[str]
    context_refs: List[str]
    created_at: str
    updated_at: str
    org_id: Optional[str] = None
    agent_id: Optional[str] = None
    parent_id: Optional[str] = None
    result_summary: Optional[str] = None


@dataclass
class L3TaskTree:
    task: L3Task
    children: List["L3TaskTree"] = field(default_factory=list)


@dataclass
class GoalTree:
    goal: MemoryUnit
    tasks: List[L3TaskTree] = field(default_factory=list)


@dataclass
class IngestRequest:
    content: str
    content_type: str = "text"
    org_id: Optional[str] = None
    level: Optional[int] = None
    parent_id: Optional[str] = None
    task_status: Optional[str] = None
    task_progress: Optional[float] = None


@dataclass
class BatchIngestRequest:
    events: List[IngestRequest] = field(default_factory=list)


@dataclass
class IngestResponse:
    status: str
    event_id: str
    write_path: Optional[str] = None


@dataclass
class BatchIngestResponse:
    status: str
    event_ids: List[str] = field(default_factory=list)
    count: int = 0


@dataclass
class RetrieveRequest:
    query: str
    limit: int = 10
    enable_arbitration: Optional[bool] = None
    min_score: Optional[float] = None
    token_budget: Optional[int] = None
    graph_depth: int = 1
    start_time: Optional[str] = None
    end_time: Optional[str] = None
    as_of: Optional[str] = None
    org_id: Optional[str] = None
    agent_id: Optional[str] = None
    image: Optional[str] = None
    audio: Optional[str] = None
    video: Optional[str] = None


@dataclass
class RetrieveResultItem:
    unit: Dict[str, Any]
    score: float


@dataclass
class RetrieveResponse:
    stream_id: str
    query: str
    results: List[Dict[str, Any]] = field(default_factory=list)
    query_time_ms: int = 0


@dataclass
class MemoryContextRequest:
    user_id: str
    query: str
    limit: int = 12
    enable_arbitration: Optional[bool] = None
    min_score: Optional[float] = None
    token_budget: Optional[int] = None
    graph_depth: int = 1
    start_time: Optional[str] = None
    end_time: Optional[str] = None
    as_of: Optional[str] = None
    org_id: Optional[str] = None
    agent_id: Optional[str] = None
    format: Optional[str] = None
    image: Optional[str] = None
    audio: Optional[str] = None
    video: Optional[str] = None


@dataclass
class MemoryContextResponse:
    query: str
    format: str
    strategy: str
    token_budget: int
    used_token_estimate: int
    matched_count: int
    included_count: int
    truncated: bool
    context: str
    hits: List[Dict[str, Any]] = field(default_factory=list)
    query_time_ms: int = 0


@dataclass
class UpdateTaskStatusRequest:
    status: TaskStatus
    progress: Optional[float] = None
    result_summary: Optional[str] = None


@dataclass
class AddEdgeRequest:
    source_id: str
    target_id: str
    relation: str
    weight: Optional[float] = None


@dataclass
class JoinClusterRequest:
    node_id: int
    address: Optional[str] = None


@dataclass
class SemanticMemoryPreviewRequest:
    instruction: str
    org_id: Optional[str] = None
    mode: Optional[str] = None
    forget_mode: str = "logical"
    limit: int = 10


@dataclass
class SemanticMemoryExecuteRequest:
    plan_id: str
    confirm: bool
    org_id: Optional[str] = None
    reviewer: Optional[str] = None
    note: Optional[str] = None
