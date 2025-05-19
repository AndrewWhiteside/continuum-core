# 🪐 CONTINUUM v0.1

> _"Memory is the fuel of intelligence."_

Continuum is a lightweight, extensible memory API for AI agents and Copilot extensions, designed to serve as a memory layer for MCP-based tools. This v0.1 release includes an in-memory store to help you quickly bootstrap memory features without the overhead of a full vector database.

---

## 🚀 Features

- 🧠 In-memory memory store (session-based)
- 📥 `POST /memory` to store new messages
- 📤 `GET /memory?session=...` to retrieve session memory
- 🩺 `/healthz` endpoint
- 🐳 Dockerized
- ☁️ Kubernetes-ready manifests
- 🔧 Easy to plug into GitHub Copilot Agent (via MCP tool)

---

## 🌌 Architecture

```
┌────────────────────────────┐
│     GitHub Copilot        │
│     (MCP Agent)           │
└────────────┬──────────────┘
             │
     Tool call via HTTP
             │
┌────────────▼──────────────┐
│     CONTINUUM API         │
│  (store & retrieve logic)│
└────────────┬──────────────┘
             │
      In-Memory Store 🧠
             │
     (Qdrant plug-in coming soon)
```

---

## 🛠️ Endpoints

### `POST /memory`

Stores a memory for a session.

#### Example Request

```json
{
  "session": "caepe-abc123",
  "message": "Deployed staging version of frontend"
}
```

---

### `GET /memory?session=caepe-abc123`

Retrieves all messages for the session.

#### Example Response

```json
{
  "memories": [
    "Deployed staging version of frontend",
    "Rolled back to previous release"
  ]
}
```

---

## 🐳 Docker

Build and run locally:

```bash
docker build -t continuum .
docker run -p 8080:8080 continuum
```

---

## ☸️ Kubernetes

Deploy to your cluster:

```bash
kubectl apply -f k8s/deployment.yaml
kubectl apply -f k8s/service.yaml
kubectl apply -f k8s/ingress.yaml
```

---

## 📦 Project Structure

```bash
continuum-core/
├── Dockerfile
├── go.mod
├── main.go
├── internal/
│   └── memory/       # In-memory store logic
├── pkg/
│   └── api/          # HTTP handlers
└── k8s/              # Kubernetes manifests
```

---

## 🧪 Coming Soon

- 🧲 Qdrant-backed vector store
- 🔍 Semantic memory retrieval
- 🔐 Auth and session TTL
- 💾 Redis / Disk-based fallback

---

## 🛸 License

MIT. Free to use, extend, and enhance. Keep it weird. Build smart memory.

---

Built with ❤️ by [Andrew Whiteside](https://github.com/AndrewWhiteside)