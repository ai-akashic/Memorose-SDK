package com.memorose.client;

import com.fasterxml.jackson.databind.ObjectMapper;
import java.net.URI;
import java.net.URLEncoder;
import java.nio.charset.StandardCharsets;
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;
import java.time.Duration;
import java.util.Map;
import java.util.List;
import com.fasterxml.jackson.core.type.TypeReference;
import com.fasterxml.jackson.databind.JsonNode;

import com.memorose.client.models.*;
import com.memorose.client.exceptions.MemoroseAPIException;

public class MemoroseClient {
    private final String endpoint;
    private final String apiKey;
    private final HttpClient httpClient;
    private final ObjectMapper objectMapper;

    public MemoroseClient(String endpoint, String apiKey) {
        this.endpoint = endpoint.replaceAll("/$", "");
        this.apiKey = apiKey;
        this.httpClient = HttpClient.newBuilder()
                .connectTimeout(Duration.ofSeconds(10))
                .build();
        this.objectMapper = new ObjectMapper();
    }

    private String enc(String value) {
        if (value == null) return "";
        return URLEncoder.encode(value, StandardCharsets.UTF_8).replace("+", "%20");
    }

    private void handleHttpError(HttpResponse<String> response) throws MemoroseAPIException {
        if (response.statusCode() >= 400) {
            String errorMsg = response.body();
            try {
                JsonNode jsonNode = objectMapper.readTree(errorMsg);
                if (jsonNode.has("error")) {
                    errorMsg = jsonNode.get("error").asText();
                }
            } catch (Exception ignored) {
            }
            throw new MemoroseAPIException(response.statusCode(), errorMsg, response.body());
        }
    }

    private <T> T executeRequest(HttpRequest request, Class<T> responseType) throws Exception {
        HttpResponse<String> response = httpClient.send(request, HttpResponse.BodyHandlers.ofString());
        handleHttpError(response);
        
        String body = response.body();
        if (body == null || body.trim().isEmpty()) {
            return null;
        }
        return objectMapper.readValue(body, responseType);
    }

    private <T> T executeRequestRef(HttpRequest request, TypeReference<T> typeRef) throws Exception {
        HttpResponse<String> response = httpClient.send(request, HttpResponse.BodyHandlers.ofString());
        handleHttpError(response);
        
        String body = response.body();
        if (body == null || body.trim().isEmpty()) {
            return null;
        }
        return objectMapper.readValue(body, typeRef);
    }

    public IngestResponse ingestEvent(String userId, String streamId, IngestRequest requestData) throws Exception {
        String json = objectMapper.writeValueAsString(requestData);
        HttpRequest request = HttpRequest.newBuilder()
                .uri(URI.create(String.format("%s/v1/users/%s/streams/%s/events", this.endpoint, enc(userId), enc(streamId))))
                .header("x-api-key", this.apiKey)
                .header("Content-Type", "application/json")
                .timeout(Duration.ofSeconds(30))
                .POST(HttpRequest.BodyPublishers.ofString(json))
                .build();
        return executeRequest(request, IngestResponse.class);
    }

    public BatchIngestResponse ingestEventsBatch(String userId, String streamId, BatchIngestRequest requestData) throws Exception {
        String json = objectMapper.writeValueAsString(requestData);
        HttpRequest request = HttpRequest.newBuilder()
                .uri(URI.create(String.format("%s/v1/users/%s/streams/%s/events/batch", this.endpoint, enc(userId), enc(streamId))))
                .header("x-api-key", this.apiKey)
                .header("Content-Type", "application/json")
                .timeout(Duration.ofSeconds(30))
                .POST(HttpRequest.BodyPublishers.ofString(json))
                .build();
        return executeRequest(request, BatchIngestResponse.class);
    }

    public RetrieveResponse retrieveMemory(String userId, String streamId, RetrieveRequest requestData) throws Exception {
        String json = objectMapper.writeValueAsString(requestData);
        HttpRequest request = HttpRequest.newBuilder()
                .uri(URI.create(String.format("%s/v1/users/%s/streams/%s/retrieve", this.endpoint, enc(userId), enc(streamId))))
                .header("x-api-key", this.apiKey)
                .header("Content-Type", "application/json")
                .timeout(Duration.ofSeconds(30))
                .POST(HttpRequest.BodyPublishers.ofString(json))
                .build();
        return executeRequest(request, RetrieveResponse.class);
    }

    public MemoryContextResponse buildMemoryContext(MemoryContextRequest requestData) throws Exception {
        String json = objectMapper.writeValueAsString(requestData);
        HttpRequest request = HttpRequest.newBuilder()
                .uri(URI.create(this.endpoint + "/v1/memory/context"))
                .header("x-api-key", this.apiKey)
                .header("Content-Type", "application/json")
                .timeout(Duration.ofSeconds(30))
                .POST(HttpRequest.BodyPublishers.ofString(json))
                .build();
        return executeRequest(request, MemoryContextResponse.class);
    }

    public void deleteMemory(String userId, String memoryId) throws Exception {
        HttpRequest request = HttpRequest.newBuilder()
                .uri(URI.create(String.format("%s/v1/users/%s/memories/%s", this.endpoint, enc(userId), enc(memoryId))))
                .header("x-api-key", this.apiKey)
                .timeout(Duration.ofSeconds(15))
                .DELETE()
                .build();
        HttpResponse<String> response = httpClient.send(request, HttpResponse.BodyHandlers.ofString());
        handleHttpError(response);
    }

    public Map<String, Object> previewSemanticMemory(String userId, SemanticMemoryPreviewRequest requestData) throws Exception {
        String json = objectMapper.writeValueAsString(requestData);
        HttpRequest request = HttpRequest.newBuilder()
                .uri(URI.create(String.format("%s/v1/users/%s/memories/semantic/preview", this.endpoint, enc(userId))))
                .header("x-api-key", this.apiKey)
                .header("Content-Type", "application/json")
                .timeout(Duration.ofSeconds(30))
                .POST(HttpRequest.BodyPublishers.ofString(json))
                .build();
        return executeRequestRef(request, new TypeReference<Map<String, Object>>(){});
    }

    public Map<String, Object> executeSemanticMemory(String userId, SemanticMemoryExecuteRequest requestData) throws Exception {
        String json = objectMapper.writeValueAsString(requestData);
        HttpRequest request = HttpRequest.newBuilder()
                .uri(URI.create(String.format("%s/v1/users/%s/memories/semantic/execute", this.endpoint, enc(userId))))
                .header("x-api-key", this.apiKey)
                .header("Content-Type", "application/json")
                .timeout(Duration.ofSeconds(30))
                .POST(HttpRequest.BodyPublishers.ofString(json))
                .build();
        return executeRequestRef(request, new TypeReference<Map<String, Object>>(){});
    }

    public List<GoalTree> getTaskTree(String userId, String streamId) throws Exception {
        HttpRequest request = HttpRequest.newBuilder()
                .uri(URI.create(String.format("%s/v1/users/%s/streams/%s/tasks/tree", this.endpoint, enc(userId), enc(streamId))))
                .header("x-api-key", this.apiKey)
                .timeout(Duration.ofSeconds(10))
                .GET()
                .build();
        return executeRequestRef(request, new TypeReference<List<GoalTree>>(){});
    }

    public List<GoalTree> getAllTaskTrees(String userId) throws Exception {
        HttpRequest request = HttpRequest.newBuilder()
                .uri(URI.create(String.format("%s/v1/users/%s/tasks/tree", this.endpoint, enc(userId))))
                .header("x-api-key", this.apiKey)
                .timeout(Duration.ofSeconds(10))
                .GET()
                .build();
        return executeRequestRef(request, new TypeReference<List<GoalTree>>(){});
    }

    public List<L3Task> getReadyTasks(String userId) throws Exception {
        HttpRequest request = HttpRequest.newBuilder()
                .uri(URI.create(String.format("%s/v1/users/%s/tasks/ready", this.endpoint, enc(userId))))
                .header("x-api-key", this.apiKey)
                .timeout(Duration.ofSeconds(10))
                .GET()
                .build();
        return executeRequestRef(request, new TypeReference<List<L3Task>>(){});
    }

    public L3Task updateTaskStatus(String userId, String taskId, UpdateTaskStatusRequest requestData) throws Exception {
        String json = objectMapper.writeValueAsString(requestData);
        HttpRequest request = HttpRequest.newBuilder()
                .uri(URI.create(String.format("%s/v1/users/%s/tasks/%s/status", this.endpoint, enc(userId), enc(taskId))))
                .header("x-api-key", this.apiKey)
                .header("Content-Type", "application/json")
                .timeout(Duration.ofSeconds(15))
                .PUT(HttpRequest.BodyPublishers.ofString(json))
                .build();
        return executeRequest(request, L3Task.class);
    }

    public Map<String, Object> addEdge(String userId, AddEdgeRequest requestData) throws Exception {
        String json = objectMapper.writeValueAsString(requestData);
        HttpRequest request = HttpRequest.newBuilder()
                .uri(URI.create(String.format("%s/v1/users/%s/graph/edges", this.endpoint, enc(userId))))
                .header("x-api-key", this.apiKey)
                .header("Content-Type", "application/json")
                .timeout(Duration.ofSeconds(10))
                .POST(HttpRequest.BodyPublishers.ofString(json))
                .build();
        return executeRequestRef(request, new TypeReference<Map<String, Object>>(){});
    }

    public Map<String, Object> getPendingCount() throws Exception {
        HttpRequest request = HttpRequest.newBuilder()
                .uri(URI.create(this.endpoint + "/v1/status/pending"))
                .header("x-api-key", this.apiKey)
                .timeout(Duration.ofSeconds(5))
                .GET()
                .build();
        return executeRequestRef(request, new TypeReference<Map<String, Object>>(){});
    }

    public Map<String, Object> listOrganizationKnowledge(String orgId, Map<String, String> query) throws Exception {
        StringBuilder uri = new StringBuilder(String.format("%s/v1/organizations/%s/knowledge", this.endpoint, enc(orgId)));
        if (query != null && !query.isEmpty()) {
            uri.append("?");
            boolean first = true;
            for (Map.Entry<String, String> entry : query.entrySet()) {
                if (!first) uri.append("&");
                uri.append(enc(entry.getKey())).append("=").append(enc(entry.getValue()));
                first = false;
            }
        }
        HttpRequest request = HttpRequest.newBuilder()
                .uri(URI.create(uri.toString()))
                .header("x-api-key", this.apiKey)
                .timeout(Duration.ofSeconds(10))
                .GET()
                .build();
        return executeRequestRef(request, new TypeReference<Map<String, Object>>(){});
    }

    public Map<String, Object> getOrganizationKnowledge(String orgId, String knowledgeId) throws Exception {
        HttpRequest request = HttpRequest.newBuilder()
                .uri(URI.create(String.format("%s/v1/organizations/%s/knowledge/%s", this.endpoint, enc(orgId), enc(knowledgeId))))
                .header("x-api-key", this.apiKey)
                .timeout(Duration.ofSeconds(10))
                .GET()
                .build();
        return executeRequestRef(request, new TypeReference<Map<String, Object>>(){});
    }

    public Map<String, Object> getOrganizationKnowledgeMetrics(String orgId) throws Exception {
        HttpRequest request = HttpRequest.newBuilder()
                .uri(URI.create(String.format("%s/v1/organizations/%s/knowledge/metrics", this.endpoint, enc(orgId))))
                .header("x-api-key", this.apiKey)
                .timeout(Duration.ofSeconds(10))
                .GET()
                .build();
        return executeRequestRef(request, new TypeReference<Map<String, Object>>(){});
    }

    public Map<String, Object> initializeCluster() throws Exception {
        HttpRequest request = HttpRequest.newBuilder()
                .uri(URI.create(this.endpoint + "/v1/cluster/initialize"))
                .header("x-api-key", this.apiKey)
                .timeout(Duration.ofSeconds(15))
                .POST(HttpRequest.BodyPublishers.noBody())
                .build();
        return executeRequestRef(request, new TypeReference<Map<String, Object>>(){});
    }

    public Map<String, Object> joinCluster(JoinClusterRequest requestData) throws Exception {
        String json = objectMapper.writeValueAsString(requestData);
        HttpRequest request = HttpRequest.newBuilder()
                .uri(URI.create(this.endpoint + "/v1/cluster/join"))
                .header("x-api-key", this.apiKey)
                .header("Content-Type", "application/json")
                .timeout(Duration.ofSeconds(15))
                .POST(HttpRequest.BodyPublishers.ofString(json))
                .build();
        return executeRequestRef(request, new TypeReference<Map<String, Object>>(){});
    }

    public void leaveCluster(String nodeId) throws Exception {
        HttpRequest request = HttpRequest.newBuilder()
                .uri(URI.create(String.format("%s/v1/cluster/nodes/%s", this.endpoint, enc(nodeId))))
                .header("x-api-key", this.apiKey)
                .timeout(Duration.ofSeconds(15))
                .DELETE()
                .build();
        
        HttpResponse<String> response = httpClient.send(request, HttpResponse.BodyHandlers.ofString());
        handleHttpError(response);
    }
}
