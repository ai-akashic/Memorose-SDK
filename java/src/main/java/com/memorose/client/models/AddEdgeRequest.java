package com.memorose.client.models;

import com.fasterxml.jackson.annotation.JsonInclude;

@JsonInclude(JsonInclude.Include.NON_NULL)
public class AddEdgeRequest {
    public String source_id;
    public String target_id;
    public String relation;
    public Double weight;

    public AddEdgeRequest() {}
}
