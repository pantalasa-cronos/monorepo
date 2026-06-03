package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type response struct {
	Service string `json:"service"`
	Message string `json:"message"`
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(response{
			Service: "backend-go",
			Message: "hello from pantalasa-cronos monorepo backend-go",
		})
	})

	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
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

// QA monorepo path-filter probe S1 — touches services/backend-go only (2026-06-03)

// QA S2 probe — backend-go leg of multi-component PR
