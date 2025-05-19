package api

import (
    "encoding/json"
    "net/http"
    "github.com/andrewwhiteside/continuum-core/internal/memory"
)

var store = memory.NewStore()

type memoryRequest struct {
    Session string `json:"session"`
    Message string `json:"message"`
}

func RegisterHandlers(mux *http.ServeMux) {
    mux.HandleFunc("/memory", handleMemory)
    mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("ok"))
    })
}

func handleMemory(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodPost:
        var req memoryRequest
        if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
            http.Error(w, "bad request", http.StatusBadRequest)
            return
        }
        store.Store(req.Session, req.Message)
        w.WriteHeader(http.StatusNoContent)

    case http.MethodGet:
        session := r.URL.Query().Get("session")
        memories := store.Retrieve(session)
        json.NewEncoder(w).Encode(map[string]interface{}{
            "memories": memories,
        })

    default:
        http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
    }
}