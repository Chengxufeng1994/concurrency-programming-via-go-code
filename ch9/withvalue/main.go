package main

import (
	"context"
	"fmt"
)

func main() {
	ctx1 := context.WithValue(context.Background(), "key1", "value1")
	ctx2 := context.WithValue(ctx1, "key2", "value2")
	ctx3 := context.WithValue(ctx2, "key3", "value3")
	ctx4 := context.WithValue(ctx3, "key4", "value4")

	fmt.Println(ctx4.Value("key1"))
}
