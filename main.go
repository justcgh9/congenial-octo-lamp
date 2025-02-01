package main

import (
	"fmt"
	"os"
	"stella-implementation-in-go/parser"
	"stella-implementation-in-go/visitor"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

func handle_expr_context(ctx parser.IExprContext) {
	switch ctx := ctx.(type) {
	default:
		print("unsupported syntax\n")
        fmt.Printf("%T", ctx)
	case *parser.ConstTrueContext:
		print("I see true\n")
	case *parser.IfContext:
		handle_expr_context(ctx.GetCondition())
		handle_expr_context(ctx.GetThenExpr())
		handle_expr_context(ctx.GetElseExpr())
	case *parser.VarContext:
		print("I see variable ", ctx.GetName().GetText(), "\n")
    case *parser.TerminatingSemicolonContext:

	}
}

func handle_decl_context(ctx parser.IDeclContext) {
	switch ctx := ctx.(type) {
	default:
		print("unsupported syntax\n")
	case *parser.DeclFunContext:
		print("Declare function ", ctx.GetName().GetText(), "\n")
		fmt.Println(ctx.Get_paramDecl().GetParamType().GetText())
		fmt.Printf("%s\n", ctx.GetReturnType().GetText())
		fmt.Printf("%s\n", ctx.GetReturnExpr().GetText())
		fmt.Printf("%T\n", ctx.GetReturnExpr())
		handle_expr_context(ctx.GetReturnExpr())
	}
}

func handle_program_context(ctx parser.IProgramContext) {
	for _, decl := range ctx.GetDecls() {
		fmt.Printf("%T\n", ctx)
		handle_decl_context(decl)
	}
}

// TODO: проверить в песочнице, примет ли функция в качестве параметра функцию -> функцию -> значение в качестве функции от двух аргументов и наоборот.

func main() {
	fileName := "examples/0"
	content, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("error reading from file ", fileName, " : ", err.Error())
	}
	input := antlr.NewInputStream(string(content))
	lexer := parser.NewStellaLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewStellaParser(stream)

	var v parser.StellaParserVisitor
	v = &visitor.TypeCheckVisitor{}
	tree := p.Program()
	v.VisitProgram(tree.(*parser.ProgramContext))
	// handle_program_context(tree)

}
