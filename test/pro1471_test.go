package test

import (
	"log"
	"testing"

	"github.com/snowcrystall/leetcode-go/problems"
)

func TestGetStrongest(t *testing.T) {
	log.Printf("expact:%v", []int{11, 8, 6, 6, 7})
	log.Printf("get:%v", problems.GetStrongest([]int{6, 7, 11, 7, 6, 8}, 5))
	log.Printf("expact:%v", []int{5, 5})
	log.Printf("get:%v", problems.GetStrongest([]int{1, 1, 3, 5, 5}, 2))
	log.Printf("expact:%v", []int{513})
	log.Printf("get:%v", problems.GetStrongest([]int{513}, 1))
}
