package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/EIYARO-Project/core-sdk/client"
)

func main() {
	client := client.NewStdClient("http://localhost:9888", "")
	response, err := client.Get("net-info")
	if err != nil {
		fmt.Printf("Error getting net-info: %s", err)
		os.Exit(1)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error reading body: %s", err)
		os.Exit(1)
	}

	var netInfo interface{}
	if err := json.Unmarshal(body, &netInfo); err != nil {
		fmt.Printf("Error decoding net-info: %s", err)
		os.Exit(1)
	}

	if json, err := json.MarshalIndent(netInfo, "", "    "); err != nil {
		fmt.Printf("Error encoding net-info: %s", err)
		os.Exit(1)
	} else {
		fmt.Printf("Net Info: %s\n", string(json))
	}

}
