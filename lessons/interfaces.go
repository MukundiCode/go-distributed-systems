package lessons

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
type Rect struct {
	w   int
	h   int
	idx int
}

func (c Rect) Area() int {
	return c.w * c.h
}

func (c Rect) Name() string {
	return "rect"
}

func (c Rect) String() string {
	return fmt.Sprintf("rect %dx%d area=%d", c.w, c.h, c.Area())
}

func (c Rect) Idx() int {
	return c.idx
}

type Square struct {
	s   int
	idx int
}

func (c Square) Area() int {
	return c.s * c.s
}

func (c Square) Name() string {
	return "square"
}

func (c Square) String() string {
	return fmt.Sprintf("square %d area=%d", c.s, c.Area())
}

func (c Square) Idx() int {
	return c.idx
}

type Tri struct {
	b   int
	h   int
	idx int
}

func (c Tri) Area() int {
	return (c.b * c.h) / 2
}

func (c Tri) Name() string {
	return "tri"
}

func (c Tri) String() string {
	return fmt.Sprintf("tri %dx%d area=%d", c.b, c.h, c.Area())
}

func (c Tri) Idx() int {
	return c.idx
}

type Unknown struct {
	name string
	idx  int
}

func (Unknown) Area() int {
	return 0
}

func (c Unknown) Name() string {
	return c.name
}

func (c Unknown) String() string {
	return fmt.Sprintf("unknown: " + c.name)
}

func (c Unknown) Idx() int {
	return c.idx
}

type Shape interface {
	Area() int
	Name() string
	Idx() int
}

type byShape []Shape

func (s byShape) Len() int { return len(s) }
func (s byShape) Less(i, j int) bool {
	_, ok := s[j].(Unknown)
	if ok {
		_, ok = s[i].(Unknown)
		if ok {
			return s[i].Idx() > s[j].Idx()
		}
	}
	if s[i].Area() == s[j].Area() {
		if s[i].Name() == s[j].Name() {
			return s[i].Idx() > s[j].Idx()
		}
		return s[i].Name() > s[j].Name()
	}
	return s[i].Area() < s[j].Area()
}
func (s byShape) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

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
	num := len(lines)
	shapes := make([]Shape, num)
	rectCount, sqCount, triCount := 0, 0, 0

	for i, val := range lines {
		name := val[0]
		switch name {
		case "rect":
			w, _ := strconv.Atoi(val[1])
			h, _ := strconv.Atoi(val[2])
			shapes[i] = Rect{w, h, i}
			rectCount++
		case "square":
			s, _ := strconv.Atoi(val[1])
			shapes[i] = Square{s, i}
			sqCount++
		case "tri":
			b, _ := strconv.Atoi(val[1])
			h, _ := strconv.Atoi(val[2])
			shapes[i] = Tri{b, h, i}
			triCount++
		default:
			shapes[i] = Unknown{name, i}
		}

	}

	sort.Sort(byShape(shapes))

	shouldPrintTotals := true
	for i := len(shapes) - 1; i >= 0; i-- {
		switch shapes[i].(type) {
		case Rect, Tri, Square:
			fmt.Println(shapes[i])
		case Unknown:
			if shouldPrintTotals {
				fmt.Printf("rects=%d squares=%d tris=%d\n", rectCount, sqCount, triCount)
			}
			shouldPrintTotals = false
			fmt.Println(shapes[i])
		}
	}
	if shouldPrintTotals {
		fmt.Printf("rects=%d squares=%d tris=%d\n", rectCount, sqCount, triCount)
	}
}
