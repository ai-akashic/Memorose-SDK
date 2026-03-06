import requests
from typing import List, Dict, Any, Optional
from dataclasses import asdict

from .types import (
    IngestRequest, IngestResponse,
    RetrieveRequest, RetrieveResponse,
    UpdateTaskStatusRequest, AddEdgeRequest,
    JoinClusterRequest,
    GoalTree, L3Task
)

class MemoroseAPIError(Exception):
    """Exception raised when the Memorose API returns an error response."""
    def __init__(self, status_code: int, message: str, raw_response: str):
        self.status_code = status_code
        self.message = message
        self.raw_response = raw_response
        super().__init__(f"API Error {status_code}: {message}")

class MemoroseClient:
    """
    Python client for Memorose hybrid AI memory storage engine.
    """
    def __init__(self, endpoint: str, api_key: str, timeout: int = 10):
        self.endpoint = endpoint.rstrip('/')
        self.api_key = api_key
        self.timeout = timeout
        self.session = requests.Session()
        self.session.headers.update({
            "Authorization": f"Bearer {self.api_key}",
            "Content-Type": "application/json"
        })

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
        """Remove None values from dataclass dictionary."""
        return {k: v for k, v in asdict(data).items() if v is not None}

    def ingest_event(self, user_id: str, app_id: str, stream_id: str, event_data: IngestRequest) -> IngestResponse:
        """Ingest an event stream to form memory context."""
        url = f"{self.endpoint}/v1/users/{user_id}/apps/{app_id}/streams/{stream_id}/events"
        response = self.session.post(url, json=self._clean_dict(event_data), timeout=self.timeout)
        data = self._handle_response(response)
        return IngestResponse(**data) if data else IngestResponse("", "")

    def retrieve_memory(self, user_id: str, app_id: str, stream_id: str, query_data: RetrieveRequest) -> RetrieveResponse:
        """Retrieve contextual memory."""
        url = f"{self.endpoint}/v1/users/{user_id}/apps/{app_id}/streams/{stream_id}/retrieve"
        response = self.session.post(url, json=self._clean_dict(query_data), timeout=self.timeout)
        data = self._handle_response(response)
        return RetrieveResponse(**data) if data else RetrieveResponse(stream_id, query_data.query, [])

    def get_task_tree(self, user_id: str, app_id: str, stream_id: str) -> List[GoalTree]:
        """Get the task tree associated with a specific stream."""
        url = f"{self.endpoint}/v1/users/{user_id}/apps/{app_id}/streams/{stream_id}/tasks/tree"
        response = self.session.get(url, timeout=self.timeout)
        data = self._handle_response(response)
        return [GoalTree(**item) for item in data] if data else []

    def get_all_task_trees(self, user_id: str) -> List[GoalTree]:
        """Get all task trees for a specific user."""
        url = f"{self.endpoint}/v1/users/{user_id}/tasks/tree"
        response = self.session.get(url, timeout=self.timeout)
        data = self._handle_response(response)
        return [GoalTree(**item) for item in data] if data else []

    def get_ready_tasks(self, user_id: str) -> List[L3Task]:
        """Get tasks that are ready to be executed for a specific user."""
        url = f"{self.endpoint}/v1/users/{user_id}/tasks/ready"
        response = self.session.get(url, timeout=self.timeout)
        data = self._handle_response(response)
        return [L3Task(**item) for item in data] if data else []

    def update_task_status(self, user_id: str, task_id: str, status_data: UpdateTaskStatusRequest) -> L3Task:
        """Update the status of a specific task."""
        url = f"{self.endpoint}/v1/users/{user_id}/tasks/{task_id}/status"
        response = self.session.put(url, json=self._clean_dict(status_data), timeout=self.timeout)
        data = self._handle_response(response)
        return L3Task(**data) if data else None

    def add_edge(self, user_id: str, edge_data: AddEdgeRequest) -> Dict[str, Any]:
        """Add a graph edge explicitly for a user."""
        url = f"{self.endpoint}/v1/users/{user_id}/graph/edges"
        response = self.session.post(url, json=self._clean_dict(edge_data), timeout=self.timeout)
        return self._handle_response(response) or {}

    def get_pending_count(self) -> Dict[str, Any]:
        """Get the pending background tasks status."""
        url = f"{self.endpoint}/v1/status/pending"
        response = self.session.get(url, timeout=self.timeout)
        return self._handle_response(response) or {}

    def initialize_cluster(self) -> Dict[str, Any]:
        """Initialize the Memorose distributed cluster."""
        url = f"{self.endpoint}/v1/cluster/initialize"
        response = self.session.post(url, timeout=self.timeout)
        return self._handle_response(response) or {}

    def join_cluster(self, join_data: JoinClusterRequest) -> Dict[str, Any]:
        """Join an existing Memorose distributed cluster."""
        url = f"{self.endpoint}/v1/cluster/join"
        response = self.session.post(url, json=self._clean_dict(join_data), timeout=self.timeout)
        return self._handle_response(response) or {}

    def leave_cluster(self, node_id: str) -> bool:
        """Remove a node from the cluster."""
        url = f"{self.endpoint}/v1/cluster/nodes/{node_id}"
        response = self.session.delete(url, timeout=self.timeout)
        self._handle_response(response)
        return True
