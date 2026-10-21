package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// serve prints one item. An item that starts with "+" also queues the rest
// of its text as a brand-new urgent item.
func serve(x string, urgent chan<- string) {
	fmt.Println("served:", x)
	if strings.HasPrefix(x, "+") {
		urgent <- x[1:]
	}
}

// drain serves every item: urgent before normal before low, re-checking
// urgent before each pick. normal and low are closed once loaded; urgent
// is never closed. Return when nothing is left.
func drain(urgent chan string, normal, low <-chan string) {

	for urgent != nil {
		select {
		case u := <-urgent:
			serve(u, urgent)
			continue
		default:
			if normal == nil && low == nil {
				urgent = nil
			}

		}

		select {
		case u, ok := <-normal:
			if !ok {
				normal = nil
				break
			}
			serve(u, urgent)
			continue
		default:
		}

		select {
		case u, ok := <-low:
			if !ok {
				low = nil
				break
			}
			serve(u, urgent)
		default:

		}
	}

}

// awaitLate prints "late: <x>" if a value arrives on late within 100ms,
// otherwise "idle".
func awaitLate(late <-chan string) {
	select {
	case l := <-late:
		fmt.Println("late: " + l)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("idle")
	}
}

func main() {
	var lines []string
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		if line := strings.TrimSpace(sc.Text()); line != "" {
			lines = append(lines, line)
		}
	}
	urgent := make(chan string, 1024)
	normal := make(chan string, len(lines)+1)
	low := make(chan string, len(lines)+1)
	delay, lateVal, hasLate := 0, "", false
	for _, line := range lines {
		f := strings.Fields(line)
		switch f[0] {
		case "U":
			urgent <- f[1]
		case "N":
			normal <- f[1]
		case "L":
			low <- f[1]
		case "late":
			delay, _ = strconv.Atoi(f[1])
			lateVal, hasLate = f[2], true
		}
	}
	close(normal)
	close(low)

	drain(urgent, normal, low)

	late := make(chan string, 1)
	if hasLate {
		go func() {
			time.Sleep(time.Duration(delay) * time.Millisecond)
			late <- lateVal
		}()
	}
	awaitLate(late)
}
