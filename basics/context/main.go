package main

import (
	"context"
	"fmt"
	"time"
)

func doingSomething(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("Exiting")
		return
	default:
		time.Sleep(10 * time.Second)
	}

}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	go doingSomething(ctx)

	time.Sleep(3 * time.Second)
}
