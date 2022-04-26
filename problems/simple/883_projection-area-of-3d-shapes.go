package simple

func projectionArea(grid [][]int) int {
	shadowXY := 0
	shadowXZ := 0
	shadowYZ := 0
	ylinemax := map[int]int{}
	for _, xline := range grid {
		xlinemax := 0
		for y, v := range xline {
			if v != 0 {
				shadowXY += 1
			}
			if v > xlinemax {
				xlinemax = v
			}
			if ylinemax[y] < v {
				ylinemax[y] = v
			}
		}
		shadowXZ += xlinemax
	}
	for _, v := range ylinemax {
		shadowYZ += v
	}
	return shadowXZ + shadowXY + shadowYZ
}
