package main

import (
	"fmt"

	initialize "Society-Synergy/init"
)

func init() {
	initialize.InitializeSetup()
}

func main() {
	fmt.Println("Hello world")
	initialize.StartServer()
}
