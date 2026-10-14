package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// TODO: the Shape interface (Area() int and Name() string), the three
// shape types, their String methods, and the sortable slice type.

// num returns word i of a split line as an int (0 if the word is missing).
func num(f []string, i int) int {
	if i >= len(f) {
		return 0
	}
	n, _ := strconv.Atoi(f[i])
	return n
}

// keeps the sort import alive until you use it
var _ = sort.Stable

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 1024*1024), 1024*1024)
	var lines [][]string // each input line split into words: {"rect", "3", "4"}
	for sc.Scan() {
		f := strings.Fields(sc.Text())
		if len(f) > 0 {
			lines = append(lines, f)
		}
	}

	// TODO: turn each line into a Shape (or remember its unknown kind),
	// sort, print the shapes, the count line, then the unknown kinds.
	_ = lines
	_ = num
	fmt.Println("TODO")
}
