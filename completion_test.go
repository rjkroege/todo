package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParseTaskPaperItemForTags(t *testing.T) {
	for _, tc := range []struct {
		entry string
		want  []string
	}{
		{
			entry: "@foo",
			want:  []string{"@foo"},
		},
		{
			entry: "bar @foo",
			want:  []string{"@foo"},
		},
		{
			entry: "bar @foo hello",
			want:  []string{"@foo"},
		},
		{
			entry: "bar @foo hello @bang",
			want:  []string{"@foo", "@bang"},
		},
		{
			entry: "bar @foo(a, b, c) hello @bang",
			want:  []string{"@foo(a, b, c)", "@bang"},
		},
		{
			entry: "foo@mu",
			want:  []string{},
		},
		{
			entry: "moo @blah(jkjk, jkjk)",
			want:  []string{"@blah(jkjk, jkjk)"},
		},
		{
			entry: "cow @jk_jkjk@(jk)",
			want:  []string{},
		},
		{
			entry: "milk @a(@, jk, #)",
			want:  []string{"@a(@, jk, #)"},
		},
		{
			// This matches incorrectly.
			entry: "cream @blah@",
			want:  []string{},
		},
		{
			// This matches incorrectly.
			entry: "@aa @bb",
			want:  []string{"@aa", "@bb"},
		},
		{
			// This matches incorrectly.
			entry: "@a @b",
			want:  []string{"@a", "@b"},
		},
	} {
		t.Run(tc.entry, func(t *testing.T) {
			got := parseTaskPaperItemForTags(tc.entry)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("dump mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
