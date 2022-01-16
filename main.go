package main

import (
	"flag"
	"log"
)

var todofile = flag.String("todofile", "Documents/todo.taskpaper", "path to .taskpaper file")

func main() {
	log.Println("hello")

	// Define arguments here

	flag.Parse()

	tpf, err := getTaskPaperFilePath(*todofile)
	if err != nil {
		// TODO(rjk): some way to give up in a way that makes things helpful in Alfred.
		log.Fatal(err)
	}

	if flag.NArg() > 0 {
		log.Println("running with arguments so need to do something different")
		return
	}

	log.Println("running without arguments so need to just open TaskPaper")

	runTaskPaper(tpf)

	// TODO(rjk): decide how I want to log for Alfred debugging?

}
