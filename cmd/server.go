package main

import (
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
	log.Println("start")
	d := func(request *http.Request) {
		url := *request.URL
		url.Scheme = "https"
		url.Host = "slack.com"

		req, err := http.NewRequest(request.Method, url.String(), request.Body)
		if err != nil {
			log.Fatal(err.Error())
		}
		req.Header = request.Header
		*request = *req
	}
	rp := &httputil.ReverseProxy{Director: d}
	s := http.Server{
		Addr:    ":8080",
		Handler: rp,
	}
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err.Error())
	}
}
