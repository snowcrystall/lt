package test

import (
	"log"
	"testing"

	"github.com/snowcrystall/leetcode-go/problems"
)

func TestVisit(t *testing.T) {
	h := problems.Constructor("leetcode.com")
	h.Visit("google.com")
	h.Visit("facebook.com")
	h.Visit("youtube.com")
	log.Println(h.Back(1))
	log.Println(h.Back(1))
	h.Visit("linkedin.com")
	log.Println(h.Forward(2))
	log.Println(h.Forward(3))
}
