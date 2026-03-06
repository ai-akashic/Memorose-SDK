from dataclasses import dataclass, field
from typing import Optional, List, Dict, Any, Union

TaskStatus = Union[str, Dict[str, str]]

@dataclass
class Asset:
    storage_key: str
    original_name: str
    asset_type: str
    metadata: Dict[str, str] = field(default_factory=dict)

@dataclass
class TaskMetadata:
    status: TaskStatus
    progress: float

@dataclass
class MemoryUnit:
    id: str
    user_id: str
    app_id: str
    stream_id: str
    memory_type: str
    content: str
    keywords: List[str]
    importance: float
    level: int
    transaction_time: str
    last_accessed_at: str
    access_count: int
    references: List[str]
    assets: List[Asset] = field(default_factory=list)
    org_id: Optional[str] = None
    agent_id: Optional[str] = None
    embedding: Optional[List[float]] = None
    valid_time: Optional[str] = None
    task_metadata: Optional[TaskMetadata] = None

@dataclass
class L3Task:
    task_id: str
    user_id: str
    app_id: str
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
    children: List['L3TaskTree'] = field(default_factory=list)

@dataclass
class GoalTree:
    goal: MemoryUnit
    tasks: List[L3TaskTree] = field(default_factory=list)


@dataclass
class IngestRequest:
    content: str
    content_type: str = "text"
    level: Optional[int] = None
    parent_id: Optional[str] = None
    task_status: Optional[str] = None
    task_progress: Optional[float] = None

@dataclass
class IngestResponse:
    status: str
    event_id: str

@dataclass
class RetrieveRequest:
    query: str
    include_vector: bool = False
    enable_arbitration: bool = False
    min_score: Optional[float] = None
    graph_depth: int = 1
    start_time: Optional[str] = None
    end_time: Optional[str] = None
    as_of: Optional[str] = None
    agent_id: Optional[str] = None

@dataclass
class RetrieveResponse:
    stream_id: str
    query: str
    results: List[Dict[str, Any]] = field(default_factory=list)

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
