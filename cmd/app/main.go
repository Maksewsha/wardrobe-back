package main

import (
	cfg "github.com/maksewsha/wardrobe-back/cfg"
	application "github.com/maksewsha/wardrobe-back/internal/app"
)

func main() {
	cfg := cfg.New()
	application.Start(cfg)
}
