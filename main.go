package main

import (
	"fmt"

	"github.com/hueyjj/jaccard-string/jaccard"
	"github.com/hueyjj/jaccard-string/set"
)

func main() {
	// Intersect example
	s1, s2 := "abcdefgh", "hgxxxxxfedab"
	fmt.Printf("Intersect(%s, %s): %s\n", s1, s2, set.Intersect(s1, s2))
	// > Intersect(abcdefgh, hgxxxxxfedab): hgfedab

	// Union example
	s1, s2 = "helllo", "helllostr"
	fmt.Printf("Union(%s, %s): %s\n", s1, s2, set.Union(s1, s2))
	// > Union(helllo, helllostr): helostr

	// Using the Coefficient function
	s1, s2 = "qwerty", "foobarqe"
	fmt.Printf("Jaccard coefficient (%s, %s): %f\n", s1, s2, jaccard.Coefficient(s1, s2))
	// > Jaccard coefficient (qwerty, foobarqe): 0.300000

	s1, s2 = "hello", "hello world"
	fmt.Printf("Jaccard coefficient (%s, %s): %f\n", s1, s2, jaccard.Coefficient(s1, s2))
	// > Jaccard coefficient (hello, hello world): 0.500000

	s1, s2 = "hello world", "world hello"
	fmt.Printf("Jaccard coefficient (%s, %s): %f\n", s1, s2, jaccard.Coefficient(s1, s2))
	// > Jaccard coefficient (hello world, world hello): 1.000000
}
