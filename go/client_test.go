package memorose

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIngestEvent(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"event_id":"e1","status":"success"}`))
	}))
	defer server.Close()

	client := NewClient(server.URL, "test_key")
	req := IngestRequest{
		Content: "user signed up",
	}
	resp, err := client.IngestEvent(context.Background(), "user1", "app1", "stream1", req)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if resp.EventID != "e1" {
		t.Errorf("expected EventID e1, got %v", resp.EventID)
	}
	if resp.Status != "success" {
		t.Errorf("expected status success, got %v", resp.Status)
	}
}

func TestRetrieveMemory(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"stream_id":"stream1","query":"what happened?","results":[{"id":"m1","content":"user signed up"}]}`))
	}))
	defer server.Close()

	client := NewClient(server.URL, "test_key")
	req := RetrieveRequest{
		Query: "what happened?",
	}
	resp, err := client.RetrieveMemory(context.Background(), "user1", "app1", "stream1", req)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if resp.StreamID != "stream1" {
		t.Errorf("expected StreamID stream1, got %v", resp.StreamID)
	}
	if resp.Query != "what happened?" {
		t.Errorf("expected query 'what happened?', got %v", resp.Query)
	}
	if len(resp.Results) != 1 {
		t.Fatalf("expected 1 result, got %v", len(resp.Results))
	}
	if resp.Results[0]["id"] != "m1" {
		t.Errorf("expected memory ID m1, got %v", resp.Results[0]["id"])
	}
}
