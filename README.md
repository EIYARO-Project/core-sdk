# EIYARO Go SDK
A software developer kit for the `EIYARO` cryptocurrency.

## Introduction

This `SDK` makes a distinction between endpoints that return a single data type and endpoints that manage `CRUD` like resources.

For example, [Net Info](#net-info) and [Difficulty](#difficulty) are single data type endpoints. These do not have more that one endpoint to manage their returned data type.

On the other hand, [Access Tokens](#access-tokens) is a `CRUD` like resource due to the fact that it has endpoints to list, view, edit, add and delete the corresponding resource.

## Client
In order to have [Dependency Injection](https://en.wikipedia.org/wiki/Dependency_injection), there's an [interface](client/clientinterface.go#L7) for the `HTTP` Client.

The `SDK` provides an implementation using Go's standard library.

### Example
```go
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
```

## API
The `EIYARO` node provides an `API` containing many endpoints. These endpoints are documented in the [API Reference](https://eiyaro.org/api-reference.html).

Below are some examples of the different ways one can access the `API`. 

We will only mention enough examples for you to have an idea of how it works. Eventually full documentation will follow.

### Net Info
This endpoint is retrieved via a `GET` method.

#### Example
```go
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
		fmt.Printf("NetInfo: %s\n", netInfo.StringIndent())
	}
}
```

### Difficulty
This endpoint is retrieved via a `POST` method

#### Example
```go
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
```

## Access Tokens

> **NOTE**\
> Under development