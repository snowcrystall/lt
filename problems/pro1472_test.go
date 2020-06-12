package main

import (
	"log"
	"testing"
)

func TestVisit(t *testing.T) {
	h := Constructor("leetcode.com")
	h.Visit("google.com")
	h.Visit("facebook.com")
	h.Visit("youtube.com")
	log.Printf(h.Back(1))
	log.Printf(h.Back(1))
	h.Visit("linkedin.com")
	log.Printf(h.Forward(2))
	log.Printf(h.Forward(3))
}
