package main

import (
	"encoding/json"
	"log/syslog"
	"net/http"
	"strings"
)

var logger *syslog.Writer
var err error

func main() {
	logger, err = syslog.New(syslog.LOG_WARNING, "httpheaderecho")
	if err != nil {
		panic("Cannot attach to syslog")
	}
	http.HandleFunc("/", echoHandler)
	logger.Crit(http.ListenAndServe(":8080", nil).Error())
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	headers := map[string]string{}
	for k, v := range r.Header {
		headers[k] = strings.Join(v, ", ")
	}
	headers["Remote-Addr"] = strings.Split(r.RemoteAddr, ":")[0]
	headersDump, err := json.Marshal(headers)
	if err != nil {
		logger.Err(err.Error())
	}
	w.Write(headersDump)
}
