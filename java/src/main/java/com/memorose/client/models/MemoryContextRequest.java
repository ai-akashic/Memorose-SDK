package com.memorose.client.models;

import com.fasterxml.jackson.annotation.JsonInclude;

@JsonInclude(JsonInclude.Include.NON_NULL)
public class MemoryContextRequest {
    public String user_id;
    public String query;
    public Integer limit;
    public Boolean enable_arbitration;
    public Double min_score;
    public Integer token_budget;
    public Integer graph_depth;
    public String start_time;
    public String end_time;
    public String as_of;
    public String org_id;
    public String agent_id;
    public String format;
    public String image;
    public String audio;
    public String video;
}
