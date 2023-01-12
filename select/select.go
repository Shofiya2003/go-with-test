package racer

import (
	"fmt"
	"net/http"
	"time"
)

const tenSecondTimeOut = 10 * time.Second

func Racer(url1, url2 string) (winner string, err error) {
	return ConfigurableRacer(url1, url2, tenSecondTimeOut)
}

func ConfigurableRacer(url1, url2 string, timeout time.Duration) (winner string, err error) {
	select {
	case <-ping(url1):
		return url1, nil
	case <-ping(url2):
		return url2, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out while waiting for %s and %s", url1, url2)
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})

	go func() {
		http.Get(url)
		close(ch)
	}()

	return ch
}

func measureResponseTime(url string) time.Duration {
	startTime1 := time.Now()
	http.Get(url)
	timeDuration1 := time.Since(startTime1)
	return timeDuration1
}
