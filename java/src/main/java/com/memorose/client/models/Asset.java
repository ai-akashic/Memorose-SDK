package com.memorose.client.models;

import com.fasterxml.jackson.annotation.JsonInclude;
import java.util.Map;

@JsonInclude(JsonInclude.Include.NON_NULL)
public class Asset {
    public String storage_key;
    public String original_name;
    public String asset_type;
    public Map<String, String> metadata;

    public Asset() {}
}
