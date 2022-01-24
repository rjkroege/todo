package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func prependToTaskPaper(tpf string, message []string) {
	// log.Printf("Append %#v to the taskpaper here", message)

	ofd, err := os.Open(tpf)
	if err != nil {
		log.Fatalf("can't open %q: %v", tpf, err)
	}

	tmp := tpf + ".tmp"
	nfd, err := os.Create(tmp)
	if err != nil {
		log.Fatalf("can't create %q: %v", tmp, err)
	}

	newitem := "- " + strings.Join(message, " ") + "\n"

	if _, err := nfd.WriteString(newitem); err != nil {
		nfd.Close()
		os.Remove(tmp)
		log.Fatalf("can't write to %q: %v", tmp, err)
	}

	if _, err := io.Copy(nfd, ofd); err != nil {
		nfd.Close()
		os.Remove(tmp)
		log.Fatalf("can't copy %q to %q: %v", tpf, tmp, err)
	}

	nfd.Close()
	ofd.Close()

	oldtpf := tpf + ".old"
	if err := os.Link(tpf, oldtpf); err != nil {
		os.Remove(tmp)
		log.Fatalf("can't link %q to %q: %v", tpf, oldtpf, err)
	}

	if err := os.Remove(tpf); err != nil {
		os.Remove(tmp)
		log.Fatalf("can't Remove %q to %q: %v", tpf, oldtpf, err)
	}

	if err := os.Link(tmp, tpf); err != nil {
		log.Fatalf("can't link %q to %q: %v", tmp, tpf, err)
	}

	os.Remove(tmp)
	os.Remove(oldtpf)
}
