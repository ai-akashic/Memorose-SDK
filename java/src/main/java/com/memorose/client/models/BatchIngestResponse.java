package com.memorose.client.models;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import java.util.List;

@JsonIgnoreProperties(ignoreUnknown = true)
public class BatchIngestResponse {
    public String status;
    public List<String> event_ids;
    public Integer count;
}
