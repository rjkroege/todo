package main

import (
	"flag"
	"log"
	"os"
	"strings"
	//	aw "github.com/deanishe/awgo"
)

var todofile = flag.String("todofile", "Documents/todo.taskpaper", "path to .taskpaper file")
var actionflag = flag.Bool("action", false, "should we ")

func main() {
	log.Println("hello", os.Args)

	cmdname := os.Args[0]

	flag.Parse()

	tpf, err := getTaskPaperFilePath(*todofile)
	if err != nil {
		// TODO(rjk): some way to give up in a way that makes things helpful in Alfred.
		log.Fatal(err)
	}

	if strings.HasSuffix(cmdname, "todo") {
		log.Println("this is the query path")

		genAlfredResult(tpf, flag.Args())

	} else {
		log.Println("this is the action path")

		// TODO(rjk): revisit this when I have the filter working correctly.
		if len(os.Args) == 1 {
			log.Println("this is the action path, no args")
			log.Println("running without arguments so need to just open TaskPaper")

			runTaskPaper(tpf)
		} else {
			log.Println("running action path with arguments so need to prepend")
			prependToTaskPaper(tpf, os.Args)
		}

	}

}
