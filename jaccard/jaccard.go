package jaccard

import (
	"github.com/hueyjj/jaccard-string/set"
)

// Coefficient uses Jaccard Coefficient to determine the similarity between
// two strings. The similarity ranges between [0, 1.0], where 1 being a complete
// similar and 0 being not similar at all.
// J(A, B) = |A & B| / |A | B|
func Coefficient(s1, s2 string) float64 {
	return float64(len(set.Intersect(s1, s2))) / float64(len(set.Union(s1, s2)))
}
