package main

import (
	"dklang/lexer"
	"dklang/parser"
	"os"

	"github.com/sanity-io/litter"
)

func main() {

	if len(os.Args) < 2 {
		panic("Usage: main <filename>")
	}
	filename := os.Args[1]

	buf, err := os.ReadFile(filename)
	if err != nil {
		panic("File not found")
	}
	s := string(buf)

	tokens := lexer.Tokenize(s)
	//litter.Dump(tokens)

	program := parser.ParseProgram(tokens)
	litter.Dump(program)

	if program != nil {
		program.Run()
	}

}
