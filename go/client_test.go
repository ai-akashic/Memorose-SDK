package memorose

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAddMemory(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"id":"1","content":"test memory","metadata":{}}`))
	}))
	defer server.Close()

	client := NewClient(server.URL, "test_key")
	mem, err := client.AddMemory(context.Background(), "test memory", map[string]interface{}{"key": "value"})

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if mem.ID != "1" {
		t.Errorf("expected ID 1, got %v", mem.ID)
	}
	if mem.Content != "test memory" {
		t.Errorf("expected content 'test memory', got %v", mem.Content)
	}
}

func TestSearchMemories(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"id":"1","content":"test memory","metadata":{}}]`))
	}))
	defer server.Close()

	client := NewClient(server.URL, "test_key")
	mems, err := client.SearchMemories(context.Background(), "test query", 5)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(mems) != 1 {
		t.Fatalf("expected 1 result, got %d", len(mems))
	}
	if mems[0].ID != "1" {
		t.Errorf("expected ID 1, got %v", mems[0].ID)
	}
}

func TestDeleteMemory(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client := NewClient(server.URL, "test_key")
	err := client.DeleteMemory(context.Background(), "1")

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}
