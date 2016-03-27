package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

func where(w http.ResponseWriter, r *http.Request) {
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	host, _ := net.LookupAddr(ip)
	fmt.Fprintf(w, "IP: %s\nHost: %s", ip, host[0])
}

func main() {
	http.HandleFunc("/", where)
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
