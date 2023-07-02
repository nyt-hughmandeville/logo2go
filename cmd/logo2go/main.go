package main

import (
	"fmt"
	"os"
	"strings"
)

const eof = -1

type lexer struct {
	input string // the string being scanned.
	start int    // start position of this item.
	pos   int    // current position in the input.
	width int    // width of last rune read from input.
	out   string
}

func (l *lexer) run() {
	for state := lexProcedure; state != nil; {
		state = state(l)
	}
	l.out += "}\n"
}

// Check if current position is at a space.
func (l *lexer) isSpace() bool {
	return l.input[l.pos] == ' ' || l.input[l.pos] == '\t' || l.input[l.pos] == '\n'
}

func (l *lexer) ignore() {
	l.start = l.pos
}

func (l *lexer) next() rune {
	if l.pos >= len(l.input) {
		l.width = 0
		return eof
	}
	r, w := rune(l.input[l.pos]), 1
	l.width = w
	l.pos += w
	return r
}

// stateFn represents the state of the scanner
// as a function that returns the next state.
type stateFn func(*lexer) stateFn

func lexProcedure(l *lexer) stateFn {
	if strings.HasPrefix(l.input[l.start:], "FD ") {
		l.out += "logo.Forward("
		l.start += len("FD ")
		l.pos = l.start
		lexNumber(l)
		l.out += ")\n"
		return lexProcedure
	}
	return nil
}

func lexNumber(l *lexer) stateFn {
	// Trim leading spaces.
	for l.isSpace() {
		l.ignore()
	}
	for r := l.next(); (r >= '0' && r <= '9') || r == '.' || r == '-' || r == '+'; r = l.next() {
		l.out += string(r)
		continue
	}
	l.start = l.pos
	return nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: %s [file]\n", os.Args[0])
		os.Exit(2)
	}
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Printf("Problem opening file: %s\n", err)
		os.Exit(1)
	}

	l := &lexer{
		input: string(data),
		out:   "package main\n\nimport \"github.com/nyt-hughmandeville/logo2go/pkg/logo\"\n\nfunc main() {\n",
	}
	l.run()

	fmt.Println(l.out)

}
