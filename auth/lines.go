package auth

import "strings"

// iterSeq is an iterator over sequences of individual values.
// When called as seq(yield), seq calls yield(v) for each value v in the sequence,
// stopping early if yield returns false.
// See the [iter] package documentation for more details.
type iterSeq[V any] func(yield func(V) bool)

// Lines returns an iterator over the newline-terminated lines in the string s.
// The lines yielded by the iterator include their terminating newlines.
// If s is empty, the iterator yields no lines at all.
// If s does not end in a newline, the final yielded line will not end in a newline.
// It returns a single-use iterator.
func stringLines(s string) iterSeq[string] {
	return func(yield func(string) bool) {
		for len(s) > 0 {
			var line string
			if i := strings.IndexByte(s, '\n'); i >= 0 {
				line, s = s[:i+1], s[i+1:]
			} else {
				line, s = s, ""
			}
			if !yield(line) {
				return
			}
		}
	}
}
