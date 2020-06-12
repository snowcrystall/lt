package main

// https://leetcode.com/problems/paint-house-iii/

func minCost(houses []int, cost [][]int, m int, n int, target int) int {
	// dg[i][j][k] ,i is the num of houses, j is the num of colors, k is the num of neighborhood
	dg := make([][][]int, m)
	for i := 0; i < m; i++ {
		dg[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			dg[i][j] = make([]int, target)
			for k := 0; k < target; k++ {
				dg[i][j][k] = -1
			}
		}
	}

	mincost := -1

	if houses[0] == 0 {
		for j := 0; j < n; j++ {
			dg[0][j][0] = cost[0][j]
		}
	} else {
		dg[0][houses[0]-1][0] = 0
	}
	for i := 1; i < m; i++ {
		if houses[i] == 0 {
			for j := 0; j < n; j++ {
				for k := 0; k < target; k++ {
					for j2 := 0; j2 < n; j2++ {

						if j == j2 {
							if dg[i-1][j2][k] != -1 && (dg[i][j][k] == -1 || dg[i][j][k] > dg[i-1][j2][k]+cost[i][j]) {
								dg[i][j][k] = dg[i-1][j2][k] + cost[i][j]
							}
						} else {
							if k > 0 && dg[i-1][j2][k-1] != -1 && (dg[i][j][k] == -1 || dg[i][j][k] > dg[i-1][j2][k-1]+cost[i][j]) {
								dg[i][j][k] = dg[i-1][j2][k-1] + cost[i][j]
							}
						}

					}
				}
			}
		} else {
			for j := 0; j < n; j++ {
				for k := 0; k < target; k++ {
					if houses[i]-1 == j {
						if dg[i-1][j][k] != -1 && (dg[i][j][k] == -1 || dg[i][j][k] > dg[i-1][j][k]) {
							dg[i][houses[i]-1][k] = dg[i-1][j][k]
						}
					} else {
						if k > 0 && dg[i-1][j][k-1] != -1 && (dg[i][houses[i]-1][k] == -1 || dg[i][houses[i]-1][k] > dg[i-1][j][k-1]) {
							dg[i][houses[i]-1][k] = dg[i-1][j][k-1]
						}
					}
				}
			}
		}
	}
	for j := 0; j < n; j++ {
		if dg[m-1][j][target-1] != -1 {
			if mincost == -1 {
				mincost = dg[m-1][j][target-1]
			} else {
				if dg[m-1][j][target-1] < mincost {
					mincost = dg[m-1][j][target-1]
				}
			}
		}
	}
	return mincost
}
