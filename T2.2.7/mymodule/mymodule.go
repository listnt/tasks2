package mymodule

func Or(channels ...<-chan interface{}) <-chan interface{} {
	out := make(chan interface{})

	for _, ch := range channels {
		go func(ch1 <-chan interface{}) {
			_, ok := <-ch1
			if !ok {
				close(out)
			}
		}(ch)
	}
	return out
}
