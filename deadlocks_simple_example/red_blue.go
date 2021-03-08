package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	lock1 = sync.Mutex{}
	lock2 = sync.Mutex{}
)

func blueRobot() {
	for {
		fmt.Println("Blue: Acquiring lock1")
		lock1.Lock()
		fmt.Println("Blue: Acquiring lock2")
		lock2.Lock()
		fmt.Println("Blue: Both locks Acquired")
		lock1.Unlock()
		lock2.Unlock()
		fmt.Println("Blue: Locks have been released")
	}
}

func redRobot() {
	for {
		fmt.Println("Red:  Acquiring lock1")
		lock1.Lock()
		fmt.Println("Red:  Acquiring lock2")
		lock2.Lock()
		fmt.Println("Red:  Both locks Acquired")
		lock1.Unlock()
		lock2.Unlock()
		fmt.Println("Red:  Locks have been released")
	}
}

func main() {
	go blueRobot()
	go redRobot()
	time.Sleep(20 * time.Second)
	fmt.Println("Done")
}
