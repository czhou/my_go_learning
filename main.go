package main

import (
	"fmt"
	"time"
)

func main() {
	startTime := time.Now();
	resultChan := make(chan int, 2);
	go sum(1, 100000000, resultChan);
	//go sum(1, 50000000, resultChan);
	//go sum(50000001, 100000000, resultChan);
	//var sum int;
	//for i := 0; i <= 10000; i++{
	//	sum += i;
	//}
	//result1, result2 := <- resultChan, <- resultChan;
	result1 := <- resultChan;
	endTime := time.Now();
	fmt.Println(startTime, endTime);
	fmt.Println(endTime.Sub(startTime));
	//fmt.Println(result1, result2, result1 + result2);
	fmt.Println(result1);
}

func sum(start int, end int, resultChan chan int){
	var sum int;
	for i := start; i <= end; i++{
		sum += i;
	}
	resultChan <- sum;
}