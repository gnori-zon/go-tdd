package _select

import (
	"fmt"
	"net/http"
	"time"
)

type ErrTimeout struct {
	urls    []string
	timeout time.Duration
}

func (e ErrTimeout) Error() string {
	return fmt.Sprintf("timeout %s for %v", e.timeout, e.urls)
}

var tenSecondTimeout = 10 * time.Second

func Racer(lhsUrl, rhsUrl string) (winner string, err error) {
	return ConfigurableRacer(lhsUrl, rhsUrl, tenSecondTimeout)
}

func ConfigurableRacer(lhsUrl, rhsUrl string, timeout time.Duration) (winner string, err error) {
	select {
	case <-ping(lhsUrl):
		return lhsUrl, nil
	case <-ping(rhsUrl):
		return rhsUrl, nil
	case <-time.After(timeout):
		return "", ErrTimeout{urls: []string{lhsUrl, rhsUrl}, timeout: timeout}
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
