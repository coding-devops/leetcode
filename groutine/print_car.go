package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func test(tch chan struct{}) {
	fmt.Println("123123")
	<-tch
	fmt.Println("1ppp")

}
func test2() {
	time.Sleep(time.Second * 3)
}
func main() {
	//catch := make(chan struct{}, 1)
	//go test(catch)
	////go test2()
	//time.Sleep(time.Second * 1)
	//catch <- struct{}{}
	//time.Sleep(time.Second * 3)
	wg := sync.WaitGroup{}

	wg.Add(3)

	catch := make(chan struct{}, 1)
	dogch := make(chan struct{}, 1)
	fishch := make(chan struct{}, 1)

	go cat(&wg, catch, fishch)
	go dog(&wg, dogch, catch)
	go fish(&wg, fishch, dogch)
	fishch <- struct{}{}
	wg.Wait()
}

func cat(group *sync.WaitGroup, catch chan struct{}, fish chan struct{}) {
	var count uint64 = 0
	for {
		if count >= 5 {
			group.Done()
			return
		}
		<-fish
		fmt.Println("cat")
		atomic.AddUint64(&count, 1)
		catch <- struct{}{}
	}
}
func dog(group *sync.WaitGroup, dogch chan struct{}, catch chan struct{}) {
	var count uint64 = 0
	for {
		if count >= 5 {
			group.Done()
			return
		}
		<-catch
		fmt.Println("dog")
		atomic.AddUint64(&count, 1)
		dogch <- struct{}{}
	}
}
func fish(group *sync.WaitGroup, fish chan struct{}, dogch chan struct{}) {
	var count uint64 = 0
	for {
		if count >= 5 {
			group.Done()
			return
		}
		<-dogch
		fmt.Println("fish")
		fmt.Println("")
		atomic.AddUint64(&count, 1)
		fish <- struct{}{}
	}
}
