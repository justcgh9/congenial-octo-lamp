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
}

func TypeCheck(exp, given interface{}) {

	if exp.(env.Type).Type() == given.(env.Type).Type() {
		return
	}

	switch exp.(type) {
	default:
		fmt.Println("ERROR_UNEXPECTED_TYPE_FOR_EXPRESSION. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())
		os.Exit(1)
	case env.Func:
		ef := exp.(env.Func)
		_, ok := given.(env.Tuple)
		if ok {
			fmt.Println("ERROR_UNEXPECTED_TUPLE. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())
			os.Exit(1)
		}

		ft, ok := given.(env.Func)

		if !ok || !ft.IsAnonymous || ft.Return.Type() != ef.Return.Type() {
			fmt.Println("ERROR_UNEXPECTED_TYPE_FOR_EXPRESSION. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())
			os.Exit(1)
		}

		if len(ef.Args) != len(ft.Args) {
			fmt.Println("UNEXPECTED_NUMBER_OF_PARAMETERS_IN_LAMBDA. Expected ", len(ef.Args), " , given ", len(ft.Args))
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

		_, ok = given.(env.Tuple)
		if ok {
			fmt.Println("ERROR_UNEXPECTED_TUPLE. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())
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

		_, ok = given.(env.Tuple)
		if ok {
			fmt.Println("ERROR_UNEXPECTED_TUPLE. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())
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
		
	case env.Unit:
		ft, ok := given.(env.Func)
		if ok && ft.IsAnonymous {
			fmt.Println("ERROR_UNEXPECTED_LAMBDA. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())
			os.Exit(1)
		}

		_, ok = given.(env.Tuple)
		if ok {
			fmt.Println("ERROR_UNEXPECTED_TUPLE. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())
			os.Exit(1)
		}


		fmt.Println("ERROR_UNEXPECTED_TYPE_FOR_EXPRESSION. Expected ", exp.(env.Type).Type(), " , given ", given.(env.Type).Type())
		os.Exit(1)
		
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

	args := make([]env.Type, 0, len(ctx.GetParamDecls()))
	for _, val := range ctx.GetParamDecls() {
		// fmt.Printf("%T %s\n", val, val.GetName().GetText())
		args = append(args, val.Accept(v).(env.Type))
	}
	returnType := ctx.GetReturnType().Accept(v).(env.Type)
	
	v.env.Put(ctx.GetName().GetText(), env.Func{
		Args: args,
		Return: returnType,
	})
	// fmt.Printf("%T\n", ctx.GetReturnType())
	v.env.Push()

	for _, val := range ctx.GetParamDecls() {
		v.env.Put(val.GetName().GetText(), val.Accept(v).(env.Type))
	}

	ans := ctx.GetReturnExpr().Accept(v).(env.Type)
	// fmt.Printf("\n---------\n%T\n----------\n", ctx.GetReturnExpr())
	
	TypeCheck(returnType, ans)
	
	// fmt.Println(args)

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
	return v.VisitChildren(ctx)
}

func (v *TypeCheckVisitor) VisitDeclExceptionVariant(ctx *parser.DeclExceptionVariantContext) interface{} {
	return v.VisitChildren(ctx)
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
	nType := ctx.GetN().Accept(v).(env.Type)
	TypeCheck(env.Nat{}, nType)
	
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
	return v.VisitChildren(ctx)
}

func (v *TypeCheckVisitor) VisitGreaterThan(ctx *parser.GreaterThanContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeCheckVisitor) VisitEqual(ctx *parser.EqualContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeCheckVisitor) VisitThrow(ctx *parser.ThrowContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeCheckVisitor) VisitMultiply(ctx *parser.MultiplyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeCheckVisitor) VisitConstMemory(ctx *parser.ConstMemoryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeCheckVisitor) VisitList(ctx *parser.ListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeCheckVisitor) VisitTryCatch(ctx *parser.TryCatchContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeCheckVisitor) VisitTryCastAs(ctx *parser.TryCastAsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeCheckVisitor) VisitHead(ctx *parser.HeadContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeCheckVisitor) VisitTerminatingSemicolon(ctx *parser.TerminatingSemicolonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeCheckVisitor) VisitNotEqual(ctx *parser.NotEqualContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeCheckVisitor) VisitConstUnit(ctx *parser.ConstUnitContext) interface{} {
	//TODO

	return env.Unit{}
}

func (v *TypeCheckVisitor) VisitSequence(ctx *parser.SequenceContext) interface{} {
	return v.VisitChildren(ctx)
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

	returnType := ctx.GetReturnExpr().Accept(v).(env.Type)
	v.env.Pop()
	return env.Func{
		Args: []env.Type{ctx.GetParamDecls()[0].Accept(v).(env.Type)},
		Return: returnType,
		IsAnonymous: true,
	}
}

func (v *TypeCheckVisitor) VisitConstInt(ctx *parser.ConstIntContext) interface{} {
	//TODO

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
		fmt.Println("ERROR_UNEXPECTED_TYPE_FOR_EXPRESSION")
		os.Exit(1)
	}

	theN, elsE := ctx.GetThenExpr().Accept(v).(env.Type), ctx.GetThenExpr().Accept(v).(env.Type)
	if theN.Type() != elsE.Type() {
		fmt.Println("ERROR_UNEXPECTED_TYPE_FOR_EXPRESSION")
		os.Exit(1)
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

	
	if len(funcType.Args) != len(ctx.GetArgs()) {
		fmt.Println("ERROR_INCORRECT_NUMBER_OF_ARGUMENTS")
		os.Exit(1)
	}


	for i, val := range ctx.GetArgs() {
		TypeCheck(funcType.Args[i], val.Accept(v))
		// if val.Accept(v).(env.Type).Type() != funcType.Args[i].Type() {
		// 	fmt.Println("ERROR_UNEXPECTED_TYPE_FOR_EXPRESSION")
		// 	os.Exit(1)
		// }
	}

	return funcType.Return
}

func (v *TypeCheckVisitor) VisitDeref(ctx *parser.DerefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeCheckVisitor) VisitIsEmpty(ctx *parser.IsEmptyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeCheckVisitor) VisitPanic(ctx *parser.PanicContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeCheckVisitor) VisitLessThanOrEqual(ctx *parser.LessThanOrEqualContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeCheckVisitor) VisitSucc(ctx *parser.SuccContext) interface{} {
	//TODO

	s := ctx.Expr().Accept(v).(env.Type)
	TypeCheck(env.Nat{}, s)
	// if  s.Type() == "Nat" {
	// 	return env.Nat{}
	// }

	// fmt.Println("ERROR_UNEXPECTED_TYPE_FOR_EXPRESSION. Expected: Nat, got: ", s.Type())
	// os.Exit(1)

	return env.Nat{}
}

func (v *TypeCheckVisitor) VisitInl(ctx *parser.InlContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeCheckVisitor) VisitGreaterThanOrEqual(ctx *parser.GreaterThanOrEqualContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeCheckVisitor) VisitInr(ctx *parser.InrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeCheckVisitor) VisitMatch(ctx *parser.MatchContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeCheckVisitor) VisitLogicNot(ctx *parser.LogicNotContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeCheckVisitor) VisitParenthesisedExpr(ctx *parser.ParenthesisedExprContext) interface{} {
	return ctx.GetExpr_().Accept(v)
}

func (v *TypeCheckVisitor) VisitTail(ctx *parser.TailContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeCheckVisitor) VisitRecord(ctx *parser.RecordContext) interface{} {
	return v.VisitChildren(ctx)
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
	return v.VisitChildren(ctx)
}

func (v *TypeCheckVisitor) VisitPred(ctx *parser.PredContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeCheckVisitor) VisitTypeAsc(ctx *parser.TypeAscContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeCheckVisitor) VisitNatRec(ctx *parser.NatRecContext) interface{} {
	//TODO

	TypeCheck(env.Nat{}, ctx.GetN().Accept(v))
	// _, ok := ctx.GetN().Accept(v).(env.Nat)
	// if !ok {
	// 	fmt.Println("ERROR_UNEXPECTED_TYPE_FOR_EXPRESSION")
	// 	os.Exit(1)
	// }

	zType := ctx.GetInitial().Accept(v).(env.Type)
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
	// stepType, ok := ctx.GetStep().Accept(v).(env.Func)
	// if !ok {
	// 	// fmt.Printf("%T\n", ctx.GetInitial())
	// 	fmt.Println("ERROR_UNEXPECTED_TYPE_FOR_EXPRESSION")
	// 	os.Exit(1)
	// }

	// if len(stepType.Args) != 1 || stepType.Args[0].Type() != "Nat" {
	// 	fmt.Println("ERROR_UNEXPECTED_TYPE_FOR_EXPRESSION")
	// 	os.Exit(1)
	// }

	// secondStepType, ok := stepType.Return.(env.Func)
	// if !ok {
	// 	fmt.Println("ERROR_UNEXPECTED_TYPE_FOR_EXPRESSION")
	// 	os.Exit(1)
	// }

	// if len(secondStepType.Args) != 1 || secondStepType.Args[0].Type() != zType.Type() || secondStepType.Return.Type() != zType.Type() {
	// 	fmt.Println("ERROR_UNEXPECTED_TYPE_FOR_EXPRESSION")
	// 	os.Exit(1)
	// }

	// fmt.Printf("%T\n", stepType.Return.(env.Func).Args[0])

	return zType
}

func (v *TypeCheckVisitor) VisitUnfold(ctx *parser.UnfoldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeCheckVisitor) VisitRef(ctx *parser.RefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeCheckVisitor) VisitDotTuple(ctx *parser.DotTupleContext) interface{} {
	// TODO
	tuple, ok := ctx.GetExpr_().Accept(v).(env.Tuple)
	if !ok {
		fmt.Println("ERROR_NOT_A_TUPLE")
		os.Exit(1)
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
	return v.VisitChildren(ctx)
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
		v.env.Put(b.Name, b.T)
	}
	t := ctx.GetBody().Accept(v).(env.Type)
	v.env.Pop()
	return t
}

func (v *TypeCheckVisitor) VisitAssign(ctx *parser.AssignContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeCheckVisitor) VisitTuple(ctx *parser.TupleContext) interface{} {
	// TODO
	t := ctx.GetExprs()
	types := make([]env.Type, len(t))
	for i, val := range t {
		types[i] = val.Accept(v).(env.Type)
	}

	return env.Tuple {
		Elements: types,
	}
}

func (v *TypeCheckVisitor) VisitConsList(ctx *parser.ConsListContext) interface{} {
	return v.VisitChildren(ctx)
}

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
	return v.VisitChildren(ctx)
}

func (v *TypeCheckVisitor) VisitPatternCons(ctx *parser.PatternConsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeCheckVisitor) VisitPatternTuple(ctx *parser.PatternTupleContext) interface{} {
	return v.VisitChildren(ctx)
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
	return v.VisitChildren(ctx)
}

func (v *TypeCheckVisitor) VisitPatternInr(ctx *parser.PatternInrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeCheckVisitor) VisitPatternTrue(ctx *parser.PatternTrueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeCheckVisitor) VisitPatternInl(ctx *parser.PatternInlContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeCheckVisitor) VisitPatternVar(ctx *parser.PatternVarContext) interface{} {
	// TODO

	return env.Binding {
		Name: ctx.GetName().GetText(),
	}
}

func (v *TypeCheckVisitor) VisitParenthesisedPattern(ctx *parser.ParenthesisedPatternContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeCheckVisitor) VisitPatternSucc(ctx *parser.PatternSuccContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeCheckVisitor) VisitPatternFalse(ctx *parser.PatternFalseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeCheckVisitor) VisitPatternUnit(ctx *parser.PatternUnitContext) interface{} {
	return v.VisitChildren(ctx)
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
	return v.VisitChildren(ctx)
}

func (v *TypeCheckVisitor) VisitTypeRec(ctx *parser.TypeRecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeCheckVisitor) VisitTypeAuto(ctx *parser.TypeAutoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeCheckVisitor) VisitTypeSum(ctx *parser.TypeSumContext) interface{} {
	return v.VisitChildren(ctx)
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
	return v.VisitChildren(ctx)
}

func (v *TypeCheckVisitor) VisitTypeList(ctx *parser.TypeListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeCheckVisitor) VisitRecordFieldType(ctx *parser.RecordFieldTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeCheckVisitor) VisitVariantFieldType(ctx *parser.VariantFieldTypeContext) interface{} {
	return v.VisitChildren(ctx)
}





