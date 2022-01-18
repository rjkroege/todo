package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"

	"github.com/felixge/taskpaper"
)

func recurse(item *taskpaper.Item, tags map[string]struct{}) {
	tagsinentry := parseTaskPaperItemForTags(item.Content)
	for _, ti := range tagsinentry {
		tags[ti] = struct{}{}
	}

	for _, ci := range item.Children {
		recurse(ci, tags)
	}
}

// getTaskPaperTags extracts the TaskPaper file tags from tpf.
func getTaskPaperTags(tpf string) ([]string, error) {
	log.Println("getting tags from", tpf)

	contents, err := os.ReadFile(tpf)
	if err != nil {
		return []string{}, fmt.Errorf("todo can't read taskpaper file %q: %v", tpf, err)
	}

	root, err := taskpaper.Unmarshal(contents)
	if err != nil {
		return []string{}, fmt.Errorf("todo can't parse taskpaper file %q: %v", tpf, err)
	}

	tags := make(map[string]struct{})

	recurse(root, tags)

	taglist := make([]string, 0, len(tags))
	for k := range tags {
		// Filter out tags that aren't interesting: done markers
		if !strings.HasPrefix(k, "@done") && !strings.HasPrefix(k, "@project") && !strings.HasPrefix(k, "@nongoal") && !strings.HasPrefix(k, "@search") {
			taglist = append(taglist, k)
		}
	}

	return taglist, nil
}

const (
	START   = iota
	TAG     = iota
	BRACKET = iota
	TEXT    = iota
)

func parseTaskPaperItemForTags(entry string) []string {
	tags := make([]string, 0)

	state := START
	tagstart := 0

	for i, r := range entry {
		switch {
		case state == START && r == '@':
			state = TAG
			tagstart = i
		case state == START && r != '@':
			state = TEXT
		case state == TAG && unicode.IsSpace(r):
			state = START
			tags = append(tags, entry[tagstart:i])
		case state == TAG && r == '(':
			state = BRACKET
		case state == TAG && !(unicode.IsLetter(r) || unicode.IsNumber(r) || r == '_'):
			state = TEXT
		case state == BRACKET && r == ')':
			state = TEXT
			tags = append(tags, entry[tagstart:i+1])
		case state == TEXT && unicode.IsSpace(r):
			state = START
		}
	}

	switch {
	case state == TAG:
		tags = append(tags, entry[tagstart:])
	}

	return tags
}
