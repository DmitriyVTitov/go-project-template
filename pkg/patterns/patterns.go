package patterns

import "sync"

// FanIn pattern generic implementation.
func FanIn[T any](chans ...chan T) chan T {
	out := make(chan T)

	var wg sync.WaitGroup

	for _, ch := range chans {
		go func(ch chan T) {
			for val := range ch {
				out <- val
			}

			wg.Done()
		}(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
