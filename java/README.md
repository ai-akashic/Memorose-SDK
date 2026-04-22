# Memorose Java SDK

Java client for the current Memorose `/v1` runtime.

## Installation

```xml
<dependency>
    <groupId>dev.memorose.sdk</groupId>
    <artifactId>memorose-client</artifactId>
    <version>0.1.1</version>
</dependency>
```

## Usage

```java
import com.memorose.client.MemoroseClient;
import com.memorose.client.models.IngestRequest;
import com.memorose.client.models.RetrieveRequest;
import com.memorose.client.models.RetrieveResponse;

public class Main {
    public static void main(String[] args) throws Exception {
        MemoroseClient client = new MemoroseClient("http://127.0.0.1:3000", "your_api_key");
        String streamId = "11111111-1111-1111-1111-111111111111";

        client.ingestEvent("user_123", streamId, new IngestRequest("Dylan prefers concise summaries."));

        RetrieveRequest request = new RetrieveRequest("What does Dylan prefer?");
        request.limit = 5;
        RetrieveResponse response = client.retrieveMemory("user_123", streamId, request);

        System.out.println(response.results);
    }
}
```

The SDK sends `x-api-key` and mirrors the server `/v1` REST routes.
