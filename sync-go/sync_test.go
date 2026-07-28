package syncgo

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing 3 tims leaves the counter at 3", func(t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, 3, counter)
	})

	t.Run("incrementing concurrently", func(t *testing.T) {
		wantedCount := 1000
		counter := NewCounter()

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for range wantedCount {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}

		wg.Wait()

		assertCounter(t, wantedCount, counter)
	})
}

func assertCounter(t testing.TB, want int, counter *Counter) {
	t.Helper()
	if want != counter.Value() {
		t.Errorf("Wanted counter to be %d, but it is %d", want, counter.Value())
	}
}
