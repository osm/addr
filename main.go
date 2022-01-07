package main

import (
	"flag"
	"fmt"
	"net/http"
	"strings"
)

func main() {
	port := flag.String("p", "4000", "listen port")
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		ip := req.Header.Get("X-Real-IP")
		if ip == "" {
			ip = req.RemoteAddr[0:strings.LastIndex(req.RemoteAddr, ":")]
		}

		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, ip+"\r\n")

	})

	http.ListenAndServe(":"+*port, nil)
}
