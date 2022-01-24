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
		log.Printf("this is the action path. real args %#v args: %#v", os.Args, flag.Args())
		if *actionflag {
			log.Println("this is the action path, no args")
			log.Println("running without arguments so need to just open TaskPaper")

			runTaskPaper(tpf)
		} else {
			prependToTaskPaper(tpf, flag.Args())
		}
	}
}
