package main

import (
	"second-spot-backend/internal/bootstrap"
)

func main() {
	if err := bootstrap.Start(); err != nil {
		panic(err)
	}
}
