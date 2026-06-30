package main

// backend-go depends on the shared lib/common library (a sibling subdirectory).
// Because that dependency is declared in backend-go's component `paths`, a
// change to lib/common re-evaluates this service in Lunar even though nothing
// under services/backend-go/ changed.

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/pantalasa-cronos/monorepo/lib/common"
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
			Message: common.Greeting(),
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
