from dataclasses import asdict, is_dataclass
from typing import Any, Dict, List, Optional

import requests

from .types import (
    AddEdgeRequest,
    BatchIngestRequest,
    BatchIngestResponse,
    GoalTree,
    IngestRequest,
    IngestResponse,
    JoinClusterRequest,
    L3Task,
    MemoryContextRequest,
    MemoryContextResponse,
    RetrieveRequest,
    RetrieveResponse,
    SemanticMemoryExecuteRequest,
    SemanticMemoryPreviewRequest,
    UpdateTaskStatusRequest,
)


class MemoroseAPIError(Exception):
    def __init__(self, status_code: int, message: str, raw_response: str):
        self.status_code = status_code
        self.message = message
        self.raw_response = raw_response
        super().__init__(f"API Error {status_code}: {message}")


class MemoroseClient:
    def __init__(self, endpoint: str, api_key: str, timeout: int = 10):
        self.endpoint = endpoint.rstrip("/")
        self.api_key = api_key
        self.timeout = timeout
        self.session = requests.Session()
        self.session.headers.update(
            {
                "x-api-key": self.api_key,
                "Content-Type": "application/json",
            }
        )

    def _handle_response(self, response: requests.Response) -> Any:
        if response.status_code >= 400:
            error_message = response.text
            try:
                error_data = response.json()
                if "error" in error_data:
                    error_message = error_data["error"]
            except ValueError:
                pass
            raise MemoroseAPIError(response.status_code, error_message, response.text)

        if not response.content:
            return None
        return response.json()

    def _clean_dict(self, data: Any) -> Dict[str, Any]:
        if is_dataclass(data):
            data = asdict(data)
        return {key: value for key, value in data.items() if value is not None}

    def ingest_event(
        self, user_id: str, stream_id: str, event_data: IngestRequest
    ) -> IngestResponse:
        url = f"{self.endpoint}/v1/users/{user_id}/streams/{stream_id}/events"
        response = self.session.post(
            url, json=self._clean_dict(event_data), timeout=self.timeout
        )
        data = self._handle_response(response)
        return IngestResponse(**data) if data else IngestResponse("", "")

    def ingest_events_batch(
        self, user_id: str, stream_id: str, payload: BatchIngestRequest
    ) -> BatchIngestResponse:
        url = f"{self.endpoint}/v1/users/{user_id}/streams/{stream_id}/events/batch"
        response = self.session.post(
            url, json=self._clean_dict(payload), timeout=self.timeout
        )
        data = self._handle_response(response)
        return BatchIngestResponse(**data) if data else BatchIngestResponse("")

    def retrieve_memory(
        self, user_id: str, stream_id: str, query_data: RetrieveRequest
    ) -> RetrieveResponse:
        url = f"{self.endpoint}/v1/users/{user_id}/streams/{stream_id}/retrieve"
        response = self.session.post(
            url, json=self._clean_dict(query_data), timeout=self.timeout
        )
        data = self._handle_response(response)
        return (
            RetrieveResponse(**data)
            if data
            else RetrieveResponse(stream_id=stream_id, query=query_data.query)
        )

    def build_memory_context(
        self, request_data: MemoryContextRequest
    ) -> MemoryContextResponse:
        url = f"{self.endpoint}/v1/memory/context"
        response = self.session.post(
            url, json=self._clean_dict(request_data), timeout=self.timeout
        )
        data = self._handle_response(response)
        return MemoryContextResponse(**data)

    def delete_memory(self, user_id: str, memory_id: str) -> bool:
        url = f"{self.endpoint}/v1/users/{user_id}/memories/{memory_id}"
        response = self.session.delete(url, timeout=self.timeout)
        self._handle_response(response)
        return True

    def preview_semantic_memory(
        self, user_id: str, request_data: SemanticMemoryPreviewRequest
    ) -> Dict[str, Any]:
        url = f"{self.endpoint}/v1/users/{user_id}/memories/semantic/preview"
        response = self.session.post(
            url, json=self._clean_dict(request_data), timeout=self.timeout
        )
        return self._handle_response(response) or {}

    def execute_semantic_memory(
        self, user_id: str, request_data: SemanticMemoryExecuteRequest
    ) -> Dict[str, Any]:
        url = f"{self.endpoint}/v1/users/{user_id}/memories/semantic/execute"
        response = self.session.post(
            url, json=self._clean_dict(request_data), timeout=self.timeout
        )
        return self._handle_response(response) or {}

    def get_task_tree(self, user_id: str, stream_id: str) -> List[GoalTree]:
        url = f"{self.endpoint}/v1/users/{user_id}/streams/{stream_id}/tasks/tree"
        response = self.session.get(url, timeout=self.timeout)
        data = self._handle_response(response)
        return [GoalTree(**item) for item in data] if data else []

    def get_all_task_trees(self, user_id: str) -> List[GoalTree]:
        url = f"{self.endpoint}/v1/users/{user_id}/tasks/tree"
        response = self.session.get(url, timeout=self.timeout)
        data = self._handle_response(response)
        return [GoalTree(**item) for item in data] if data else []

    def get_ready_tasks(self, user_id: str) -> List[L3Task]:
        url = f"{self.endpoint}/v1/users/{user_id}/tasks/ready"
        response = self.session.get(url, timeout=self.timeout)
        data = self._handle_response(response)
        return [L3Task(**item) for item in data] if data else []

    def update_task_status(
        self, user_id: str, task_id: str, status_data: UpdateTaskStatusRequest
    ) -> L3Task:
        url = f"{self.endpoint}/v1/users/{user_id}/tasks/{task_id}/status"
        response = self.session.put(
            url, json=self._clean_dict(status_data), timeout=self.timeout
        )
        data = self._handle_response(response)
        return L3Task(**data) if data else None

    def add_edge(self, user_id: str, edge_data: AddEdgeRequest) -> Dict[str, Any]:
        url = f"{self.endpoint}/v1/users/{user_id}/graph/edges"
        response = self.session.post(
            url, json=self._clean_dict(edge_data), timeout=self.timeout
        )
        return self._handle_response(response) or {}

    def get_pending_count(self) -> Dict[str, Any]:
        url = f"{self.endpoint}/v1/status/pending"
        response = self.session.get(url, timeout=self.timeout)
        return self._handle_response(response) or {}

    def list_organization_knowledge(
        self, org_id: str, query: Optional[Dict[str, Any]] = None
    ) -> Dict[str, Any]:
        url = f"{self.endpoint}/v1/organizations/{org_id}/knowledge"
        response = self.session.get(url, params=query or {}, timeout=self.timeout)
        return self._handle_response(response) or {}

    def get_organization_knowledge(self, org_id: str, knowledge_id: str) -> Dict[str, Any]:
        url = f"{self.endpoint}/v1/organizations/{org_id}/knowledge/{knowledge_id}"
        response = self.session.get(url, timeout=self.timeout)
        return self._handle_response(response) or {}

    def get_organization_knowledge_metrics(self, org_id: str) -> Dict[str, Any]:
        url = f"{self.endpoint}/v1/organizations/{org_id}/knowledge/metrics"
        response = self.session.get(url, timeout=self.timeout)
        return self._handle_response(response) or {}

    def initialize_cluster(self) -> Dict[str, Any]:
        url = f"{self.endpoint}/v1/cluster/initialize"
        response = self.session.post(url, timeout=self.timeout)
        return self._handle_response(response) or {}

    def join_cluster(self, join_data: JoinClusterRequest) -> Dict[str, Any]:
        url = f"{self.endpoint}/v1/cluster/join"
        response = self.session.post(
            url, json=self._clean_dict(join_data), timeout=self.timeout
        )
        return self._handle_response(response) or {}

    def leave_cluster(self, node_id: str) -> bool:
        url = f"{self.endpoint}/v1/cluster/nodes/{node_id}"
        response = self.session.delete(url, timeout=self.timeout)
        self._handle_response(response)
        return True
