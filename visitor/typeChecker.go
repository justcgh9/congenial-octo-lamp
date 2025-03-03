package visitor

import (
	"fmt"
	"os"
	"stella-implementation-in-go/env"
	"stella-implementation-in-go/parser"
	"strconv"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

var (
	inlFlag bool = false
	inrFlag bool = false
)

type Visitor struct {
	*antlr.BaseParseTreeVisitor
	env             	env.Env 		// our linked list of contexts
	checking        	bool			// checking or inferring the type
	checkingForType 	env.Type		// Type we are checking for
	matchType			env.Type
	subtyping 			int
	ambiguousAsBottom	bool
}

func (v *Visitor) Visit(tree antlr.ParseTree) interface{} {
	panic("unimplemented")
}

func (v *Visitor) VisitAbstraction(ctx *parser.AbstractionContext) interface{} {
	v.env.Push()
	defer v.env.Pop()

	if v.checking && v.subtyping == 1 {
		
		expected, ok := v.checkingForType.(env.Func);
		defer func(){v.checkingForType = expected}()

		if  !ok {
			// fmt.Println(v.checkingForType.Type())
			v.err("ERROR_UNEXPECTED_LAMBDA")
		}

		v.subtyping = -1

		for i, param := range ctx.GetParamDecls() {
			v.checkingForType = expected.Args[i]
			param.Accept(v)
		}

		v.subtyping = 1

		v.checkingForType = expected.Return
		ctx.GetReturnExpr().Accept(v)

		return expected
	}

	if v.checking && v.subtyping == -1 {
		
		expected, ok := v.checkingForType.(env.Func);
		defer func(){v.checkingForType = expected}()

		if  !ok {
			// fmt.Println(v.checkingForType.Type())
			v.err("ERROR_UNEXPECTED_LAMBDA")
		}

		v.subtyping = 1

		for i, param := range ctx.GetParamDecls() {
			v.checkingForType = expected.Args[i]
			param.Accept(v)
		}

		v.subtyping = -1

		v.checkingForType = expected.Return
		ctx.GetReturnExpr().Accept(v)

		return expected
	}

	if v.checking {

		expected, ok := v.checkingForType.(env.Func);
		defer func(){v.checkingForType = expected}()

		if  !ok {
			// fmt.Println(v.checkingForType.Type())
			v.err("ERROR_UNEXPECTED_LAMBDA")
		}

		for i, param := range ctx.GetParamDecls() {
			v.checkingForType = expected.Args[i]
			param.Accept(v)
		}

		v.checkingForType = expected.Return
		ctx.GetReturnExpr().Accept(v)

		return expected
	}

	args := make([]env.Type, 0, len(ctx.GetParamDecls()))
	for _, decl := range ctx.GetParamDecls() {
		args = append(args, decl.Accept(v).(env.Type))
	}
	
	return env.Func{
		Args: args,
		Return: ctx.GetReturnExpr().Accept(v).(env.Type),
		IsAnonymous: true,
	}
}

func (v *Visitor) VisitAdd(ctx *parser.AddContext) interface{} {
	panic("unimplemented")
}

func (v *Visitor) VisitAnExtension(ctx *parser.AnExtensionContext) interface{} {
	panic("unimplemented")
}

func (v *Visitor) VisitApplication(ctx *parser.ApplicationContext) interface{} {

	ch, ty := v.checking, v.checkingForType
	defer func ()  {v.checking, v.checkingForType = ch, ty}()

	v.checking = false
	fn, ok := ctx.GetFun().Accept(v).(env.Func)
	if !ok {
		v.err("ERROR_NOT_A_FUNCTION")
	}

	for i, param := range fn.Args {
		v.checking = true
		v.checkingForType = param

		ctx.GetArgs()[i].Accept(v)
	}

	v.checking, v.checkingForType = ch, ty

	if ch && v.subtyping == 1 && v.isSubtype(fn.Return, v.checkingForType) {
		return v.checkingForType
	}


	if ch && v.subtyping == -1 && v.isSubtype(v.checkingForType, fn.Return) {
		return fn.Return
	}

	if ch {
		// TypeCheck(v.checkingForType, fn.Return)
		if v.checkingForType.Type() != fn.Return.Type() {
			v.err("ERROR_UNEXPECTED_TYPE_FOR_EXPRESSION")
		}
		return v.checkingForType
	}

	return fn.Return
}

func (v *Visitor) VisitAssign(ctx *parser.AssignContext) interface{} {
	defer func (ch bool, ty env.Type)  {v.checking, v.checkingForType = ch, ty}(v.checking, v.checkingForType)

	v.checking = false
	lhs, ok := ctx.GetLhs().Accept(v).(env.Reference)
	if !ok {
		v.err("ERROR_NOT_A_REFERENCE")
	}

	v.checking = true
	v.checkingForType = lhs.UnderlyingType

	ctx.GetRhs().Accept(v)

	return env.Unit{}
}

func (v *Visitor) VisitBinding(ctx *parser.BindingContext) interface{} {
	panic("unimplemented")
}

func (v *Visitor) VisitChildren(node antlr.RuleNode) interface{} {
	panic("unimplemented")
}

func (v *Visitor) VisitConsList(ctx *parser.ConsListContext) interface{} {
	if v.checking {
		lst, ok := v.checkingForType.(env.List)
		if !ok {
			v.err("ERROR_UNEXPECTED_LIST")
		}
		v.checkingForType = lst.T
		ctx.GetHead().Accept(v)
		v.checkingForType = lst
		ctx.GetTail().Accept(v)

		return lst
	}

	t := ctx.GetHead().Accept(v).(env.Type)

	v.checking = true 
	chType := v.checkingForType
	v.checkingForType = env.List{
		T: t,		
	}
	ctx.GetTail().Accept(v)

	v.checkingForType = chType
	v.checking = false
	return env.List{
		T: t,
		IsLiteral: true,
	}
}

func (v *Visitor) VisitConstFalse(ctx *parser.ConstFalseContext) interface{} {
	if !v.checking {
		return env.Bool{}
	} else {
		if v.subtyping == 1 && v.isSubtype(env.Bool{}, v.checkingForType) {
			return v.checkingForType
		}
	
	
		if v.subtyping == -1 && v.isSubtype(v.checkingForType, env.Bool{}) {
			return env.Bool{}
		}
		// TypeCheck(v.checkingForType, env.Bool{})
		if v.checkingForType.Type() != (env.Bool{}).Type() {
			v.err("ERROR_UNEXPECTED_TYPE_FOR_EXPRESSION")
		}
	}

	return env.Bool{}
}

func (v *Visitor) VisitConstInt(ctx *parser.ConstIntContext) interface{} {
	if !v.checking {
		return env.Nat{}
	} else {
		if v.subtyping == 1 && v.isSubtype(env.Nat{}, v.checkingForType) {
			return v.checkingForType
		}
	
	
		if v.subtyping == -1 && v.isSubtype(v.checkingForType, env.Nat{}) {
			return env.Nat{}
		}
		// TypeCheck(v.checkingForType, env.Nat{})
		if v.checkingForType.Type() != (env.Nat{}).Type() {
			v.err("ERROR_UNEXPECTED_TYPE_FOR_EXPRESSION")
		}
	}

	return env.Nat{}
}

func (v *Visitor) VisitConstMemory(ctx *parser.ConstMemoryContext) interface{} {
	if v.checking {
		t, ok := v.checkingForType.(env.Reference)
		if !ok {
			v.err("ERROR_UNEXPECTED_MEMORY_ADDRESS")
		}

		return t
	}

	if v.ambiguousAsBottom { return env.Bottom{} }
	v.err("ERROR_AMBIGUOUS_REFERENCE_TYPE")
	return nil
}

func (v *Visitor) VisitConstTrue(ctx *parser.ConstTrueContext) interface{} {
	if !v.checking {
		return env.Bool{}
	} else {
		if v.subtyping == 1 && v.isSubtype(env.Bool{}, v.checkingForType) {
			return v.checkingForType
		}
	
	
		if v.subtyping == -1 && v.isSubtype(v.checkingForType, env.Bool{}) {
			return env.Bool{}
		}
		// TypeCheck(v.checkingForType, env.Bool{})
		if v.checkingForType.Type() != (env.Bool{}).Type() {
			v.err("ERROR_UNEXPECTED_TYPE_FOR_EXPRESSION")
		}
	}

	return env.Bool{}
}

func (v *Visitor) VisitConstUnit(ctx *parser.ConstUnitContext) interface{} {
	if !v.checking {
		return env.Unit{}
	} else {
		if v.subtyping == 1 && v.isSubtype(env.Unit{}, v.checkingForType) {
			return v.checkingForType
		}
	
	
		if v.subtyping == -1 && v.isSubtype(v.checkingForType, env.Unit{}) {
			return env.Unit{}
		}
		// TypeCheck(v.checkingForType, env.Unit{})
		if v.checkingForType.Type() != (env.Unit{}).Type() {
			v.err("ERROR_UNEXPECTED_TYPE_FOR_EXPRESSION")
		}
	}

	return env.Unit{}
}

func (v *Visitor) VisitDeclExceptionType(ctx *parser.DeclExceptionTypeContext) interface{} {
	t := ctx.GetExceptionType().Accept(v).(env.Type)
	v.env.Put(exceptionType, t)
	return t
}

func (v *Visitor) VisitDeclExceptionVariant(ctx *parser.DeclExceptionVariantContext) interface{} {
	panic("unimplemented")
}

func (v *Visitor) VisitDeclFun(ctx *parser.DeclFunContext) interface{} {

	v.env.Push()

	// -------------- Check local functions

	for _, decl := range ctx.GetLocalDecls() {

		val, ok := decl.(*parser.DeclFunContext)
		if !ok {continue}

		args := make([]env.Type, 0, len(val.GetParamDecls()))

		for _, val := range val.GetParamDecls() {
			v.checking = false
			args = append(args, val.Accept(v).(env.Type))
		}

		returnType := val.GetReturnType().Accept(v).(env.Type)
		
		v.env.Put(val.GetName().GetText(), env.Func{
			Args: args,
			Return: returnType,
		})
	}

	for _, decl := range ctx.GetLocalDecls() {
		decl.Accept(v)
	}

	// ------------------- Check body of given function

	for _, param := range ctx.GetParamDecls() {
		v.checking = false
		param.Accept(v)
	}

	ch, ty := v.checking, v.checkingForType
	defer func ()  {		
		v.checking = ch
		v.checkingForType = ty
	} ()

	v.checking = true
	v.checkingForType = ctx.GetReturnType().Accept(v).(env.Type)
	ctx.GetReturnExpr().Accept(v)

	v.env.Pop()

	return nil
}

func (v *Visitor) VisitDeclFunGeneric(ctx *parser.DeclFunGenericContext) interface{} {
	panic("unimplemented")
}

func (v *Visitor) VisitDeclTypeAlias(ctx *parser.DeclTypeAliasContext) interface{} {
	panic("unimplemented")
}

func (v *Visitor) VisitDeref(ctx *parser.DerefContext) interface{} {

	if v.checking && v.subtyping == 1 {
		v.checking = false

		t, ok := ctx.GetExpr_().Accept(v).(env.Reference)
		if !ok || !v.isSubtype(t.UnderlyingType, v.checkingForType) {
			v.err("ERROR_UNEXPECTED_TYPE_FOR_EXPRESSION")
		}
		
		v.checking = true

		return v.checkingForType
	}

	if v.checking && v.subtyping == -1 {
		v.checking = false

		t, ok := ctx.GetExpr_().Accept(v).(env.Reference)
		if !ok || !v.isSubtype(v.checkingForType, t.UnderlyingType) {
			v.err("ERROR_UNEXPECTED_TYPE_FOR_EXPRESSION")
		}
		
		v.checking = true

		return v.checkingForType
	}

	if v.checking {
		v.checkingForType = env.Reference{
			UnderlyingType: v.checkingForType,
		}

		ctx.GetExpr_().Accept(v)

		v.checkingForType = v.checkingForType.(env.Reference).UnderlyingType
		return v.checkingForType
	}

	t, ok := ctx.GetExpr_().Accept(v).(env.Reference)
	if !ok {
		v.err("ERROR_NOT_A_REFERENCE")
	}

	return t.UnderlyingType
}

func (v *Visitor) VisitDivide(ctx *parser.DivideContext) interface{} {
	panic("unimplemented")
}

func (v *Visitor) VisitDotRecord(ctx *parser.DotRecordContext) interface{} {
	
	ch, ty := v.checking, v.checkingForType
	defer func ()  {v.checking, v.checkingForType = ch, ty} ()

	v.checking = false
	record, ok := ctx.GetExpr_().Accept(v).(env.Record)
	if !ok {
		v.err("ERROR_NOT_A_RECORD")
	}
	
	mty, ok := record.Elements[ctx.GetLabel().GetText()]
	if !ok {
		v.err("ERROR_UNEXPECTED_FIELD_ACCESS")
	}

	if ch && v.subtyping == 1 && v.isSubtype(mty, ty) {
		return ty
	}


	if ch && v.subtyping == -1 && v.isSubtype(ty, mty) {
		return mty
	}

	if ch {
		// TypeCheck(ty, mty)
		if ty.Type() != mty.Type() {
			v.err("ERROR_UNEXPECTED_TYPE_FOR_EXPRESSION")
		}
		return ty
	}

	return mty
}

func (v *Visitor) VisitDotTuple(ctx *parser.DotTupleContext) interface{} {
	
	ch, ty := v.checking, v.checkingForType
	defer func ()  {v.checking, v.checkingForType = ch, ty} ()

	v.checking = false
	tuple, ok := ctx.GetExpr_().Accept(v).(env.Tuple)
	if !ok {
		v.err("ERROR_NOT_A_TUPLE")
	}
	
	idx, _ := strconv.Atoi(ctx.GetIndex().GetText())

	if idx > len(tuple.Elements) || idx < 1{
		v.err("ERROR_TUPLE_INDEX_OUT_OF_BOUNDS")
	}

	if ch && v.subtyping == 1 && v.isSubtype(tuple, ty) {
		return ty
	}


	if ch && v.subtyping == -1 && v.isSubtype(ty, tuple) {
		return tuple
	}

	if ch {
		// TypeCheck(ty, tuple.Elements[idx - 1])
		if ty.Type() != tuple.Elements[idx - 1].Type() {
			v.err("ERROR_UNEXPECTED_TYPE_FOR_EXPRESSION")
		}
		return ty
	}

	return tuple.Elements[idx - 1]
}

func (v *Visitor) VisitEqual(ctx *parser.EqualContext) interface{} {
	panic("unimplemented")
}

func (v *Visitor) VisitErrorNode(node antlr.ErrorNode) interface{} {
	panic("unimplemented")
}

func (v *Visitor) VisitFix(ctx *parser.FixContext) interface{} {
	if v.checking {
		ty := v.checkingForType

		v.checkingForType = env.Func{
			Args: []env.Type{ty},
			Return: ty,
		}

		ctx.GetExpr_().Accept(v)

		v.checkingForType = ty
		return ty
	}

	t, ok := ctx.GetExpr_().Accept(v).(env.Func)
	if !ok {
		v.err("ERROR_NOT_A_FUNCTION")
	}

	if len(t.Args) != 1 {
		v.err("ERROR_INCORRECT_NUMBER_OF_ARGUMENTS")
	}
	
	// TypeCheck(t.Args[0], t.Return)
	if t.Args[0].Type() != t.Return.Type() {
		v.err("ERROR_UNEXPECTED_TYPE_FOR_EXPRESSION")
	}

	return t.Return
}

func (v *Visitor) VisitFold(ctx *parser.FoldContext) interface{} {
	panic("unimplemented")
}

func (v *Visitor) VisitGreaterThan(ctx *parser.GreaterThanContext) interface{} {
	panic("unimplemented")
}

func (v *Visitor) VisitGreaterThanOrEqual(ctx *parser.GreaterThanOrEqualContext) interface{} {
	panic("unimplemented")
}

func (v *Visitor) VisitHead(ctx *parser.HeadContext) interface{} {
	ch, ty := v.checking, v.checkingForType
	defer func() {v.checking, v.checkingForType = ch, ty} ()

	if v.checking {
		v.checkingForType = env.List{
			T: v.checkingForType,
		}
	}
	t, ok := ctx.GetList().Accept(v).(env.List)
	if !ok {
		v.err("ERROR_NOT_A_LIST")
	}

	if ch && v.subtyping == 1 && v.isSubtype(t.T, ty) {
		return ty
	}


	if ch && v.subtyping == -1 && v.isSubtype(ty, t.T) {
		return t.T
	}

	if ch {
		// TypeCheck(t.T, ty)
		if t.T.Type() != ty.Type() {
			v.err("ERROR_UNEXPECTED_TYPE_FOR_EXPRESSION")
		}
		return ty
	}

	return t.T
}

func (v *Visitor) VisitIf(ctx *parser.IfContext) interface{} {
	
	ch, ty := v.checking, v.checkingForType

	v.checking, v.checkingForType = true, env.Bool{}

	_ = ctx.GetCondition().Accept(v)

	v.checking, v.checkingForType = ch, ty

	if v.checking {

		ctx.GetThenExpr().Accept(v)

		ctx.GetElseExpr().Accept(v)

		return v.checkingForType
	}

	t := ctx.GetThenExpr().Accept(v).(env.Type)
	
	v.checking = true
	v.checkingForType = t
	
	ctx.GetElseExpr().Accept(v)

	v.checkingForType = ty
	v.checking = false
	
	return t
}

func (v *Visitor) VisitInl(ctx *parser.InlContext) interface{} {
	if !v.checking {
		if v.ambiguousAsBottom { return env.Sum{
			Right: env.Bottom{},
			Left: ctx.GetExpr_().Accept(v).(env.Type),
		} }
		v.err("ERROR_AMBIGUOUS_SUM_TYPE")
	}

	t, ok := v.checkingForType.(env.Sum)
	if !ok {
		v.err("ERROR_UNEXPECTED_INJECTION")
	}

	defer func(ch bool, ty env.Type) {v.checking, v.checkingForType = ch, ty} (v.checking, v.checkingForType)
	v.checkingForType = t.Left
	ctx.GetExpr_().Accept(v)

	return t
}

func (v *Visitor) VisitInlineAnnotation(ctx *parser.InlineAnnotationContext) interface{} {
	panic("unimplemented")
}

func (v *Visitor) VisitInr(ctx *parser.InrContext) interface{} {
	if !v.checking {
		if v.ambiguousAsBottom { return env.Sum{
			Left: env.Bottom{},
			Right: ctx.GetExpr_().Accept(v).(env.Type),
		} }
		v.err("ERROR_AMBIGUOUS_SUM_TYPE")
	}

	t, ok := v.checkingForType.(env.Sum)
	if !ok {
		v.err("ERROR_UNEXPECTED_INJECTION")
	}
	
	defer func(ch bool, ty env.Type) {v.checking, v.checkingForType = ch, ty} (v.checking, v.checkingForType)
	v.checkingForType = t.Right
	ctx.GetExpr_().Accept(v)

	return t
}

func (v *Visitor) VisitIsEmpty(ctx *parser.IsEmptyContext) interface{} {
	ch, ty := v.checking, v.checkingForType
	defer func() {v.checking, v.checkingForType = ch, ty} ()

	v.checking = false
	_, ok := ctx.GetList().Accept(v).(env.List)
	if !ok {
		v.err("ERROR_NOT_A_LIST")
	}

	if ch && v.subtyping == 1 && v.isSubtype(env.Bool{}, ty) {
		return ty
	}


	if ch && v.subtyping == -1 && v.isSubtype(ty, env.Bool{}) {
		return env.Bool{}
	}

	if ch {
		// TypeCheck(env.Bool{}, ty)
		if (env.Bool{}).Type() != ty.Type() {
			v.err("ERROR_UNEXPECTED_TYPE_FOR_EXPRESSION")
		}
		return ty
	}

	return env.Bool{}
}

func (v *Visitor) VisitIsZero(ctx *parser.IsZeroContext) interface{} {
	ch, ty := v.checking, v.checkingForType

	v.checking = true
	v.checkingForType = env.Nat{}

	ctx.GetN().Accept(v)

	v.checking = ch
	v.checkingForType = ty

	if ch && v.subtyping == 1 && v.isSubtype(env.Bool{}, v.checkingForType) {
		return v.checkingForType
	}


	if ch && v.subtyping == -1 && v.isSubtype(v.checkingForType, env.Bool{}) {
		return env.Bool{}
	}
	
	if v.checking {
		// TypeCheck(v.checkingForType, env.Bool{})
		if v.checkingForType.Type() != (env.Bool{}).Type() {
			v.err("ERROR_UNEXPECTED_TYPE_FOR_EXPRESSION")
		}
	}
	
	return env.Bool{}
}

func (v *Visitor) VisitLabelledPattern(ctx *parser.LabelledPatternContext) interface{} {
	panic("unimplemented")
}

func (v *Visitor) VisitLanguageCore(ctx *parser.LanguageCoreContext) interface{} {
	panic("unimplemented")
}

func (v *Visitor) VisitLessThan(ctx *parser.LessThanContext) interface{} {
	panic("unimplemented")
}

func (v *Visitor) VisitLessThanOrEqual(ctx *parser.LessThanOrEqualContext) interface{} {
	panic("unimplemented")
}

func (v *Visitor) VisitLet(ctx *parser.LetContext) interface{} {
	
	v.env.Push()
	defer v.env.Pop()

	ch, ty := v.checking, v.checkingForType
	defer func() {v.checking, v.checkingForType = ch, ty} ()

	v.checking = false
	for _, pat := range ctx.GetPatternBindings() {
		pat.Accept(v)
	}

	if ch {
		v.checking = true
		v.checkingForType = ty
		ctx.GetBody().Accept(v)
	}
	
	return ctx.GetBody().Accept(v)

}

func (v *Visitor) VisitLetRec(ctx *parser.LetRecContext) interface{} {
	panic("unimplemented")
}

func (v *Visitor) VisitList(ctx *parser.ListContext) interface{} {
	if v.checking {

		lst, ok := v.checkingForType.(env.List)
		if !ok {

			if len(ctx.GetExprs()) <= 0 && ! v.ambiguousAsBottom {
				// if v.ambiguousAsBottom { return env.Bottom{} }
				v.err("ERROR_AMBIGUOUS_LIST_TYPE")
			}

			v.err("ERROR_UNEXPECTED_LIST")
		}

		for _, expr := range ctx.GetExprs() {
			v.checkingForType = lst.T
			expr.Accept(v)
		}

		v.checkingForType = lst

		return v.checkingForType
	}

	if len(ctx.GetExprs()) <= 0 {
		if v.ambiguousAsBottom { return env.List{
			T: env.Bottom{},
		}}
		v.err("ERROR_AMBIGUOUS_LIST_TYPE")
	}

	t := ctx.GetExprs()[0].Accept(v).(env.Type)

	defer func(t env.Type) {v.checkingForType = t} (v.checkingForType)
	
	v.checking = true
	v.checkingForType = t

	for _, expr := range ctx.GetExprs() {
		// TypeCheck(t, expr.Accept(v))
		expr.Accept(v)
	}

	v.checking = false

	return env.List{
		T: t,
		IsLiteral: true,
	}
	
}

func (v *Visitor) VisitLogicAnd(ctx *parser.LogicAndContext) interface{} {
	panic("unimplemented")
}

func (v *Visitor) VisitLogicNot(ctx *parser.LogicNotContext) interface{} {
	panic("unimplemented")
}

func (v *Visitor) VisitLogicOr(ctx *parser.LogicOrContext) interface{} {
	panic("unimplemented")
}

func (v *Visitor) VisitMatch(ctx *parser.MatchContext) interface{} {

	ch, ty := v.checking, v.checkingForType
	defer func(mty env.Type) {v.checking, v.checkingForType, v.matchType = ch, ty, mty} (v.matchType)
	
	if len(ctx.GetCases()) <= 0 {
		v.err("ERROR_ILLEGAL_EMPTY_MATCHING")
	}
	v.checking = false
	matchType := ctx.GetExpr_().Accept(v).(env.Type)


	v.checking = true
	v.matchType = matchType

	l, r := inlFlag, inrFlag
	defer func(){inlFlag, inrFlag = l, r} ()
	inlFlag = false
	inrFlag = false
	
	for _, expr := range ctx.GetCases() {
		expr.Accept(v)
	}

	if _, ok := matchType.(env.Sum); ok && ! (inlFlag && inrFlag)  {
		v.err("ERROR_NONEXHAUSTIVE_MATCH_PATTERNS")
	}

	return ty
}

func (v *Visitor) VisitMatchCase(ctx *parser.MatchCaseContext) interface{} {
	if !v.checking{
		v.err("strange match case")
	}
	
	v.env.Push()
	defer v.env.Pop()


	ctx.GetPattern_().Accept(v)
	// v.checking = false
	ctx.GetExpr_().Accept(v)
	// TypeCheck(v.checkingForType, t)
	// v.checking = true

	return v.checkingForType
}

func (v *Visitor) VisitMultiply(ctx *parser.MultiplyContext) interface{} {
	panic("unimplemented")
}

func (v *Visitor) VisitNatRec(ctx *parser.NatRecContext) interface{} {
	ch, ty := v.checking, v.checkingForType

	v.checking = true
	v.checkingForType = env.Nat{}
	ctx.GetN().Accept(v)

	v.checking = ch
	v.checkingForType = ty

	z := ctx.GetInitial().Accept(v).(env.Type)

	v.checking = true
	v.checkingForType = env.Func{
		Args: []env.Type{env.Nat{}},
		Return: env.Func{
			Args: []env.Type{z},
			Return: z,
		},
	}

	ctx.GetStep().Accept(v)

	v.checking = ch
	v.checkingForType = ty

	if ch && v.subtyping == 1 && v.isSubtype(z, v.checkingForType) {
		return v.checkingForType
	}


	if ch && v.subtyping == -1 && v.isSubtype(v.checkingForType, z) {
		return z
	}
	
	if v.checking {
		// TypeCheck(v.checkingForType, z)
		if v.checkingForType.Type() != z.Type() {
			v.err("ERROR_UNEXPECTED_TYPE_FOR_EXPRESSION")
		}
	}

	return z
}

func (v *Visitor) VisitNotEqual(ctx *parser.NotEqualContext) interface{} {
	panic("unimplemented")
}

func (v *Visitor) VisitPanic(ctx *parser.PanicContext) interface{} {
	if v.checking {
		return v.checkingForType
	}

	if v.ambiguousAsBottom { return env.Bottom{} }
	v.err("ERROR_AMBIGUOUS_PANIC_TYPE")
	return nil
}

func (v *Visitor) VisitParamDecl(ctx *parser.ParamDeclContext) interface{} {
	
	ch := v.checking
	v.checking = false
	
	nm, ty := ctx.GetName().GetText(), ctx.GetParamType().Accept(v).(env.Type)

	v.checking = ch

	if v.checking && v.subtyping == 1 && v.isSubtype(ty, v.checkingForType) {
		return v.checkingForType
	}

	if v.checking && v.subtyping == -1 && v.isSubtype(v.checkingForType, ty) {
		return ty
	}
	
	if v.checking && v.subtyping != 0 {
		v.err("ERROR_UNEXPECTED_TYPE_FOR_EXPRESSION")
	}

	if v.checking && ty.Type() != v.checkingForType.Type() {
		v.err("ERROR_UNEXPECTED_TYPE_FOR_PARAMETER")
	}

	v.env.Put(nm, ty)

	return ty
}

func (v *Visitor) VisitParenthesisedExpr(ctx *parser.ParenthesisedExprContext) interface{} {
	return ctx.GetExpr_().Accept(v)
}

func (v *Visitor) VisitParenthesisedPattern(ctx *parser.ParenthesisedPatternContext) interface{} {
	panic("unimplemented")
}

func (v *Visitor) VisitPatternAsc(ctx *parser.PatternAscContext) interface{} {
	panic("unimplemented")
}

func (v *Visitor) VisitPatternBinding(ctx *parser.PatternBindingContext) interface{} {
	
	if v.checking {
		ctx.GetRhs().Accept(v)
		ctx.GetPat().Accept(v)
		return v.checkingForType
	}
	
	ty := v.checkingForType
	defer func() {v.checkingForType = ty} ()
	
	v.checkingForType = ctx.GetRhs().Accept(v).(env.Type)
	ctx.GetPat().Accept(v)
	val := v.checkingForType

	return val
}

func (v *Visitor) VisitPatternCastAs(ctx *parser.PatternCastAsContext) interface{} {
	panic("unimplemented")
}

func (v *Visitor) VisitPatternCons(ctx *parser.PatternConsContext) interface{} {
	panic("unimplemented")
}

func (v *Visitor) VisitPatternFalse(ctx *parser.PatternFalseContext) interface{} {
	panic("unimplemented")
}

func (v *Visitor) VisitPatternInl(ctx *parser.PatternInlContext) interface{} {
	t, ok := v.matchType.(env.Sum)
	if !ok {
		v.err("ERROR_UNEXPECTED_PATTERN_FOR_TYPE")
	}

	defer func(t env.Type, c bool) {v.matchType, v.checking = t, c} (v.matchType, v.checking)
	v.checking = false

	v.matchType = t.Left
	ctx.GetPattern_().Accept(v)
	inlFlag = true
	return nil
}

func (v *Visitor) VisitPatternInr(ctx *parser.PatternInrContext) interface{} {
	t, ok := v.matchType.(env.Sum)
	if !ok {
		fmt.Printf("%T\n", v.matchType)
		v.err("ERROR_UNEXPECTED_PATTERN_FOR_TYPE")
	}

	defer func(t env.Type, c bool) {v.matchType, v.checking = t, c} (v.matchType, v.checking)
	v.checking = false

	v.matchType = t.Right
	ctx.GetPattern_().Accept(v)
	inrFlag = true
	return nil
}

func (v *Visitor) VisitPatternInt(ctx *parser.PatternIntContext) interface{} {
	panic("unimplemented")
}

func (v *Visitor) VisitPatternList(ctx *parser.PatternListContext) interface{} {
	panic("unimplemented")
}

func (v *Visitor) VisitPatternRecord(ctx *parser.PatternRecordContext) interface{} {
	panic("unimplemented")
}

func (v *Visitor) VisitPatternSucc(ctx *parser.PatternSuccContext) interface{} {
	panic("unimplemented")
}

func (v *Visitor) VisitPatternTrue(ctx *parser.PatternTrueContext) interface{} {
	panic("unimplemented")
}

func (v *Visitor) VisitPatternTuple(ctx *parser.PatternTupleContext) interface{} {
	panic("unimplemented")
}

func (v *Visitor) VisitPatternUnit(ctx *parser.PatternUnitContext) interface{} {
	panic("unimplemented")
}

func (v *Visitor) VisitPatternVar(ctx *parser.PatternVarContext) interface{} {
	if v.checking || v.matchType == nil || v.matchType.Type() == "" {
		v.env.Put(ctx.GetName().GetText(), v.checkingForType)
		return v.checkingForType
	}

	v.env.Put(ctx.GetName().GetText(), v.matchType)
	return v.matchType
}

func (v *Visitor) VisitPatternVariant(ctx *parser.PatternVariantContext) interface{} {
	t, ok := v.matchType.(env.Variant)
	if !ok {
		v.err("ERROR_UNEXPECTED_PATTERN_FOR_TYPE")
	}

	defer func(t env.Type, c bool) {v.matchType, v.checking = t, c} (v.matchType, v.checking)
	v.checking = false

	v.matchType, ok = t.Elements[ctx.GetLabel().GetText()]
	if !ok {
		v.err("ERROR_UNEXPECTED_PATTERN_FOR_TYPE")
	}
	ctx.GetPattern_().Accept(v)
	return nil
}

func (v *Visitor) VisitPred(ctx *parser.PredContext) interface{} {
	panic("unimplemented")
}

func (v *Visitor) VisitProgram(ctx *parser.ProgramContext) interface{} {
	v.env.Push()

	defer v.env.Pop()
	if len(ctx.GetExtensions()) > 0 && strings.Contains(ctx.Get_extension().GetText(), "structural-subtyping") {
		// fmt.Println(ctx.Get_extension().GetText())
		v.subtyping = 1
	}

	v.ambiguousAsBottom = len(ctx.GetExtensions()) > 0 && strings.Contains(ctx.Get_extension().GetText(), "ambiguous-type-as-bottom")

	for _, decl := range ctx.GetDecls() {
		v.env.Push()
		val, ok := decl.(*parser.DeclFunContext)
		if !ok {continue}

		args := make([]env.Type, 0, len(val.GetParamDecls()))

		for _, val := range val.GetParamDecls() {
			args = append(args, val.Accept(v).(env.Type))
		}

		returnType := val.GetReturnType().Accept(v).(env.Type)
		
		v.env.Pop()
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
	

	return nil
}

func (v *Visitor) VisitRecord(ctx *parser.RecordContext) interface{} {
	
	defer func(c bool, t env.Type) {
		v.checking = c
		v.checkingForType = t
	} (v.checking, v.checkingForType)

	if v.checking && v.subtyping == 1 {

		record, ok := v.checkingForType.(env.Record)
		if !ok {
			v.err("ERROR_UNEXPECTED_RECORD")
		}

		labels := make([]string, len(ctx.GetBindings()))
		for i, binding := range ctx.GetBindings() {
			labels[i] = binding.GetName().GetText()
		}

		for key, _ := range record.Elements {
			if !contains(key, labels...) {
				v.err("ERROR_MISSING_RECORD_FIELDS")
			}
		}

		for _, binding := range ctx.GetBindings() {
			t, ok := record.Elements[binding.GetName().GetText()]
			if !ok {
				v.checking = false
				binding.GetRhs().Accept(v)
				v.checking = true
				continue
			}

			v.checkingForType = t
			binding.GetRhs().Accept(v)
		}

		return record
	}

	if v.checking && v.subtyping == -1 {
		record, ok := v.checkingForType.(env.Record)
		if !ok {
			v.err("ERROR_UNEXPECTED_RECORD")
		}

		for _, binding := range ctx.GetBindings() {

			t, ok := record.Elements[binding.GetName().GetText()]
			if !ok {
				v.err("ERROR_MISSING_RECORD_FIELDS")
			}

			v.checkingForType = t
			binding.GetRhs().Accept(v)
		}

		return record
	}

	if v.checking {
		record, ok := v.checkingForType.(env.Record)
		if !ok {
			v.err("ERROR_UNEXPECTED_RECORD")
		}

		if len(record.Elements) > len(ctx.GetBindings()) {
			v.err("ERROR_MISSING_RECORD_FIELDS")
		}

		if len(record.Elements) < len(ctx.GetBindings()) {
			v.err("ERROR_UNEXPECTED_RECORD_FIELDS")
		}

		for _, binding := range ctx.GetBindings() {
			t, ok := record.Elements[binding.GetName().GetText()]
			if !ok {
				v.err("ERROR_UNEXPECTED_RECORD_FIELDS")
			}

			v.checkingForType = t
			binding.GetRhs().Accept(v)
		}

		v.checkingForType = record

		return record
	}

	record := env.Record{
		Elements: make(map[string]env.Type, len(ctx.GetBindings())),
		IsLiteral: true,
	}
	
	for _, binding := range ctx.GetBindings() {
		record.Elements[binding.GetName().GetText()] = eraseLiterals(binding.GetRhs().Accept(v).(env.Type))
	}

	return record
}

func (v *Visitor) VisitRecordFieldType(ctx *parser.RecordFieldTypeContext) interface{} {
	panic("unimplemented")
}

func (v *Visitor) VisitRef(ctx *parser.RefContext) interface{} {

	defer func(i int) {v.subtyping = i} (v.subtyping)

	v.subtyping = 2

	if v.checking {
		t, ok := v.checkingForType.(env.Reference)
		if !ok {
			v.err("ERROR_UNEXPECTED_REFERENCE")
		}

		v.checkingForType = t.UnderlyingType
		ctx.GetExpr_().Accept(v)

		v.checkingForType = t
		return t
	}

	return env.Reference{
		UnderlyingType: ctx.GetExpr_().Accept(v).(env.Type),
	}
}

func (v *Visitor) VisitSequence(ctx *parser.SequenceContext) interface{} {
	ch, ty := v.checking, v.checkingForType
	defer func(t env.Type, c bool) {v.checkingForType, v.checking = t, c} (v.checkingForType, v.checking)

	v.checking = true
	v.checkingForType = env.Unit{}

	ctx.GetExpr1().Accept(v)

	if ch {
		v.checkingForType = ty
		ctx.GetExpr2().Accept(v)

		return ty
	}

	v.checking = false
	return ctx.GetExpr2().Accept(v)
}

func (v *Visitor) VisitStart_Expr(ctx *parser.Start_ExprContext) interface{} {
	panic("unimplemented")
}

func (v *Visitor) VisitStart_Program(ctx *parser.Start_ProgramContext) interface{} {
	panic("unimplemented")
}

func (v *Visitor) VisitStart_Type(ctx *parser.Start_TypeContext) interface{} {
	panic("unimplemented")
}

func (v *Visitor) VisitSubtract(ctx *parser.SubtractContext) interface{} {
	panic("unimplemented")
}

func (v *Visitor) VisitSucc(ctx *parser.SuccContext) interface{} {
	
	ch, ty := v.checking, v.checkingForType

	v.checking = true
	v.checkingForType = env.Nat{}

	ctx.GetN().Accept(v)

	v.checking = ch
	v.checkingForType = ty

	if ch && v.subtyping == 1 && v.isSubtype(env.Nat{}, v.checkingForType) {
		return v.checkingForType
	}


	if ch && v.subtyping == -1 && v.isSubtype(v.checkingForType, env.Nat{}) {
		return env.Nat{}
	}
	
	if v.checking {
		// TypeCheck(v.checkingForType, env.Nat{})
		if v.checkingForType.Type() != (env.Nat{}).Type() {
			v.err("ERROR_UNEXPECTED_TYPE_FOR_EXPRESSION")
		}
	}

	return env.Nat{}
}

func (v *Visitor) VisitTail(ctx *parser.TailContext) interface{} {
	ch, ty := v.checking, v.checkingForType
	defer func() {v.checking, v.checkingForType = ch, ty} ()

	t, ok := ctx.GetList().Accept(v).(env.List)
	if !ok {
		v.err("ERROR_NOT_A_LIST")
	}

	return t
}

func (v *Visitor) VisitTerminal(node antlr.TerminalNode) interface{} {
	panic("unimplemented")
}

func (v *Visitor) VisitTerminatingSemicolon(ctx *parser.TerminatingSemicolonContext) interface{} {
	return ctx.GetExpr_().Accept(v)
}

func (v *Visitor) VisitThrow(ctx *parser.ThrowContext) interface{} {

	if v.env.Check(exceptionType).Type() == "" {
		v.err("ERROR_EXCEPTION_TYPE_NOT_DECLARED")
	}

	if !v.checking {
		if v.ambiguousAsBottom { return env.Bottom{} }
		v.err("ERROR_AMBIGUOUS_THROW_TYPE")
	}

	t := v.checkingForType

	defer func(t env.Type) {v.checkingForType = t} (v.checkingForType)
	v.checkingForType = v.env.Check(exceptionType)
    ctx.GetExpr_().Accept(v)

	return t
}

func (v *Visitor) VisitTryCastAs(ctx *parser.TryCastAsContext) interface{} {
	panic("unimplemented")
}

func (v *Visitor) VisitTryCatch(ctx *parser.TryCatchContext) interface{} {
	
	if v.env.Check(exceptionType).Type() == "" {
		v.err("ERROR_EXCEPTION_TYPE_NOT_DECLARED")
	}
	v.env.Push()
	defer v.env.Pop()
	if v.checking {
		ctx.GetTryExpr().Accept(v)

		t := v.checkingForType
		v.checkingForType = v.env.Check(exceptionType)

		ctx.GetPat().Accept(v)

		v.checkingForType = t

		ctx.GetFallbackExpr().Accept(v)
	}

	ty := ctx.GetTryExpr().Accept(v).(env.Type)


	t := v.checkingForType
	v.checkingForType = v.env.Check(exceptionType)
	v.checking = true

	ctx.GetPat().Accept(v)

	v.checkingForType = ty

	ctx.GetFallbackExpr().Accept(v)

	v.checking = false
	v.checkingForType = t

	return ty
}

func (v *Visitor) VisitTryWith(ctx *parser.TryWithContext) interface{} {

	if v.env.Check(exceptionType).Type() == "" {
		v.err("ERROR_EXCEPTION_TYPE_NOT_DECLARED")
	}
	v.env.Push()
	defer v.env.Pop()
	if v.checking {
		ctx.GetTryExpr().Accept(v)

		ctx.GetFallbackExpr().Accept(v)
	}

	t := ctx.GetTryExpr().Accept(v).(env.Type)
	ch, ty := v.checking, v.checkingForType
	defer func() {v.checking, v.checkingForType = ch, ty} ()

	v.checkingForType = t
	v.checking = true

	ctx.GetFallbackExpr().Accept(v)
	return t
}

func (v *Visitor) VisitTuple(ctx *parser.TupleContext) interface{} {
	if v.checking {
		ty := v.checkingForType
		defer func() {v.checkingForType = ty} ()

		tuple, ok := ty.(env.Tuple)
		if !ok {
			v.err("ERROR_UNEXPECTED_TUPLE")
		}

		if len(tuple.Elements) != len(ctx.GetExprs()) {
			v.err("UNEXPECTED_TUPLE_LENGTH")
		}

		for i, expr := range ctx.GetExprs() {
			v.checkingForType = tuple.Elements[i]
			expr.Accept(v)
		}

		return tuple
	}

	elements := make([]env.Type, 0, len(ctx.GetExprs()))

	for _, expr := range ctx.GetExprs() {
		elements = append(elements, eraseLiterals(expr.Accept(v).(env.Type)))
	}
	
	return env.Tuple {
		Elements: elements,
		IsLiteral: true,
	}
}

func (v *Visitor) VisitTypeAbstraction(ctx *parser.TypeAbstractionContext) interface{} {
	panic("unimplemented")
}

func (v *Visitor) VisitTypeApplication(ctx *parser.TypeApplicationContext) interface{} {
	panic("unimplemented")
}

func (v *Visitor) VisitTypeAsc(ctx *parser.TypeAscContext) interface{} {
	ch, ty := v.checking, v.checkingForType
	defer func ()  {
		v.checking, v.checkingForType = ch, ty
	} ()

	v.checkingForType = ctx.GetType_().Accept(v).(env.Type)
	v.checking = true
	
	ctx.GetExpr_().Accept(v)
	// fmt.Println(v.checkingForType.Type())

	if ch && v.subtyping == 1 && v.isSubtype(v.checkingForType, ty) {
		return ty
	}


	if ch && v.subtyping == -1 && v.isSubtype(ty, v.checkingForType) {
		return v.checkingForType
	}
	
	if ch {
		// TypeCheck(ty, v.checkingForType)
		if v.checkingForType.Type() != ty.Type() {
			v.err("ERROR_UNEXPECTED_TYPE_FOR_EXPRESSION")
		}
	}

	// fmt.Printf("%T\n", ctx.GetType_())
	val := v.checkingForType
	return val
}

func (v *Visitor) VisitTypeAuto(ctx *parser.TypeAutoContext) interface{} {
	panic("unimplemented")
}

func (v *Visitor) VisitTypeBool(ctx *parser.TypeBoolContext) interface{} {
	return env.Bool{}
}

func (v *Visitor) VisitTypeBottom(ctx *parser.TypeBottomContext) interface{} {
	return env.Bottom{}
}

func (v *Visitor) VisitTypeCast(ctx *parser.TypeCastContext) interface{} {
	panic("unimplemented")
}

func (v *Visitor) VisitTypeForAll(ctx *parser.TypeForAllContext) interface{} {
	panic("unimplemented")
}

func (v *Visitor) VisitTypeFun(ctx *parser.TypeFunContext) interface{} {
	
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

func (v *Visitor) VisitTypeList(ctx *parser.TypeListContext) interface{} {
	return env.List{
		T: ctx.GetType_().Accept(v).(env.Type),
	}
}

func (v *Visitor) VisitTypeNat(ctx *parser.TypeNatContext) interface{} {
	return env.Nat{}
}

func (v *Visitor) VisitTypeParens(ctx *parser.TypeParensContext) interface{} {
	return ctx.GetType_().Accept(v)
}

func (v *Visitor) VisitTypeRec(ctx *parser.TypeRecContext) interface{} {
	panic("unimplemented")
}

func (v *Visitor) VisitTypeRecord(ctx *parser.TypeRecordContext) interface{} {
	
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

func (v *Visitor) VisitTypeRef(ctx *parser.TypeRefContext) interface{} {
	return env.Reference{
		UnderlyingType: ctx.GetType_().Accept(v).(env.Type),
	}
}

func (v *Visitor) VisitTypeSum(ctx *parser.TypeSumContext) interface{} {
	return env.Sum{
		Left: ctx.GetLeft().Accept(v).(env.Type),
		Right: ctx.GetRight().Accept(v).(env.Type),
	}
}

func (v *Visitor) VisitTypeTop(ctx *parser.TypeTopContext) interface{} {
	return env.Top{}
}

func (v *Visitor) VisitTypeTuple(ctx *parser.TypeTupleContext) interface{} {
	
	t := ctx.GetTypes()
	
	types := make([]env.Type, len(t))
	
	for i, val := range t {
		types[i] = val.Accept(v).(env.Type)
	}

	return env.Tuple {
		Elements: types,
	}
}

func (v *Visitor) VisitTypeUnit(ctx *parser.TypeUnitContext) interface{} {
	return env.Unit{}
}

func (v *Visitor) VisitTypeVar(ctx *parser.TypeVarContext) interface{} {
	panic("unimplemented")
}

func (v *Visitor) VisitTypeVariant(ctx *parser.TypeVariantContext) interface{} {
	bindings := ctx.GetFieldTypes()
	
	var t = env.Variant{
		Elements: make(map[string]env.Type, len(bindings)),
	}

	for _, binding := range bindings {
		binding := binding.Accept(v).(env.Binding)

		t.Elements[binding.Name] = binding.T
	}

	return t
}

func (v *Visitor) VisitUnfold(ctx *parser.UnfoldContext) interface{} {
	panic("unimplemented")
}

func (v *Visitor) VisitVar(ctx *parser.VarContext) interface{} {
	tp := v.env.Check(ctx.GetName().GetText())
	if tp.Type() == "" {
		fmt.Println("ERROR_UNDEFINED_VARIABLE.: ", ctx.GetName().GetText())
		os.Exit(1)
	}

	if !v.checking {return tp}

	if v.subtyping == 1 && v.isSubtype(tp, v.checkingForType) {
		return v.checkingForType
	}

	if v.subtyping == -1 && v.isSubtype(v.checkingForType, tp) {
		return tp
	}

	// TypeCheck(v.checkingForType, tp)
	if v.checkingForType.Type() != tp.Type() {
		v.err("ERROR_UNEXPECTED_TYPE_FOR_EXPRESSION")
	}

	return tp
}

func (v *Visitor) VisitVariant(ctx *parser.VariantContext) interface{} {
	if !v.checking && (v.subtyping != 1 && v.subtyping != -1) {
		if v.ambiguousAsBottom { return env.Bottom{} }
		v.err("ERROR_AMBIGUOUS_VARIANT_TYPE")
	}

	if !v.checking {
		return env.Variant{
			Elements: map[string]env.Type{
				ctx.GetLabel().GetText(): ctx.GetRhs().Accept(v).(env.Type),
			},
		}
	}

	variant, ok := v.checkingForType.(env.Variant)
	if !ok {
		v.err("ERROR_UNEXPECTED_VARIANT")
	}

	labelType, ok := variant.Elements[ctx.GetLabel().GetText()]
	if !ok {
		v.err("ERROR_UNEXPECTED_VARIANT_LABEL")
	}

	defer func(t env.Type) {
		v.checkingForType = t
	} (v.checkingForType)

	v.checkingForType = labelType
	ctx.GetRhs().Accept(v)

	return variant
}

func (v *Visitor) VisitVariantFieldType(ctx *parser.VariantFieldTypeContext) interface{} {
	return env.Binding{
		Name: ctx.GetLabel().GetText(),
		T: ctx.GetType_().Accept(v).(env.Type),
	}
}
