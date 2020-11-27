package main

import (
	"fmt"
	"log"
	"net/http"
	"path"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	fmt.Println("FakePI")

	port := 8090
	base := "./resources"

	fmt.Printf("Listening on %d\n", port)
	fmt.Printf("Serving up from: %s\n", base)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		res, err := NewResource(path.Join(base, r.RequestURI))
		if err != nil {
			log.Print("Failed to load resource: ", err)
			w.WriteHeader(404)
			return
		}

		for _, header := range res.Headers {
			w.Header().Add(header.Key, header.Value)
		}
		w.WriteHeader(res.Status)
		_, err = w.Write(res.Body)
		if err != nil {
			log.Print("Failed to write output", err)
		}
	})

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), r)
	if err != nil {
		log.Panic("Failed to start server", err)
	}
}
