package memorose

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type APIError struct {
	StatusCode int
	Message    string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("API Error %d: %s", e.StatusCode, e.Message)
}

type Client struct {
	endpoint   string
	apiKey     string
	httpClient *http.Client
}

func NewClient(endpoint, apiKey string) *Client {
	return NewClientWithTimeout(endpoint, apiKey, 10*time.Second)
}

func NewClientWithTimeout(endpoint, apiKey string, timeout time.Duration) *Client {
	return &Client{
		endpoint:   strings.TrimRight(endpoint, "/"),
		apiKey:     apiKey,
		httpClient: &http.Client{Timeout: timeout},
	}
}

func (c *Client) do(ctx context.Context, method, path string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, method, c.endpoint+path, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("Content-Type", "application/json")
	return c.httpClient.Do(req)
}

func (c *Client) requestJSON(ctx context.Context, method, path string, reqBody interface{}, res interface{}) error {
	var bodyReader io.Reader
	if reqBody != nil {
		data, err := json.Marshal(reqBody)
		if err != nil {
			return err
		}
		bodyReader = bytes.NewReader(data)
	}

	resp, err := c.do(ctx, method, path, bodyReader)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		bodyBytes, _ := io.ReadAll(resp.Body)
		var errRes ErrorResponse
		if err := json.Unmarshal(bodyBytes, &errRes); err == nil && errRes.Error != "" {
			return &APIError{StatusCode: resp.StatusCode, Message: errRes.Error}
		}
		return &APIError{StatusCode: resp.StatusCode, Message: string(bodyBytes)}
	}

	if res != nil {
		if err := json.NewDecoder(resp.Body).Decode(res); err != nil && err != io.EOF {
			return err
		}
	}
	return nil
}

func (c *Client) IngestEvent(ctx context.Context, userID, appID, streamID string, req IngestRequest) (*IngestResponse, error) {
	path := fmt.Sprintf("/v1/users/%s/apps/%s/streams/%s/events", userID, appID, streamID)
	var res IngestResponse
	err := c.requestJSON(ctx, http.MethodPost, path, req, &res)
	return &res, err
}

func (c *Client) RetrieveMemory(ctx context.Context, userID, appID, streamID string, req RetrieveRequest) (*RetrieveResponse, error) {
	path := fmt.Sprintf("/v1/users/%s/apps/%s/streams/%s/retrieve", userID, appID, streamID)
	var res RetrieveResponse
	err := c.requestJSON(ctx, http.MethodPost, path, req, &res)
	return &res, err
}

func (c *Client) GetTaskTree(ctx context.Context, userID, appID, streamID string) ([]GoalTree, error) {
	path := fmt.Sprintf("/v1/users/%s/apps/%s/streams/%s/tasks/tree", userID, appID, streamID)
	var res []GoalTree
	err := c.requestJSON(ctx, http.MethodGet, path, nil, &res)
	return res, err
}

func (c *Client) GetAllTaskTrees(ctx context.Context, userID string) ([]GoalTree, error) {
	path := fmt.Sprintf("/v1/users/%s/tasks/tree", userID)
	var res []GoalTree
	err := c.requestJSON(ctx, http.MethodGet, path, nil, &res)
	return res, err
}

func (c *Client) GetReadyTasks(ctx context.Context, userID string) ([]L3Task, error) {
	path := fmt.Sprintf("/v1/users/%s/tasks/ready", userID)
	var res []L3Task
	err := c.requestJSON(ctx, http.MethodGet, path, nil, &res)
	return res, err
}

func (c *Client) UpdateTaskStatus(ctx context.Context, userID, taskID string, req UpdateTaskStatusRequest) (*L3Task, error) {
	path := fmt.Sprintf("/v1/users/%s/tasks/%s/status", userID, taskID)
	var res L3Task
	err := c.requestJSON(ctx, http.MethodPut, path, req, &res)
	return &res, err
}

func (c *Client) AddEdge(ctx context.Context, userID string, req AddEdgeRequest) (*AddEdgeResponse, error) {
	path := fmt.Sprintf("/v1/users/%s/graph/edges", userID)
	var res AddEdgeResponse
	err := c.requestJSON(ctx, http.MethodPost, path, req, &res)
	return &res, err
}

func (c *Client) GetPendingCount(ctx context.Context) (*PendingCountResponse, error) {
	var res PendingCountResponse
	err := c.requestJSON(ctx, http.MethodGet, "/v1/status/pending", nil, &res)
	return &res, err
}

func (c *Client) InitializeCluster(ctx context.Context) (*ClusterResponse, error) {
	var res ClusterResponse
	err := c.requestJSON(ctx, http.MethodPost, "/v1/cluster/initialize", nil, &res)
	return &res, err
}

func (c *Client) JoinCluster(ctx context.Context, req JoinClusterRequest) (*ClusterResponse, error) {
	var res ClusterResponse
	err := c.requestJSON(ctx, http.MethodPost, "/v1/cluster/join", req, &res)
	return &res, err
}

func (c *Client) LeaveCluster(ctx context.Context, nodeID string) error {
	path := fmt.Sprintf("/v1/cluster/nodes/%s", nodeID)
	return c.requestJSON(ctx, http.MethodDelete, path, nil, nil)
}
