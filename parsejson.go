package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Node struct {
	Name string `json:"name"`
	IP   string `json:"ip"`
	X    int    `json:"x"`
	Y    int    `json:"y"`
}

func parseJSON(filename string) []Node {

	jsonFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Successfully opened file: ", filename)
		defer jsonFile.Close()
	}
	nodes := make([]Node, 100)
	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &nodes)
	return nodes

}
