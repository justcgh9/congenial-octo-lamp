package checker

import (
	"fmt"
	"os"
)

func UnexpectedTypeForExpression() {
	fmt.Println("ERROR_UNEXPECTED_TYPE_FOR_EXPRESSION")
	os.Exit(1)
}

func UndefinedVariable() {
	fmt.Println("ERROR_UNDEFINED_VARIABLE")
	os.Exit(1)
}

func NotAReference() {
	fmt.Println("ERROR_NOT_A_REFERENCE")
	os.Exit(1)
}

func MissingMain() {
	fmt.Println("ERROR_MISSING_MAIN")
	os.Exit(1)
}

func InfiniteType() {
	fmt.Println("ERROR_OCCURS_CHECK_INFINITE_TYPE")
	os.Exit(1)
}

func UndeclaredExceptionType() {
	fmt.Println("ERROR_EXCEPTION_TYPE_NOT_DECLARED")
	os.Exit(1)
}

func MissingRecordFields() {
	fmt.Println("ERROR_MISSING_RECORD_FIELDS")
	os.Exit(1)
}