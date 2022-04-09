package test

import (
	"testing"

	"github.com/snowcrystall/leetcode-go/problems"
)

func TestReachingPoints(t *testing.T) {
	/*if true != problems.ReachingPoints(1, 1, 1, 1) {
		t.Errorf("error for ReachingPoints(1,1,1,1,)")
	}
	if true != problems.ReachingPoints(1, 1, 3, 5) {
		t.Errorf("error for ReachingPoints(1,1,3,5,)")
	}
	if false != problems.ReachingPoints(4, 1, 3, 5) {
		t.Errorf("error for ReachingPoints(1,1,3,5,)")
	}
	if false != problems.ReachingPoints(35, 13, 455955547, 420098884) {
		t.Errorf("error for ReachingPoints(35, 13, 455955547, 420098884)")
	}
	if false != problems.ReachingPoints(1, 1, 2, 2) {
		t.Errorf("error for ReachingPoints(1,1,2,2,)")
	}*/
	if false != problems.ReachingPoints(1, 17, 999999995, 17) {
		t.Errorf("error for ReachingPoints(1, 17, 999999995, 17)")
	}

}
