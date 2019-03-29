package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type Ban struct {
	visitIPs map[string]struct{}
}

func NewBan() *Ban {
	return &Ban{visitIPs: make(map[string]struct{})}
}

//判断ip是否存在
func (o *Ban) visit(ip string) bool {
	mapMutex.Lock()
	defer mapMutex.Unlock()
	if _, ok := o.visitIPs[ip]; ok {
		return true
	}
	o.visitIPs[ip] = struct{}{}
	go o.invalidAfter1Min(ip)
	return false
}
func (o *Ban) invalidAfter1Min(ip string) {
	time.Sleep(time.Minute * 1)
	mapMutex.Lock()
	visitIPs := o.visitIPs
	delete(visitIPs, ip)
	o.visitIPs = visitIPs
	mapMutex.Unlock()
}

var mapMutex *sync.Mutex //声明锁

func main() {
	mapMutex = new(sync.Mutex) //创建一个锁
	var success int64          //声明
	ban := NewBan()
	wg := new(sync.WaitGroup)
	for i := 0; i < 1000; i++ {
		for j := 0; j < 100; j++ {
			wg.Add(1)
			ipEnd := j
			go func() {
				defer wg.Done()
				ip := fmt.Sprintf("192.168.1.%d", ipEnd)
				fmt.Println(ip)
				if !ban.visit(ip) {
					atomic.AddInt64(&success, 1)
				}
			}()
		}
	}
	wg.Wait()
	fmt.Println("success:", success)

}
