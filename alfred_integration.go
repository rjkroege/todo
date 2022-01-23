package main

import (
	"log"
	//	aw "github.com/deanishe/awgo"

	"encoding/json"
	"os"
	"strings"
	"sort"
)

func genAlfredResult(tpf string, args []string) {

	var result Result

	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "	")


	log.Printf("all args %#v", args)

	pargs := []string{}
	for _, a := range args {
		// Probably want a tokenizer instead.
		v := strings.Split(a, " ")

		for _, s := range(v) { 

		if strings.HasPrefix(s, "@") {
			pargs = append(pargs, s)
		}

		}
	}

	log.Printf("@ args %#v", args)


	if len(pargs) > 0 {
		tags, err := getTaskPaperTags(tpf)
		if err != nil {
			log.Printf("can't read tags from %q: %v", tpf, tags)
		}

		th :=  make(map[string]int)
		
		for _, pr := range pargs {
			for _, t := range tags {
				if strings.HasPrefix(t, pr) &&  pr != t {
					if s, ok := th[t]; !ok ||  s < len(pr) {
						th[t] = len(pr)
					}
				}
			}
		}

		for t, v := range th {
			result.Items = append(result.Items, &Item{
				Uid:   t,
				Title: t,
				// Maybe use a StringBuilder
				// Must be smarter
				Arg:          strings.Join(args, " ") + t + " ",
				// TODO(rjk): be smarter...
				Autocomplete: strings.Join(args, " ") + t + " ",
				relevance: v,
			})
		}


		sort.Sort(result.Items)

		
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
	relevance int
}

type Result struct {
	Items ItemCollection  `json:"items"`
}

type ItemCollection []*Item

func (c ItemCollection) Len() int {
	return len(c)
}

func (c ItemCollection) Less(i,  j int) bool {
	return c[i].relevance < c[j].relevance
}

func (c ItemCollection) Swap(i, j int) {
	tmp := c[i]
	c[i] = c[j]
	c[j] = tmp
}

var _ = ItemCollection(nil)
