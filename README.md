# jaccard-string
jaccard-string determines how similar two strings are using Jaccard Index (https://en.wikipedia.org/wiki/Jaccard_index). 

"hello world" and "world hello" is 100% similar (coefficient = 1.0).

"hello world" and "hello ddddd" is 50% similar (coefficient = 0.5).

"hello world" and "zzzzzzz" is 0% similar (coefficient = 0.0).

# Quickstart
```bash
$ go get github.com/hueyjj/jaccard-string
$ go build
$ ./jaccard-string.exe
Intersect(abcdefgh, hgxxxxxfedab): hgfedab
Union(helllo, helllostr): helostr
Jaccard coefficient (qwerty, foobarqe): 0.300000
Jaccard coefficient (hello, hello world): 0.500000
Jaccard coefficient (hello world, world hello): 1.000000
```

# Example usage
```golang
package main

import (
    "github.com/hueyjj/jaccard-string/set"
    "github.com/hueyjj/jaccard-string/jaccard"
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
```

# Test
```bash
$ go test
PASS
ok      github.com/hueyjj/jaccard-string        0.280s
```