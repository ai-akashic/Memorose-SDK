package com.memorose.client.models;

import com.fasterxml.jackson.annotation.JsonInclude;

@JsonInclude(JsonInclude.Include.NON_NULL)
public class SemanticMemoryPreviewRequest {
    public String instruction;
    public String org_id;
    public String mode;
    public String forget_mode;
    public Integer limit;
}
