package com.memorose.client.models;

import com.fasterxml.jackson.annotation.JsonInclude;

@JsonInclude(JsonInclude.Include.NON_NULL)
public class SemanticMemoryExecuteRequest {
    public String plan_id;
    public String org_id;
    public boolean confirm;
    public String reviewer;
    public String note;
}
