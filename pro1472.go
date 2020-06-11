package main

type BrowserHistory struct {
	history []string
	cur     int
}

func Constructor(homepage string) BrowserHistory {
	return BrowserHistory{[]string{homepage}, 0}
}

func (this *BrowserHistory) Visit(url string) {
	this.history = this.history[:this.cur+1]
	this.cur++
	this.history = append(this.history, url)
}

func (this *BrowserHistory) Back(steps int) string {
	if steps > this.cur {
		this.cur = 0
		return this.history[0]
	} else {
		this.cur -= steps
		return this.history[this.cur]
	}
}

func (this *BrowserHistory) Forward(steps int) string {
	if len(this.history[:])-this.cur > steps {
		this.cur += steps
		return this.history[this.cur]
	} else {
		this.cur = len(this.history[:]) - 1
		return this.history[this.cur]
	}

}
