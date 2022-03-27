package http

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

type Handler struct {
	Router *mux.Router
	Service PostService
	Server *http.Server
}

type Response struct {
	Message string `json:"message"`
}

func NewHandler(service PostService) *Handler {
	h := &Handler{
		Service: service,
	}

	h.Router = mux.NewRouter()

	h.MapRoutes()

	h.Server = &http.Server{
		Addr: "0.0.0.0:8080",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      h.Router,
	}

	return h
}

func (h *Handler) MapRoutes() {
	h.Router.HandleFunc("/api/v1/post/{id}", h.GetPost).Methods("GET")
	h.Router.HandleFunc("/api/v1/post", h.CreatePost).Methods("POST")	
	h.Router.HandleFunc("/api/v1/post/{id}", h.UpdatePost).Methods("PUT")
	h.Router.HandleFunc("/api/v1/post/{id}", h.DeletePost).Methods("DELETE")
}

func (h *Handler) Serve() error {
	go func() {
		if err := h.Server.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	h.Server.Shutdown(ctx)

	log.Println("shutting down gracefully")
	return nil
}