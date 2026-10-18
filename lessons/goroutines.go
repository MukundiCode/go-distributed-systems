package lessons

import "fmt"

// generate must send 2..limit on out, then close it.
func generate(out chan<- int, limit int) {
	for i := 2; i <= limit; i++ {
		out <- i
	}
	close(out)
}

// filter must forward every number that p does not divide, then close out
// once in is closed.
func filter(in <-chan int, out chan<- int, p int) {
	for i := range in {
		if i%p != 0 {
			out <- i
		}
	}
	close(out)
}

func main() {
	var limit int
	fmt.Scan(&limit)

	head := make(chan int)
	go func() {
		generate(head, limit)
	}()

	count := 0
	// TODO: take the next prime from head and print it, count it, start a
	// filter for it that reads from head, and make the filter's output the
	// new head. Stop when head is closed.
	i, ok := <-head
	for ok {
		count++
		fmt.Println(i)
		newHead := make(chan int)
		go filter(head, newHead, i)
		head = newHead
		i, ok = <-head
	}

	fmt.Println("count:", count)
}
