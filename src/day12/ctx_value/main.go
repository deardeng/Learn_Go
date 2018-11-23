package main

import (
	"context"
	"fmt"
)

func process(ctx context.Context) {
	ret, ok := ctx.Value("trace_id").(int)
	if !ok {
		ret = 8888
	}
	fmt.Printf("ret : %d\n", ret)

	s, _ := ctx.Value("session").(string)
	fmt.Printf("session:%s\n", s)
}

func main() {
	ctx := context.WithValue(context.Background(), "trace_id", 1234)
	ctx = context.WithValue(ctx, "session", "abcde")
	process(ctx)
}
