package com.memorose.client.models;

import com.fasterxml.jackson.annotation.JsonInclude;

@JsonInclude(JsonInclude.Include.NON_NULL)
public class TaskMetadata {
    public Object status;
    public Double progress;

    public TaskMetadata() {}
}
