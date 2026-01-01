// Test version, not of significant importance
package main

import (
	"fmt"
	"log"

	"github.com/freedom-sketch/sub2go/internal/api"
)

func main() {
	url := "https://example.com"

	encrypted, err := api.EncryptString(url)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(encrypted)
}
