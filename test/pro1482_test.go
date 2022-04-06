package test

import (
	"log"
	"testing"

	"github.com/snowcrystall/leetcode-go/problems"
)

func TestMinDays(t *testing.T) {
	log.Printf("expact:3")
	log.Printf("get:%d", problems.MinDays([]int{1, 10, 3, 10, 2}, 3, 1))
	log.Printf("expact:-1")
	log.Printf("get:%d", problems.MinDays([]int{1, 10, 3, 10, 2}, 3, 2))
	log.Printf("expact:12")
	log.Printf("get:%d", problems.MinDays([]int{7, 7, 7, 7, 12, 7, 7}, 2, 3))

}
