package com.memorose.client.models;

import com.fasterxml.jackson.annotation.JsonInclude;
import java.util.List;

@JsonInclude(JsonInclude.Include.NON_NULL)
public class BatchIngestRequest {
    public List<IngestRequest> events;

    public BatchIngestRequest() {}
}
