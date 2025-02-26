package _test

import "testing"

func TestRun(t *testing.T) {

}

type BrowserHistory struct {
	urls      []string
	currIndex int
}

func Constructor(homePage string) BrowserHistory {
	return BrowserHistory{
		urls:      []string{homePage},
		currIndex: 0,
	}
}
func (this *BrowserHistory) Visit(url string) {
	for len(this.urls) > this.currIndex+1 {
		this.urls = this.urls[:len(this.urls)-1]
	}
	this.urls = append(this.urls, url)
	this.currIndex++
}

func (this *BrowserHistory) Back(steps int) string {
	this.currIndex = max(this.currIndex-steps, 0)
	return this.urls[this.currIndex]
}

func (this *BrowserHistory) Forward(steps int) string {
	this.currIndex = min(this.currIndex+steps, len(this.urls)-1)
	return this.urls[this.currIndex]
}
