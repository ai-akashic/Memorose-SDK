package com.memorose.client.models;

import com.fasterxml.jackson.annotation.JsonInclude;

@JsonInclude(JsonInclude.Include.NON_NULL)
public class RetrieveRequest {
    public String query;
    public Boolean include_vector;
    public Boolean enable_arbitration;
    public Double min_score;
    public Integer graph_depth;
    public String start_time;
    public String end_time;
    public String as_of;
    public String agent_id;

    public RetrieveRequest() {}
    public RetrieveRequest(String query) { this.query = query; }
}
