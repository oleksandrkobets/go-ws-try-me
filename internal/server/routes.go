package server

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", s.HelloWorldHandler)
	r.Get("/another_route/", s.AnotherRouteGetHandler)
	r.Post("/another_route/", s.AnotherRoutePostHandler)

	return r
}

func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}

func (s *Server) AnotherRouteGetHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "why are you runnin?"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}

func (s *Server) AnotherRoutePostHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	b, err := io.ReadAll(r.Body)

	var unmarshaledJson map[string]string

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", b)

	var unmarshalingError = json.Unmarshal(b, &unmarshaledJson)
	if unmarshalingError != nil {
		panic(unmarshalingError)
	}

	fmt.Println(unmarshaledJson)

	resp["message"] = "Success!"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}
