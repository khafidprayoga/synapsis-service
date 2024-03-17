package main

import (
	"fmt"
	"time"
)

func main() {
	at := time.UnixMilli(1710649092076)
	fmt.Println(at.Format(time.RFC3339))
}
