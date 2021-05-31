package factorial

import (
	"fmt"
	"time"
)

func Factorial() {
	defer elapsed()()
	number := 6
	//workerCount := 1

	jobs := make(chan int, number)
	result := make(chan int, number)

	//for i := 0; i < workerCount; i++ {
	//	go worker(jobs, result)
	//}

	go worker(jobs, result)
	//go worker(jobs, result)
	//go worker(jobs, result)
	//go worker(jobs, result)
	//go worker(jobs, result)
	//go worker(jobs, result)

	for i := 0; i <= number; i++ {
		jobs <- i
		//fmt.Println(i)
	}
	close(jobs)

	//for res := range result{
	//	fmt.Println(i)
	//	fmt.Println(res)
	//}

	for i := len(result); i >= 0; i-- {
		fmt.Println(<-result)
	}
}

func worker(jobs <-chan int, result chan<- int) {
	for job := range jobs {
		result <- factorial(job)
	}
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func elapsed() func() {
	start := time.Now()
	return func() {
		fmt.Printf("Calculation took %v\n", time.Since(start))
	}
}
