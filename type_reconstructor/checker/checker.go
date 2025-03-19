package checker

import (
	// "fmt"
	"fmt"
	"stella-implementation-in-go/parser"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

var constraints []Constraint

var typeVarCounter = 0

func NewTypeVar() Type {
	typeVarCounter++
    return Type{
        Kind: UnificationVar,
        Data: TypeVar{ID: typeVarCounter},
    }
}


type TypeReconstructionVisitor struct {
	*antlr.BaseParseTreeVisitor
	Stack Stack
	Env Env
	ExceptionType Type
	PropagatedType Type
}

func (v *TypeReconstructionVisitor) VisitStart_Program(ctx *parser.Start_ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeReconstructionVisitor) VisitStart_Expr(ctx *parser.Start_ExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeReconstructionVisitor) VisitStart_Type(ctx *parser.Start_TypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeReconstructionVisitor) VisitProgram(ctx *parser.ProgramContext) interface{} {
	v.Env.Push()
	defer v.Env.Pop()

	for _, decl := range ctx.GetDecls() {
		v.Env.Push()
		val, ok := decl.(*parser.DeclFunContext)
		if !ok {continue}

		args := make([]Type, 0, len(val.GetParamDecls()))

		for _, val := range val.GetParamDecls() {
			args = append(args, val.Accept(v).(Type))
		}

		returnType := val.GetReturnType().Accept(v).(Type)
		
		v.Env.Pop()
		v.Env.Put(val.GetName().GetText(), NewFunc(
			args[0],
			returnType,
		))
	}

	for _, decl := range ctx.GetDecls() {
		decl.Accept(v)
		for _, constraint := range constraints {
			fmt.Println(constraint.Left, "=", constraint.Right)
		}
		_, err := unifyConstraints(constraints)
		if err != nil {
			fmt.Println(err.Error())
			UnexpectedTypeForExpression()
		}
	}	

	val:= v.Env.Check("main")
	if val.Kind != Func  {
		MissingMain()
	}

	return v.VisitChildren(ctx)
}

func (v *TypeReconstructionVisitor) VisitLanguageCore(ctx *parser.LanguageCoreContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeReconstructionVisitor) VisitAnExtension(ctx *parser.AnExtensionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeReconstructionVisitor) VisitDeclFun(ctx *parser.DeclFunContext) interface{} {
	
	v.Env.Push()
	defer v.Env.Pop()

	givenFunc  := v.Env.Check(ctx.GetName().GetText()).Data.(FuncType)

	argType    := givenFunc.Param 
	v.Env.Put(ctx.GetParamDecls()[0].GetName().GetText(),argType)

	returnType := ctx.GetReturnExpr().Accept(v).(Type)

	function := NewFunc(
		argType,
		returnType,		
	)
	
	AddConstraint(
		NewConstraint(
			returnType,
			givenFunc.Return,
		),
		&constraints,
	)
	
	return function
	
}

func (v *TypeReconstructionVisitor) VisitDeclFunGeneric(ctx *parser.DeclFunGenericContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeReconstructionVisitor) VisitDeclTypeAlias(ctx *parser.DeclTypeAliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeReconstructionVisitor) VisitDeclExceptionType(ctx *parser.DeclExceptionTypeContext) interface{} {
	
	v.ExceptionType = ctx.GetExceptionType().Accept(v).(Type)

	return nil
}

func (v *TypeReconstructionVisitor) VisitDeclExceptionVariant(ctx *parser.DeclExceptionVariantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeReconstructionVisitor) VisitInlineAnnotation(ctx *parser.InlineAnnotationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeReconstructionVisitor) VisitParamDecl(ctx *parser.ParamDeclContext) interface{} {

	paramType := ctx.GetParamType().Accept(v).(Type)

	v.Env.Put(
		ctx.GetName().GetText(),
		paramType,
	)

	return paramType
}

func (v *TypeReconstructionVisitor) VisitFold(ctx *parser.FoldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeReconstructionVisitor) VisitAdd(ctx *parser.AddContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeReconstructionVisitor) VisitIsZero(ctx *parser.IsZeroContext) interface{} {
	contentType := ctx.GetN().Accept(v).(Type)

	switch contentType.Kind {
	default:
		UnexpectedTypeForExpression()
	case Int:
		break
	case UnificationVar:
		AddConstraint(
			NewConstraint(contentType, NewInt()),
			&constraints,
		)
	}

	return NewBool()
}

func (v *TypeReconstructionVisitor) VisitVar(ctx *parser.VarContext) interface{} {
	variable := v.Env.Check(ctx.GetName().GetText())
	if variable.Kind == "" {
		UndefinedVariable()
	}
	return variable
}

func (v *TypeReconstructionVisitor) VisitTypeAbstraction(ctx *parser.TypeAbstractionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeReconstructionVisitor) VisitDivide(ctx *parser.DivideContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeReconstructionVisitor) VisitLessThan(ctx *parser.LessThanContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeReconstructionVisitor) VisitDotRecord(ctx *parser.DotRecordContext) interface{} {
	record := ctx.GetExpr_().Accept(v).(Type)

	label := ctx.GetLabel().GetText()

	switch record.Kind{
	default:
		UnexpectedTypeForExpression()
	case Record:

		res, ok := record.Data.(RecordType).Content[label]
		if !ok {
			MissingRecordFields()
		}
		return res

	case UnificationVar:
		x := NewTypeVar()

		AddConstraint(
			NewConstraint(
				record,
				NewRecord(map[string]Type{
					label: x,
				}),
			),
			&constraints,
		)
		
		return x
	}

	return v.VisitChildren(ctx)
}

func (v *TypeReconstructionVisitor) VisitGreaterThan(ctx *parser.GreaterThanContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeReconstructionVisitor) VisitEqual(ctx *parser.EqualContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeReconstructionVisitor) VisitThrow(ctx *parser.ThrowContext) interface{} {

	if v.ExceptionType.Kind == "" {
		UndeclaredExceptionType()
	}

	exception := ctx.GetExpr_().Accept(v).(Type)

	AddConstraint(
		NewConstraint(
			exception,
			v.ExceptionType,
		),
		&constraints,
	)

	return NewTypeVar()
}

func (v *TypeReconstructionVisitor) VisitMultiply(ctx *parser.MultiplyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeReconstructionVisitor) VisitConstMemory(ctx *parser.ConstMemoryContext) interface{} {
	return NewReference(
		NewTypeVar(),
	)
}

func (v *TypeReconstructionVisitor) VisitList(ctx *parser.ListContext) interface{} {
	if len(ctx.GetExprs()) == 0 {
		return NewList(NewTypeVar())
	}

	exprs := ctx.GetExprs()
	contentType := exprs[0].Accept(v).(Type)

	for _, expr := range exprs[1:] {
		exprType := expr.Accept(v).(Type)
		AddConstraint(NewConstraint(
			contentType,
			exprType,
		), &constraints)
	}

	return NewList(contentType)
}

func (v *TypeReconstructionVisitor) VisitTryCatch(ctx *parser.TryCatchContext) interface{} {

	v.Env.Push()
	defer v.Env.Pop()

	try := ctx.GetTryExpr().Accept(v).(Type)

	x := ctx.GetPat().Accept(v).(Type)

	AddConstraint(
		NewConstraint(
			x,
			v.ExceptionType,
		),
		&constraints,
	)

	AddConstraint(
		NewConstraint(
			try,
			ctx.GetFallbackExpr().Accept(v).(Type),
		),
		&constraints,
	)

	return try
}

func (v *TypeReconstructionVisitor) VisitTryCastAs(ctx *parser.TryCastAsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeReconstructionVisitor) VisitHead(ctx *parser.HeadContext) interface{} {
	list := ctx.GetList().Accept(v).(Type)

	fresh := NewTypeVar()

	AddConstraint(
		NewConstraint(
			list,
			NewList(fresh),
		),
		&constraints,
	)

	return fresh
}

func (v *TypeReconstructionVisitor) VisitTerminatingSemicolon(ctx *parser.TerminatingSemicolonContext) interface{} {
	return ctx.GetExpr_().Accept(v)
}

func (v *TypeReconstructionVisitor) VisitNotEqual(ctx *parser.NotEqualContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeReconstructionVisitor) VisitConstUnit(ctx *parser.ConstUnitContext) interface{} {
	return NewUnit()
}

func (v *TypeReconstructionVisitor) VisitSequence(ctx *parser.SequenceContext) interface{} {

	left := ctx.GetExpr1().Accept(v).(Type)

	AddConstraint(
		NewConstraint(
			left,
			NewUnit(),
		),
		&constraints,
	)

	return ctx.GetExpr2().Accept(v)
}

func (v *TypeReconstructionVisitor) VisitConstFalse(ctx *parser.ConstFalseContext) interface{} {
	return NewBool()
}

func (v *TypeReconstructionVisitor) VisitAbstraction(ctx *parser.AbstractionContext) interface{} {

	v.Env.Push()
	defer v.Env.Pop()

	function := NewFunc(
		ctx.GetParamDecls()[0].Accept(v).(Type),
		ctx.GetReturnExpr().Accept(v).(Type),		
	)
	
	return function
}

func (v *TypeReconstructionVisitor) VisitConstInt(ctx *parser.ConstIntContext) interface{} {
	return NewInt()
}

func (v *TypeReconstructionVisitor) VisitVariant(ctx *parser.VariantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeReconstructionVisitor) VisitConstTrue(ctx *parser.ConstTrueContext) interface{} {
	return NewBool()
}

func (v *TypeReconstructionVisitor) VisitSubtract(ctx *parser.SubtractContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeReconstructionVisitor) VisitTypeCast(ctx *parser.TypeCastContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeReconstructionVisitor) VisitIf(ctx *parser.IfContext) interface{} {
	
	predicate := ctx.GetCondition().Accept(v).(Type)

	ifBlock := ctx.GetThenExpr().Accept(v).(Type)
	elseBlock := ctx.GetElseExpr().Accept(v).(Type)

	switch predicate.Kind {
	default:
		UnexpectedTypeForExpression()
	case Bool:
		break
	case UnificationVar:
		AddConstraint(
			NewConstraint(predicate, NewBool()),
			&constraints,
		)
	}

	AddConstraint(
		NewConstraint(ifBlock, elseBlock),
		&constraints,
	)

	return ifBlock
}

func (v *TypeReconstructionVisitor) VisitApplication(ctx *parser.ApplicationContext) interface{} {
	functionType := ctx.GetFun().Accept(v).(Type)
	argType 	 := ctx.GetArgs()[0].Accept(v).(Type)

	//fmt.Println("application", typeVarCounter + 1)
	x := NewTypeVar()

	AddConstraint(
		NewConstraint(
			functionType,
			NewFunc(argType, x),
		),
		&constraints,
	)

	return x
}

func (v *TypeReconstructionVisitor) VisitDeref(ctx *parser.DerefContext) interface{} {
	
	ref := ctx.GetExpr_().Accept(v).(Type)

	switch ref.Kind {
	default:
		NotAReference()
	case Reference:
		return ref.Data.(ReferenceType).Content
	case UnificationVar:
		deref := NewTypeVar()
		AddConstraint(
			NewConstraint(
				ref,
				NewReference(
					deref,
				),
			),
			&constraints,
		)
		return deref
	}

	return v.VisitChildren(ctx)
}

func (v *TypeReconstructionVisitor) VisitIsEmpty(ctx *parser.IsEmptyContext) interface{} {
	list := ctx.GetList().Accept(v).(Type)

	fresh := NewList(NewTypeVar())

	AddConstraint(
		NewConstraint(
			list,
			fresh,
		),
		&constraints,
	)

	return NewBool()
}

func (v *TypeReconstructionVisitor) VisitPanic(ctx *parser.PanicContext) interface{} {
	return NewTypeVar()
}

func (v *TypeReconstructionVisitor) VisitLessThanOrEqual(ctx *parser.LessThanOrEqualContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeReconstructionVisitor) VisitSucc(ctx *parser.SuccContext) interface{} {

	contentType := ctx.GetN().Accept(v).(Type)

	switch contentType.Kind {
	default:
		UnexpectedTypeForExpression()
	case Int:
		break
	case UnificationVar:
		AddConstraint(
			NewConstraint(contentType, NewInt()),
			&constraints,
		)
	}

	return NewInt()
}

func (v *TypeReconstructionVisitor) VisitInl(ctx *parser.InlContext) interface{} {
	return NewSum(
		ctx.GetExpr_().Accept(v).(Type),
		NewTypeVar(),
	)
}

func (v *TypeReconstructionVisitor) VisitGreaterThanOrEqual(ctx *parser.GreaterThanOrEqualContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeReconstructionVisitor) VisitInr(ctx *parser.InrContext) interface{} {
	return NewSum(
		NewTypeVar(),
		ctx.GetExpr_().Accept(v).(Type),
	)
}

func (v *TypeReconstructionVisitor) VisitMatch(ctx *parser.MatchContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeReconstructionVisitor) VisitLogicNot(ctx *parser.LogicNotContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeReconstructionVisitor) VisitParenthesisedExpr(ctx *parser.ParenthesisedExprContext) interface{} {
	return ctx.GetExpr_().Accept(v)
}

func (v *TypeReconstructionVisitor) VisitTail(ctx *parser.TailContext) interface{} {
	list := ctx.GetList().Accept(v).(Type)

	fresh := NewList(NewTypeVar())

	AddConstraint(
		NewConstraint(
			list,
			fresh,
		),
		&constraints,
	)

	return fresh
}

func (v *TypeReconstructionVisitor) VisitRecord(ctx *parser.RecordContext) interface{} {
	types := map[string]Type{}

	for _, expr := range ctx.GetBindings() {
		types[expr.GetName().GetText()] = expr.GetRhs().Accept(v).(Type)
	}

	return NewRecord(types)
}

func (v *TypeReconstructionVisitor) VisitLogicAnd(ctx *parser.LogicAndContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeReconstructionVisitor) VisitTypeApplication(ctx *parser.TypeApplicationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeReconstructionVisitor) VisitLetRec(ctx *parser.LetRecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeReconstructionVisitor) VisitLogicOr(ctx *parser.LogicOrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeReconstructionVisitor) VisitTryWith(ctx *parser.TryWithContext) interface{} {

	if v.ExceptionType.Kind == "" {
		UndeclaredExceptionType()
	}

	tryExpr  := ctx.GetTryExpr().Accept(v).(Type)
	fallback := ctx.GetFallbackExpr().Accept(v).(Type)

	AddConstraint(
		NewConstraint(
			tryExpr,
			fallback,
		),
		&constraints,
	)

	return tryExpr
}

func (v *TypeReconstructionVisitor) VisitPred(ctx *parser.PredContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeReconstructionVisitor) VisitTypeAsc(ctx *parser.TypeAscContext) interface{} {

	AscribingType := ctx.GetType_().Accept(v).(Type)

	AscribingExpr := ctx.GetExpr_().Accept(v).(Type)

	AddConstraint(
		NewConstraint(
			AscribingExpr,
			AscribingType,
		),
		&constraints,
	)

	return AscribingType
}

func (v *TypeReconstructionVisitor) VisitNatRec(ctx *parser.NatRecContext) interface{} {

	n := ctx.GetN().Accept(v).(Type)

	AddConstraint(
		NewConstraint(
			n,
			NewInt(),
		),
		&constraints,
	)

	z := ctx.GetInitial().Accept(v).(Type)

	s := ctx.GetStep().Accept(v).(Type)

	AddConstraint(
		NewConstraint(
			s,
			NewFunc(
				NewInt(),
				NewFunc(z, z),
			),
		),
		&constraints,
	)

	return z
}

func (v *TypeReconstructionVisitor) VisitUnfold(ctx *parser.UnfoldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeReconstructionVisitor) VisitRef(ctx *parser.RefContext) interface{} {
	return NewReference(
		ctx.GetExpr_().Accept(v).(Type),
	)
}

func (v *TypeReconstructionVisitor) VisitDotTuple(ctx *parser.DotTupleContext) interface{} {

	tuple := ctx.GetExpr_().Accept(v).(Type)

	index, _ := strconv.Atoi(ctx.GetIndex().GetText())

	switch tuple.Kind{
	default:
		UnexpectedTypeForExpression()
	case Tuple:
		return tuple.Data.(TupleType).Content[index - 1]
	case UnificationVar:
		x := NewTypeVar()
		y := NewTypeVar()

		AddConstraint(
			NewConstraint(
				tuple,
				NewTuple([]Type{x, y}),
			),
			&constraints,
		)

		if index == 1 {return x}
		if index == 2 {return y}
	}

	return v.VisitChildren(ctx)
}

func (v *TypeReconstructionVisitor) VisitFix(ctx *parser.FixContext) interface{} {

	x := NewTypeVar()

	function := ctx.GetExpr_().Accept(v).(Type)

	AddConstraint(
		NewConstraint(
			function,
			NewFunc(x, x),
		),
		&constraints,
	)

	return x
}

func (v *TypeReconstructionVisitor) VisitLet(ctx *parser.LetContext) interface{} {

	v.Env.Push()
	defer v.Env.Pop()

	for _, binding := range ctx.GetPatternBindings() {
		_ = binding.Accept(v).(Type)
	}

	return ctx.GetBody().Accept(v)
}

func (v *TypeReconstructionVisitor) VisitAssign(ctx *parser.AssignContext) interface{} {

	ref := ctx.GetLhs().Accept(v).(Type)

	val := ctx.GetRhs().Accept(v).(Type)

	switch ref.Kind {
	default:
		NotAReference()
	case Reference:
		AddConstraint(
			NewConstraint(ref.Data.(ReferenceType).Content, val),
			&constraints,
		)
	case UnificationVar:
		AddConstraint(
			NewConstraint(ref, NewReference(val)),
			&constraints,
		)
	}

	return NewUnit()
}

func (v *TypeReconstructionVisitor) VisitTuple(ctx *parser.TupleContext) interface{} {

	types := []Type{}

	for _, expr := range ctx.GetExprs() {
		types = append(types, expr.Accept(v).(Type))
	}

	return NewTuple(types)
}

func (v *TypeReconstructionVisitor) VisitConsList(ctx *parser.ConsListContext) interface{} {

	head := ctx.GetHead().Accept(v).(Type)
	list := ctx.GetTail().Accept(v).(Type)

	AddConstraint(
		NewConstraint(
			NewList(head),
			list,
		),
		&constraints,
	)

	return list
}

func (v *TypeReconstructionVisitor) VisitPatternBinding(ctx *parser.PatternBindingContext) interface{} {
	
	x := ctx.GetPat().Accept(v).(Type)

	AddConstraint(
		NewConstraint(
			x,
			ctx.GetRhs().Accept(v).(Type),
		),
		&constraints,
	)

	return x
}

func (v *TypeReconstructionVisitor) VisitBinding(ctx *parser.BindingContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeReconstructionVisitor) VisitMatchCase(ctx *parser.MatchCaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeReconstructionVisitor) VisitPatternCons(ctx *parser.PatternConsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeReconstructionVisitor) VisitPatternTuple(ctx *parser.PatternTupleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeReconstructionVisitor) VisitPatternList(ctx *parser.PatternListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeReconstructionVisitor) VisitPatternRecord(ctx *parser.PatternRecordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeReconstructionVisitor) VisitPatternVariant(ctx *parser.PatternVariantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeReconstructionVisitor) VisitPatternAsc(ctx *parser.PatternAscContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeReconstructionVisitor) VisitPatternInt(ctx *parser.PatternIntContext) interface{} {
	return NewInt()
}

func (v *TypeReconstructionVisitor) VisitPatternInr(ctx *parser.PatternInrContext) interface{} {
	return NewSum(
		NewTypeVar(),
		ctx.GetPattern_().Accept(v).(Type),
	)
}

func (v *TypeReconstructionVisitor) VisitPatternTrue(ctx *parser.PatternTrueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeReconstructionVisitor) VisitPatternInl(ctx *parser.PatternInlContext) interface{} {
	return NewSum(
		ctx.GetPattern_().Accept(v).(Type),
		NewTypeVar(),
	)
}

func (v *TypeReconstructionVisitor) VisitPatternVar(ctx *parser.PatternVarContext) interface{} {

	x := NewTypeVar()

	v.Env.Put(ctx.GetName().GetText(), x)

	return x
}

func (v *TypeReconstructionVisitor) VisitParenthesisedPattern(ctx *parser.ParenthesisedPatternContext) interface{} {
	return ctx.GetPattern_().Accept(v)
}

func (v *TypeReconstructionVisitor) VisitPatternSucc(ctx *parser.PatternSuccContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeReconstructionVisitor) VisitPatternFalse(ctx *parser.PatternFalseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeReconstructionVisitor) VisitPatternUnit(ctx *parser.PatternUnitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeReconstructionVisitor) VisitPatternCastAs(ctx *parser.PatternCastAsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeReconstructionVisitor) VisitLabelledPattern(ctx *parser.LabelledPatternContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeReconstructionVisitor) VisitTypeTuple(ctx *parser.TypeTupleContext) interface{} {
	types := []Type{}

	for _, expr := range ctx.GetTypes() {
		types = append(types, expr.Accept(v).(Type))
	}

	return NewTuple(types)
}

func (v *TypeReconstructionVisitor) VisitTypeTop(ctx *parser.TypeTopContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeReconstructionVisitor) VisitTypeBool(ctx *parser.TypeBoolContext) interface{} {
	return NewBool()
}

func (v *TypeReconstructionVisitor) VisitTypeRef(ctx *parser.TypeRefContext) interface{} {
	return NewReference(
		ctx.GetType_().Accept(v).(Type),
	)
}

func (v *TypeReconstructionVisitor) VisitTypeRec(ctx *parser.TypeRecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeReconstructionVisitor) VisitTypeAuto(ctx *parser.TypeAutoContext) interface{} {
	//fmt.Println("type auto", typeVarCounter + 1)
	return NewTypeVar()
}

func (v *TypeReconstructionVisitor) VisitTypeSum(ctx *parser.TypeSumContext) interface{} {
	return NewSum(
		ctx.GetLeft().Accept(v).(Type),
		ctx.GetRight().Accept(v).(Type),
	)
}

func (v *TypeReconstructionVisitor) VisitTypeVar(ctx *parser.TypeVarContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeReconstructionVisitor) VisitTypeVariant(ctx *parser.TypeVariantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeReconstructionVisitor) VisitTypeUnit(ctx *parser.TypeUnitContext) interface{} {
	return NewUnit()
}

func (v *TypeReconstructionVisitor) VisitTypeNat(ctx *parser.TypeNatContext) interface{} {
	return NewInt()
}

func (v *TypeReconstructionVisitor) VisitTypeBottom(ctx *parser.TypeBottomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeReconstructionVisitor) VisitTypeParens(ctx *parser.TypeParensContext) interface{} {
	return ctx.GetType_().Accept(v)
}

func (v *TypeReconstructionVisitor) VisitTypeFun(ctx *parser.TypeFunContext) interface{} {
	return NewFunc(
		ctx.GetParamTypes()[0].Accept(v).(Type),
		ctx.GetReturnType().Accept(v).(Type),
	)
}

func (v *TypeReconstructionVisitor) VisitTypeForAll(ctx *parser.TypeForAllContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeReconstructionVisitor) VisitTypeRecord(ctx *parser.TypeRecordContext) interface{} {

	content := map[string]Type{}

	for _, binding := range ctx.GetFieldTypes() {
		content[binding.GetLabel().GetText()] = binding.GetType_().Accept(v).(Type)
	}

	return NewRecord(content)
}

func (v *TypeReconstructionVisitor) VisitTypeList(ctx *parser.TypeListContext) interface{} {
	return NewList(ctx.GetType_().Accept(v).(Type))
}

func (v *TypeReconstructionVisitor) VisitRecordFieldType(ctx *parser.RecordFieldTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeReconstructionVisitor) VisitVariantFieldType(ctx *parser.VariantFieldTypeContext) interface{} {
	return v.VisitChildren(ctx)
}
