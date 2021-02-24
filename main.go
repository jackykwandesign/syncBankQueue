package main

import (
	"fmt"
	"time"

	"jacky.com/lockqueue/bankqueue"
)

var balance int = 10000000
var waitTime int = 0
var totalTx int = 0

func Deposit(num int, tx string) {
	bq := bankqueue.GetInstance()
	bq.Start()
	currentBalance := balance
	fmt.Println(fmt.Sprintf("%s: Before My balance have %d", tx, currentBalance))
	if currentBalance < num {
		fmt.Println(fmt.Sprintf("%s: Reject! Balance %d, can't deposit %d", tx, currentBalance, num))
		return
	} else {
		currentBalance -= num
		fmt.Println(fmt.Sprintf("%s: Accept! Deposit %d, Remain %d", tx, num, currentBalance))
		balance = currentBalance
	}
	time.Sleep(time.Duration(33333) * time.Nanosecond) //simulate db insert time
	// postgresql 30K when 200 million data
	// https://blog.timescale.com/blog/timescaledb-vs-6a696248104e/
	bq.End()
}

func testA(workerID int) {
	for {
		start := time.Now()
		Deposit(1000, fmt.Sprintf("%d", workerID))
		duration := time.Since(start)
		fmt.Println(duration.Nanoseconds())
		waitTime = waitTime + int(duration.Nanoseconds())
		totalTx++
		fmt.Println(fmt.Sprintf("Total waiting time: %v, Total Tx %d, Average Wait: %v ms", waitTime, totalTx, (waitTime/totalTx)/1000000))
		// time.Sleep(time.Duration(workerID*100) * time.Millisecond)
	}
}

// func testB() {
// 	for {
// 		Deposit(1000, "B")
// 		time.Sleep(1 * time.Millisecond)
// 	}
// }
func genWorker() {
	for i := 0; i < 100; i++ {
		go testA(i)
	}
}
func main() {

	genWorker()
	select {}

}
