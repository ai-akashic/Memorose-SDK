# Memorose Java SDK

Java client for the Memorose Hybrid AI Memory Storage Engine.

## Installation

Add the following to your `pom.xml`:

```xml
<dependency>
    <groupId>dev.memorose.sdk</groupId>
    <artifactId>memorose-client</artifactId>
    <version>0.1.0</version>
</dependency>
```

## Usage

```java
import com.memorose.client.MemoroseClient;
import java.util.Map;
import java.util.HashMap;
import java.util.List;

public class Main {
    public static void main(String[] args) {
        try {
            MemoroseClient client = new MemoroseClient("http://localhost:8000", "your_api_key");

            // Add a memory
            Map<String, Object> metadata = new HashMap<>();
            metadata.put("source", "wikipedia");
            
            Map<String, Object> memory = client.addMemory("The capital of Italy is Rome.", metadata);
            System.out.println("Added memory ID: " + memory.get("id"));

            // Search memories
            List<Map<String, Object>> results = client.searchMemories("What is the capital of Italy?", 5);
            System.out.println("Found " + results.size() + " results.");

            // Delete a memory
            client.deleteMemory((String) memory.get("id"));
            System.out.println("Memory deleted.");

        } catch (Exception e) {
            e.printStackTrace();
        }
    }
}
```
