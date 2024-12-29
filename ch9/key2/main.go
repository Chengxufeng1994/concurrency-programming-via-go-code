package main

import (
	"context"
	"fmt"

	"github.com/Chengxufeng1994/concurrency-programming-via-go-code/ch9/key2/p1"
	"github.com/Chengxufeng1994/concurrency-programming-via-go-code/ch9/key2/p2"
)

func main() {
	ctx := context.Background()
	ctx = p1.NewContext(ctx, "123")
	ctx = p2.NewContext(ctx, true)
	fmt.Println(p1.GetContext(ctx))
	fmt.Println(p2.GetContext(ctx))
}
