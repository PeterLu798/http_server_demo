package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/golang/glog"
)

func main() {
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/healthz", healthz)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func healthz(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "200")
	w.Header().Set("Message", "ok")
	uri := r.URL.String()
	method := r.Method
	ip := r.Header.Get("X-Real-Ip")
	if ip == "" {
		ip = r.Header.Get("X-Forwarded-For")
		if ip == "" {
			ip = r.RemoteAddr
		}
	}
	glog.V(2).Info("Uri: %s  Method: %s Client IP %s", uri, method, ip)
	fmt.Printf("Uri: %s  Method: %s Client IP %s", uri, method, ip)
}

func headers(res http.ResponseWriter, req *http.Request) {
	header := req.Header
	fmt.Println(header)
	for k, vals := range header {
		res.Header().Set(k, vals[0])
	}
	VERSION := os.Getenv("VERSION")
	if VERSION != "" {
		res.Header().Set("VERSION", VERSION)
	} else {
		res.Header().Set("VERSION", "No Found")
	}
}
