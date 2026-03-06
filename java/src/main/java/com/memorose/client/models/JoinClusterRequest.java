package com.memorose.client.models;

import com.fasterxml.jackson.annotation.JsonInclude;

@JsonInclude(JsonInclude.Include.NON_NULL)
public class JoinClusterRequest {
    public int node_id;
    public String address;

    public JoinClusterRequest() {}
    public JoinClusterRequest(int node_id) { this.node_id = node_id; }
}
