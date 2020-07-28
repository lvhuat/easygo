package taketime

import "time"

func Do(f func()) time.Duration {
	since := time.Now()
	f()
	return time.Now().Sub(since)
}
