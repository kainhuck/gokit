//package main
//
//import (
//	"fmt"
//  "github.com/kainhuck/gokit/sync"
//)
//
//func main() {
//	var w sync.WaitGroupWrapper
//
//	w.Wrap(func() {
//		fmt.Println("foo")
//	})
//
//	w.Wrap(func() {
//		fmt.Println("bar")
//	})
//
//	w.Wait()
//
//}


package sync

import "sync"

type WaitGroupWrapper struct {
	sync.WaitGroup
}

func (wg *WaitGroupWrapper) Wrap(f func()) {
	wg.Add(1)
	go func() {
		f()
		wg.Done()
	}()
}
