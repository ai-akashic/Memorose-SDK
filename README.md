# Memorose SDKs

[![Maven Central](https://img.shields.io/maven-central/v/dev.memorose.sdk/memorose-client?color=blue)](https://central.sonatype.com/artifact/dev.memorose.sdk/memorose-client)
[![PyPI](https://img.shields.io/pypi/v/memorose?color=blue)](https://pypi.org/project/memorose/)
[![Go Reference](https://pkg.go.dev/badge/github.com/memorose/memorose-sdk/go.svg)](https://pkg.go.dev/github.com/memorose/memorose-sdk/go)
[![npm version](https://badge.fury.io/js/memorose-client.svg)](https://badge.fury.io/js/memorose-client)

Official multi-language client libraries for the **Memorose Hybrid AI Memory Storage Engine (CortexDB)**.

These SDKs allow you to effortlessly integrate powerful, multi-tenant memory capabilities (Vector, Text, Graph) into your LLM applications.

## 📦 Supported Languages

We currently provide official SDKs for the following languages:

| Language | Package Manager | Package Name / Dependency | Documentation |
| :--- | :--- | :--- | :--- |
| **Python** | PyPI | `memorose` | [Python Docs](./python/README.md) |
| **Java** | Maven Central | `dev.memorose.sdk:memorose-client` | [Java Docs](./java/README.md) |
| **Go** | Go Modules | `github.com/memorose/memorose-sdk/go` | [Go Docs](./go/README.md) |
| **Node.js** | npm | `memorose-client` | [Node.js Docs](./node/README.md) |

---

## 🚀 Quick Start

All SDKs follow a similar, intuitive API design. The only credential you need to start interacting with your Memorose server is your `api_key`.

### Python
```python
# pip install memorose
from memorose import MemoroseClient

client = MemoroseClient(endpoint="http://localhost:8000", api_key="your_api_key")
memory = client.add_memory(content="The capital of France is Paris.", metadata={"source": "wikipedia"})
```

### Java
```xml
<!-- pom.xml -->
<dependency>
    <groupId>dev.memorose.sdk</groupId>
    <artifactId>memorose-client</artifactId>
    <version>0.1.0</version>
</dependency>
```
```java
MemoroseClient client = new MemoroseClient("http://localhost:8000", "your_api_key");
```

### Go
```bash
go get github.com/memorose/memorose-sdk/go@v0.1.0
```
```go
import "github.com/memorose/memorose-sdk/go"

client := memorose.NewClient("http://localhost:8000", "your_api_key")
```

### Node.js / TypeScript
```typescript
// npm install memorose-client
import { MemoroseClient } from 'memorose-client';

const client = new MemoroseClient('http://localhost:8000', 'your_api_key');
```

---

## 🛠️ Development & Contributing

This repository is structured as a monorepo. Each language SDK has its own directory and isolated build/test pipeline.

*   `java/`: Maven project for the Java SDK.
*   `python/`: Standard Python package (managed via `pyproject.toml`).
*   `go/`: Go module.
*   `node/`: TypeScript/Node.js package.

If you find a bug or want to request a feature, please [open an issue](https://github.com/memorose/memorose-sdk/issues).

## 📄 License

This project is licensed under the Apache License 2.0 - see the [LICENSE](../LICENSE) file for details.
