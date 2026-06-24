package main

// Touches backend-go so the path-filtered code collector fires for this
// sub-component alongside the per-service CI job (LUNAR_COMPONENT attribution test).
// Re-run after reverting the snyk TEMP pin that wedged the cronos hub manifest.

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type response struct {
	Service string `json:"service"`
	Message string `json:"message"`
	Version string `json:"version"`
}

const version = "0.1.1"

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(response{
			Service: "backend-go",
			Message: "hello from pantalasa-cronos monorepo backend-go",
			Version: version,
		})
	})

	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	})

	mux.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]string{"version": version})
	})

	addr := os.Getenv("ADDR")
	if addr == "" {
		addr = ":8080"
	}

	log.Printf("backend-go listening on %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("listen: %v", err)
	}
}
