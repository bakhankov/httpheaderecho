package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", echoHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	headers := map[string]string{}
	for k, v := range r.Header {
		headers[k] = strings.Join(v, ", ")
	}
	headers["Remote-addr"] = r.RemoteAddr
	headersDump, err := json.Marshal(headers)
	if err != nil {
		fmt.Println(err.Error())
	}
	w.Write(headersDump)
}
