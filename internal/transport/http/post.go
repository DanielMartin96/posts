package http

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/DanielMartin96/posts/internal/post"
	"github.com/gorilla/mux"
)

type PostService interface {
	GetPost(ctx context.Context, ID string) (post.Post, error)
	CreatePost(ctx context.Context, pst post.Post) (post.Post, error)
	UpdatePost(ctx context.Context, ID string, pst post.Post) (post.Post, error)
	DeletePost(ctx context.Context, ID string) error
}

func (h *Handler) GetPost(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	pst, err := h.Service.GetPost(r.Context(), id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(pst); err != nil {
		panic(err)
	}
}

type PostPostRequest struct {
	Slug   string `json:"slug" validate:"required"`
	Author string `json:"author" validate:"required"`
	Body   string `json:"body" validate:"required"`
}

func postFromPostPostRequest(u PostPostRequest) post.Post {
	return post.Post{
		Slug:   u.Slug,
		Author: u.Author,
		Body:   u.Body,
	}
}

func (h *Handler) CreatePost(w http.ResponseWriter, r *http.Request) {
	var postPstReq PostPostRequest
	if err := json.NewDecoder(r.Body).Decode(&postPstReq); err != nil {
		return
	}

	pst := postFromPostPostRequest(postPstReq)
	pst, err := h.Service.CreatePost(r.Context(), pst)
	if err != nil {
		fmt.Printf("failed creating post: %w", err)
		return
	}
	if err := json.NewEncoder(w).Encode(pst); err != nil {
		panic(err)
	}
}

type UpdatePostRequest struct {
	Slug   string `json:"slug" validate:"required"`
	Author string `json:"author" validate:"required"`
	Body   string `json:"body" validate:"required"`
}

func postFromUpdatePostRequest(u UpdatePostRequest) post.Post {
	return post.Post{
		Slug:   u.Slug,
		Author: u.Author,
		Body:   u.Body,
	}
}

func (h *Handler) UpdatePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postID := vars["id"]

	var updatePstRequest UpdatePostRequest
	if err := json.NewDecoder(r.Body).Decode(&updatePstRequest); err != nil {
		return
	}

	pst := postFromUpdatePostRequest(updatePstRequest)

	pst, err := h.Service.UpdatePost(r.Context(), postID, pst)
	if err != nil {
		fmt.Println("error updating post")
		return
	}
	if err := json.NewEncoder(w).Encode(pst); err != nil {
		panic(err)
	}
}

func (h *Handler) DeletePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postID := vars["id"]

	if postID == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := h.Service.DeletePost(r.Context(), postID)
	if err != nil {
		return
	}

	if err := json.NewEncoder(w).Encode(Response{Message: "Successfully Deleted"}); err != nil {
		panic(err)
	}
}