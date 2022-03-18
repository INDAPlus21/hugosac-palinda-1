package main

import "fmt"

// sum the numbers in a and send the result on res.
func sum(a []int, res chan<- int) {
	s := 0
	for _, n := range a {
		s += n
	}
	res <- s
}

// concurrently sum the array a.
func ConcurrentSum(a []int) int {
	n := len(a)
	ch := make(chan int)
	go sum(a[:n/2], ch)
	go sum(a[n/2:], ch)

	return <-ch + <-ch
}

func main() {
	// example call
	a := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(ConcurrentSum(a))
}
