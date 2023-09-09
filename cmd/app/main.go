package main

import (
	"fmt"
	"os"

	cfg "github.com/maksewsha/wardrobe-back/cfg"
	application "github.com/maksewsha/wardrobe-back/internal/app"
)

func main() {
	cfg, err := cfg.New()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error to create config: %v\n", err)
		os.Exit(1)
	}
	application.Start(cfg)
}
