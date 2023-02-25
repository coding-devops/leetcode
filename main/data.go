package main

import "fmt"

var active = make(chan struct{}, 4)
var jobs = make(chan int, 10)

// func main() {
// 	go func() {
// 		for i := 0; i < 10; i++ {
// 			jobs <- (i + 1)
// 		}
// 		close(jobs)
// 		fmt.Println("222")
// 	}()
// 	time.Sleep(time.Second)
// 	i := <-jobs
// 	fmt.Println(i)
// 	var wg sync.WaitGroup

// 	for j := range jobs {
// 		wg.Add(1)
// 		go func(j int) {
// 			active <- struct{}{}
// 			log.Printf("handle job: %d\n", j)
// 			time.Sleep(2 * time.Second)
// 			<-active
// 			wg.Done()
// 		}(j)
// 	}
// 	wg.Wait()
// }

//从一个已关闭的channel接受
// func main() {

// 	ctx, cancel := context.WithCancel(context.Background())

// 	gen := func(ctx context.Context) {
// 		for {
// 			select {
// 			case <-ctx.Done():
// 				fmt.Println("123")
// 				return
// 			case <-work():
// 				fmt.Println("gg")
// 			}
// 			fmt.Println("ff")
// 		}

// 	}
// 	go gen(ctx)
// 	time.Sleep(time.Second * 2)
// 	cancel()
// 	fmt.Println("cancel，此时应该输出123")
// 	time.Sleep(time.Second * 10)
// }

// func work() chan int {
// 	c := make(chan int, 1)
// 	n := 0
// 	for {
// 		time.Sleep(time.Second * 1)
// 		n++
// 		if n == 5 {
// 			c <- 1
// 			break
// 		}
// 	}
// 	return c
// }

func main() {
	a := []int{112, 2, 3, 4, 12, 0}
	quickSort(a, 0, len(a)-1)

	for _, v := range a {
		fmt.Println(v)
	}
	// fmt.Println("length is ", len(a))
	// sort.Ints(a)
	// for _, v := range a {
	// 	fmt.Println(v)
	// }
}
func quickSort(a []int, low int, high int) {
	if len(a) <= 0 {
		return
	}
	i := low //基准数
	j := high
	if i >= j {
		return
	}
	tmp := a[i]
	for i < j {
		for i < j {
			if a[j] < tmp {
				a[i] = a[j]
				i++
				break
			} else {
				j--
			}
		}
		for i < j {
			if a[i] > tmp {
				a[j] = a[i]
				j--
				break
			} else {
				i++
			}
		}
	}
	a[i] = tmp

	quickSort(a, low, i-1)
	quickSort(a, i+1, high)
}

type IntSlice []int

func (x IntSlice) Len() int           { return len(x) }
func (x IntSlice) Less(i, j int) bool { return x[i] < x[j] }
func (x IntSlice) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }
