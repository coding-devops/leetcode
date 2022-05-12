package main

import (
	"fmt"
	"time"
)

//func main() {
//	fmt.Println(os.Hostname())
//	conn, _ := net.Dial("udp4", "1.2.3.4:56")
//	fmt.Println(conn.LocalAddr().String())
//	//ticker := time.NewTicker(time.Duration(3) * time.Second)
//	//for range ticker.C {
//	//	fmt.Println("sss")
//	//}
//	//t := time.Second * 5
//	//d := &Defender{tick: time.NewTicker(t)}
//	//err := d.Execute()
//	//if err != nil {
//	//	return
//	//}
//}

type Defender struct {
	tick      *time.Ticker
	Round     int
	rootPath  string
	rpcServer string
	healthKey string
	hostName  string
	times     int
}

func (d *Defender) Execute() error {
	for {
		select {
		case <-d.tick.C:
			fmt.Println("1")
		}
	}
}

type Ticker struct {
	C <-chan time.Time // The channel on which the ticks are delivered.
}
