package com.memorose.client.models;

import com.fasterxml.jackson.annotation.JsonInclude;

@JsonInclude(JsonInclude.Include.NON_NULL)
public class UpdateTaskStatusRequest {
    public Object status;
    public Double progress;
    public String result_summary;

    public UpdateTaskStatusRequest() {}
    public UpdateTaskStatusRequest(Object status) { this.status = status; }
}
