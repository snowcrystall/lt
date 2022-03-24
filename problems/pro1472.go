package problems

type BrowserHistory struct {
	history []string
	cur     int
}

func Constructor(homepage string) BrowserHistory {
	return BrowserHistory{[]string{homepage}, 0}
}

func (b *BrowserHistory) Visit(url string) {
	b.history = b.history[:b.cur+1]
	b.cur++
	b.history = append(b.history, url)
}

func (b *BrowserHistory) Back(steps int) string {
	if steps > b.cur {
		b.cur = 0
		return b.history[0]
	} else {
		b.cur -= steps
		return b.history[b.cur]
	}
}

func (b *BrowserHistory) Forward(steps int) string {
	if len(b.history[:])-b.cur > steps {
		b.cur += steps
		return b.history[b.cur]
	} else {
		b.cur = len(b.history[:]) - 1
		return b.history[b.cur]
	}

}
