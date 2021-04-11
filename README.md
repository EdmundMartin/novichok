# Novichok

Novichok is a Golang API Client for the Chess.com API.

Currently, a small number of methods are supported:
```
PlayerProfile
PlayerStats
GameArchive
```

Example Usage:
```go
package main

import (
	"context"
	"fmt"
	"github.com/EdmundMartin/novichok"
	"time"
)

func main() {
	client := novichok.NewChessComClient(20 * time.Second)
	res, _ := client.GetGameArchive(context.Background(), "EJM979", time.Now())
	fmt.Println(res)
}
```