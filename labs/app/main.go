package main

import (
    "encoding/json"
    "log"
    "net/http"
    "os"
    "time"
)

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    mux := http.NewServeMux()
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        if err := json.NewEncoder(w).Encode(map[string]string{"status": "ok", "course": "ACE"}); err != nil {
            log.Printf("response error: %v", err)
        }
    })
    server := &http.Server{Addr: ":" + port, Handler: mux, ReadHeaderTimeout: 5 * time.Second}
    log.Printf("listening on port %s", port)
    log.Fatal(server.ListenAndServe())
}
