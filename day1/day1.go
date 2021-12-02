package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	fmt.Println("Getting input data")

	content, err := ioutil.ReadFile("./data")

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(content))
}
