package main

import (
	"github.com/nounoursheureux/goraph"
	"io/ioutil"
	"log"
)

func handle(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	var dot goraph.DotTranslater
	var dotsrc, err = goraph.ConvertFile("graph", &dot)
	handle(err)

	handle(ioutil.WriteFile("graph.dot", []byte(dotsrc), 0666))
}
