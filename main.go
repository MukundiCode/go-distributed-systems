package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

// Both wrappers embed io.Writer, so until you give them their own Write
// they pass text through untouched.
type CountingWriter struct {
	io.Writer
	Bytes int
	Lines int
}

type UpperWriter struct {
	io.Writer
}

func (c *CountingWriter) Write(p []byte) (int, error) {
	c.Bytes += len(p)
	c.Lines += bytes.Count(p, []byte("\n"))

	return c.Writer.Write(p)
}

func (u UpperWriter) Write(p []byte) (int, error) {
	upperBytes := bytes.ToUpper(p)

	n, err := u.Writer.Write(upperBytes)
	if err != nil {
		return 0, err
	}

	if n != len(upperBytes) {
		return 0, io.ErrShortWrite
	}

	return len(p), nil
}

// TODO: Write on CountingWriter
// TODO: Write on UpperWriter

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 1024*1024), 1024*1024)
	sc.Scan()
	mode := sc.Text()
	var lines []string
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}

	var w io.Writer = os.Stdout
	var cw *CountingWriter
	switch mode {
	case "upper":
		w = UpperWriter{Writer: w}
	case "count":
		cw = &CountingWriter{Writer: w}
		w = cw
	case "upper+count":
		cw = &CountingWriter{Writer: w}
		w = UpperWriter{Writer: cw}
	case "count+upper":
		cw = &CountingWriter{Writer: UpperWriter{Writer: w}}
		w = cw
	}

	for _, l := range lines {
		fmt.Fprintln(w, l)
	}
	if cw != nil {
		fmt.Printf("bytes=%d lines=%d\n", cw.Bytes, cw.Lines)
	}
}
