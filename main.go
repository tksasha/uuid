package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	fmt.Println( //nolint:forbidigo
		uuid.NewString(),
	)
}
