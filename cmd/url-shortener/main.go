package main

import (
	"fmt"
	"github.com/dinerozz/url_shortener/internal/config"
)

func main() {
	cfg := config.MustLoad()

	fmt.Println(cfg)

	// TODO: init logger: slog

	// TODO: init storage: sqlite3

	// TODO: init router: chi, "chi render"

	// TODO: run server
}
