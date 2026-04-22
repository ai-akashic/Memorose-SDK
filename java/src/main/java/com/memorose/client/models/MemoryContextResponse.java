package com.memorose.client.models;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import java.util.List;
import java.util.Map;

@JsonIgnoreProperties(ignoreUnknown = true)
public class MemoryContextResponse {
    public String query;
    public String format;
    public String strategy;
    public Integer token_budget;
    public Integer used_token_estimate;
    public Integer matched_count;
    public Integer included_count;
    public Boolean truncated;
    public String context;
    public List<Map<String, Object>> hits;
    public Long query_time_ms;
}
