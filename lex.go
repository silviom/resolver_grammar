package main

import (
	"fmt"
	"regexp"
)

const EOF = 0

type tokenDef struct {
	re    *regexp.Regexp
	token int
}

var definitions = []tokenDef{

	{re: regexp.MustCompile(`^[_a-zA-Z0-9]+`), token: NAME},
	{re: regexp.MustCompile(`^{{2}`), token: BEGIN},
	{re: regexp.MustCompile(`^}{2}`), token: END},
	{re: regexp.MustCompile(`^\(`), token: OPEN},
	{re: regexp.MustCompile(`^,`), token: COMMA},
	{re: regexp.MustCompile(`^\)`), token: CLOSE},
	{re: regexp.MustCompile(`^\s`), token: SPACE},
	{re: regexp.MustCompile(`^\S`), token: NOT_SPACE},
}

func (l *interpreter) Lex(lval *resolverSymType) int {
	lval.String = ""

	// Remove new lines
	for ; len(l.input) > 0 && l.input[0] == '\n'; l.input = l.input[1:] {
	}

	// Check if the input has ended.
	if len(l.input) == 0 {
		fmt.Printf("LEX: EOF\n")
		return EOF
	}

	for _, def := range definitions {
		result := def.re.FindString(l.input)
		length := len(result)
		if length > 0 {
			fmt.Printf("LEX: `%s` '%s' [%d]\n", def.re, result, length)
			lval.String = result
			l.input = l.input[length:]
			return def.token
		}
	}

	// Otherwise return the next letter.
	letter := int(l.input[0])
	l.input = l.input[1:]
	return letter
}
