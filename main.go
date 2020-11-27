package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"path"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"

	"github.com/sazzer/fakepi/internal"
)

func main() {
	fmt.Println("FakePI")

	port := flag.Int("port", 8000, "Port to listen on")
	base := flag.String("resources", ".", "Directory containing resources to serve")
	flag.Parse()

	fmt.Printf("Listening on %d\n", *port)
	fmt.Printf("Serving up from: %s\n", *base)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		res, err := internal.NewResource(path.Join(*base, r.RequestURI))
		if err != nil {
			// We failed to load the resource. Try loading a slightly different one to see.
			res, err = internal.NewResource(path.Join(*base, r.RequestURI, "_index"))
		}

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

	err := http.ListenAndServe(fmt.Sprintf(":%d", *port), r)
	if err != nil {
		log.Panic("Failed to start server", err)
	}
}
