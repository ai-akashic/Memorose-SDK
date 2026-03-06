package com.memorose.client.models;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import java.util.List;
import java.util.Map;

@JsonIgnoreProperties(ignoreUnknown = true)
public class RetrieveResponse {
    public String stream_id;
    public String query;
    public List<Map<String, Object>> results;
}
