package main

import "fmt"

func main() {
	var i int
	var value string
	for i, value = range languages {
		fmt.Println("language", i, ":", value)
	}
	fmt.Println("product with key 2", products[2])
}
