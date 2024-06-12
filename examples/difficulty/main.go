package main

import (
	"fmt"

	"github.com/EIYARO-Project/core-sdk/api"
	"github.com/EIYARO-Project/core-sdk/client"
)

func main() {
	client := client.NewStdClient("http://localhost:9888", "")
	api := api.NewApi(client)
	difficulty, err := api.Difficulty(0, "")
	if err != nil {
		fmt.Printf("Error getting Difficulty: %s", err)
	} else {
		fmt.Printf("Difficulty: %s\n", difficulty.StringIndent())
	}
}
