package main

type expr interface{}

type stmtsAST struct {
	list []*stmtAST
}

type stmtAST struct {
	before   string
	resolver *resolverAST
	after    string
}

type resolverAST struct {
	name      string
	arguments *argsAST
}

type argsAST struct {
	list []*argAST
}

type argAST struct {
	before   string
	resolver *resolverAST
	after    string
}
