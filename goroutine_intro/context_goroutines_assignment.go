package main

import (
	"fmt"
	"context"
	"time"
)

func main() {
	ctx, err := context.WithDeadline(context.Background(), time.Now().Add(time.Second*2))
	defer err()

	go func() {
		fmt.Print("Task started...")
		time.Sleep(time.Second*5)
		fmt.Println("go routine survived")
	}()

	select {
	case <- ctx.Done():
		fmt.Print(ctx.Err())
	}
}
