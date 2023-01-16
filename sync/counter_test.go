package counter

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		wanted := 1000
		counter := Counter{}

		var wg sync.WaitGroup
		wg.Add(wanted)
		for i := 0; i < wanted; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}

		wg.Wait()

		assertCounter(t, &counter, wanted)
	})
}

func assertCounter(t testing.TB, counter *Counter, value int) {
	t.Helper()
	if counter.Value() != value {
		t.Errorf("wanted %d got %d ", value, counter.Value())
	}

}
