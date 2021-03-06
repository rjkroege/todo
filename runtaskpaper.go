package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func getTaskPaperFilePath(tpf string) (string, error) {
	if !filepath.IsAbs(tpf) {
		h, err := os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf("getTaskPaperFilePath no home dir %v", err)
		}
		tpf = filepath.Join(h, tpf)
	}

	if ext := filepath.Ext(tpf); ext != ".taskpaper" {
		return "", fmt.Errorf("getTaskPaperFilePath wrong extension %q", ext)
	}
	return tpf, nil
}

func runTaskPaper(tpf string) {
	taskpapercmd := exec.Command("/usr/bin/open", tpf)

	err := taskpapercmd.Run()
	if err != nil {
		log.Fatalf("todo can't open TaskPaper %q: %v", tpf, err)
	}
}
