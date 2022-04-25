package main

import (
	"fmt"
	"log"
)

func main() {
	ans, err := Problem474()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("474: %v\n", ans)
}
