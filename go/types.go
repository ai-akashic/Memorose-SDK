package memorose

type TaskStatus interface{}

type Asset struct {
	StorageKey   string            `json:"storage_key"`
	OriginalName string            `json:"original_name"`
	AssetType    string            `json:"asset_type"`
	Description  *string           `json:"description,omitempty"`
	Metadata     map[string]string `json:"metadata,omitempty"`
}

type TaskMetadata struct {
	Status   TaskStatus `json:"status"`
	Progress float64    `json:"progress"`
}

type MemoryUnit struct {
	ID                   string                   `json:"id"`
	OrgID                *string                  `json:"org_id,omitempty"`
	UserID               string                   `json:"user_id"`
	AgentID              *string                  `json:"agent_id,omitempty"`
	StreamID             string                   `json:"stream_id"`
	MemoryType           string                   `json:"memory_type"`
	Domain               *string                  `json:"domain,omitempty"`
	NamespaceKey         *string                  `json:"namespace_key,omitempty"`
	SharePolicy          map[string]interface{}   `json:"share_policy,omitempty"`
	Content              string                   `json:"content"`
	Embedding            []float64                `json:"embedding,omitempty"`
	Visible              *bool                    `json:"visible,omitempty"`
	MaterializationState *string                  `json:"materialization_state,omitempty"`
	MaterializedAt       *string                  `json:"materialized_at,omitempty"`
	Keywords             []string                 `json:"keywords"`
	Importance           float64                  `json:"importance"`
	Level                int                      `json:"level"`
	TransactionTime      string                   `json:"transaction_time"`
	ValidTime            *string                  `json:"valid_time,omitempty"`
	LastAccessedAt       string                   `json:"last_accessed_at"`
	AccessCount          int                      `json:"access_count"`
	References           []string                 `json:"references"`
	Assets               []Asset                  `json:"assets"`
	ExtractedFacts       []map[string]interface{} `json:"extracted_facts,omitempty"`
	TaskMetadata         *TaskMetadata            `json:"task_metadata,omitempty"`
}

type L3Task struct {
	TaskID        string     `json:"task_id"`
	OrgID         *string    `json:"org_id,omitempty"`
	UserID        string     `json:"user_id"`
	AgentID       *string    `json:"agent_id,omitempty"`
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
	Goal  map[string]interface{} `json:"goal"`
	Tasks []L3TaskTree           `json:"tasks"`
}

type IngestRequest struct {
	Content      string   `json:"content"`
	ContentType  string   `json:"content_type,omitempty"`
	OrgID        *string  `json:"org_id,omitempty"`
	Level        *int     `json:"level,omitempty"`
	ParentID     *string  `json:"parent_id,omitempty"`
	TaskStatus   *string  `json:"task_status,omitempty"`
	TaskProgress *float64 `json:"task_progress,omitempty"`
}

type BatchIngestRequest struct {
	Events []IngestRequest `json:"events"`
}

type IngestResponse struct {
	Status  string `json:"status"`
	EventID string `json:"event_id"`
}

type BatchIngestResponse struct {
	Status   string   `json:"status"`
	EventIDs []string `json:"event_ids"`
	Count    int      `json:"count"`
}

type RetrieveRequest struct {
	Query             string   `json:"query"`
	Limit             int      `json:"limit,omitempty"`
	EnableArbitration bool     `json:"enable_arbitration,omitempty"`
	MinScore          *float64 `json:"min_score,omitempty"`
	TokenBudget       *int     `json:"token_budget,omitempty"`
	GraphDepth        int      `json:"graph_depth,omitempty"`
	StartTime         *string  `json:"start_time,omitempty"`
	EndTime           *string  `json:"end_time,omitempty"`
	AsOf              *string  `json:"as_of,omitempty"`
	OrgID             *string  `json:"org_id,omitempty"`
	AgentID           *string  `json:"agent_id,omitempty"`
	Image             *string  `json:"image,omitempty"`
	Audio             *string  `json:"audio,omitempty"`
	Video             *string  `json:"video,omitempty"`
}

type RetrieveResponse struct {
	StreamID    string                   `json:"stream_id"`
	Query       string                   `json:"query"`
	Results     []map[string]interface{} `json:"results"`
	QueryTimeMs int64                    `json:"query_time_ms"`
}

type MemoryContextRequest struct {
	UserID            string   `json:"user_id"`
	Query             string   `json:"query"`
	Limit             int      `json:"limit,omitempty"`
	EnableArbitration bool     `json:"enable_arbitration,omitempty"`
	MinScore          *float64 `json:"min_score,omitempty"`
	TokenBudget       *int     `json:"token_budget,omitempty"`
	GraphDepth        int      `json:"graph_depth,omitempty"`
	StartTime         *string  `json:"start_time,omitempty"`
	EndTime           *string  `json:"end_time,omitempty"`
	AsOf              *string  `json:"as_of,omitempty"`
	OrgID             *string  `json:"org_id,omitempty"`
	AgentID           *string  `json:"agent_id,omitempty"`
	Format            *string  `json:"format,omitempty"`
	Image             *string  `json:"image,omitempty"`
	Audio             *string  `json:"audio,omitempty"`
	Video             *string  `json:"video,omitempty"`
}

type MemoryContextResponse struct {
	Query             string                   `json:"query"`
	Format            string                   `json:"format"`
	Strategy          string                   `json:"strategy"`
	TokenBudget       int                      `json:"token_budget"`
	UsedTokenEstimate int                      `json:"used_token_estimate"`
	MatchedCount      int                      `json:"matched_count"`
	IncludedCount     int                      `json:"included_count"`
	Truncated         bool                     `json:"truncated"`
	Context           string                   `json:"context"`
	Hits              []map[string]interface{} `json:"hits"`
	QueryTimeMs       int64                    `json:"query_time_ms"`
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

type SemanticMemoryPreviewRequest struct {
	Instruction string  `json:"instruction"`
	OrgID       *string `json:"org_id,omitempty"`
	Mode        *string `json:"mode,omitempty"`
	ForgetMode  string  `json:"forget_mode,omitempty"`
	Limit       int     `json:"limit,omitempty"`
}

type SemanticMemoryExecuteRequest struct {
	PlanID   string  `json:"plan_id"`
	OrgID    *string `json:"org_id,omitempty"`
	Confirm  bool    `json:"confirm"`
	Reviewer *string `json:"reviewer,omitempty"`
	Note     *string `json:"note,omitempty"`
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
	Status  string                   `json:"status"`
	NodeID  *int                     `json:"node_id,omitempty"`
	Message *string                  `json:"message,omitempty"`
	Shards  []map[string]interface{} `json:"shards,omitempty"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
