package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	srv := &http.Server{
		Addr: ":80",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer r.Body.Close()
			reqBody, err := ioutil.ReadAll(r.Body)
			if err != nil {
				w.WriteHeader(400)
			}
			if _, err := w.Write(reqBody); err != nil {
				w.WriteHeader(500)
			}
		}),
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
