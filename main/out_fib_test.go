package main

import (
	ctx "context"
	"fmt"
	"sync"
	"testing"
	"time"
	"unsafe"
)

//输出前n位fib
func TestFib(t *testing.T) {
	var a int = 1
	var b int = 0
	for i := 0; i < 10; i++ {
		t.Log(a)
		temp := b
		b = a
		a = a + temp
	}
}

//快速交换两个值
func TestSwitchPara(t *testing.T) {
	a := 1
	b := 2
	a, b = b, a
	t.Log(a)
}

//递归计算前n位fib数的和

// 数组比较
func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 2, 3, 4}
	c := [...]int{1, 2, 3, 4, 5}
	t.Log(a == b)
	t.Log(c == c)
}

const (
	Readable = 1 << iota
	Writeable
	R_W
)

//按位清零
func TestBitClear(t *testing.T) {
	a := 7            // 0111
	a = a &^ Readable //  将可读位置零
	t.Log(a)
}

//循环
func TestLoop(t *testing.T) {
	for {
		t.Log("111")
		break
	}
}

//集合
// 数组和切片

func TestArrayAndSlice(t *testing.T) {

	//var a [3]int //默认初始化为0值
	b := [3]int{1, 2, 3} //声明的同时初始化
	b[0] = 1
	b[1] = 2
	b[2] = 3

}

//数组的遍历

func TestArrayTravel(t *testing.T) {
	b := [3]int{1, 2, 3} //声明的同时初始化
	for _, e := range b {
		t.Log(e)
	}
}

//数组截取
func Test_sectionArray(t *testing.T) {
	arr3 := [...]int{1, 2, 3, 4}
	ar_sec := arr3[1:]
	t.Log(ar_sec)

	ar_sec2 := arr3[:] //这里返回的其实是一个切片
	t.Log(ar_sec2)
}

// 切片 : 共享存储结构
//	长度已知的情况可以用数组。
//	切片的方便之处主要来自于自动增长
//	但切片的自动增长会导致内存分配和数据复制，以及未来相关的GC开销。

// cap存在的意义是什么？
//cap是指后面对应的连续存储空间的大小，你可以向slice中append元素，没有超过cap及不会引发slice扩容。
func TestSlice(t *testing.T) {

	var s0 []int

	t.Log(len(s0), cap(s0)) // 0 0  // len 初始化元素个数。和最大长度

	s0 = append(s0, 1)

	t.Log(len(s0), cap(s0))

	s1 := make([]int, 3, 7)

	t.Log(len(s1), cap(s1))

	t.Log(s1[0], s1[1], s1[2])

	s1 = append(s1, 1)
	t.Log(s1[0], s1[1], s1[2], s1[3])

}

// 切片中的共享存储
func TestSliceShareMemory(t *testing.T) {

	year := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S"}

	slice_01 := year[3:7] //D E F G

	t.Log(slice_01, "长度:", len(slice_01), "容量:", cap(slice_01))

	slice_02 := year[5:11] //F G H I J K

	t.Log(slice_02, "长度:", len(slice_02), "容量:", cap(slice_02))

	slice_02[0] = "OJBK"

	t.Log("slice_02修改后的 slice_01信息", slice_01, "长度:", len(slice_01), "容量:", cap(slice_01))
}

//通过 切片变量 和 nil的比较，来判断切片变量是否被初始化了

func TestSliceInit(t *testing.T) {
	//var a []int
	//var b = make([]int,0,0)
	//c := []int{}
	//t.Log(a, len(a), cap(a))
	//t.Log(b, len(b), cap(b))
	//t.Log(c, len(c), cap(c))
	//t.Log(a==nil, b==nil, c==nil)

	a := []int{1, 2, 3, 4, 5, 6}
	a = append(a, 3)
	t.Log(a, len(a), cap(a))
}

//可变长参数

func Test_arg(t *testing.T) {
	Sum_arg(1, 2, 3, 4)
}
func Sum_arg(ok ...int) {

	for _, v := range ok {
		fmt.Println(v)
	}
}

//延迟执行
func Test_defer(t *testing.T) {
	defer func() {
		t.Log("Sssss")
	}()

	t.Log("1")
}

type Employee struct {
}

type First_Object struct {
	id             string
	age            string
	String_pointer func() string
}

func Test_object(t *testing.T) {

	e := First_Object{"1", "2", func() string {
		return "method in struct"
	}}

	t.Logf("e 的内存地址: %x", unsafe.Pointer(&e.age))
	t.Log(e.String_pointer())
	//t.Log("*********************")
	//e.String_value()

}

// pointer
//func (test_struct First_Object) String_pointer() string {
//	fmt.Printf("test_struct pointer 的内存地址: %x\n", unsafe.Pointer(&test_struct.age))
//	return "pointer is running"
//}

//func (test_struct First_Object) String_value() string {
//	fmt.Printf("test_struct value 的内存地址: %x\n", unsafe.Pointer(&test_struct))
//	return "value is running"
//}

func TestMap(t *testing.T) {

	map01 := map[string]string{} //默认空串儿

	t.Log(map01["1"])

}

//golang 异常处理 error
// go 语言没有异常机制，只有自定义error

func TestGroutine(t *testing.T) {

	counter := 0
	for i := 0; i < 10000; i++ {
		go func() {
			counter++
		}()
	}
	time.Sleep(time.Second * 10)
	fmt.Println(counter)

}

type uninit struct {
	age int
}

func TestCounterWaitGroup(t *testing.T) {
	var wg sync.WaitGroup
	counter := 0
	for i := 0; i < 5000000; i++ {
		wg.Add(1) // 每启动一个协程增加一个等待
		go func() {
			counter++
			wg.Done()
		}()
		wg.Wait()
	}
	t.Logf("counter = %d", counter)
}

//阻塞式 channal
func TestChannal(t *testing.T) {
	select {
	case chan01 := <-AsynTask(): //这个地方理论上除了方法栈的切换耗时 无其他阻塞操作
		t.Log(chan01)
	case <-time.After(time.Second * 1): //这个地方会阻塞1s
		t.Log("service 执行超时了")
	default: //无阻塞，所以目前 是在 case1 和 default 之间切换
		t.Log("default executed")
	}
}
func service() string {
	fmt.Println("业务方法执行中 3s。。。")
	time.Sleep(time.Second * 3)
	fmt.Println("业务方法执行完毕 。。。")
	return "done"
}
func AsynTask() chan string {
	chan01 := make(chan string)
	go func() {
		time.Sleep(time.Second * 1)
		chan01 <- "s" //会阻塞这里等待有人把通道里的东西取走
		fmt.Println("测试阻塞")
	}()
	fmt.Println("协程正在运行但是还没输出结果。现在程序return channal了")
	return chan01
}

//golang全都是值传递
func TestChannleClose(t *testing.T) {
	var group sync.WaitGroup
	ch := make(chan int, 10)
	fmt.Printf("main goroutine chan int address: %x\n", unsafe.Pointer(&ch))
	fmt.Printf("main goroutine waitGroup address: %x\n", unsafe.Pointer(&group))
	group.Add(1)
	go func(ch1 chan int, wg *sync.WaitGroup) {
		fmt.Printf("producer chan int address: %x\n", unsafe.Pointer(&ch1))
		fmt.Printf("producer goroutine waitGroup address: %x\n", unsafe.Pointer(wg))
		for i := 1; i < 11; i++ {
			ch1 <- i
		}
		close(ch1)
		wg.Done()
	}(ch, &group)
	group.Add(1)
	go func(ch1 *chan int, wg *sync.WaitGroup) {
		fmt.Printf("consumer chan int address: %x\n", unsafe.Pointer(ch1))
		fmt.Printf("consumer goroutine waitGroup address: %x\n", unsafe.Pointer(wg))
		for {
			if v, ok := <-*ch1; ok {
				fmt.Println(v)
			} else {
				break
			}
		}
		wg.Done()
	}(&ch, &group)
	group.Wait()
}

// 猜测，channal close 操作看起来也是异步的
// 当buffer channal中的数据尚未被读取完时，channal是无法关闭的
func pruducer(c chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
		fmt.Println("关闭通道")
		wg.Done()
	}()
}

func consumer(c chan int, w *sync.WaitGroup) {
	go func() {
		for {
			time.Sleep(time.Second * 3)
			if i, ok := <-c; ok {
				fmt.Println(i)
			} else {
				break
			}
		}
		w.Done()
	}()
}

func TestChannelClose(t *testing.T) {
	wg := &sync.WaitGroup{}
	c := make(chan int, 10)
	wg.Add(2)
	pruducer(c, wg)
	consumer(c, wg)
	wg.Wait()
}

// waitGroup 值传递测试
// 发布者
func publisher(ch chan int) {
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()
}

// 接受者
func receiver(ch chan int, wg *sync.WaitGroup) {
	go func() {
		//for {
		for i := 0; i < 10; i++ {
			fmt.Println(<-ch)
		}

		wg.Done()
		fmt.Println("消费完了")
	}()
}

func TestPubRe(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int)
	publisher(ch)
	wg.Add(1)
	receiver(ch, &wg)
	wg.Wait()
	fmt.Println("主进程完了")
}

//******************************************************
//利用channel 的关闭时的广播 实现 任务的取消
func CancelTask01(chan01 chan struct{}) chan struct{} {
	chan01 <- struct{}{}
	return chan01
}
func CancelTask02(chan01 chan struct{}) {
	close(chan01)
}

func mainTask(ctx01 ctx.Context) bool {
	select {
	case <-ctx01.Done():
		fmt.Println("true 了")
		return true
	default:
		fmt.Println("false 了")
		return false
	}
}

func TestCancelTaskChannel(t *testing.T) {
	ctx01, cancel := ctx.WithCancel(ctx.Background()) // 获取根context ,通过context 机制实现消息通知,取消任务

	for i := 0; i < 5; i++ {
		go func(i int, ctx ctx.Context) {
			for {
				if mainTask(ctx) {
					fmt.Println(i, "Done 了")
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
		}(i, ctx01)
	}
	cancel()
	time.Sleep(time.Second * 1)
}

// close channel 时 等待接受消息的阻塞中的 channel会接收到什么呢？
func TestCloseChannel(t *testing.T) {
	chan01 := make(chan string)

	go func() {
		if _, ok := <-chan01; ok { //这里会阻塞 直到有人塞值进去
			fmt.Println("ok")
		} else {
			fmt.Println("not ok")
		}
	}()
	fmt.Println("先睡1s")
	time.Sleep(time.Second * 1)
	fmt.Println("此时关闭channel")
	close(chan01)

	//chan01 <- "strng"
	time.Sleep(time.Second * 1)

}
