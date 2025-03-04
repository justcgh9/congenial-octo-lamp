package visitor

import (
	"fmt"
	"os"
	"stella-implementation-in-go/env"
	"stella-implementation-in-go/parser"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

type TypeCheckVisitor struct {
	*antlr.BaseParseTreeVisitor
	env env.Env
	checking bool
	checkingForType env.Type
	patterns map[string]bool
	curPattern string
}

const (
	exceptionType = "-exceptionType"
)

func (v *TypeCheckVisitor) GeneratePatternOptions(x interface{}, pattern string) {
	switch x := x.(type) {
	default:
	case env.Bool:
		v.patterns[pattern + "true"] = false
		v.patterns[pattern + "false"] = false
	case env.Nat:
		v.patterns[pattern + "nat"] = false
		v.patterns[pattern + "nat 0"] = false
	case env.Sum:
		v.GeneratePatternOptions(x.Left, pattern + "inl ")
		v.GeneratePatternOptions(x.Right, pattern + "inr ")
	case env.Func:
		v.patterns[pattern + x.Type()] = false
	case env.Unit:
		v.patterns[pattern + "unit"] = false
	case env.List:
		if x.T == nil || x.T.Type() == "" {
			fmt.Println("ERROR_AMBIGUOUS_LIST_TYPE")
			os.Exit(1)
		}
		v.patterns[pattern + "[]"] = false
		v.patterns[pattern + "[]" + x.T.Type()] = false
	
	case env.Tuple:
		v.patterns[pattern + x.Type()] = false
	}
}

func TypeCheck(exp, given interface{}, args ...string) {

	if exp.(env.Type).Type() == given.(env.Type).Type() {
		return
	}

	switch exp.(type) {
	default:
		fmt.Println("ERROR_UNEXPECTED_TYPE_FOR_EXPRESSION. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())
		os.Exit(1)
	case env.Func:
		ef := exp.(env.Func)
		tuple, ok := given.(env.Tuple)
		if ok {
			if tuple.IsLiteral {
				fmt.Println("ERROR_UNEXPECTED_TUPLE. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())				
			} else {
				fmt.Println("ERROR_UNEXPECTED_TYPE_FOR_EXPRESSION. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())
			}
			os.Exit(1)
		}

		if throw, ok := given.(*env.Throw); ok {
			if isAny(throw.UnderlyingType) {
				throw.UnderlyingType = exp.(env.Type)
				return
			}

			TypeCheck(exp, throw.UnderlyingType)
			return
		}

		val, ok := given.(env.List)
		if ok {
			if _, ok := isAmbiguousType(val); ok {
				fmt.Println("ERROR_AMBIGUOUS_LIST_TYPE")
			} else  {
				if val.IsLiteral {
					fmt.Println("ERROR_UNEXPECTED_LIST. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())				
				} else {
					fmt.Println("ERROR_UNEXPECTED_TYPE_FOR_EXPRESSION. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())
				}
			}
			os.Exit(1)
		}

		_, ok = given.(env.Sum)
		if ok {
			fmt.Println("ERROR_UNEXPECTED_INJECTION. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())
			os.Exit(1)
		}

		record, ok := given.(env.Record)
		if ok {
			if record.IsLiteral {
				fmt.Println("ERROR_UNEXPECTED_RECORD. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())				
			} else {
				fmt.Println("ERROR_UNEXPECTED_TYPE_FOR_EXPRESSION. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())
			}
			os.Exit(1)
		}

		reference, ok := given.(env.Reference)
		if ok {
			if _, ok := isAmbiguousType(reference); ok {
				fmt.Println("ERROR_AMBIGUOUS_REFERENCE_TYPE")
			} else {
				if reference.IsLiteral {
					fmt.Println("ERROR_UNEXPECTED_MEMORY_ADDRESS. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())				
				} else {
					fmt.Println("ERROR_UNEXPECTED_TYPE_FOR_EXPRESSION. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())
				}
			}
			os.Exit(1)
		}

		ft, ok := given.(env.Func)

		if !ok || !ft.IsAnonymous || ft.Return.Type() != ef.Return.Type() {
			fmt.Println("ERROR_UNEXPECTED_TYPE_FOR_EXPRESSION. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())
			os.Exit(1)
		}

		if len(ef.Args) != len(ft.Args) {
			fmt.Println("ERROR_UNEXPECTED_NUMBER_OF_PARAMETERS_IN_LAMBDA. Expected ", len(ef.Args), " , given ", len(ft.Args))
			os.Exit(1)
		}

		fmt.Println("ERROR_UNEXPECTED_TYPE_FOR_PARAMETER")
		os.Exit(1)

	case env.Nat:
		ft, ok := given.(env.Func)
		if ok && ft.IsAnonymous {
			fmt.Println("ERROR_UNEXPECTED_LAMBDA. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())
			os.Exit(1)
		}

		if throw, ok := given.(*env.Throw); ok {
			if isAny(throw.UnderlyingType) {
				throw.UnderlyingType = exp.(env.Type)
				return
			}

			TypeCheck(exp, throw.UnderlyingType)
			return
		}

		val, ok := given.(env.List)
		if ok {
			if _, ok := isAmbiguousType(val); ok {
				fmt.Println("ERROR_AMBIGUOUS_LIST_TYPE")
			} else  {
				if val.IsLiteral {
					fmt.Println("ERROR_UNEXPECTED_LIST. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())				
				} else {
					fmt.Println("ERROR_UNEXPECTED_TYPE_FOR_EXPRESSION. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())
				}
			}
			os.Exit(1)
		}

		_, ok = given.(env.Sum)
		if ok {
			fmt.Println("ERROR_UNEXPECTED_INJECTION. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())
			os.Exit(1)
		}

		tuple, ok := given.(env.Tuple)
		if ok {
			if tuple.IsLiteral {
				fmt.Println("ERROR_UNEXPECTED_TUPLE. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())				
			} else {
				fmt.Println("ERROR_UNEXPECTED_TYPE_FOR_EXPRESSION. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())
			}
			os.Exit(1)
		}

		record, ok := given.(env.Record)
		if ok {
			if record.IsLiteral {
				fmt.Println("ERROR_UNEXPECTED_RECORD. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())				
			} else {
				fmt.Println("ERROR_UNEXPECTED_TYPE_FOR_EXPRESSION. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())
			}
			os.Exit(1)
		}

		reference, ok := given.(env.Reference)
		if ok {
			if _, ok := isAmbiguousType(reference); ok {
				fmt.Println("ERROR_AMBIGUOUS_REFERENCE_TYPE")
			} else {
				if reference.IsLiteral {
					fmt.Println("ERROR_UNEXPECTED_MEMORY_ADDRESS. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())				
				} else {
					fmt.Println("ERROR_UNEXPECTED_TYPE_FOR_EXPRESSION. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())
				}
			}
			os.Exit(1)
		}

		fmt.Println("ERROR_UNEXPECTED_TYPE_FOR_EXPRESSION. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())
		os.Exit(1)

	case env.Bool:
		ft, ok := given.(env.Func)
		if ok && ft.IsAnonymous {
			fmt.Println("ERROR_UNEXPECTED_LAMBDA. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())
			os.Exit(1)
		}

		if throw, ok := given.(*env.Throw); ok {
			if isAny(throw.UnderlyingType) {
				throw.UnderlyingType = exp.(env.Type)
				return
			}

			TypeCheck(exp, throw.UnderlyingType)
			return
		}

		val, ok := given.(env.List)
		if ok {
			if _, ok := isAmbiguousType(val); ok {
				fmt.Println("ERROR_AMBIGUOUS_LIST_TYPE")
			} else  {
				if val.IsLiteral {
					fmt.Println("ERROR_UNEXPECTED_LIST. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())				
				} else {
					fmt.Println("ERROR_UNEXPECTED_TYPE_FOR_EXPRESSION. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())
				}
			}
			os.Exit(1)
		}

		record, ok := given.(env.Record)
		if ok {
			if record.IsLiteral {
				fmt.Println("ERROR_UNEXPECTED_RECORD. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())				
			} else {
				fmt.Println("ERROR_UNEXPECTED_TYPE_FOR_EXPRESSION. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())
			}
			os.Exit(1)
		}

		reference, ok := given.(env.Reference)
		if ok {
			if _, ok := isAmbiguousType(reference); ok {
				fmt.Println("ERROR_AMBIGUOUS_REFERENCE_TYPE")
			} else {
				if reference.IsLiteral {
					fmt.Println("ERROR_UNEXPECTED_MEMORY_ADDRESS. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())				
				} else {
					fmt.Println("ERROR_UNEXPECTED_TYPE_FOR_EXPRESSION. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())
				}
			}
			os.Exit(1)
		}

		_, ok = given.(env.Sum)
		if ok {
			fmt.Println("ERROR_UNEXPECTED_INJECTION. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())
			os.Exit(1)
		}

		tuple, ok := given.(env.Tuple)
		if ok {
			if tuple.IsLiteral {
				fmt.Println("ERROR_UNEXPECTED_TUPLE. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())				
			} else {
				fmt.Println("ERROR_UNEXPECTED_TYPE_FOR_EXPRESSION. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())
			}
			os.Exit(1)
		}


		fmt.Println("ERROR_UNEXPECTED_TYPE_FOR_EXPRESSION. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())
		os.Exit(1)
		

	case env.Tuple:
		ft, ok := given.(env.Func)
		if ok && ft.IsAnonymous {
			fmt.Println("ERROR_UNEXPECTED_LAMBDA. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())
			os.Exit(1)
		}

		if throw, ok := given.(*env.Throw); ok {
			if isAny(throw.UnderlyingType) {
				throw.UnderlyingType = exp.(env.Type)
				return
			}

			TypeCheck(exp, throw.UnderlyingType)
			return
		}

		val, ok := given.(env.List)
		if ok {
			if _, ok := isAmbiguousType(val); ok {
				fmt.Println("ERROR_AMBIGUOUS_LIST_TYPE")
			} else  {
				if val.IsLiteral {
					fmt.Println("ERROR_UNEXPECTED_LIST. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())				
				} else {
					fmt.Println("ERROR_UNEXPECTED_TYPE_FOR_EXPRESSION. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())
				}
			}
			os.Exit(1)
		}

		_, ok = given.(env.Sum)
		if ok {
			fmt.Println("ERROR_UNEXPECTED_INJECTION. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())
			os.Exit(1)
		}

		record, ok := given.(env.Record)
		if ok {
			if record.IsLiteral {
				fmt.Println("ERROR_UNEXPECTED_RECORD. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())				
			} else {
				fmt.Println("ERROR_UNEXPECTED_TYPE_FOR_EXPRESSION. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())
			}
			os.Exit(1)
		}

		reference, ok := given.(env.Reference)
		if ok {
			if _, ok := isAmbiguousType(reference); ok {
				fmt.Println("ERROR_AMBIGUOUS_REFERENCE_TYPE")
			} else {
				if reference.IsLiteral {
					fmt.Println("ERROR_UNEXPECTED_MEMORY_ADDRESS. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())				
				} else {
					fmt.Println("ERROR_UNEXPECTED_TYPE_FOR_EXPRESSION. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())
				}
			}
			os.Exit(1)
		}

		et := exp.(env.Tuple)
		tt, ok := given.(env.Tuple)

		if !ok {
			fmt.Println("ERROR_UNEXPECTED_TYPE_FOR_EXPRESSION. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())
			os.Exit(1)
		}

		if len(et.Elements) != len(tt.Elements) {
			fmt.Println("ERROR_UNEXPECTED_TUPLE_LENGTH. Expected ", len(et.Elements), " , given ", len(tt.Elements))
			os.Exit(1)
		}

		for i := range et.Elements {
			TypeCheck(et.Elements[i], tt.Elements[i])
		}

		fmt.Println("ERROR_UNEXPECTED_TYPE_FOR_EXPRESSION. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())
		os.Exit(1)
	
	case env.Record:
		ft, ok := given.(env.Func)
		if ok && ft.IsAnonymous {
			fmt.Println("ERROR_UNEXPECTED_LAMBDA. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())
			os.Exit(1)
		}

		if throw, ok := given.(*env.Throw); ok {
			if isAny(throw.UnderlyingType) {
				throw.UnderlyingType = exp.(env.Type)
				return
			}

			TypeCheck(exp, throw.UnderlyingType)
			return
		}

		tuple, ok := given.(env.Tuple)
		if ok {
			if tuple.IsLiteral {
				fmt.Println("ERROR_UNEXPECTED_TUPLE. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())				
			} else {
				fmt.Println("ERROR_UNEXPECTED_TYPE_FOR_EXPRESSION. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())
			}
			os.Exit(1)
		}

		val, ok := given.(env.List)
		if ok {
			if _, ok := isAmbiguousType(val); ok {
				fmt.Println("ERROR_AMBIGUOUS_LIST_TYPE")
			} else  {
				if val.IsLiteral {
					fmt.Println("ERROR_UNEXPECTED_LIST. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())				
				} else {
					fmt.Println("ERROR_UNEXPECTED_TYPE_FOR_EXPRESSION. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())
				}
			}
			os.Exit(1)
		}

		_, ok = given.(env.Sum)
		if ok {
			fmt.Println("ERROR_UNEXPECTED_INJECTION. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())
			os.Exit(1)
		}

		reference, ok := given.(env.Reference)
		if ok {
			if _, ok := isAmbiguousType(reference); ok {
				fmt.Println("ERROR_AMBIGUOUS_REFERENCE_TYPE")
			} else {
				if reference.IsLiteral {
					fmt.Println("ERROR_UNEXPECTED_MEMORY_ADDRESS. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())				
				} else {
					fmt.Println("ERROR_UNEXPECTED_TYPE_FOR_EXPRESSION. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())
				}
			}
			os.Exit(1)
		}

		et := exp.(env.Record)
		record, ok := given.(env.Record)
		if !ok{
			fmt.Println("ERROR_UNEXPECTED_TYPE_FOR_EXPRESSION. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())
			os.Exit(1)
		}

		expectedLength := len(et.Elements)
		givenLength := len(record.Elements)

		if expectedLength < givenLength {
			fmt.Println("ERROR_UNEXPECTED_RECORD_FIELDS")
			os.Exit(1)
		}

		if expectedLength > givenLength {
			fmt.Println("ERROR_MISSING_RECORD_FIELDS")
			os.Exit(1)
		}

		for key, value := range et.Elements {
			givenElement, ok := record.Elements[key]
			if !ok {
				fmt.Println("ERROR_MISSING_RECORD_FIELDS")
				os.Exit(1)
			}

			TypeCheck(value, givenElement)
		}

		fmt.Println("ERROR_UNEXPECTED_TYPE_FOR_EXPRESSION. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())
		os.Exit(1)
	
	case env.Reference:
		ft, ok := given.(env.Func)
		if ok && ft.IsAnonymous {
			fmt.Println("ERROR_UNEXPECTED_LAMBDA. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())
			os.Exit(1)
		}

		if throw, ok := given.(*env.Throw); ok {
			if isAny(throw.UnderlyingType) {
				throw.UnderlyingType = exp.(env.Type)
				return
			}

			TypeCheck(exp, throw.UnderlyingType)
			return
		}

		tuple, ok := given.(env.Tuple)
		if ok {
			if tuple.IsLiteral {
				fmt.Println("ERROR_UNEXPECTED_TUPLE. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())				
			} else {
				fmt.Println("ERROR_UNEXPECTED_TYPE_FOR_EXPRESSION. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())
			}
			os.Exit(1)
		}

		val, ok := given.(env.List)
		if ok {
			if _, ok := isAmbiguousType(val); ok {
				fmt.Println("ERROR_AMBIGUOUS_LIST_TYPE")
			} else  {
				if val.IsLiteral {
					fmt.Println("ERROR_UNEXPECTED_LIST. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())				
				} else {
					fmt.Println("ERROR_UNEXPECTED_TYPE_FOR_EXPRESSION. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())
				}
			}
			os.Exit(1)
		}

		record, ok := given.(env.Record)
		if ok {
			if record.IsLiteral {
				fmt.Println("ERROR_UNEXPECTED_RECORD. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())				
			} else {
				fmt.Println("ERROR_UNEXPECTED_TYPE_FOR_EXPRESSION. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())
			}
			os.Exit(1)
		}

		_, ok = given.(env.Sum)
		if ok {
			fmt.Println("ERROR_UNEXPECTED_INJECTION. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())
			os.Exit(1)
		}

		reference, ok := given.(env.Reference)
		if ok {
			if _, ok := isAmbiguousType(reference); ok && contains("strictRef", args...) {
				fmt.Println("ERROR_AMBIGUOUS_REFERENCE_TYPE")
				os.Exit(1)
			}
			exp := exp.(env.Reference)
			TypeCheck(exp.UnderlyingType, reference.UnderlyingType, args...)
		}


		fmt.Println("ERROR_UNEXPECTED_TYPE_FOR_EXPRESSION. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())
		os.Exit(1)
	
		
	case env.Unit:
		ft, ok := given.(env.Func)
		if ok && ft.IsAnonymous {
			fmt.Println("ERROR_UNEXPECTED_LAMBDA. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())
			os.Exit(1)
		}

		if throw, ok := given.(*env.Throw); ok {
			if isAny(throw.UnderlyingType) {
				throw.UnderlyingType = exp.(env.Type)
				return
			}

			TypeCheck(exp, throw.UnderlyingType)
			return
		}

		tuple, ok := given.(env.Tuple)
		if ok {
			if tuple.IsLiteral {
				fmt.Println("ERROR_UNEXPECTED_TUPLE. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())				
			} else {
				fmt.Println("ERROR_UNEXPECTED_TYPE_FOR_EXPRESSION. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())
			}
			os.Exit(1)
		}

		val, ok := given.(env.List)
		if ok {
			if _, ok := isAmbiguousType(val); ok {
				fmt.Println("ERROR_AMBIGUOUS_LIST_TYPE")
			} else  {
				if val.IsLiteral {
					fmt.Println("ERROR_UNEXPECTED_LIST. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())				
				} else {
					fmt.Println("ERROR_UNEXPECTED_TYPE_FOR_EXPRESSION. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())
				}
			}
			os.Exit(1)
		}

		record, ok := given.(env.Record)
		if ok {
			if record.IsLiteral {
				fmt.Println("ERROR_UNEXPECTED_RECORD. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())				
			} else {
				fmt.Println("ERROR_UNEXPECTED_TYPE_FOR_EXPRESSION. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())
			}
			os.Exit(1)
		}

		reference, ok := given.(env.Reference)
		if ok {
			if _, ok := isAmbiguousType(reference); ok {
				fmt.Println("ERROR_AMBIGUOUS_REFERENCE_TYPE")
			} else {
				if reference.IsLiteral {
					fmt.Println("ERROR_UNEXPECTED_MEMORY_ADDRESS. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())				
				} else {
					fmt.Println("ERROR_UNEXPECTED_TYPE_FOR_EXPRESSION. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())
				}
			}
			os.Exit(1)
		}

		_, ok = given.(env.Sum)
		if ok {
			fmt.Println("ERROR_UNEXPECTED_INJECTION. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())
			os.Exit(1)
		}


		fmt.Println("ERROR_UNEXPECTED_TYPE_FOR_EXPRESSION. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())
		os.Exit(1)
	
	case env.Sum:
		et := exp.(env.Sum)
		ft, ok := given.(env.Func)
		if ok && ft.IsAnonymous {
			fmt.Println("ERROR_UNEXPECTED_LAMBDA. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())
			os.Exit(1)
		}

		tuple, ok := given.(env.Tuple)
		if ok {
			if tuple.IsLiteral {
				fmt.Println("ERROR_UNEXPECTED_TUPLE. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())				
			} else {
				fmt.Println("ERROR_UNEXPECTED_TYPE_FOR_EXPRESSION. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())
			}
			os.Exit(1)
		}

		if throw, ok := given.(*env.Throw); ok {
			if isAny(throw.UnderlyingType) {
				throw.UnderlyingType = exp.(env.Type)
				return
			}

			TypeCheck(exp, throw.UnderlyingType)
			return
		}

		val, ok := given.(env.List)
		if ok {
			if _, ok := isAmbiguousType(val); ok {
				fmt.Println("ERROR_AMBIGUOUS_LIST_TYPE")
			} else  {
				if val.IsLiteral {
					fmt.Println("ERROR_UNEXPECTED_LIST. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())				
				} else {
					fmt.Println("ERROR_UNEXPECTED_TYPE_FOR_EXPRESSION. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())
				}
			}
			os.Exit(1)
		}

		record, ok := given.(env.Record)
		if ok {
			if record.IsLiteral {
				fmt.Println("ERROR_UNEXPECTED_RECORD. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())				
			} else {
				fmt.Println("ERROR_UNEXPECTED_TYPE_FOR_EXPRESSION. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())
			}
			os.Exit(1)
		}

		reference, ok := given.(env.Reference)
		if ok {
			if _, ok := isAmbiguousType(reference); ok {
				fmt.Println("ERROR_AMBIGUOUS_REFERENCE_TYPE")
			} else {
				if reference.IsLiteral {
					fmt.Println("ERROR_UNEXPECTED_MEMORY_ADDRESS. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())				
				} else {
					fmt.Println("ERROR_UNEXPECTED_TYPE_FOR_EXPRESSION. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())
				}
			}
			os.Exit(1)
		}

		st, ok := given.(env.Sum) 
		if !ok {
			fmt.Println("ERROR_UNEXPECTED_TYPE_FOR_EXPRESSION. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())
			os.Exit(1)
		}
		
		if st.Left.Type() != et.Left.Type() {
			if isAny(et.Left) || isAny(st.Left) {
				if contains("strictSum", args...) {
					// fmt.Println(et.Left.Type(), st.Left.Type())
					fmt.Println("ERROR_AMBIGUOUS_SUM_TYPE")
					os.Exit(1)
				}
			} else {
				TypeCheck(et.Left, st.Left, args...)
			}
		}

		if st.Right.Type() != et.Right.Type() {
			if isAny(et.Right) || isAny(st.Right) {
				if contains("strictSum", args...) {
					fmt.Println(et.Right.Type(), st.Right.Type())
					fmt.Println("ERROR_AMBIGUOUS_SUM_TYPE")
					os.Exit(1)
				}
			} else {
				TypeCheck(et.Right, st.Right, args...)
			}
		}


		return

	case env.List:

		ft, ok := given.(env.Func)
		if ok && ft.IsAnonymous {
			fmt.Println("ERROR_UNEXPECTED_LAMBDA. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())
			os.Exit(1)
		}

		if throw, ok := given.(*env.Throw); ok {
			if isAny(throw.UnderlyingType) {
				throw.UnderlyingType = exp.(env.Type)
				return
			}

			TypeCheck(exp, throw.UnderlyingType)
			return
		}

		tuple, ok := given.(env.Tuple)
		if ok {
			if tuple.IsLiteral {
				fmt.Println("ERROR_UNEXPECTED_TUPLE. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())				
			} else {
				fmt.Println("ERROR_UNEXPECTED_TYPE_FOR_EXPRESSION. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())
			}
			os.Exit(1)
		}

		_, ok = given.(env.Sum)
		if ok {
			fmt.Println("ERROR_UNEXPECTED_INJECTION. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())
			os.Exit(1)
		}

		record, ok := given.(env.Record)
		if ok {
			if record.IsLiteral {
				fmt.Println("ERROR_UNEXPECTED_RECORD. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())				
			} else {
				fmt.Println("ERROR_UNEXPECTED_TYPE_FOR_EXPRESSION. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())
			}
			os.Exit(1)
		}

		reference, ok := given.(env.Reference)
		if ok {
			if _, ok := isAmbiguousType(reference); ok {
				fmt.Println("ERROR_AMBIGUOUS_REFERENCE_TYPE")
			} else {
				if reference.IsLiteral {
					fmt.Println("ERROR_UNEXPECTED_MEMORY_ADDRESS. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())				
				} else {
					fmt.Println("ERROR_UNEXPECTED_TYPE_FOR_EXPRESSION. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())
				}
			}
			os.Exit(1)
		}

		gt, ok := given.(env.List)
		if !ok {
			fmt.Println("ERROR_UNEXPECTED_TYPE_FOR_EXPRESSION. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())
			os.Exit(1)
		}

		if (gt.T == nil || gt.T.Type() == "") {
			if contains("strictList", args...) {
				fmt.Println("ERROR_AMBIGUOUS_LIST_TYPE")
				os.Exit(1)
			}

			return
		}

	case *env.Throw:
		throw := exp.(*env.Throw)
		given := given.(env.Type)

		if given, ok := given.(*env.Throw); ok {

			if isAny(throw.UnderlyingType) {
				throw.UnderlyingType = given.UnderlyingType
				return
			}

			if isAny(given.UnderlyingType) {
				given.UnderlyingType = throw.UnderlyingType
				return
			} 

			TypeCheck(throw.UnderlyingType, given.UnderlyingType)
			return
		}

		if isAny(throw.UnderlyingType) {
			throw.UnderlyingType = given
			return
		}

		TypeCheck(throw.UnderlyingType, given)
		return
	}
	

	fmt.Println("ERROR_UNEXPECTED_TYPE_FOR_EXPRESSION. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())
	os.Exit(1)
}

func (v *TypeCheckVisitor) VisitStart_Program(ctx *parser.Start_ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeCheckVisitor) VisitStart_Expr(ctx *parser.Start_ExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeCheckVisitor) VisitStart_Type(ctx *parser.Start_TypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeCheckVisitor) VisitProgram(ctx *parser.ProgramContext) interface{} {
	//TODO check the order of functions
	v.env.Push()

	for _, decl := range ctx.GetDecls() {
		val, ok := decl.(*parser.DeclFunContext)
		if !ok {continue}
		args := make([]env.Type, 0, len(val.GetParamDecls()))
		for _, val := range val.GetParamDecls() {
			args = append(args, val.Accept(v).(env.Type))
		}
		returnType := val.GetReturnType().Accept(v).(env.Type)
		
		v.env.Put(val.GetName().GetText(), env.Func{
			Args: args,
			Return: returnType,
		})
	}

	for _, decl := range ctx.GetDecls() {
		decl.Accept(v)
	}

	val, ok := v.env.Check("main").(env.Func)
	if !ok  {
		fmt.Println("ERROR_MISSING_MAIN")
		os.Exit(1)	
	}

	if len(val.Args) != 1 {
		fmt.Println("ERROR_INCORRECT_ARITY_OF_MAIN")
		os.Exit(1)
	}
	
	v.env.Pop()

	return v.VisitChildren(ctx)
}

func (v *TypeCheckVisitor) VisitLanguageCore(ctx *parser.LanguageCoreContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeCheckVisitor) VisitAnExtension(ctx *parser.AnExtensionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeCheckVisitor) VisitDeclFun(ctx *parser.DeclFunContext) interface{} {
	//TODO


	returnType := ctx.GetReturnType().Accept(v).(env.Type)

	v.env.Push()
	v.env.Put("0", returnType)

	decls := ctx.GetLocalDecls()

	for _, decl := range decls {
		val, ok := decl.(*parser.DeclFunContext)
		if !ok {
			fmt.Println("ERROR_NOT_A_FUNCTION")
		}
		args := make([]env.Type, 0, len(val.GetParamDecls()))
		for _, value := range val.GetParamDecls() {
			args = append(args, value.Accept(v).(env.Type))
		}
		returnType := val.GetReturnType().Accept(v).(env.Type)

		v.env.Put(val.GetName().GetText(), env.Func{
			Args: args,
			Return: returnType,
		})
	}

	for _, decl := range decls {
		decl.Accept(v)
	}

	for _, val := range ctx.GetParamDecls() {
		v.env.Put(val.GetName().GetText(), val.Accept(v).(env.Type))
	}

	prev := v.checking
	prevType := v.checkingForType
	v.checking = true
	v.checkingForType = returnType

	ans := ctx.GetReturnExpr().Accept(v).(env.Type)
	TypeCheck(returnType, ans)

	v.checkingForType = prevType
	v.checking = prev
	v.env.Pop()
	return v.VisitChildren(ctx)
}

func (v *TypeCheckVisitor) VisitDeclFunGeneric(ctx *parser.DeclFunGenericContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeCheckVisitor) VisitDeclTypeAlias(ctx *parser.DeclTypeAliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeCheckVisitor) VisitDeclExceptionType(ctx *parser.DeclExceptionTypeContext) interface{} {
	t := ctx.GetExceptionType().Accept(v).(env.Type)
	v.env.Put(exceptionType, t)
	return t
}

func (v *TypeCheckVisitor) VisitDeclExceptionVariant(ctx *parser.DeclExceptionVariantContext) interface{} {
	return v.Visit(ctx)
}

func (v *TypeCheckVisitor) VisitInlineAnnotation(ctx *parser.InlineAnnotationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeCheckVisitor) VisitParamDecl(ctx *parser.ParamDeclContext) interface{} {
	//TODO

	return ctx.GetParamType().Accept(v)
}

func (v *TypeCheckVisitor) VisitFold(ctx *parser.FoldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeCheckVisitor) VisitAdd(ctx *parser.AddContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeCheckVisitor) VisitIsZero(ctx *parser.IsZeroContext) interface{} {
	//TODO
	prev := v.checking
	prevType := v.checkingForType
	v.checking = true
	v.checkingForType = env.Nat{}

	s := ctx.Expr().Accept(v).(env.Type)
	TypeCheck(env.Nat{}, s)

	v.checkingForType = prevType
	v.checking = prev
	
	return env.Bool{}
}

func (v *TypeCheckVisitor) VisitVar(ctx *parser.VarContext) interface{} {
	//TODO
	tp := v.env.Check(ctx.GetName().GetText())
	if tp.Type() == "" {
		fmt.Println("ERROR_UNDEFINED_VARIABLE.: ", ctx.GetName().GetText())
		os.Exit(1)
	}
	return tp
}

func (v *TypeCheckVisitor) VisitTypeAbstraction(ctx *parser.TypeAbstractionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeCheckVisitor) VisitDivide(ctx *parser.DivideContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeCheckVisitor) VisitLessThan(ctx *parser.LessThanContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeCheckVisitor) VisitDotRecord(ctx *parser.DotRecordContext) interface{} {

	label := ctx.GetLabel().GetText()
	record, ok := ctx.GetExpr_().Accept(v).(env.Record)
	if !ok {
		fmt.Println("ERROR_NOT_A_RECORD")
		os.Exit(1)
	}

	for _, t := range record.Elements {
		throwAmbiguousType(t)
	}

	item, ok := record.Elements[label]
	if !ok {
		fmt.Println("ERROR_UNEXPECTED_FIELD_ACCESS")
		os.Exit(1)
	}

	return item
}

func (v *TypeCheckVisitor) VisitGreaterThan(ctx *parser.GreaterThanContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeCheckVisitor) VisitEqual(ctx *parser.EqualContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeCheckVisitor) VisitThrow(ctx *parser.ThrowContext) interface{} {
	t := v.env.Check(exceptionType)
	_, ok := t.(env.Erroneos)
	if ok {
		fmt.Println("ERROR_EXCEPTION_TYPE_NOT_DECLARED")
		os.Exit(1)
	}

	TypeCheck(t, ctx.GetExpr_().Accept(v))

	return &env.Throw{
		UnderlyingType: env.Any{},
	}
}

func (v *TypeCheckVisitor) VisitMultiply(ctx *parser.MultiplyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeCheckVisitor) VisitConstMemory(ctx *parser.ConstMemoryContext) interface{} {
	return env.Reference{
		UnderlyingType: env.Any{},
		IsLiteral: true,
	}
}

func (v *TypeCheckVisitor) VisitList(ctx *parser.ListContext) interface{} {
	
	exprs := ctx.GetExprs()

	// fmt.Println(ctx.GetText())

	if len(exprs) <= 0 {
		return env.List{
			IsLiteral: true,
		}
	}

	t := exprs[0].Accept(v).(env.Type)

	for _, val := range exprs {
		t2 := val.Accept(v).(env.Type)
		// fmt.Println(t.Type(), val.Accept(v).(env.Type).Type())
		TypeCheck(t, t2)
	}

	return env.List{
		T: t,
		IsLiteral: true,
	}
}

func (v *TypeCheckVisitor) VisitTryCatch(ctx *parser.TryCatchContext) interface{} {
	t := v.env.Check(exceptionType)
	_, ok := t.(env.Erroneos)
	if ok {
		fmt.Println("ERROR_EXCEPTION_TYPE_NOT_DECLARED")
		os.Exit(1)
	}

	tryType := ctx.GetTryExpr().Accept(v).(env.Type)
	_ = tryType

	v.env.Push()
	v.env.Put("-1", env.Erroneos{})

	pat := ctx.GetPat().Accept(v)
	if val, ok := pat.(env.Binding); ok {
		v.env.Put(val.Name, val.T)
	}

	expr := ctx.GetFallbackExpr().Accept(v)
	TypeCheck(tryType, expr)

	f, fok := tryType.(env.Sum)
	s, sok := expr.(env.Sum)
	if fok && sok {
		return ascriptSumTypes(f, s)
	}
	

	v.env.Pop()
	return tryType
}

func (v *TypeCheckVisitor) VisitTryCastAs(ctx *parser.TryCastAsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeCheckVisitor) VisitHead(ctx *parser.HeadContext) interface{} {
	
	throwAmbiguousType(ctx.GetList().Accept(v))

	t, ok := ctx.GetList().Accept(v).(env.List); 
	if !ok {
		fmt.Println("ERROR_NOT_A_LIST")
		os.Exit(1)
	}

	if t.T == nil {
		fmt.Println("ERROR_AMBIGUOUS_LIST_TYPE")
		os.Exit(1)
	}
	
	return t.T
}

func (v *TypeCheckVisitor) VisitTerminatingSemicolon(ctx *parser.TerminatingSemicolonContext) interface{} {
	return ctx.GetExpr_().Accept(v)
}

func (v *TypeCheckVisitor) VisitNotEqual(ctx *parser.NotEqualContext) interface{} {
	return v.VisitChildren(ctx)
}

/*

*/

func (v *TypeCheckVisitor) VisitConstUnit(ctx *parser.ConstUnitContext) interface{} {
	//TODO

	return env.Unit{}
}

func (v *TypeCheckVisitor) VisitSequence(ctx *parser.SequenceContext) interface{} {
	// fmt.Printf("%T %T\n", ctx.GetExpr1(), ctx.GetExpr2())
	TypeCheck(env.Unit{}, ctx.GetExpr1().Accept(v))
	return ctx.GetExpr2().Accept(v)
}

func (v *TypeCheckVisitor) VisitConstFalse(ctx *parser.ConstFalseContext) interface{} {
	//TODO

	return env.Bool{}
}

func (v *TypeCheckVisitor) VisitAbstraction(ctx *parser.AbstractionContext) interface{} {
	//TODO

	
	v.env.Push()
	for i := range ctx.GetParamDecls() {
		v.env.Put(ctx.GetParamDecls()[i].GetName().GetText(), ctx.GetParamDecls()[i].Accept(v).(env.Type))
	}

	prev := v.checking
	v.checking = false

	returnType := ctx.GetReturnExpr().Accept(v).(env.Type)

	v.checking = prev

	v.env.Pop()

	res := env.Func{
		Return: returnType,
		IsAnonymous: true,
	}

	args := make([]env.Type, len(ctx.GetParamDecls()))
	for i, val := range ctx.GetParamDecls() {
		args[i] = val.Accept(v).(env.Type)
	}

	res.Args = args
	return res
}

func (v *TypeCheckVisitor) VisitConstInt(ctx *parser.ConstIntContext) interface{} {
	//TODO

	if val, _ := strconv.Atoi(ctx.GetN().GetText()); val < 0 {
		fmt.Println("ERROR_ILLEGAL_NEGATIVE_LITERAL")	
	}

	return env.Nat{}
}

func (v *TypeCheckVisitor) VisitVariant(ctx *parser.VariantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeCheckVisitor) VisitConstTrue(ctx *parser.ConstTrueContext) interface{} {
	//TODO

	return env.Bool{}
}

func (v *TypeCheckVisitor) VisitSubtract(ctx *parser.SubtractContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeCheckVisitor) VisitTypeCast(ctx *parser.TypeCastContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeCheckVisitor) VisitIf(ctx *parser.IfContext) interface{} {
	// TODO
    
	_, ok := ctx.GetCondition().Accept(v).(env.Bool)
	if !ok {
		TypeCheck(env.Bool{}, ctx.GetCondition().Accept(v))
		os.Exit(1)
	}

	theN, elsE := ctx.GetThenExpr().Accept(v).(env.Type), ctx.GetElseExpr().Accept(v).(env.Type)
	
	if !v.checking {
		throwAmbiguousType(theN)	
	} else {
		TypeCheck(v.checkingForType, theN)
	}

	TypeCheck(theN, elsE)

	f, fok := theN.(env.Sum)
	s, sok := elsE.(env.Sum)
	if fok && sok {
		return ascriptSumTypes(f, s)
	}

	
	
	if theN, ok := theN.(*env.Throw); ok && !isAny(theN.UnderlyingType) {
		return theN.UnderlyingType
	}

	return theN
}

func (v *TypeCheckVisitor) VisitApplication(ctx *parser.ApplicationContext) interface{} {
	// TODO

	// fmt.Printf("%T %T\n", ctx.GetFun(), ctx.GetArgs())
	funcType, ok := ctx.GetFun().Accept(v).(env.Func)

	if !ok {
		// fmt.Printf("%T\n", ctx.GetFun())
		fmt.Println("ERROR_NOT_A_FUNCTION")
		os.Exit(1)
	}

	throwAmbiguousType(funcType.Return)
	
	if len(funcType.Args) != len(ctx.GetArgs()) {
		fmt.Println("ERROR_INCORRECT_NUMBER_OF_ARGUMENTS")
		os.Exit(1)
	}


	for i, val := range ctx.GetArgs() {
		if funcType.IsAnonymous {
			TypeCheck(funcType.Args[i], val.Accept(v))
		} else {
			TypeCheck(funcType.Args[i], val.Accept(v))
		}
		// if val.Accept(v).(env.Type).Type() != funcType.Args[i].Type() {
		// 	fmt.Println("ERROR_UNEXPECTED_TYPE_FOR_EXPRESSION")
		// 	os.Exit(1)
		// }
	}

	return funcType.Return
}

func (v *TypeCheckVisitor) VisitDeref(ctx *parser.DerefContext) interface{} {
	ref, ok := ctx.GetExpr_().Accept(v).(env.Reference)
	if !ok {
		fmt.Println("ERROR_NOT_A_REFERENCE")
		os.Exit(1)
	}

	// if isAmbiguousType(ref) {
	// 	fmt.Println("ERROR_AMBIGUOUS_REFERENCE_TYPE")
	// 	os.Exit(1)
	// }

	return ref.UnderlyingType	
}

func (v *TypeCheckVisitor) VisitIsEmpty(ctx *parser.IsEmptyContext) interface{} {

	throwAmbiguousType(ctx.GetList().Accept(v))
	
	t, ok := ctx.GetList().Accept(v).(env.List); 
	if !ok {
		fmt.Println("ERROR_NOT_A_LIST")
		os.Exit(1)
	}

	if t.T == nil {
		fmt.Println("ERROR_AMBIGUOUS_LIST_TYPE")
		os.Exit(1)
	}
	
	return env.Bool{}
}

func (v *TypeCheckVisitor) VisitPanic(ctx *parser.PanicContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeCheckVisitor) VisitLessThanOrEqual(ctx *parser.LessThanOrEqualContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeCheckVisitor) VisitSucc(ctx *parser.SuccContext) interface{} {
	//TODO

	prev := v.checking
	prevType := v.checkingForType
	v.checking = true
	v.checkingForType = env.Nat{}

	s := ctx.Expr().Accept(v).(env.Type)
	TypeCheck(env.Nat{}, s)

	v.checkingForType = prevType
	v.checking = prev

	
	// if  s.Type() == "Nat" {
	// 	return env.Nat{}
	// }

	// fmt.Println("ERROR_UNEXPECTED_TYPE_FOR_EXPRESSION. Expected: Nat, got: ", s.Type())
	// os.Exit(1)

	return env.Nat{}
}

func (v *TypeCheckVisitor) VisitInl(ctx *parser.InlContext) interface{} {
	return env.Sum {
		Left: ctx.GetExpr_().Accept(v).(env.Type),
		Right: env.Any{},
	}
}

func (v *TypeCheckVisitor) VisitGreaterThanOrEqual(ctx *parser.GreaterThanOrEqualContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeCheckVisitor) VisitInr(ctx *parser.InrContext) interface{} {
	return env.Sum {
		Left: env.Any{},
		Right: ctx.GetExpr_().Accept(v).(env.Type),
	}
}

func (v *TypeCheckVisitor) VisitMatch(ctx *parser.MatchContext) interface{} {
	// fmt.Printf("%T\n", ctx.GetExpr_())

	v.env.Push()
	v.env.Put("-1", ctx.GetExpr_().Accept(v).(env.Type))

	cases := ctx.GetCases()

	if len(cases) <= 0 {
		fmt.Println("ERROR_ILLEGAL_EMPTY_MATCHING")
		os.Exit(1)
	}


	buff := v.patterns
	v.patterns = make(map[string]bool)
	v.GeneratePatternOptions(ctx.GetExpr_().Accept(v), "")

	// t := cases[0].Accept(v).(env.Type)

	// for _, val := range ctx.GetCases() {
	// 	fmt.Printf("%T %T\n", val.GetPattern_(), val.GetExpr_())
	// 	val.GetPattern_().Accept(v)
	// }

	for i := 0; i < len(cases); i++ {
		// fmt.Println(v.env.Check("0").Type(), cases[i].Accept(v).(env.Type).Type())
		TypeCheck(v.env.Check("0"), eraseLiterals(cases[i].Accept(v)))
	}

	for k, v := range v.patterns{
		if !v {
			fmt.Println("ERROR_NONEXHAUSTIVE_MATCH_PATTERNS. Key", k, "is not matched")
			os.Exit(1)
		}
	}

	v.patterns = buff
	v.env.Pop()
	return v.env.Check("0")
}

func (v *TypeCheckVisitor) VisitLogicNot(ctx *parser.LogicNotContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeCheckVisitor) VisitParenthesisedExpr(ctx *parser.ParenthesisedExprContext) interface{} {
	return ctx.GetExpr_().Accept(v)
}

func (v *TypeCheckVisitor) VisitTail(ctx *parser.TailContext) interface{} {

	throwAmbiguousType(ctx.GetList().Accept(v))
	
	t, ok := ctx.GetList().Accept(v).(env.List); 
	if !ok {
		fmt.Println("ERROR_NOT_A_LIST")
		os.Exit(1)
	}

	if t.T == nil {
		fmt.Println("ERROR_AMBIGUOUS_LIST_TYPE")
		os.Exit(1)
	}
	// t.IsLiteral = false
	return t
}

func (v *TypeCheckVisitor) VisitRecord(ctx *parser.RecordContext) interface{} {
	
	record := env.Record {
		Elements: make(map[string]env.Type, len(ctx.GetBindings())),
		IsLiteral: true,
	}

	for _, val := range ctx.GetBindings() {
		name := val.GetName().GetText()
		fieldType := val.GetRhs().Accept(v).(env.Type)
		// if isAmbiguousSumType(fieldType) {
		// 	// fmt.Println(fieldType.Type())
		// 	fmt.Println("ERROR_AMBIGUOUS_SUM_TYPE")
		// 	os.Exit(1)
		// }
	
		// if isAmbiguousType(fieldType) {
		// 	fmt.Println("ERROR_AMBIGUOUS_LIST_TYPE")
		// 	os.Exit(1)
		// }
	
		// if isAmbiguousType(fieldType) {
		// 	fmt.Println("ERROR_AMBIGUOUS_REFERENCE_TYPE")
		// 	os.Exit(1)
		// }
		funcType, ok := fieldType.(env.Func)
		if ok {
			funcType.IsAnonymous = false
			record.Elements[name] = funcType
			continue	
		}
		fieldType = eraseLiterals(fieldType)
		record.Elements[name] = fieldType
	}
	
	return record
}

func (v *TypeCheckVisitor) VisitLogicAnd(ctx *parser.LogicAndContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeCheckVisitor) VisitTypeApplication(ctx *parser.TypeApplicationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeCheckVisitor) VisitLetRec(ctx *parser.LetRecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeCheckVisitor) VisitLogicOr(ctx *parser.LogicOrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeCheckVisitor) VisitTryWith(ctx *parser.TryWithContext) interface{} {
	v.env.Push()
	tryType := ctx.GetTryExpr().Accept(v).(env.Type)
	_ = tryType

	expr := ctx.GetFallbackExpr().Accept(v)
	TypeCheck(tryType, expr)

	f, fok := tryType.(env.Sum)
	s, sok := expr.(env.Sum)
	if fok && sok {
		return ascriptSumTypes(f, s)
	}
	
	v.env.Pop()
	return expr
}

func (v *TypeCheckVisitor) VisitPred(ctx *parser.PredContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeCheckVisitor) VisitTypeAsc(ctx *parser.TypeAscContext) interface{} {
	
	expected := ctx.GetType_().Accept(v).(env.Type)
	given := ctx.GetExpr_().Accept(v).(env.Type)

	TypeCheck(expected, given)

	eraseLiterals(expected)

	return expected
}

func (v *TypeCheckVisitor) VisitNatRec(ctx *parser.NatRecContext) interface{} {
	//TODO
	prev := v.checking
	prevType := v.checkingForType
	v.checking = true
	v.checkingForType = env.Nat{}

	TypeCheck(env.Nat{}, ctx.GetN().Accept(v))

	v.checking = prev
	v.checkingForType = prevType

	zType := ctx.GetInitial().Accept(v).(env.Type)

	v.checking = true
	v.checkingForType = env.Func{
		Args: []env.Type{env.Nat{}},
		Return: env.Func{
			Args: []env.Type{zType},
			Return: zType,
		},
	}

	TypeCheck(
		env.Func{
			Args: []env.Type{env.Nat{}},
			Return: env.Func{
				Args: []env.Type{zType},
				Return: zType,
			},
		},
		ctx.GetStep().Accept(v),
	)

	v.checkingForType = prevType
	v.checking = prev

	return zType
}

func (v *TypeCheckVisitor) VisitUnfold(ctx *parser.UnfoldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeCheckVisitor) VisitRef(ctx *parser.RefContext) interface{} {
	return env.Reference{
		UnderlyingType: ctx.GetExpr_().Accept(v).(env.Type),
	}
}

func (v *TypeCheckVisitor) VisitDotTuple(ctx *parser.DotTupleContext) interface{} {
	// TODO
	tuple, ok := ctx.GetExpr_().Accept(v).(env.Tuple)
	if !ok {
		fmt.Println("ERROR_NOT_A_TUPLE")
		os.Exit(1)
	}

	for _, t := range tuple.Elements {
		throwAmbiguousType(t)
	}

	idx, err := strconv.Atoi(ctx.GetIndex().GetText())
	if err != nil {
		fmt.Println("ERROR_ILLEGAL_EXPRESSION")
		os.Exit(1)
	}

	if idx > len(tuple.Elements) || idx < 1 {
		fmt.Println("ERROR_TUPLE_INDEX_OUT_OF_BOUNDS")
		os.Exit(1)
	}

	return tuple.Elements[idx-1]
}

func (v *TypeCheckVisitor) VisitFix(ctx *parser.FixContext) interface{} {
	funcType, ok := ctx.GetExpr_().Accept(v).(env.Func)
	if !ok {
		fmt.Println("ERROR_NOT_A_FUNCTION")
		os.Exit(1)
	}

	if len(funcType.Args) != 1 {
		fmt.Println("ERROR_INCORRECT_NUMBER_OF_ARGUMENTS")
		os.Exit(1)
	}

	TypeCheck(funcType.Args[0], funcType.Return)

	return funcType.Return
}

func (v *TypeCheckVisitor) VisitLet(ctx *parser.LetContext) interface{} {
	// TODO

	pts := ctx.GetPatternBindings()
	v.env.Push()
	for _, val := range pts {
		b, ok := val.Accept(v).(env.Binding)
		if ! ok {
			fmt.Println("ERROR_UNEXPECTED_TYPE_FOR_EXPRESSION")
			os.Exit(1)
		}
		throwAmbiguousType(b.T)
		v.env.Put(b.Name, b.T)
	}
	t := ctx.GetBody().Accept(v).(env.Type)
	v.env.Pop()
	return t
}

func (v *TypeCheckVisitor) VisitAssign(ctx *parser.AssignContext) interface{} {

	ref, ok := ctx.GetLhs().Accept(v).(env.Reference)
	if !ok {
		fmt.Println("ERROR_NOT_A_REFERENCE")
		os.Exit(1)
	}

	TypeCheck(ref.UnderlyingType, ctx.GetRhs().Accept(v))

	return env.Unit{}
}

func (v *TypeCheckVisitor) VisitTuple(ctx *parser.TupleContext) interface{} {
	// TODO
	t := ctx.GetExprs()
	types := make([]env.Type, len(t))
	for i, val := range t {
		types[i] = eraseLiterals(val.Accept(v).(env.Type))
		// if isAmbiguousSumType(types[i]) {
		// 	// fmt.Println(types[i].Type())
		// 	fmt.Println("ERROR_AMBIGUOUS_SUM_TYPE")
		// 	os.Exit(1)
		// }
	
		// if isAmbiguousType(types[i]) {
		// 	fmt.Println("ERROR_AMBIGUOUS_LIST_TYPE")
		// 	os.Exit(1)
		// }
	
		// if isAmbiguousType(types[i]) {
		// 	fmt.Println("ERROR_AMBIGUOUS_REFERENCE_TYPE")
		// 	os.Exit(1)
		// }
	}

	return env.Tuple {
		Elements: types,
		IsLiteral: true,
	}
}

func (v *TypeCheckVisitor) VisitConsList(ctx *parser.ConsListContext) interface{} {

	t := ctx.GetHead().Accept(v).(env.Type)

	TypeCheck(env.List{
		T: t,		
	}, ctx.GetTail().Accept(v).(env.Type))

	return env.List{
		T: t,
		IsLiteral: true,
	}
}
// if consTailType.elementType != undefined && consTailType.elementType != consHeadType {error} 



func (v *TypeCheckVisitor) VisitPatternBinding(ctx *parser.PatternBindingContext) interface{} {
	// TODO
	t := ctx.GetPat().Accept(v).(env.Binding)
	t.T = ctx.GetRhs().Accept(v).(env.Type)
	return t
}

func (v *TypeCheckVisitor) VisitBinding(ctx *parser.BindingContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeCheckVisitor) VisitMatchCase(ctx *parser.MatchCaseContext) interface{} {

	pt := ctx.GetPattern_()
	expr := ctx.GetExpr_()

	v.curPattern = ""

	myPt := pt.Accept(v)

	switch myPt := myPt.(type) {
	default:
		return expr.Accept(v)
	case env.Binding:
		temp := v.env.Check(myPt.Name)
		v.env.Put(myPt.Name, myPt.T)
		res := expr.Accept(v)
		v.env.Put(myPt.Name, temp)
		return res
	case []env.Binding:
		var temp []env.Type
		for _, val := range myPt {
			temp = append(temp, v.env.Check(val.Name))
			v.env.Put(val.Name, val.T)
		}
		res := expr.Accept(v)
		for i, val := range myPt {
			v.env.Put(val.Name, temp[i])
		}
		return res
	}
}

func (v *TypeCheckVisitor) VisitPatternCons(ctx *parser.PatternConsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeCheckVisitor) VisitPatternTuple(ctx *parser.PatternTupleContext) interface{} {
	tuple, ok := v.env.Check("-1").(env.Tuple); 
	if !ok {
		fmt.Println("ERROR_UNEXPECTED_PATTERN_FOR_TYPE")
		os.Exit(1)
	}
	exprs := ctx.GetPatterns()
	if len(exprs) != len(tuple.Elements) {
		fmt.Println("ERROR_UNEXPECTED_PATTERN_FOR_TYPE")
		os.Exit(1)
	}

	var bindings []env.Binding
	buffPattern := v.curPattern
	v.curPattern += "-*#blabla"

	for i := range exprs {
		v.env.Put("-1", tuple.Elements[i])
		switch t := exprs[i].(type) {
		default:
			t.Accept(v)
		case *parser.PatternVarContext:
			bindings = append(bindings, t.Accept(v).(env.Binding))
		case *parser.PatternTupleContext:
			bindings = append(bindings, t.Accept(v).([]env.Binding)...)
		case *parser.PatternListContext:
			bindings = append(bindings, t.Accept(v).([]env.Binding)...)
		}
	}

	v.curPattern = buffPattern
	addIfExists(v.patterns, v.curPattern + tuple.Type())
	v.env.Put("-1", tuple)
	return bindings
}

func (v *TypeCheckVisitor) VisitPatternList(ctx *parser.PatternListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeCheckVisitor) VisitPatternRecord(ctx *parser.PatternRecordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeCheckVisitor) VisitPatternVariant(ctx *parser.PatternVariantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeCheckVisitor) VisitPatternAsc(ctx *parser.PatternAscContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeCheckVisitor) VisitPatternInt(ctx *parser.PatternIntContext) interface{} {

	if v.env.Check("-1").Type() == "" {
		//Catch pattern
		TypeCheck(env.Nat{}, v.env.Check(exceptionType))
		return env.Nat{}
	}

	// Pattern matching
	if _, ok := v.env.Check("-1").(env.Nat); !ok {
		fmt.Println("ERROR_UNEXPECTED_PATTERN_FOR_TYPE")
		os.Exit(1)
	}
	val, err := strconv.Atoi(ctx.GetN().GetText())
	if err != nil {
		fmt.Println("ERROR_UNEXPECTED_PATTERN_FOR_TYPE")
		os.Exit(1)
	}

	if val == 0 {
		addIfExists(v.patterns, v.curPattern + "nat 0")
		// v.patterns[v.curPattern + "nat 0"] = true
	}

	return env.Nat{}
}

func (v *TypeCheckVisitor) VisitPatternInr(ctx *parser.PatternInrContext) interface{} {
	
	temp := v.env.Check("-1")

	if temp.Type() == "" {
		//catch pattern
		exc, ok := v.env.Check(exceptionType).(env.Sum)
		if !ok {
			fmt.Println("ERROR_UNEXPECTED_PATTERN_FOR_TYPE")
			os.Exit(1)
		}

		v.env.Push()
		v.env.Put(exceptionType, exc.Right)
		x := ctx.GetPattern_().Accept(v)
		v.env.Pop()
		return x
	}

	if _, ok := temp.(env.Sum); !ok {
		fmt.Println("ERROR_UNEXPECTED_PATTERN_FOR_TYPE")
		os.Exit(1)
	}

	v.env.Put("-1", temp.(env.Sum).Right)
	buffPattern := v.curPattern
	v.curPattern += "inr "
	pattern := ctx.GetPattern_().Accept(v)	
	v.env.Put("-1", temp)
	// fmt.Printf("%T\n", ctx.GetPattern_())
	v.curPattern = buffPattern
	return pattern
}

func (v *TypeCheckVisitor) VisitPatternTrue(ctx *parser.PatternTrueContext) interface{} {
	if _, ok := v.env.Check("-1").(env.Bool); !ok {
		fmt.Println("ERROR_UNEXPECTED_PATTERN_FOR_TYPE")
		os.Exit(1)
	}
	// if _, ok := v.patterns[v.curPattern + "true"]; !ok {
	// 	fmt.Println("ERROR_UNEXPECTED_PATTERN_FOR_TYPE")
	// 	os.Exit(1)
	// }
	addIfExists(v.patterns, v.curPattern + "true")
	return env.Bool{
		Val: "true",
	}
}

func (v *TypeCheckVisitor) VisitPatternInl(ctx *parser.PatternInlContext) interface{} {
	temp := v.env.Check("-1")

	if temp.Type() == "" {
		//catch pattern
		exc, ok := v.env.Check(exceptionType).(env.Sum)
		if !ok {
			fmt.Println("ERROR_UNEXPECTED_PATTERN_FOR_TYPE")
			os.Exit(1)
		}

		v.env.Push()
		v.env.Put(exceptionType, exc.Right)
		x := ctx.GetPattern_().Accept(v)
		v.env.Pop()
		return x
	}

	if _, ok := temp.(env.Sum); !ok {
		fmt.Println("ERROR_UNEXPECTED_PATTERN_FOR_TYPE")
		os.Exit(1)
	}

	v.env.Put("-1", temp.(env.Sum).Left)
	buffPattern := v.curPattern
	v.curPattern += "inl "
	pattern := ctx.GetPattern_().Accept(v)	
	v.env.Put("-1", temp)
	// fmt.Printf("%T\n", ctx.GetPattern_())
	v.curPattern = buffPattern
	return pattern
}

func (v *TypeCheckVisitor) VisitPatternVar(ctx *parser.PatternVarContext) interface{} {
	// TODO
	// fmt.Println(ctx.GetName().GetText())
	if val := v.env.Check("-1"); val.Type() != "" {
		for key := range v.patterns {
			if startsWith(v.curPattern, key) {
				addIfExists(v.patterns, key)
			}
		}

		return env.Binding {
			Name: ctx.GetName().GetText(),
			T: val,
		}
	}

	if t := v.env.Check(exceptionType); t.Type() != "" {
		return env.Binding {
			Name: ctx.GetName().GetText(),
			T: t,
		}
	}

	return env.Binding {
		Name: ctx.GetName().GetText(),
	}
}

func (v *TypeCheckVisitor) VisitParenthesisedPattern(ctx *parser.ParenthesisedPatternContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeCheckVisitor) VisitPatternSucc(ctx *parser.PatternSuccContext) interface{} {
	if _, ok := v.env.Check("-1").(env.Nat); !ok {
		fmt.Println("ERROR_UNEXPECTED_PATTERN_FOR_TYPE")
		os.Exit(1)
	}
	switch x := ctx.GetPattern_().(type) {
	default:
		fmt.Println("ERROR_UNEXPECTED_PATTERN_FOR_TYPE")
		os.Exit(1)
	case *parser.PatternVarContext:
		r := x.Accept(v)
		// if !ok {}
		removeIfExists(v.patterns, v.curPattern + "nat 0")
		return r
	case *parser.PatternIntContext:
		// return x.Accept(v)
		r := x.Accept(v)
		// if !ok {}
		removeIfExists(v.patterns, v.curPattern + "nat 0")
		return r
	}
	return v.VisitChildren(ctx)
}

func (v *TypeCheckVisitor) VisitPatternFalse(ctx *parser.PatternFalseContext) interface{} {
	if _, ok := v.env.Check("-1").(env.Bool); !ok {
		fmt.Println("ERROR_UNEXPECTED_PATTERN_FOR_TYPE")
		os.Exit(1)
	}

	// if _, ok := v.patterns[v.curPattern + "false"]; !ok {
	// 	fmt.Println("ERROR_UNEXPECTED_PATTERN_FOR_TYPE")
	// 	os.Exit(1)
	// }
	// v.patterns[v.curPattern + "false"] = true
	addIfExists(v.patterns, v.curPattern + "false")
	return env.Bool{
		Val: "false",
	}
}

func (v *TypeCheckVisitor) VisitPatternUnit(ctx *parser.PatternUnitContext) interface{} {
	if _, ok := v.env.Check("-1").(env.Unit); !ok {
		fmt.Println("ERROR_UNEXPECTED_PATTERN_FOR_TYPE")
		os.Exit(1)
	}
	// v.patterns[v.curPattern + "unit"] = true
	addIfExists(v.patterns, v.curPattern + "unit")
	return env.Unit{}
}

func (v *TypeCheckVisitor) VisitPatternCastAs(ctx *parser.PatternCastAsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeCheckVisitor) VisitLabelledPattern(ctx *parser.LabelledPatternContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeCheckVisitor) VisitTypeTuple(ctx *parser.TypeTupleContext) interface{} {
	//TODO
	t := ctx.GetTypes()
	types := make([]env.Type, len(t))
	for i, val := range t {
		types[i] = val.Accept(v).(env.Type)
	}

	return env.Tuple {
		Elements: types,
	}
}

func (v *TypeCheckVisitor) VisitTypeTop(ctx *parser.TypeTopContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeCheckVisitor) VisitTypeBool(ctx *parser.TypeBoolContext) interface{} {
	//TODO
	return env.Bool{}
}

func (v *TypeCheckVisitor) VisitTypeRef(ctx *parser.TypeRefContext) interface{} {
	return env.Reference{
		UnderlyingType: ctx.GetType_().Accept(v).(env.Type),
	}
}

func (v *TypeCheckVisitor) VisitTypeRec(ctx *parser.TypeRecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeCheckVisitor) VisitTypeAuto(ctx *parser.TypeAutoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeCheckVisitor) VisitTypeSum(ctx *parser.TypeSumContext) interface{} {
	l := ctx.GetLeft().Accept(v).(env.Type)
	r := ctx.GetRight().Accept(v).(env.Type)
	// fmt.Printf("%T %T\n", l, r)
	return env.Sum{
		Left: l,
		Right: r,
	}
}

func (v *TypeCheckVisitor) VisitTypeVar(ctx *parser.TypeVarContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeCheckVisitor) VisitTypeVariant(ctx *parser.TypeVariantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeCheckVisitor) VisitTypeUnit(ctx *parser.TypeUnitContext) interface{} {
	//TODO

	return env.Unit{}
}

func (v *TypeCheckVisitor) VisitTypeNat(ctx *parser.TypeNatContext) interface{} {
	//TODO 

	return env.Nat{}
}

func (v *TypeCheckVisitor) VisitTypeBottom(ctx *parser.TypeBottomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeCheckVisitor) VisitTypeParens(ctx *parser.TypeParensContext) interface{} {
	// fmt.Printf("%T\n", ctx.GetType_())
	return ctx.GetType_().Accept(v).(env.Type)	
}

func (v *TypeCheckVisitor) VisitTypeFun(ctx *parser.TypeFunContext) interface{} {
	// TODO
	rawArgs := ctx.GetParamTypes()
	rawReturnType := ctx.GetReturnType()

	args := make([]env.Type, len(rawArgs))
	for i, val := range rawArgs {
		args[i] = val.Accept(v).(env.Type)
	}

	returnType := rawReturnType.Accept(v).(env.Type)

	return env.Func{
		Args: args,
		Return: returnType,
	}
}

func (v *TypeCheckVisitor) VisitTypeForAll(ctx *parser.TypeForAllContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeCheckVisitor) VisitTypeRecord(ctx *parser.TypeRecordContext) interface{} {

	record := env.Record{
		Elements: make(map[string]env.Type, len(ctx.GetFieldTypes())),
	}

	for _, val := range ctx.GetFieldTypes() {
		name := val.GetLabel().GetText()
		fieldType := val.GetType_().Accept(v).(env.Type)

		record.Elements[name] = fieldType
	}

	return record
}

func (v *TypeCheckVisitor) VisitTypeList(ctx *parser.TypeListContext) interface{} {
	return env.List{
		T: ctx.GetType_().Accept(v).(env.Type),
	}
}

func (v *TypeCheckVisitor) VisitRecordFieldType(ctx *parser.RecordFieldTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeCheckVisitor) VisitVariantFieldType(ctx *parser.VariantFieldTypeContext) interface{} {
	return v.VisitChildren(ctx)
}





