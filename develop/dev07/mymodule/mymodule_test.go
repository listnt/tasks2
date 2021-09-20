package mymodule

import (
	"testing"
	"time"
)

func TestCase1(t *testing.T){
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}
	start := time.Now()
	<-Or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)
	t1:=time.Since(start)
	if t1<time.Second || t1>2*time.Second{
		t.Error("expected something around 1 sec got",t1)
	}
}