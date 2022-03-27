package main

import (
	"fmt"
)

func Run() error  {
	fmt.Println("Setting Up Our App")

	var err error

	return nil
}

func main()  {
	if err := Run(); err != nil {
		fmt.Println("fatal error")
	}
}