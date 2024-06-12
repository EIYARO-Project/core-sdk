package main

import (
	"fmt"

	"github.com/EIYARO-Project/core-sdk/api"
	"github.com/EIYARO-Project/core-sdk/client"
)

func main() {
	client := client.NewStdClient("http://localhost:9888", "")
	api := api.NewApi(client)
	netInfo, err := api.NetInfo()
	if err != nil {
		fmt.Printf("Error getting NetInfo: %s", err)
	} else {
		fmt.Printf("NetInfo: %s", netInfo.StringIndent())
	}

}
