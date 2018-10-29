// Package word provides functions to count word frequency and length.
package word

import "strings"

// UseCount return a map of the counts of each word used in a string
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

// Count returns the number of words in a string
func Count(s string) int {
	return len(strings.Fields(s))
}