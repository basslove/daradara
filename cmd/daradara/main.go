package main

import (
	"context"
	server "github.com/basslove/daradara/internal/api/infrastructure/api"
	"github.com/basslove/daradara/internal/api/logger"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	ctx := logger.NewContext(context.Background(), logger.BaseLogger)
	server.Run(ctx)
}
