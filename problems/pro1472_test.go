package problems

import (
	"log"
	"testing"
)

func TestVisit(t *testing.T) {
	h := Constructor("leetcode.com")
	h.Visit("google.com")
	h.Visit("facebook.com")
	h.Visit("youtube.com")
	log.Println(h.Back(1))
	log.Println(h.Back(1))
	h.Visit("linkedin.com")
	log.Println(h.Forward(2))
	log.Println(h.Forward(3))
}
