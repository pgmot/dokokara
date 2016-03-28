package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

func where(w http.ResponseWriter, r *http.Request) {
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	host, _ := net.LookupAddr(ip)
	fmt.Fprintf(w, "IP: %s\nHost: %s", ip, host[0])
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	http.HandleFunc("/", where)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
