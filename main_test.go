package main

import (
	"testing"

	"github.com/hueyjj/jaccard-string/jaccard"
)

func TestMain(t *testing.T) {
	s1, s2 := "hello world", "world hello"
	coefficient := jaccard.Coefficient(s1, s2)
	if coefficient != 1.0 {
		t.Errorf("Expected jaccard.Coefficient(%s, %s)=%f, but got %f", s1, s2, 1.0, coefficient)
	}
	s1, s2 = "magic", "magic mike"
	coefficient = jaccard.Coefficient(s1, s2)
	if coefficient != 0.625 {
		t.Errorf("Expected jaccard.Coefficient(%s, %s)=%f, but got %f", s1, s2, 0.625, coefficient)
	}
	s1, s2 = "a", "a"
	coefficient = jaccard.Coefficient(s1, s2)
	if coefficient != 1.0 {
		t.Errorf("Expected jaccard.Coefficient(%s, %s)=%f, but got %f", s1, s2, 1.0, coefficient)
	}
	s1, s2 = "random string !#@$*(#@", "zzzzzzzzzzz"
	coefficient = jaccard.Coefficient(s1, s2)
	if coefficient != 0.0 {
		t.Errorf("Expected jaccard.Coefficient(%s, %s)=%f, but got %f", s1, s2, 0.0, coefficient)
	}
}
