package main

import (
	"context"

	"github.com/anderson89marques/bank/internal/adapter/rest"
)

func main() {
	rest.Run(context.Background())
}
