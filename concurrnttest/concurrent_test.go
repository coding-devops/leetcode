package main

import (
	"errors"
	"fmt"
	"github.com/go-basic/uuid"
	"runtime"
	"sync"
	"testing"
	"time"
	"unsafe"
)

/**
这个包里面包含了一些并发相关的case
*/

/**
单例模式 只运行一次
*/
func Test_case01(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			obj := RunOnce()
			fmt.Printf("对象地址%x", unsafe.Pointer(obj))
			fmt.Println()
			wg.Done()
		}(&wg)
	}
	time.Sleep(time.Second * 3)
}

var once sync.Once

type SinglePartton struct {
	a int
}

var single *SinglePartton

func RunOnce() *SinglePartton {
	uuid := uuid.New()

	fmt.Println(uuid)
	once.Do(func() {
		fmt.Print("create once")
		single = new(SinglePartton)
	})
	return single
}

/**
仅需要任意任务完成即可

*/

func Test_case02(t *testing.T) {
	fmt.Println("num of groutine before run", runtime.NumGoroutine())
	ResponseAll()
	fmt.Println("num of groutine after run", runtime.NumGoroutine())
}

func ResponseOne() string {
	chan_ := make(chan string, 10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			chan_ <- WhateverOneDone(i)
		}(i)
	}
	return <-chan_
}
func WhateverOneDone(i int) string {
	time.Sleep(time.Second * 1)
	return fmt.Sprintf("%d 号任务返回", i)
}

func ResponseAll() {
	chan_ := make(chan string, 10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			chan_ <- WhateverOneDone(i)
		}(i)
	}
	responser := ""
	for j := 0; j < 10; j++ {
		responser = <-chan_
		fmt.Println(responser)
	}
}

/****
对象池
 buffered channel 实现
*/
type Object struct {
}
type ObjPool struct {
	bufferpool chan *Object
}

func Test_obj_pool(t *testing.T) {
	objpool := NewObjPool() // 初始化对象池
	var obj *Object
	var err error
	for {
		if obj, err = objpool.GetObjFromPool(time.Second * 1); err != nil {
			fmt.Printf(err.Error())
			break
		} else {
			fmt.Printf("%T ", obj)
			fmt.Println("取出了一个对象")
		}
	}
	for {
		objpool.PutObjBackToPool(obj)
	}

}

//初始化 对象池子
func NewObjPool() *ObjPool {
	objPool := &ObjPool{}
	objPool.bufferpool = make(chan *Object, 5)
	for i := 0; i < 5; i++ {
		objPool.bufferpool <- new(Object) //循环向结构里塞5个对象数据
	}
	return objPool

}

//从池子取对象 看看取多了会咋办
func (p *ObjPool) GetObjFromPool(timeout time.Duration) (*Object, error) {
	select {
	case obj := <-p.bufferpool:

		return obj, nil
	case <-time.After(timeout):
		fmt.Println("channel长度：", len(p.bufferpool))
		return nil, errors.New("取对象超时，对象池可能已经空了")
	}
}

func (p *ObjPool) PutObjBackToPool(obj *Object) (int, error) {
	select {
	case p.bufferpool <- obj:
		fmt.Println("放入channel中")
		return len(p.bufferpool), errors.New("资源使用完毕 已放回对象池")
	default:
		fmt.Println("满了")
		return len(p.bufferpool), errors.New("put obj back to pool is failed")
	}
}

// 往池子里放数据时候 报错了怎么办
