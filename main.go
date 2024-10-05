package main

import (
	f "fmt"
	"main/utils"
	_ "net/http"
)

func main() {
	f.Println("Start main package")

	f.Printf("Result: %d\n", utils.AddInt(1, 2))
	defer f.Println("Finish main package")
}
