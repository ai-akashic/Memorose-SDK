package com.memorose.client.models;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import java.util.List;

@JsonIgnoreProperties(ignoreUnknown = true)
public class L3Task {
    public String task_id;
    public String org_id;
    public String user_id;
    public String agent_id;
    public String app_id;
    public String parent_id;
    public String title;
    public String description;
    public Object status;
    public Double progress;
    public List<String> dependencies;
    public List<String> context_refs;
    public String created_at;
    public String updated_at;
    public String result_summary;

    public L3Task() {}
}
