import unittest
from unittest.mock import patch, Mock
from memorose.client import MemoroseClient

class TestMemoroseClient(unittest.TestCase):
    def setUp(self):
        self.client = MemoroseClient("http://localhost:8000", "test_key")

    @patch('requests.Session.post')
    def test_add_memory(self, mock_post):
        mock_response = Mock()
        mock_response.json.return_value = {"id": "1", "content": "test", "metadata": {}}
        mock_response.raise_for_status.return_value = None
        mock_post.return_value = mock_response

        result = self.client.add_memory("test", {"key": "value"})
        self.assertEqual(result["id"], "1")
        self.assertEqual(result["content"], "test")

    @patch('requests.Session.get')
    def test_search_memories(self, mock_get):
        mock_response = Mock()
        mock_response.json.return_value = [{"id": "1", "content": "test", "metadata": {}}]
        mock_response.raise_for_status.return_value = None
        mock_get.return_value = mock_response

        result = self.client.search_memories("test query", limit=5)
        self.assertEqual(len(result), 1)
        self.assertEqual(result[0]["id"], "1")

    @patch('requests.Session.delete')
    def test_delete_memory(self, mock_delete):
        mock_response = Mock()
        mock_response.raise_for_status.return_value = None
        mock_delete.return_value = mock_response

        result = self.client.delete_memory("1")
        self.assertTrue(result)

if __name__ == '__main__':
    unittest.main()
