package com.memorose.client.models;

import com.fasterxml.jackson.annotation.JsonInclude;

@JsonInclude(JsonInclude.Include.NON_NULL)
public class IngestRequest {
    public String content;
    public String content_type = "text";
    public Integer level;
    public String parent_id;
    public String task_status;
    public Double task_progress;

    public IngestRequest() {}
    public IngestRequest(String content) { this.content = content; }
}
