package main

import (
	"fmt"
	"os"

	"github.com/y-yagi/configure"
	houki "github.com/y-yagi/houki-go"
)

type config struct {
	Directory []string `toml:"Directory"`
}

func main() {
	var houki houki.Houki
	var cfg config

	err := configure.Load("houki", &cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	houki.RemoveDirectories(cfg.Directory)
}
