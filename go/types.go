package memorose

type TaskStatus interface{} // can be string or map[string]string

type Asset struct {
	StorageKey   string            `json:"storage_key"`
	OriginalName string            `json:"original_name"`
	AssetType    string            `json:"asset_type"`
	Metadata     map[string]string `json:"metadata,omitempty"`
}

type TaskMetadata struct {
	Status   TaskStatus `json:"status"`
	Progress float64    `json:"progress"`
}

type MemoryUnit struct {
	ID              string        `json:"id"`
	OrgID           *string       `json:"org_id,omitempty"`
	UserID          string        `json:"user_id"`
	AgentID         *string       `json:"agent_id,omitempty"`
	AppID           string        `json:"app_id"`
	StreamID        string        `json:"stream_id"`
	MemoryType      string        `json:"memory_type"`
	Content         string        `json:"content"`
	Embedding       []float64     `json:"embedding,omitempty"`
	Keywords        []string      `json:"keywords"`
	Importance      float64       `json:"importance"`
	Level           int           `json:"level"`
	TransactionTime string        `json:"transaction_time"`
	ValidTime       *string       `json:"valid_time,omitempty"`
	LastAccessedAt  string        `json:"last_accessed_at"`
	AccessCount     int           `json:"access_count"`
	References      []string      `json:"references"`
	Assets          []Asset       `json:"assets"`
	TaskMetadata    *TaskMetadata `json:"task_metadata,omitempty"`
}

type L3Task struct {
	TaskID        string     `json:"task_id"`
	OrgID         *string    `json:"org_id,omitempty"`
	UserID        string     `json:"user_id"`
	AgentID       *string    `json:"agent_id,omitempty"`
	AppID         string     `json:"app_id"`
	ParentID      *string    `json:"parent_id,omitempty"`
	Title         string     `json:"title"`
	Description   string     `json:"description"`
	Status        TaskStatus `json:"status"`
	Progress      float64    `json:"progress"`
	Dependencies  []string   `json:"dependencies"`
	ContextRefs   []string   `json:"context_refs"`
	CreatedAt     string     `json:"created_at"`
	UpdatedAt     string     `json:"updated_at"`
	ResultSummary *string    `json:"result_summary,omitempty"`
}

type L3TaskTree struct {
	Task     L3Task       `json:"task"`
	Children []L3TaskTree `json:"children"`
}

type GoalTree struct {
	Goal  MemoryUnit   `json:"goal"`
	Tasks []L3TaskTree `json:"tasks"`
}

type IngestRequest struct {
	Content      string   `json:"content"`
	ContentType  string   `json:"content_type,omitempty"`
	Level        *int     `json:"level,omitempty"`
	ParentID     *string  `json:"parent_id,omitempty"`
	TaskStatus   *string  `json:"task_status,omitempty"`
	TaskProgress *float64 `json:"task_progress,omitempty"`
}

type IngestResponse struct {
	Status  string `json:"status"`
	EventID string `json:"event_id"`
}

type RetrieveRequest struct {
	Query             string   `json:"query"`
	IncludeVector     bool     `json:"include_vector,omitempty"`
	EnableArbitration bool     `json:"enable_arbitration,omitempty"`
	MinScore          *float64 `json:"min_score,omitempty"`
	GraphDepth        int      `json:"graph_depth,omitempty"`
	StartTime         *string  `json:"start_time,omitempty"`
	EndTime           *string  `json:"end_time,omitempty"`
	AsOf              *string  `json:"as_of,omitempty"`
	AgentID           *string  `json:"agent_id,omitempty"`
}

type RetrieveResponse struct {
	StreamID string                   `json:"stream_id"`
	Query    string                   `json:"query"`
	Results  []map[string]interface{} `json:"results"`
}

type UpdateTaskStatusRequest struct {
	Status        TaskStatus `json:"status"`
	Progress      *float64   `json:"progress,omitempty"`
	ResultSummary *string    `json:"result_summary,omitempty"`
}

type AddEdgeRequest struct {
	SourceID string   `json:"source_id"`
	TargetID string   `json:"target_id"`
	Relation string   `json:"relation"`
	Weight   *float64 `json:"weight,omitempty"`
}

type JoinClusterRequest struct {
	NodeID  int     `json:"node_id"`
	Address *string `json:"address,omitempty"`
}

type AddEdgeResponse struct {
	Status string  `json:"status"`
	EdgeID *string `json:"edge_id,omitempty"`
}

type PendingCountResponse struct {
	Pending int  `json:"pending"`
	Ready   bool `json:"ready"`
}

type ClusterResponse struct {
	Status  string  `json:"status"`
	NodeID  *int    `json:"node_id,omitempty"`
	Message *string `json:"message,omitempty"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

