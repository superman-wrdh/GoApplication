package main

import (
	"fmt"
	"sync"
)

const N = 20

func main() {
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	m := make(map[int]int)

	wg.Add(N)

	for i := 0; i < N; i++ {
		go func() {
			mu.Lock()
			defer wg.Done()
			m[i] = i
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println(len(m))
	fmt.Println(m)
}
