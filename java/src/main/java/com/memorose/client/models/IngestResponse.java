package com.memorose.client.models;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;

@JsonIgnoreProperties(ignoreUnknown = true)
public class IngestResponse {
    public String status;
    public String event_id;
}
