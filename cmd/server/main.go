package main

import (
	"fmt"

	"github.com/DanielMartin96/posts/internal/database"
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
	return nil
}

func main()  {
	if err := Run(); err != nil {
		fmt.Println("fatal error")
	}
}