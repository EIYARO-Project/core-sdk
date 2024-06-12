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
	accessTokenResource, err := api.Resource("AccessToken")
	if err != nil {
		fmt.Printf("Error getting resource AccessToken: %s\n", err)
		os.Exit(1)
	}

	list, err := accessTokenResource.List()
	if err != nil {
		fmt.Printf("Error getting AccessToken list: %s\n", err)
		os.Exit(1)
	}

	fmt.Println("List:")
	for key, value := range list {
		at := value.(resources.AccessToken)
		fmt.Printf("AccessToken(%d): %s\n", key, at.StringIndent())
	}
}
