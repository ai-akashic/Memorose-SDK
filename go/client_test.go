package memorose

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"
)

type roundTripFunc func(*http.Request) (*http.Response, error)

func (fn roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return fn(req)
}

func newTestClient(handler roundTripFunc) *Client {
	client := NewClient("http://memorose.test", "test_key")
	client.httpClient = &http.Client{Transport: handler}
	return client
}

func jsonResponse(status int, body string) *http.Response {
	return &http.Response{
		StatusCode: status,
		Header:     make(http.Header),
		Body:       io.NopCloser(strings.NewReader(body)),
	}
}

func TestIngestEventUsesCurrentRouteAndAPIKeyHeader(t *testing.T) {
	var gotPath string
	var gotAPIKey string
	client := newTestClient(func(req *http.Request) (*http.Response, error) {
		gotPath = req.URL.Path
		gotAPIKey = req.Header.Get("x-api-key")
		return jsonResponse(http.StatusOK, `{"event_id":"e1","status":"accepted"}`), nil
	})

	resp, err := client.IngestEvent(context.Background(), "user1", "stream1", IngestRequest{
		Content: "user signed up",
	})

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if resp.EventID != "e1" {
		t.Fatalf("expected EventID e1, got %v", resp.EventID)
	}
	if gotPath != "/v1/users/user1/streams/stream1/events" {
		t.Fatalf("expected current stream route, got %s", gotPath)
	}
	if gotAPIKey != "test_key" {
		t.Fatalf("expected x-api-key header, got %q", gotAPIKey)
	}
}

func TestRetrieveMemoryUsesCurrentPayloadShape(t *testing.T) {
	var gotPath string
	var gotPayload RetrieveRequest
	client := newTestClient(func(req *http.Request) (*http.Response, error) {
		gotPath = req.URL.Path
		if err := json.NewDecoder(req.Body).Decode(&gotPayload); err != nil {
			t.Fatalf("decode request body: %v", err)
		}
		return jsonResponse(http.StatusOK, `{"stream_id":"stream1","query":"what happened?","results":[],"query_time_ms":9}`), nil
	})

	resp, err := client.RetrieveMemory(context.Background(), "user1", "stream1", RetrieveRequest{
		Query:      "what happened?",
		Limit:      5,
		GraphDepth: 2,
		OrgID:      strPtr("org-1"),
	})

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if resp.QueryTimeMs != 9 {
		t.Fatalf("expected query_time_ms 9, got %v", resp.QueryTimeMs)
	}
	if gotPath != "/v1/users/user1/streams/stream1/retrieve" {
		t.Fatalf("expected current retrieve route, got %s", gotPath)
	}
	if gotPayload.Limit != 5 || gotPayload.GraphDepth != 2 {
		t.Fatalf("unexpected payload: %+v", gotPayload)
	}
	if gotPayload.OrgID == nil || *gotPayload.OrgID != "org-1" {
		t.Fatalf("expected org_id in payload, got %+v", gotPayload.OrgID)
	}
}

func strPtr(value string) *string {
	return &value
}
