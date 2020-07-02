package main

import (
	"fmt"
	"sync"
)

var (
	balance int
	mutex   sync.Mutex
)

func init() {
	balance = 1000
}

func depositMoney(value int, status chan bool) {
	mutex.Lock()
	balance += value
	fmt.Printf("Deposited amount is %d\n", value)
	status <- true
	mutex.Unlock()
}

func withdrawMoney(value int, status chan bool) {
	mutex.Lock()
	balance -= value
	fmt.Printf("Withdrawn amount is %d\n", value)
	status <- true
	mutex.Unlock()
}

func main() {
	fmt.Println("Go SYNC MUTEX example")
	done := make(chan bool)
	go depositMoney(700, done)
	go withdrawMoney(400, done)
	<-done
	<-done
	fmt.Println("Current balance is : ", balance)
}
