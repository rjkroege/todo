package main

import (
	"log"
	//	aw "github.com/deanishe/awgo"

	"encoding/json"
	"os"
	"sort"
	"strings"
)

func genAlfredResult(tpf string, args []string) {

	var result Result

	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "	")

	log.Printf("all args %#v", args)

	// Alfred doesn't tokenize or divide its argument when they're passed
	// directly. So I have to split them up. I also presume that they might
	// have been already split so that I can execute this code from the
	// shell.
	pargs := []string{}
	splitargs := []string{}
	for _, a := range args {
		// TODO(rjk): Probably want a tokenizer instead.
		v := strings.Split(a, " ")

		// Should I track the elements where I have the @?
		splitargs = append(splitargs, v...)

		for _, s := range v {

			// These are the arguments with leading @.
			if strings.HasPrefix(s, "@") {
				pargs = append(pargs, s)
			}

		}
	}

	log.Printf("@ args %#v", args)
	log.Println(splitargs)

	if len(pargs) > 0 {
		tags, err := getTaskPaperTags(tpf)
		if err != nil {
			log.Printf("can't read tags from %q: %v", tpf, tags)
		}

		th := make(map[string]int)

		for _, pr := range pargs {
			for _, t := range tags {
				if strings.HasPrefix(t, pr) && pr != t {
					if s, ok := th[t]; !ok || s < len(pr) {
						th[t] = len(pr)
					}
				}
			}
		}

		for t, v := range th {

			var sb strings.Builder
			for i, a := range splitargs {
				if i > 0 {
					sb.WriteString(" ")
				}
				if strings.HasPrefix(t, a) && t != a {
					sb.WriteString(t)
				} else {
					sb.WriteString(a)
				}
			}

			// Exclude the Uid field to make sure that the items aren't re-ordered.
			result.Items = append(result.Items, &Item{
				// Uid:   t,
				Title:        t,
				Arg:          sb.String(),
				Autocomplete: sb.String() + " ",
				relevance:    v,
			})
		}

		sort.Sort(result.Items)

	}

	// Alfred requires a non-empty Item to offer it in the list. So we create
	// one that we can pass downstream. The downstream (e.g. action handler)
	// will then take different action based on the presence of the flag.
	finalarg := strings.Join(splitargs, " ")
	if finalarg == "" {
		result.Items = append(result.Items, &Item{
			// Uid:   "task",
			Title: "Open TaskPaper",
			Arg:   "-" + actionflagstring,
		})
	} else {
		result.Items = append(result.Items, &Item{
			// Uid:   "task",
			Title: "Add " + finalarg,
			Arg:   "-" + prependflagstring + " " + finalarg,
		})
	}

	if err := encoder.Encode(result); err != nil {
		log.Fatalf("can't write json %v", err)
	}

}

type Item struct {
	Uid          string `json:"uid"`
	Title        string `json:"title"`
	Subtitle     string `json:"subtitle,omitempty"`
	Arg          string `json:"arg"`
	Autocomplete string `json:"autocomplete"`
	relevance    int
}

type Result struct {
	Items ItemCollection `json:"items"`
}

type ItemCollection []*Item

func (c ItemCollection) Len() int {
	return len(c)
}

func (c ItemCollection) Less(i, j int) bool {
	return c[i].relevance > c[j].relevance
}

func (c ItemCollection) Swap(i, j int) {
	tmp := c[i]
	c[i] = c[j]
	c[j] = tmp
}

var _ = ItemCollection(nil)
