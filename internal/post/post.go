package post

import (
	"context"
	"fmt"
)

type Post struct {
	ID string
	Slug string
	Author string
	Body string
}

type PostStore interface {
	GetPost(context.Context, string) (Post, error)
	CreatePost(context.Context, Post) (Post, error)
	UpdatePost(context.Context, string, Post) (Post, error)
	DeletePost(context.Context, string) error
}

type Service struct { 
	Store PostStore
}

func NewService(store PostStore) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service) GetPost(ctx context.Context, id string) (Post, error)  {
	pst, err := s.Store.GetPost(ctx, id)
	if err != nil {
		fmt.Println("error fetching post with id")
		return Post{}, err
	}

	return pst, nil
}

func (s *Service) CreatePost(ctx context.Context, post Post) (Post, error)  {
	pst, err := s.Store.CreatePost(ctx, post)
	if err != nil {
		fmt.Println("could not create post")
		return Post{}, err
	}

	return pst, nil
}

func (s *Service) UpdatePost(ctx context.Context, id string, post Post) (Post, error)  {
	pst, err := s.Store.UpdatePost(ctx, id, post)
	if err != nil {
		fmt.Println("could not update post")
		return Post{}, err
	}

	return pst, nil
}

func (s *Service) DeletePost(ctx context.Context, id string) error  {
	return s.Store.DeletePost(ctx, id)
}