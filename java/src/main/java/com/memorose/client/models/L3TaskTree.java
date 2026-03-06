package com.memorose.client.models;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import java.util.List;

@JsonIgnoreProperties(ignoreUnknown = true)
public class L3TaskTree {
    public L3Task task;
    public List<L3TaskTree> children;

    public L3TaskTree() {}
}
