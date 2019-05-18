package set

import "strings"

// Intersect returns a string with all same characters between two strings
func Intersect(s1, s2 string) string {
	var intersect strings.Builder
	set := make(map[rune]bool)
	for _, char := range s1 {
		if _, ok := set[char]; !ok {
			set[char] = true
		}
	}
	for _, char := range s2 {
		if _, ok := set[char]; ok && set[char] {
			intersect.WriteRune(char)
			set[char] = false
		}
	}
	return intersect.String()
}

// Union returns a string with the all non-duplicate characters between two strings
func Union(s1, s2 string) string {
	var intersect strings.Builder
	set := make(map[rune]bool)
	for _, char := range s1 {
		if _, ok := set[char]; !ok {
			set[char] = true
			intersect.WriteRune(char)
		}
	}
	for _, char := range s2 {
		if _, ok := set[char]; !ok {
			set[char] = true
			intersect.WriteRune(char)
		}
	}
	return intersect.String()
}
