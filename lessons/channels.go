package lessons

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

// collatzSteps counts how many steps n needs to reach 1.
func collatzSteps(n int) int {
	steps := 0
	for n != 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
		steps++
	}
	return steps
}

func main() {
	in := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(in, &n)
	nums := make([]int, n)
	for i := range nums {
		fmt.Fscan(in, &nums[i])
	}

	results := make([]int, n)
	var wg sync.WaitGroup
	for i, v := range nums {
		wg.Add(1)
		idx := i
		go func(idx int) {
			res := collatzSteps(v)
			results[idx] = res
			wg.Done()
		}(idx)
	}
	wg.Wait()
	// TODO: wait for every goroutine here, before anything below runs

	total := 0
	for _, r := range results {
		fmt.Println(r)
		total += r
	}
	fmt.Printf("total: %d\n", total)
}
