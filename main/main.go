package main

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
	"sync"
)

var mgroup = make(map[uint64]int)
var mutex sync.Mutex

func Trace() func() {

	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		fmt.Println("报错")
	}
	funcccc := runtime.FuncForPC(pc)
	funname := funcccc.Name()
	gid := curGoroutineID()

	mutex.Lock()
	index := mgroup[gid]
	mgroup[gid] = index + 1
	mutex.Unlock()
	s := ""
	for i := 0; i <= index; i++ {
		s = s + "    "
	}
	fmt.Printf("g[%05d]:%s-> enter:%s\n", gid, s, funname)
	return func() {
		fmt.Printf("g[%05d]:%s<- exit :%s\n", gid, s, funname)
	}
}

// trace2/trace.go
var goroutineSpace = []byte("goroutine ")

// 获取当前goroutine id
func curGoroutineID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	// Parse the 4707 out of "goroutine 4707 ["
	b = bytes.TrimPrefix(b, goroutineSpace)
	i := bytes.IndexByte(b, ' ')
	if i < 0 {
		panic(fmt.Sprintf("No space found in %q", b))
	}
	b = b[:i]
	n, err := strconv.ParseUint(string(b), 10, 64)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse goroutine ID out of %q: %v", b, err))
	}
	return n
}
