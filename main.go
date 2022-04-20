package main

import (
	"fmt"
	"log"
)

func main() {
	ans, err := Problem206()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("206: %v\n", ans)
}
