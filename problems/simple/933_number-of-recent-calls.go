package simple

type RecentCounter struct {
	stack []int
}

func Constructor() RecentCounter {
	return RecentCounter{[]int{}}
}

func (this *RecentCounter) Ping(t int) int {
	this.stack = append(this.stack, t)

	for this.stack[0] < t-3000 {
		this.stack = this.stack[1:]
	}
	return len(this.stack)

}
