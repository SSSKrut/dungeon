package main

import "fmt"

func main() {
	storage := NewInMemoryStorage()

	n := Node{ID: "1", Value: "first"}
	storage.Add(n)

	node, exists := storage.Get("1")

	if exists {
		fmt.Printf("Node exists: %v\n", node)
	} else {
		fmt.Printf("Node does not exist\n")
	}

	storage.Delete("1")
}
