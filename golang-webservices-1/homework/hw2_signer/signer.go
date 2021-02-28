package main

import (
	"fmt"
	"sort"
	"strings"
	"sync"
)

type wrap func(string) string

func ExecutePipeline(jobs ...job) {
	var in, out chan interface{}
	wg := &sync.WaitGroup{}

	for _, j := range jobs {
		in = out
		out = make(chan interface{}, 100)
		wg.Add(1)

		go func(j job, in, out chan interface{}) {
			defer wg.Done()
			defer close(out)
			j(in, out)
		}(j, in, out)
	}

	wg.Wait()
}

func SingleHash(in, out chan interface{}) {
	// OverheatLock prevent
	md5Mutex := &sync.Mutex{}
	f := func(s string) string {
		h := func(s string, w wrap, out chan string) {
			defer close(out)
			out <- w(s)
		}
		// crc32(data)
		hash1 := make(chan string)
		go h(s, DataSignerCrc32, hash1)
		// crc32(md5(data))
		hash2 := make(chan string)
		go func() {
			// md5(data)
			md5Mutex.Lock()
			md5Hash := DataSignerMd5(s)
			md5Mutex.Unlock()
			go h(md5Hash, DataSignerCrc32, hash2)
		}()
		return fmt.Sprintf("%s~%s", <-hash1, <-hash2)
	}
	wrapper(f, in, out)
}

func MultiHash(in, out chan interface{}) {
	f := func(s string) string {
		type Result struct {
			th   int
			hash string
		}
		results := make([]string, 6)
		c := make(chan Result, 6)
		for th := 0; th < 6; th++ {
			go func(th int, s string) {
				c <- Result{th, DataSignerCrc32(fmt.Sprintf("%d%s", th, s))}
			}(th, s)
		}
		for th := 0; th < 6; th++ {
			result := <-c
			results[result.th] = result.hash
		}
		return strings.Join(results, "")
	}
	wrapper(f, in, out)
}

func CombineResults(in, out chan interface{}) {
	var results []string

	for data := range in {
		results = append(results, fmt.Sprintf("%v", data))
	}

	sort.Strings(results)
	out <- strings.Join(results, "_")
}

func wrapper(w wrap, in, out chan interface{}) {
	wg := &sync.WaitGroup{}

	for data := range in {
		wg.Add(1)
		go func(s string) {
			defer wg.Done()
			out <- w(s)
		}(fmt.Sprintf("%v", data))
	}

	wg.Wait()
}
