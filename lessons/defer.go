package lessons

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
)

// perform does the work of one task. It is given; it has no cleanup of its own.
func perform(f []string) string {
	switch f[0] {
	case "div":
		a, _ := strconv.Atoi(f[1])
		b, _ := strconv.Atoi(f[2])
		return strconv.Itoa(a / b)
	case "at":
		i, _ := strconv.Atoi(f[1])
		vals := f[2:]
		return vals[i]
	case "set":
		var m map[string]string
		m[f[1]] = f[2]
		return "stored"
	case "boom":
		panic(strings.Join(f[1:], " "))
	case "fail":
		panic(errors.New(strings.Join(f[1:], " ")))
	}
	return "fine"
}

func run(f []string) (result string, err error) {
	// TODO: before the work starts, register a deferred function that
	// recovers a panic and sets err from its type: runtime.Error -> "runtime",
	// any other error -> "error " + its message, string -> "string " + the string.
	defer func() {
		if r := recover(); r != nil {
			switch e := r.(type) {
			case runtime.Error:
				err = errors.New("runtime")
			case error:
				err = errors.New("error " + e.Error())
			case string:
				err = errors.New("string " + e)
			}
		}
	}()
	fmt.Println("open A")
	// TODO: make "close A" print when run returns, on every path.
	defer fmt.Println("close A")
	fmt.Println("open B")
	// TODO: make "close B" print when run returns, on every path.
	defer fmt.Println("close B")
	return perform(f), err
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	okCount, failed := 0, 0
	for sc.Scan() {
		f := strings.Fields(sc.Text())
		if len(f) == 0 {
			continue
		}
		res, err := run(f)
		if err != nil {
			fmt.Printf("error: %s\n", err)
			failed++
		} else {
			fmt.Printf("result: %s\n", res)
			okCount++
		}
	}
	fmt.Printf("ok=%d failed=%d\n", okCount, failed)
}
