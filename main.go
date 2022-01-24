package main

import (
	"flag"
	"log"
	"os"
	"strings"
	//	aw "github.com/deanishe/awgo"
)

const (
	prependflagstring = "alfredconnectionprepend"
	actionflagstring  = "alfredconnectionaction"
)

var todofile = flag.String("todofile", "Documents/todo.taskpaper", "path to .taskpaper file")
var actionflag = flag.Bool(actionflagstring, false, "run TaskPaper")
var prependflag = flag.Bool(prependflagstring, false, "prepend args to TaskPaper file")

func main() {
	log.Printf("hello %#v", os.Args)

	//	cmdname := os.Args[0]

	// Alfred doesn't parse arguments like the shell. So I have to split them
	// apart here. Note that this will make it tedious to create a todo list
	// entry containing the reserved commandline names.
	// TODO(rjk): Make the special flags less likely to overlap with with
	// anything that I'd like to type.
	ss := []string{}
	for _, s := range os.Args {
		ds := strings.Split(s, " ")
		ss = append(ss, ds...)
	}
	os.Args = ss

	flag.Parse()

	tpf, err := getTaskPaperFilePath(*todofile)
	if err != nil {
		// TODO(rjk): some way to give up in a way that makes things helpful in Alfred.
		log.Fatal(err)
	}

	switch {
	case *actionflag:
		log.Println("this is the action path, no args")
		log.Println("running without arguments so need to just open TaskPaper")
		runTaskPaper(tpf)
	case *prependflag:
		prependToTaskPaper(tpf, flag.Args())
	default:
		genAlfredResult(tpf, flag.Args())
	}
}
