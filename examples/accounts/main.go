package main

import (
	"fmt"
	"os"

	"github.com/EIYARO-Project/core-sdk/api"
	"github.com/EIYARO-Project/core-sdk/api/resources"
	"github.com/EIYARO-Project/core-sdk/client"
)

func main() {
	client := client.NewStdClient("http://localhost:9888", "")
	api := api.NewApi(client)
	accountResource, err := api.Resource("Account")
	if err != nil {
		fmt.Printf("Error getting resource Account: %s\n", err)
		os.Exit(1)
	}

	list, err := accountResource.List()
	if err != nil {
		fmt.Printf("Error getting Account list: %s\n", err)
		os.Exit(1)
	}

	fmt.Println("List:")
	for key, value := range list {
		a := value.(resources.Account)
		fmt.Printf("Account(%d): %s\n", key, a.StringIndent())
	}
}
