package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {
	port := flag.String("p", "4000", "listen port")
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, req.RemoteAddr+"\r\n")

	})

	http.ListenAndServe(":"+*port, nil)
}
