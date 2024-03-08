// Code generated from StellaParser.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // StellaParser

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

type BaseStellaParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseStellaParserVisitor) VisitStart_Program(ctx *Start_ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitStart_Expr(ctx *Start_ExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitStart_Type(ctx *Start_TypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitLanguageCore(ctx *LanguageCoreContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitAnExtension(ctx *AnExtensionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitDeclFun(ctx *DeclFunContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitDeclFunGeneric(ctx *DeclFunGenericContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitDeclTypeAlias(ctx *DeclTypeAliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitDeclExceptionType(ctx *DeclExceptionTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitDeclExceptionVariant(ctx *DeclExceptionVariantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitInlineAnnotation(ctx *InlineAnnotationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitParamDecl(ctx *ParamDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitFold(ctx *FoldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitAdd(ctx *AddContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitIsZero(ctx *IsZeroContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitVar(ctx *VarContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitTypeAbstraction(ctx *TypeAbstractionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitDivide(ctx *DivideContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitLessThan(ctx *LessThanContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitDotRecord(ctx *DotRecordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitGreaterThan(ctx *GreaterThanContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitEqual(ctx *EqualContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitThrow(ctx *ThrowContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitMultiply(ctx *MultiplyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitConstMemory(ctx *ConstMemoryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitList(ctx *ListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitTryCatch(ctx *TryCatchContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitTryCastAs(ctx *TryCastAsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitHead(ctx *HeadContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitTerminatingSemicolon(ctx *TerminatingSemicolonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitNotEqual(ctx *NotEqualContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitConstUnit(ctx *ConstUnitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitSequence(ctx *SequenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitConstFalse(ctx *ConstFalseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitAbstraction(ctx *AbstractionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitConstInt(ctx *ConstIntContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitVariant(ctx *VariantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitConstTrue(ctx *ConstTrueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitSubtract(ctx *SubtractContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitTypeCast(ctx *TypeCastContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitIf(ctx *IfContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitApplication(ctx *ApplicationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitDeref(ctx *DerefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitIsEmpty(ctx *IsEmptyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitPanic(ctx *PanicContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitLessThanOrEqual(ctx *LessThanOrEqualContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitSucc(ctx *SuccContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitInl(ctx *InlContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitGreaterThanOrEqual(ctx *GreaterThanOrEqualContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitInr(ctx *InrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitMatch(ctx *MatchContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitLogicNot(ctx *LogicNotContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitParenthesisedExpr(ctx *ParenthesisedExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitTail(ctx *TailContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitRecord(ctx *RecordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitLogicAnd(ctx *LogicAndContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitTypeApplication(ctx *TypeApplicationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitLetRec(ctx *LetRecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitLogicOr(ctx *LogicOrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitTryWith(ctx *TryWithContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitPred(ctx *PredContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitTypeAsc(ctx *TypeAscContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitNatRec(ctx *NatRecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitUnfold(ctx *UnfoldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitRef(ctx *RefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitDotTuple(ctx *DotTupleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitFix(ctx *FixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitLet(ctx *LetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitAssign(ctx *AssignContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitTuple(ctx *TupleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitConsList(ctx *ConsListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitPatternBinding(ctx *PatternBindingContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitBinding(ctx *BindingContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitMatchCase(ctx *MatchCaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitPatternCons(ctx *PatternConsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitPatternTuple(ctx *PatternTupleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitPatternList(ctx *PatternListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitPatternRecord(ctx *PatternRecordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitPatternVariant(ctx *PatternVariantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitPatternAsc(ctx *PatternAscContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitPatternInt(ctx *PatternIntContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitPatternInr(ctx *PatternInrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitPatternTrue(ctx *PatternTrueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitPatternInl(ctx *PatternInlContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitPatternVar(ctx *PatternVarContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitParenthesisedPattern(ctx *ParenthesisedPatternContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitPatternSucc(ctx *PatternSuccContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitPatternFalse(ctx *PatternFalseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitPatternUnit(ctx *PatternUnitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitPatternCastAs(ctx *PatternCastAsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitLabelledPattern(ctx *LabelledPatternContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitTypeTuple(ctx *TypeTupleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitTypeTop(ctx *TypeTopContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitTypeBool(ctx *TypeBoolContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitTypeRef(ctx *TypeRefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitTypeRec(ctx *TypeRecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitTypeSum(ctx *TypeSumContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitTypeVar(ctx *TypeVarContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitTypeVariant(ctx *TypeVariantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitTypeUnit(ctx *TypeUnitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitTypeNat(ctx *TypeNatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitTypeBottom(ctx *TypeBottomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitTypeParens(ctx *TypeParensContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitTypeFun(ctx *TypeFunContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitTypeForAll(ctx *TypeForAllContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitTypeRecord(ctx *TypeRecordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitTypeList(ctx *TypeListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitRecordFieldType(ctx *RecordFieldTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseStellaParserVisitor) VisitVariantFieldType(ctx *VariantFieldTypeContext) interface{} {
	return v.VisitChildren(ctx)
}
