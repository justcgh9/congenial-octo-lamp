// Code generated from StellaParser.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // StellaParser

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"
// A complete Visitor for a parse tree produced by StellaParser.
type StellaParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by StellaParser#start_Program.
	VisitStart_Program(ctx *Start_ProgramContext) interface{}

	// Visit a parse tree produced by StellaParser#start_Expr.
	VisitStart_Expr(ctx *Start_ExprContext) interface{}

	// Visit a parse tree produced by StellaParser#start_Type.
	VisitStart_Type(ctx *Start_TypeContext) interface{}

	// Visit a parse tree produced by StellaParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by StellaParser#LanguageCore.
	VisitLanguageCore(ctx *LanguageCoreContext) interface{}

	// Visit a parse tree produced by StellaParser#AnExtension.
	VisitAnExtension(ctx *AnExtensionContext) interface{}

	// Visit a parse tree produced by StellaParser#DeclFun.
	VisitDeclFun(ctx *DeclFunContext) interface{}

	// Visit a parse tree produced by StellaParser#DeclFunGeneric.
	VisitDeclFunGeneric(ctx *DeclFunGenericContext) interface{}

	// Visit a parse tree produced by StellaParser#DeclTypeAlias.
	VisitDeclTypeAlias(ctx *DeclTypeAliasContext) interface{}

	// Visit a parse tree produced by StellaParser#DeclExceptionType.
	VisitDeclExceptionType(ctx *DeclExceptionTypeContext) interface{}

	// Visit a parse tree produced by StellaParser#DeclExceptionVariant.
	VisitDeclExceptionVariant(ctx *DeclExceptionVariantContext) interface{}

	// Visit a parse tree produced by StellaParser#InlineAnnotation.
	VisitInlineAnnotation(ctx *InlineAnnotationContext) interface{}

	// Visit a parse tree produced by StellaParser#paramDecl.
	VisitParamDecl(ctx *ParamDeclContext) interface{}

	// Visit a parse tree produced by StellaParser#Fold.
	VisitFold(ctx *FoldContext) interface{}

	// Visit a parse tree produced by StellaParser#Add.
	VisitAdd(ctx *AddContext) interface{}

	// Visit a parse tree produced by StellaParser#IsZero.
	VisitIsZero(ctx *IsZeroContext) interface{}

	// Visit a parse tree produced by StellaParser#Var.
	VisitVar(ctx *VarContext) interface{}

	// Visit a parse tree produced by StellaParser#TypeAbstraction.
	VisitTypeAbstraction(ctx *TypeAbstractionContext) interface{}

	// Visit a parse tree produced by StellaParser#Divide.
	VisitDivide(ctx *DivideContext) interface{}

	// Visit a parse tree produced by StellaParser#LessThan.
	VisitLessThan(ctx *LessThanContext) interface{}

	// Visit a parse tree produced by StellaParser#DotRecord.
	VisitDotRecord(ctx *DotRecordContext) interface{}

	// Visit a parse tree produced by StellaParser#GreaterThan.
	VisitGreaterThan(ctx *GreaterThanContext) interface{}

	// Visit a parse tree produced by StellaParser#Equal.
	VisitEqual(ctx *EqualContext) interface{}

	// Visit a parse tree produced by StellaParser#Throw.
	VisitThrow(ctx *ThrowContext) interface{}

	// Visit a parse tree produced by StellaParser#Multiply.
	VisitMultiply(ctx *MultiplyContext) interface{}

	// Visit a parse tree produced by StellaParser#ConstMemory.
	VisitConstMemory(ctx *ConstMemoryContext) interface{}

	// Visit a parse tree produced by StellaParser#List.
	VisitList(ctx *ListContext) interface{}

	// Visit a parse tree produced by StellaParser#TryCatch.
	VisitTryCatch(ctx *TryCatchContext) interface{}

	// Visit a parse tree produced by StellaParser#TryCastAs.
	VisitTryCastAs(ctx *TryCastAsContext) interface{}

	// Visit a parse tree produced by StellaParser#Head.
	VisitHead(ctx *HeadContext) interface{}

	// Visit a parse tree produced by StellaParser#TerminatingSemicolon.
	VisitTerminatingSemicolon(ctx *TerminatingSemicolonContext) interface{}

	// Visit a parse tree produced by StellaParser#NotEqual.
	VisitNotEqual(ctx *NotEqualContext) interface{}

	// Visit a parse tree produced by StellaParser#ConstUnit.
	VisitConstUnit(ctx *ConstUnitContext) interface{}

	// Visit a parse tree produced by StellaParser#Sequence.
	VisitSequence(ctx *SequenceContext) interface{}

	// Visit a parse tree produced by StellaParser#ConstFalse.
	VisitConstFalse(ctx *ConstFalseContext) interface{}

	// Visit a parse tree produced by StellaParser#Abstraction.
	VisitAbstraction(ctx *AbstractionContext) interface{}

	// Visit a parse tree produced by StellaParser#ConstInt.
	VisitConstInt(ctx *ConstIntContext) interface{}

	// Visit a parse tree produced by StellaParser#Variant.
	VisitVariant(ctx *VariantContext) interface{}

	// Visit a parse tree produced by StellaParser#ConstTrue.
	VisitConstTrue(ctx *ConstTrueContext) interface{}

	// Visit a parse tree produced by StellaParser#Subtract.
	VisitSubtract(ctx *SubtractContext) interface{}

	// Visit a parse tree produced by StellaParser#TypeCast.
	VisitTypeCast(ctx *TypeCastContext) interface{}

	// Visit a parse tree produced by StellaParser#If.
	VisitIf(ctx *IfContext) interface{}

	// Visit a parse tree produced by StellaParser#Application.
	VisitApplication(ctx *ApplicationContext) interface{}

	// Visit a parse tree produced by StellaParser#Deref.
	VisitDeref(ctx *DerefContext) interface{}

	// Visit a parse tree produced by StellaParser#IsEmpty.
	VisitIsEmpty(ctx *IsEmptyContext) interface{}

	// Visit a parse tree produced by StellaParser#Panic.
	VisitPanic(ctx *PanicContext) interface{}

	// Visit a parse tree produced by StellaParser#LessThanOrEqual.
	VisitLessThanOrEqual(ctx *LessThanOrEqualContext) interface{}

	// Visit a parse tree produced by StellaParser#Succ.
	VisitSucc(ctx *SuccContext) interface{}

	// Visit a parse tree produced by StellaParser#Inl.
	VisitInl(ctx *InlContext) interface{}

	// Visit a parse tree produced by StellaParser#GreaterThanOrEqual.
	VisitGreaterThanOrEqual(ctx *GreaterThanOrEqualContext) interface{}

	// Visit a parse tree produced by StellaParser#Inr.
	VisitInr(ctx *InrContext) interface{}

	// Visit a parse tree produced by StellaParser#Match.
	VisitMatch(ctx *MatchContext) interface{}

	// Visit a parse tree produced by StellaParser#LogicNot.
	VisitLogicNot(ctx *LogicNotContext) interface{}

	// Visit a parse tree produced by StellaParser#ParenthesisedExpr.
	VisitParenthesisedExpr(ctx *ParenthesisedExprContext) interface{}

	// Visit a parse tree produced by StellaParser#Tail.
	VisitTail(ctx *TailContext) interface{}

	// Visit a parse tree produced by StellaParser#Record.
	VisitRecord(ctx *RecordContext) interface{}

	// Visit a parse tree produced by StellaParser#LogicAnd.
	VisitLogicAnd(ctx *LogicAndContext) interface{}

	// Visit a parse tree produced by StellaParser#TypeApplication.
	VisitTypeApplication(ctx *TypeApplicationContext) interface{}

	// Visit a parse tree produced by StellaParser#LetRec.
	VisitLetRec(ctx *LetRecContext) interface{}

	// Visit a parse tree produced by StellaParser#LogicOr.
	VisitLogicOr(ctx *LogicOrContext) interface{}

	// Visit a parse tree produced by StellaParser#TryWith.
	VisitTryWith(ctx *TryWithContext) interface{}

	// Visit a parse tree produced by StellaParser#Pred.
	VisitPred(ctx *PredContext) interface{}

	// Visit a parse tree produced by StellaParser#TypeAsc.
	VisitTypeAsc(ctx *TypeAscContext) interface{}

	// Visit a parse tree produced by StellaParser#NatRec.
	VisitNatRec(ctx *NatRecContext) interface{}

	// Visit a parse tree produced by StellaParser#Unfold.
	VisitUnfold(ctx *UnfoldContext) interface{}

	// Visit a parse tree produced by StellaParser#Ref.
	VisitRef(ctx *RefContext) interface{}

	// Visit a parse tree produced by StellaParser#DotTuple.
	VisitDotTuple(ctx *DotTupleContext) interface{}

	// Visit a parse tree produced by StellaParser#Fix.
	VisitFix(ctx *FixContext) interface{}

	// Visit a parse tree produced by StellaParser#Let.
	VisitLet(ctx *LetContext) interface{}

	// Visit a parse tree produced by StellaParser#Assign.
	VisitAssign(ctx *AssignContext) interface{}

	// Visit a parse tree produced by StellaParser#Tuple.
	VisitTuple(ctx *TupleContext) interface{}

	// Visit a parse tree produced by StellaParser#ConsList.
	VisitConsList(ctx *ConsListContext) interface{}

	// Visit a parse tree produced by StellaParser#patternBinding.
	VisitPatternBinding(ctx *PatternBindingContext) interface{}

	// Visit a parse tree produced by StellaParser#binding.
	VisitBinding(ctx *BindingContext) interface{}

	// Visit a parse tree produced by StellaParser#matchCase.
	VisitMatchCase(ctx *MatchCaseContext) interface{}

	// Visit a parse tree produced by StellaParser#PatternCons.
	VisitPatternCons(ctx *PatternConsContext) interface{}

	// Visit a parse tree produced by StellaParser#PatternTuple.
	VisitPatternTuple(ctx *PatternTupleContext) interface{}

	// Visit a parse tree produced by StellaParser#PatternList.
	VisitPatternList(ctx *PatternListContext) interface{}

	// Visit a parse tree produced by StellaParser#PatternRecord.
	VisitPatternRecord(ctx *PatternRecordContext) interface{}

	// Visit a parse tree produced by StellaParser#PatternVariant.
	VisitPatternVariant(ctx *PatternVariantContext) interface{}

	// Visit a parse tree produced by StellaParser#PatternAsc.
	VisitPatternAsc(ctx *PatternAscContext) interface{}

	// Visit a parse tree produced by StellaParser#PatternInt.
	VisitPatternInt(ctx *PatternIntContext) interface{}

	// Visit a parse tree produced by StellaParser#PatternInr.
	VisitPatternInr(ctx *PatternInrContext) interface{}

	// Visit a parse tree produced by StellaParser#PatternTrue.
	VisitPatternTrue(ctx *PatternTrueContext) interface{}

	// Visit a parse tree produced by StellaParser#PatternInl.
	VisitPatternInl(ctx *PatternInlContext) interface{}

	// Visit a parse tree produced by StellaParser#PatternVar.
	VisitPatternVar(ctx *PatternVarContext) interface{}

	// Visit a parse tree produced by StellaParser#ParenthesisedPattern.
	VisitParenthesisedPattern(ctx *ParenthesisedPatternContext) interface{}

	// Visit a parse tree produced by StellaParser#PatternSucc.
	VisitPatternSucc(ctx *PatternSuccContext) interface{}

	// Visit a parse tree produced by StellaParser#PatternFalse.
	VisitPatternFalse(ctx *PatternFalseContext) interface{}

	// Visit a parse tree produced by StellaParser#PatternUnit.
	VisitPatternUnit(ctx *PatternUnitContext) interface{}

	// Visit a parse tree produced by StellaParser#PatternCastAs.
	VisitPatternCastAs(ctx *PatternCastAsContext) interface{}

	// Visit a parse tree produced by StellaParser#labelledPattern.
	VisitLabelledPattern(ctx *LabelledPatternContext) interface{}

	// Visit a parse tree produced by StellaParser#TypeTuple.
	VisitTypeTuple(ctx *TypeTupleContext) interface{}

	// Visit a parse tree produced by StellaParser#TypeTop.
	VisitTypeTop(ctx *TypeTopContext) interface{}

	// Visit a parse tree produced by StellaParser#TypeBool.
	VisitTypeBool(ctx *TypeBoolContext) interface{}

	// Visit a parse tree produced by StellaParser#TypeRef.
	VisitTypeRef(ctx *TypeRefContext) interface{}

	// Visit a parse tree produced by StellaParser#TypeRec.
	VisitTypeRec(ctx *TypeRecContext) interface{}

	// Visit a parse tree produced by StellaParser#TypeSum.
	VisitTypeSum(ctx *TypeSumContext) interface{}

	// Visit a parse tree produced by StellaParser#TypeVar.
	VisitTypeVar(ctx *TypeVarContext) interface{}

	// Visit a parse tree produced by StellaParser#TypeVariant.
	VisitTypeVariant(ctx *TypeVariantContext) interface{}

	// Visit a parse tree produced by StellaParser#TypeUnit.
	VisitTypeUnit(ctx *TypeUnitContext) interface{}

	// Visit a parse tree produced by StellaParser#TypeNat.
	VisitTypeNat(ctx *TypeNatContext) interface{}

	// Visit a parse tree produced by StellaParser#TypeBottom.
	VisitTypeBottom(ctx *TypeBottomContext) interface{}

	// Visit a parse tree produced by StellaParser#TypeParens.
	VisitTypeParens(ctx *TypeParensContext) interface{}

	// Visit a parse tree produced by StellaParser#TypeFun.
	VisitTypeFun(ctx *TypeFunContext) interface{}

	// Visit a parse tree produced by StellaParser#TypeForAll.
	VisitTypeForAll(ctx *TypeForAllContext) interface{}

	// Visit a parse tree produced by StellaParser#TypeRecord.
	VisitTypeRecord(ctx *TypeRecordContext) interface{}

	// Visit a parse tree produced by StellaParser#TypeList.
	VisitTypeList(ctx *TypeListContext) interface{}

	// Visit a parse tree produced by StellaParser#recordFieldType.
	VisitRecordFieldType(ctx *RecordFieldTypeContext) interface{}

	// Visit a parse tree produced by StellaParser#variantFieldType.
	VisitVariantFieldType(ctx *VariantFieldTypeContext) interface{}

}