package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

func doko(w http.ResponseWriter, r *http.Request) {
	var ip string

	ip = r.Header.Get("X-Forwarded-For")
	if ip == "" {
		var err error
		ip, _, err = net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			log.Println("SplitHostPort error: ", err)
			fmt.Fprint(w, "IP: ???\nHost: ???")
			return
		}
	}
	log.Println("IP: ", ip)

	host, err := net.LookupAddr(ip)
	if err != nil {
		log.Println("LookupAddr error: ", err)
		fmt.Fprintf(w, "IP: %s\nHost: ???", ip)
		return
	}
	log.Println("Host: ", host)

	fmt.Fprintf(w, "IP: %s\nHost: %s", ip, host[0])
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	http.HandleFunc("/", doko)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
