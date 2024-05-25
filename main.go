package main

import (
	"fmt"
	"strings"
)

func main() {
	a := strings.Split("asc", ",")
	fmt.Println(len(a))
	fmt.Println(a[0])
}
