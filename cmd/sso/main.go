package main

import (
	"example.com/UserSSO/internal/config"
	"fmt"
)

func main() {
	cfg := config.MustLoad()
	fmt.Printf("%+v\n", cfg)
}
