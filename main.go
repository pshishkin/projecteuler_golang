package main

import (
	"fmt"
	"log"
)

func main() {
	ans, err := Problem510()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("510: %v\n", ans)
}
