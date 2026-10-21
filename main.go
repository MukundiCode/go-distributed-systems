package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

// Counter holds the word counts that every goroutine updates.
type Counter struct {
	mu     sync.Mutex
	counts map[string]int
}

var stopSet map[string]bool

// stopBuilds records how many times buildStopSet ran.
var stopBuilds int

var once sync.Once

// buildStopSet is the setup work. Leave it exactly as it is.
func buildStopSet() map[string]bool {
	stopBuilds++
	set := make(map[string]bool)
	for _, w := range []string{"a", "an", "and", "in", "of", "on", "or", "the", "to"} {
		set[w] = true
	}
	return set
}

// TODO: stopWords() returns the stop set, calling buildStopSet only the
// first time any goroutine asks, however many ask at once.

// TODO: Counter.Add(word) adds 1 to that word's count, safely. Call it once
// per count: R calls for each occurrence of a word.
func (c *Counter) Add(word string, r int) {
	c.mu.Lock()
	defer c.mu.Unlock()

	once.Do(func() {
		stopSet = buildStopSet()
	})
	_, isStop := stopSet[word]
	if isStop {
		return
	}
	c.counts[word] += 1 * r
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 1024*1024), 1024*1024) // lets the scanner read lines longer than 64KB
	sc.Scan()
	r, _ := strconv.Atoi(strings.TrimSpace(sc.Text()))
	var lines []string
	for sc.Scan() {
		if line := strings.TrimSpace(sc.Text()); line != "" {
			lines = append(lines, line)
		}
	}

	c := &Counter{counts: make(map[string]int)}
	var wg sync.WaitGroup

	// TODO: one goroutine per line. Each skips the stop words and adds every
	// other word of its line r times into c. Then wait for all of them.
	for _, line := range lines {
		wg.Add(1)
		go func() {
			for _, v := range strings.Fields(line) {
				c.Add(v, r)
			}
			wg.Done()
		}()
	}

	_ = r
	_ = lines
	wg.Wait()

	words := make([]string, 0, len(c.counts))
	for w := range c.counts {
		words = append(words, w)
	}
	sort.Slice(words, func(i, j int) bool {
		if c.counts[words[i]] != c.counts[words[j]] {
			return c.counts[words[i]] > c.counts[words[j]]
		}
		return words[i] < words[j]
	})
	for _, w := range words {
		fmt.Println(w, c.counts[w])
	}
	fmt.Println("stopwords built:", stopBuilds)
}
