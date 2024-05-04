package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /ping", pong)

	addr := ":8080"
	fmt.Println("Http Server running on adrr", addr)
	http.ListenAndServe(addr, mux)
}

func pong(w http.ResponseWriter, r *http.Request) {
	fmt.Println("pong_at", time.Now().Unix())
	w.Write([]byte("poooooooooooong"))
}
