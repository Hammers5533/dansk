package main

import (
	"dklang/lexer"
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
	litter.Dump(tokens)

	//var p *parser.Parser = parser.New(tokens)

	//program := p.ParseProgram()
	//litter.Dump(program)
}
