import sys
import unittest
from types import ModuleType
from pathlib import Path
from unittest.mock import Mock, patch

sys.path.insert(0, str(Path(__file__).resolve().parents[1]))

requests_stub = ModuleType("requests")


class _Session:
    def __init__(self):
        self.headers = {}

    def post(self, *args, **kwargs):
        raise NotImplementedError

    def get(self, *args, **kwargs):
        raise NotImplementedError

    def put(self, *args, **kwargs):
        raise NotImplementedError

    def delete(self, *args, **kwargs):
        raise NotImplementedError

    def update(self, payload):
        self.headers.update(payload)


requests_stub.Session = _Session
requests_stub.Response = object
sys.modules.setdefault("requests", requests_stub)

from memorose.client import MemoroseClient
from memorose.types import IngestRequest, RetrieveRequest


class TestMemoroseClient(unittest.TestCase):
    def setUp(self):
        self.client = MemoroseClient("http://localhost:8000", "test_key")

    @patch("requests.Session.post")
    def test_ingest_event_uses_stream_route_and_api_key_header(self, mock_post):
        mock_response = Mock()
        mock_response.status_code = 200
        mock_response.content = b'{"status":"accepted","event_id":"evt_1","write_path":"local_bypass"}'
        mock_response.json.return_value = {
            "status": "accepted",
            "event_id": "evt_1",
            "write_path": "local_bypass",
        }
        mock_post.return_value = mock_response

        result = self.client.ingest_event(
            "user-1",
            "stream-1",
            IngestRequest(content="hello"),
        )

        self.assertEqual(result.event_id, "evt_1")
        self.assertEqual(result.status, "accepted")
        self.assertEqual(result.write_path, "local_bypass")
        self.assertEqual(
            mock_post.call_args.args[0],
            "http://localhost:8000/v1/users/user-1/streams/stream-1/events",
        )
        self.assertEqual(
            self.client.session.headers["x-api-key"],
            "test_key",
        )
        self.assertNotIn("Authorization", self.client.session.headers)

    @patch("requests.Session.post")
    def test_retrieve_memory_uses_current_request_shape(self, mock_post):
        mock_response = Mock()
        mock_response.status_code = 200
        mock_response.content = (
            b'{"stream_id":"stream-1","query":"hello","results":[],"query_time_ms":12}'
        )
        mock_response.json.return_value = {
            "stream_id": "stream-1",
            "query": "hello",
            "results": [],
            "query_time_ms": 12,
        }
        mock_post.return_value = mock_response

        result = self.client.retrieve_memory(
            "user-1",
            "stream-1",
            RetrieveRequest(query="hello", limit=5, org_id="org-1"),
        )

        self.assertEqual(result.stream_id, "stream-1")
        self.assertEqual(result.query_time_ms, 12)
        self.assertEqual(
            mock_post.call_args.args[0],
            "http://localhost:8000/v1/users/user-1/streams/stream-1/retrieve",
        )
        self.assertEqual(
            mock_post.call_args.kwargs["json"],
            {"query": "hello", "limit": 5, "graph_depth": 1, "org_id": "org-1"},
        )


if __name__ == "__main__":
    unittest.main()
