![](assets/DALL·E%202024-02-19%2021.57.30%20-%20a%2019th%20century%20oil%20painting%20of%20the%20greek%20figure%20tyche.png)
# Tyche
A credit card information and validation library for golang.

## Installation
```bash
go get github.com/torbenconto/tyche
```

## Usage

```go
package main

import (
	"fmt"
	"github.com/torbenconto/tyche"
	"time"
)

func main() {
	card := tyche.NewCard("4111111111111111", "123", time.Date(2027, time.January, 1, 0, 0, 0, 0, time.UTC))
	fmt.Println(card.IsValid())           // true
	fmt.Println(card.Provider().String()) // Visa
}
```
