package main

import (
	"fmt"
	"os"

	"github.com/suzupala/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}

	err = cfg.SetUser("John")
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(2)
	}

	cfg, err = config.Read()
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}

	fmt.Printf("%+v", cfg)
	os.Exit(0)
}




