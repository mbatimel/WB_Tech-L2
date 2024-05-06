package main

import (
	"testing"
	"time"
)

func TestOr(t *testing.T) {
	sig := func(after time.Duration) <- chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
	}()
	return c
	}
	
	start := time.Now()
	<-or (
		sig(1*time.Second),
		sig(2*time.Second),
		sig(3*time.Second),
		sig(4*time.Second),
	)
	
	
	elapsed := time.Since(start)

	// Проверяем, что прошло не менее 5 минут и не более 2 часов
	if elapsed < 1*time.Second || elapsed > 5*time.Second {
		t.Errorf("Unexpected time elapsed: %v", elapsed)
	}
}