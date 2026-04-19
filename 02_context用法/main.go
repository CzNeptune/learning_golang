package main

import (
	"context"
	"fmt"
)

func main() {

	ctx := context.WithValue(context.Background(), "trace_id", "123456hello")
	trace_id := ctx.Value("trace_id")
	fmt.Printf("trace_id=%v\n", trace_id)

}
