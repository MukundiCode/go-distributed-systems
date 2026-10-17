package lessons

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// A blank (or whitespace-only) line.
var ErrEmpty = errors.New("empty line")

// A non-blank line that does not have exactly three fields.
var ErrFields = errors.New("wrong field count")

// An age or score outside its allowed range.
type RangeError struct {
	Field    string
	Min, Max int
	Got      int
}

func (e RangeError) Error() string {
	return fmt.Sprintf("range: %s %d not in %d..%d", e.Field, e.Got, e.Min, e.Max)
}

type ParseError struct {
	word string
}

func (e ParseError) Error() string {
	return fmt.Sprintf("parse: %s", e.word)
}

// check validates one line "name age score" and returns the name, or the FIRST
// problem found: ErrEmpty for a blank line, ErrFields for a line without exactly
// three fields, then age (parse, then range 0..120), then score (parse, then
// range 0..100).
func check(line string) (string, error) {
	// TODO: use strings.Fields and strconv.Atoi; wrap parse failures with %w
	split := strings.Fields(line)
	// check if empty
	if len(split) == 0 {
		return "", ErrEmpty
	}
	// check if fields is 3
	if len(split) != 3 {
		return "", ErrFields
	}
	// check if i1 and i2 parse
	age, err := strconv.Atoi(split[1])
	if err != nil {
		return "", ParseError{split[1]}
	}
	// check if age in range
	if age < 0 || age > 120 {
		return "", RangeError{"age", 0, 120, age}
	}

	score, err := strconv.Atoi(split[2])
	if err != nil {
		return "", ParseError{split[2]}
	}

	if score < 0 || score > 100 {
		return "", RangeError{"score", 0, 100, score}
	}

	return split[0], nil
}

// validate adds the line number. Do not change it.
func validate(n int, line string) (string, error) {
	name, err := check(line)
	if err != nil {
		return "", fmt.Errorf("line %d: %w", n, err)
	}
	return name, nil
}

// classify turns an error from validate into the line to print:
// "empty", "fields", "parse: TEXT", or "range: ...".
func classify(err error) string {
	var re RangeError
	if errors.As(err, &re) {
		return re.Error()
	}

	var pe ParseError
	if errors.As(err, &pe) {
		return pe.Error()
	}

	if errors.Is(err, ErrEmpty) {
		return "empty"
	}

	if errors.Is(err, ErrFields) {
		return "fields"
	}
	return fmt.Sprintf("%v", err)
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	n := 0
	for sc.Scan() {
		n++
		name, err := validate(n, sc.Text())
		if err != nil {
			fmt.Println(classify(err))
		} else {
			fmt.Println("ok " + name)
		}
	}
}
