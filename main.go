package main

import (
	"log"
	"flag"
)

func main() {
	log.Println("hello")

	// Define arguments here

	flag.Parse()

	if flag.NArg() > 0 {
		log.Println("running with arguments so need to do something different")
		return
	}

	log.Println("running without arguments so need to just open TaskPaper")

	// TODO(rjk): decide how I want to log for Alfred debugging?
	
}

