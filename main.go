package main

import (
	"fmt"
)

var nodes []Node

func main() {

	nodes := parseJSON("nodes.json")
	fmt.Println(nodes[0])

}
