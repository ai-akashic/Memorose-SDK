package com.memorose.client.models;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import java.util.List;

@JsonIgnoreProperties(ignoreUnknown = true)
public class MemoryUnit {
    public String id;
    public String org_id;
    public String user_id;
    public String agent_id;
    public String stream_id;
    public String memory_type;
    public String domain;
    public String namespace_key;
    public String content;
    public List<Double> embedding;
    public Boolean visible;
    public String materialization_state;
    public String materialized_at;
    public List<String> keywords;
    public Double importance;
    public Integer level;
    public String transaction_time;
    public String valid_time;
    public String last_accessed_at;
    public Integer access_count;
    public List<String> references;
    public List<Asset> assets;
    public List<java.util.Map<String, Object>> extracted_facts;
    public java.util.Map<String, Object> share_policy;
    public TaskMetadata task_metadata;
    
    public MemoryUnit() {}
}
