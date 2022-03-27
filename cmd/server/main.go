package main

import (
	"fmt"

	"github.com/DanielMartin96/posts/internal/database"
	"github.com/DanielMartin96/posts/internal/post"
	transportHTTP "github.com/DanielMartin96/posts/internal/transport/http"
)

func Run() error  {
	fmt.Println("Setting Up Our App")

	var err error
	store, err := database.NewDatabase()
	if err != nil {
		return fmt.Errorf("failed to setup connection to the database: %w", err)
	}
	err = store.MigrateDB()
	if err != nil {
		return fmt.Errorf("failed to setup database: %w", err)
	}

	postService := post.NewService(store)
	handler := transportHTTP.NewHandler(postService)
	
	if err := handler.Serve(); err != nil {
		return fmt.Errorf("failed to gracefully serve our application: %w", err)
	}

	return nil
}

func main()  {
	if err := Run(); err != nil {
		fmt.Println("fatal error")
	}
}