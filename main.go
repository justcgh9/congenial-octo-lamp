package main

import (
	"fmt"
	"os"
	"stella-implementation-in-go/parser"
	"stella-implementation-in-go/visitor"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// TODO: проверить в песочнице, примет ли функция в качестве параметра функцию -> функцию -> значение в качестве функции от двух аргументов и наоборот.

func main() {
	if len(os.Args) != 2 {
		fmt.Println("incorrect number of command line arguments, expected 1")
		os.Exit(1)
	}
	fileName := os.Args[1]
	content, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("error reading from file ", fileName, " : ", err.Error())
	}
	input := antlr.NewInputStream(string(content))
	lexer := parser.NewStellaLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewStellaParser(stream)

	v := &visitor.TypeCheckVisitor{}
	tree := p.Program()
	v.VisitProgram(tree.(*parser.ProgramContext))

}
