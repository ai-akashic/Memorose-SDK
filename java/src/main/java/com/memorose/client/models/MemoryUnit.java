package com.memorose.client.models;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import java.util.List;

@JsonIgnoreProperties(ignoreUnknown = true)
public class MemoryUnit {
    public String id;
    public String org_id;
    public String user_id;
    public String agent_id;
    public String app_id;
    public String stream_id;
    public String memory_type;
    public String content;
    public List<Double> embedding;
    public List<String> keywords;
    public Double importance;
    public Integer level;
    public String transaction_time;
    public String valid_time;
    public String last_accessed_at;
    public Integer access_count;
    public List<String> references;
    public List<Asset> assets;
    public TaskMetadata task_metadata;
    
    public MemoryUnit() {}
}
