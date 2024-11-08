package main

import (
	"DatabaseCourse/internal/postgres"
	"context"
)

func main() {
	_, err := postgres.NewPgsStorage(context.TODO())
	if err != nil {
		panic(err)
	}
}
