package main

import (
	"DatabaseCourse/internal/task6"
	"context"
)

func main() {
	t6s, err := task6.NewTask6Storage(context.TODO())
	if err != nil {
		panic(err)
	}
	menu := task6.NewTask6Menu(t6s)
	menu.Serve()
}
