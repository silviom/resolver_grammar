package main

import (
	"fmt"
	"strings"
)

func (i *interpreter) eval(e expr) string {
	switch t := e.(type) {
	case *stmtsAST:
		if t == nil {
			return ""
		}
		results := []string{}
		for _, s := range t.list {
			results = append(results, i.eval(s))
		}
		return fmt.Sprintf("[ %s ]", strings.Join(results, ", "))

	case *stmtAST:
		if t == nil {
			return ""
		}

		b := i.eval(t.before)
		if i.evaluationFailed {
			return b
		}

		r := i.eval(t.resolver)
		if i.evaluationFailed {
			return r
		}

		a := i.eval(t.after)
		if i.evaluationFailed {
			return a
		}

		return fmt.Sprintf("<b:%s r:%s a:%s>", b, r, a)

	case *resolverAST:
		if t == nil {
			return ""
		}

		n := i.eval(t.name)
		if i.evaluationFailed {
			return n
		}

		a := i.eval(t.arguments)
		if i.evaluationFailed {
			return a
		}

		return fmt.Sprintf("<n:%s a:%s>", n, a)

	case *argsAST:
		if t == nil {
			return ""
		}
		results := []string{}
		for _, a := range t.list {
			results = append(results, i.eval(a))
		}
		return fmt.Sprintf("[ %s ]", strings.Join(results, ", "))

	case *argAST:
		if t == nil {
			return ""
		}

		b := i.eval(t.before)
		if i.evaluationFailed {
			return b
		}

		r := i.eval(t.resolver)
		if i.evaluationFailed {
			return r
		}

		a := i.eval(t.after)
		if i.evaluationFailed {
			return a
		}

		return fmt.Sprintf("<b:%s r:%s a:%s>", b, r, a)

	case string:
		return t

	default:
		panic("invalid node type")
	}
}
