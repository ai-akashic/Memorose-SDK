package com.memorose.client.models;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import java.util.List;

@JsonIgnoreProperties(ignoreUnknown = true)
public class GoalTree {
    public MemoryUnit goal;
    public List<L3TaskTree> tasks;

    public GoalTree() {}
}
