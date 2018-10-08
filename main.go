package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Nodes struct {
	Nodes []Node `json:"nodes"`
}

type Node struct {
	Name string `json:"name"`
	IP   string `json:"ip"`
	X    int    `json:"x"`
	Y    int    `json:"y"`
}

func main() {

	jsonFile, err := os.Open("nodes.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened nodes.json")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var nodes Nodes
	json.Unmarshal(byteValue, &nodes)

	fmt.Println(nodes.Nodes[5])

}
