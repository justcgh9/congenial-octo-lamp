// Code generated from StellaParser.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // StellaParser

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// BaseStellaParserListener is a complete listener for a parse tree produced by StellaParser.
type BaseStellaParserListener struct{}

var _ StellaParserListener = &BaseStellaParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseStellaParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseStellaParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseStellaParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseStellaParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart_Program is called when production start_Program is entered.
func (s *BaseStellaParserListener) EnterStart_Program(ctx *Start_ProgramContext) {}

// ExitStart_Program is called when production start_Program is exited.
func (s *BaseStellaParserListener) ExitStart_Program(ctx *Start_ProgramContext) {}

// EnterStart_Expr is called when production start_Expr is entered.
func (s *BaseStellaParserListener) EnterStart_Expr(ctx *Start_ExprContext) {}

// ExitStart_Expr is called when production start_Expr is exited.
func (s *BaseStellaParserListener) ExitStart_Expr(ctx *Start_ExprContext) {}

// EnterStart_Type is called when production start_Type is entered.
func (s *BaseStellaParserListener) EnterStart_Type(ctx *Start_TypeContext) {}

// ExitStart_Type is called when production start_Type is exited.
func (s *BaseStellaParserListener) ExitStart_Type(ctx *Start_TypeContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseStellaParserListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseStellaParserListener) ExitProgram(ctx *ProgramContext) {}

// EnterLanguageCore is called when production LanguageCore is entered.
func (s *BaseStellaParserListener) EnterLanguageCore(ctx *LanguageCoreContext) {}

// ExitLanguageCore is called when production LanguageCore is exited.
func (s *BaseStellaParserListener) ExitLanguageCore(ctx *LanguageCoreContext) {}

// EnterAnExtension is called when production AnExtension is entered.
func (s *BaseStellaParserListener) EnterAnExtension(ctx *AnExtensionContext) {}

// ExitAnExtension is called when production AnExtension is exited.
func (s *BaseStellaParserListener) ExitAnExtension(ctx *AnExtensionContext) {}

// EnterDeclFun is called when production DeclFun is entered.
func (s *BaseStellaParserListener) EnterDeclFun(ctx *DeclFunContext) {}

// ExitDeclFun is called when production DeclFun is exited.
func (s *BaseStellaParserListener) ExitDeclFun(ctx *DeclFunContext) {}

// EnterDeclFunGeneric is called when production DeclFunGeneric is entered.
func (s *BaseStellaParserListener) EnterDeclFunGeneric(ctx *DeclFunGenericContext) {}

// ExitDeclFunGeneric is called when production DeclFunGeneric is exited.
func (s *BaseStellaParserListener) ExitDeclFunGeneric(ctx *DeclFunGenericContext) {}

// EnterDeclTypeAlias is called when production DeclTypeAlias is entered.
func (s *BaseStellaParserListener) EnterDeclTypeAlias(ctx *DeclTypeAliasContext) {}

// ExitDeclTypeAlias is called when production DeclTypeAlias is exited.
func (s *BaseStellaParserListener) ExitDeclTypeAlias(ctx *DeclTypeAliasContext) {}

// EnterDeclExceptionType is called when production DeclExceptionType is entered.
func (s *BaseStellaParserListener) EnterDeclExceptionType(ctx *DeclExceptionTypeContext) {}

// ExitDeclExceptionType is called when production DeclExceptionType is exited.
func (s *BaseStellaParserListener) ExitDeclExceptionType(ctx *DeclExceptionTypeContext) {}

// EnterDeclExceptionVariant is called when production DeclExceptionVariant is entered.
func (s *BaseStellaParserListener) EnterDeclExceptionVariant(ctx *DeclExceptionVariantContext) {}

// ExitDeclExceptionVariant is called when production DeclExceptionVariant is exited.
func (s *BaseStellaParserListener) ExitDeclExceptionVariant(ctx *DeclExceptionVariantContext) {}

// EnterInlineAnnotation is called when production InlineAnnotation is entered.
func (s *BaseStellaParserListener) EnterInlineAnnotation(ctx *InlineAnnotationContext) {}

// ExitInlineAnnotation is called when production InlineAnnotation is exited.
func (s *BaseStellaParserListener) ExitInlineAnnotation(ctx *InlineAnnotationContext) {}

// EnterParamDecl is called when production paramDecl is entered.
func (s *BaseStellaParserListener) EnterParamDecl(ctx *ParamDeclContext) {}

// ExitParamDecl is called when production paramDecl is exited.
func (s *BaseStellaParserListener) ExitParamDecl(ctx *ParamDeclContext) {}

// EnterFold is called when production Fold is entered.
func (s *BaseStellaParserListener) EnterFold(ctx *FoldContext) {}

// ExitFold is called when production Fold is exited.
func (s *BaseStellaParserListener) ExitFold(ctx *FoldContext) {}

// EnterAdd is called when production Add is entered.
func (s *BaseStellaParserListener) EnterAdd(ctx *AddContext) {}

// ExitAdd is called when production Add is exited.
func (s *BaseStellaParserListener) ExitAdd(ctx *AddContext) {}

// EnterIsZero is called when production IsZero is entered.
func (s *BaseStellaParserListener) EnterIsZero(ctx *IsZeroContext) {}

// ExitIsZero is called when production IsZero is exited.
func (s *BaseStellaParserListener) ExitIsZero(ctx *IsZeroContext) {}

// EnterVar is called when production Var is entered.
func (s *BaseStellaParserListener) EnterVar(ctx *VarContext) {}

// ExitVar is called when production Var is exited.
func (s *BaseStellaParserListener) ExitVar(ctx *VarContext) {}

// EnterTypeAbstraction is called when production TypeAbstraction is entered.
func (s *BaseStellaParserListener) EnterTypeAbstraction(ctx *TypeAbstractionContext) {}

// ExitTypeAbstraction is called when production TypeAbstraction is exited.
func (s *BaseStellaParserListener) ExitTypeAbstraction(ctx *TypeAbstractionContext) {}

// EnterDivide is called when production Divide is entered.
func (s *BaseStellaParserListener) EnterDivide(ctx *DivideContext) {}

// ExitDivide is called when production Divide is exited.
func (s *BaseStellaParserListener) ExitDivide(ctx *DivideContext) {}

// EnterLessThan is called when production LessThan is entered.
func (s *BaseStellaParserListener) EnterLessThan(ctx *LessThanContext) {}

// ExitLessThan is called when production LessThan is exited.
func (s *BaseStellaParserListener) ExitLessThan(ctx *LessThanContext) {}

// EnterDotRecord is called when production DotRecord is entered.
func (s *BaseStellaParserListener) EnterDotRecord(ctx *DotRecordContext) {}

// ExitDotRecord is called when production DotRecord is exited.
func (s *BaseStellaParserListener) ExitDotRecord(ctx *DotRecordContext) {}

// EnterGreaterThan is called when production GreaterThan is entered.
func (s *BaseStellaParserListener) EnterGreaterThan(ctx *GreaterThanContext) {}

// ExitGreaterThan is called when production GreaterThan is exited.
func (s *BaseStellaParserListener) ExitGreaterThan(ctx *GreaterThanContext) {}

// EnterEqual is called when production Equal is entered.
func (s *BaseStellaParserListener) EnterEqual(ctx *EqualContext) {}

// ExitEqual is called when production Equal is exited.
func (s *BaseStellaParserListener) ExitEqual(ctx *EqualContext) {}

// EnterThrow is called when production Throw is entered.
func (s *BaseStellaParserListener) EnterThrow(ctx *ThrowContext) {}

// ExitThrow is called when production Throw is exited.
func (s *BaseStellaParserListener) ExitThrow(ctx *ThrowContext) {}

// EnterMultiply is called when production Multiply is entered.
func (s *BaseStellaParserListener) EnterMultiply(ctx *MultiplyContext) {}

// ExitMultiply is called when production Multiply is exited.
func (s *BaseStellaParserListener) ExitMultiply(ctx *MultiplyContext) {}

// EnterConstMemory is called when production ConstMemory is entered.
func (s *BaseStellaParserListener) EnterConstMemory(ctx *ConstMemoryContext) {}

// ExitConstMemory is called when production ConstMemory is exited.
func (s *BaseStellaParserListener) ExitConstMemory(ctx *ConstMemoryContext) {}

// EnterList is called when production List is entered.
func (s *BaseStellaParserListener) EnterList(ctx *ListContext) {}

// ExitList is called when production List is exited.
func (s *BaseStellaParserListener) ExitList(ctx *ListContext) {}

// EnterTryCatch is called when production TryCatch is entered.
func (s *BaseStellaParserListener) EnterTryCatch(ctx *TryCatchContext) {}

// ExitTryCatch is called when production TryCatch is exited.
func (s *BaseStellaParserListener) ExitTryCatch(ctx *TryCatchContext) {}

// EnterTryCastAs is called when production TryCastAs is entered.
func (s *BaseStellaParserListener) EnterTryCastAs(ctx *TryCastAsContext) {}

// ExitTryCastAs is called when production TryCastAs is exited.
func (s *BaseStellaParserListener) ExitTryCastAs(ctx *TryCastAsContext) {}

// EnterHead is called when production Head is entered.
func (s *BaseStellaParserListener) EnterHead(ctx *HeadContext) {}

// ExitHead is called when production Head is exited.
func (s *BaseStellaParserListener) ExitHead(ctx *HeadContext) {}

// EnterTerminatingSemicolon is called when production TerminatingSemicolon is entered.
func (s *BaseStellaParserListener) EnterTerminatingSemicolon(ctx *TerminatingSemicolonContext) {}

// ExitTerminatingSemicolon is called when production TerminatingSemicolon is exited.
func (s *BaseStellaParserListener) ExitTerminatingSemicolon(ctx *TerminatingSemicolonContext) {}

// EnterNotEqual is called when production NotEqual is entered.
func (s *BaseStellaParserListener) EnterNotEqual(ctx *NotEqualContext) {}

// ExitNotEqual is called when production NotEqual is exited.
func (s *BaseStellaParserListener) ExitNotEqual(ctx *NotEqualContext) {}

// EnterConstUnit is called when production ConstUnit is entered.
func (s *BaseStellaParserListener) EnterConstUnit(ctx *ConstUnitContext) {}

// ExitConstUnit is called when production ConstUnit is exited.
func (s *BaseStellaParserListener) ExitConstUnit(ctx *ConstUnitContext) {}

// EnterSequence is called when production Sequence is entered.
func (s *BaseStellaParserListener) EnterSequence(ctx *SequenceContext) {}

// ExitSequence is called when production Sequence is exited.
func (s *BaseStellaParserListener) ExitSequence(ctx *SequenceContext) {}

// EnterConstFalse is called when production ConstFalse is entered.
func (s *BaseStellaParserListener) EnterConstFalse(ctx *ConstFalseContext) {}

// ExitConstFalse is called when production ConstFalse is exited.
func (s *BaseStellaParserListener) ExitConstFalse(ctx *ConstFalseContext) {}

// EnterAbstraction is called when production Abstraction is entered.
func (s *BaseStellaParserListener) EnterAbstraction(ctx *AbstractionContext) {}

// ExitAbstraction is called when production Abstraction is exited.
func (s *BaseStellaParserListener) ExitAbstraction(ctx *AbstractionContext) {}

// EnterConstInt is called when production ConstInt is entered.
func (s *BaseStellaParserListener) EnterConstInt(ctx *ConstIntContext) {}

// ExitConstInt is called when production ConstInt is exited.
func (s *BaseStellaParserListener) ExitConstInt(ctx *ConstIntContext) {}

// EnterVariant is called when production Variant is entered.
func (s *BaseStellaParserListener) EnterVariant(ctx *VariantContext) {}

// ExitVariant is called when production Variant is exited.
func (s *BaseStellaParserListener) ExitVariant(ctx *VariantContext) {}

// EnterConstTrue is called when production ConstTrue is entered.
func (s *BaseStellaParserListener) EnterConstTrue(ctx *ConstTrueContext) {}

// ExitConstTrue is called when production ConstTrue is exited.
func (s *BaseStellaParserListener) ExitConstTrue(ctx *ConstTrueContext) {}

// EnterSubtract is called when production Subtract is entered.
func (s *BaseStellaParserListener) EnterSubtract(ctx *SubtractContext) {}

// ExitSubtract is called when production Subtract is exited.
func (s *BaseStellaParserListener) ExitSubtract(ctx *SubtractContext) {}

// EnterTypeCast is called when production TypeCast is entered.
func (s *BaseStellaParserListener) EnterTypeCast(ctx *TypeCastContext) {}

// ExitTypeCast is called when production TypeCast is exited.
func (s *BaseStellaParserListener) ExitTypeCast(ctx *TypeCastContext) {}

// EnterIf is called when production If is entered.
func (s *BaseStellaParserListener) EnterIf(ctx *IfContext) {}

// ExitIf is called when production If is exited.
func (s *BaseStellaParserListener) ExitIf(ctx *IfContext) {}

// EnterApplication is called when production Application is entered.
func (s *BaseStellaParserListener) EnterApplication(ctx *ApplicationContext) {}

// ExitApplication is called when production Application is exited.
func (s *BaseStellaParserListener) ExitApplication(ctx *ApplicationContext) {}

// EnterDeref is called when production Deref is entered.
func (s *BaseStellaParserListener) EnterDeref(ctx *DerefContext) {}

// ExitDeref is called when production Deref is exited.
func (s *BaseStellaParserListener) ExitDeref(ctx *DerefContext) {}

// EnterIsEmpty is called when production IsEmpty is entered.
func (s *BaseStellaParserListener) EnterIsEmpty(ctx *IsEmptyContext) {}

// ExitIsEmpty is called when production IsEmpty is exited.
func (s *BaseStellaParserListener) ExitIsEmpty(ctx *IsEmptyContext) {}

// EnterPanic is called when production Panic is entered.
func (s *BaseStellaParserListener) EnterPanic(ctx *PanicContext) {}

// ExitPanic is called when production Panic is exited.
func (s *BaseStellaParserListener) ExitPanic(ctx *PanicContext) {}

// EnterLessThanOrEqual is called when production LessThanOrEqual is entered.
func (s *BaseStellaParserListener) EnterLessThanOrEqual(ctx *LessThanOrEqualContext) {}

// ExitLessThanOrEqual is called when production LessThanOrEqual is exited.
func (s *BaseStellaParserListener) ExitLessThanOrEqual(ctx *LessThanOrEqualContext) {}

// EnterSucc is called when production Succ is entered.
func (s *BaseStellaParserListener) EnterSucc(ctx *SuccContext) {}

// ExitSucc is called when production Succ is exited.
func (s *BaseStellaParserListener) ExitSucc(ctx *SuccContext) {}

// EnterInl is called when production Inl is entered.
func (s *BaseStellaParserListener) EnterInl(ctx *InlContext) {}

// ExitInl is called when production Inl is exited.
func (s *BaseStellaParserListener) ExitInl(ctx *InlContext) {}

// EnterGreaterThanOrEqual is called when production GreaterThanOrEqual is entered.
func (s *BaseStellaParserListener) EnterGreaterThanOrEqual(ctx *GreaterThanOrEqualContext) {}

// ExitGreaterThanOrEqual is called when production GreaterThanOrEqual is exited.
func (s *BaseStellaParserListener) ExitGreaterThanOrEqual(ctx *GreaterThanOrEqualContext) {}

// EnterInr is called when production Inr is entered.
func (s *BaseStellaParserListener) EnterInr(ctx *InrContext) {}

// ExitInr is called when production Inr is exited.
func (s *BaseStellaParserListener) ExitInr(ctx *InrContext) {}

// EnterMatch is called when production Match is entered.
func (s *BaseStellaParserListener) EnterMatch(ctx *MatchContext) {}

// ExitMatch is called when production Match is exited.
func (s *BaseStellaParserListener) ExitMatch(ctx *MatchContext) {}

// EnterLogicNot is called when production LogicNot is entered.
func (s *BaseStellaParserListener) EnterLogicNot(ctx *LogicNotContext) {}

// ExitLogicNot is called when production LogicNot is exited.
func (s *BaseStellaParserListener) ExitLogicNot(ctx *LogicNotContext) {}

// EnterParenthesisedExpr is called when production ParenthesisedExpr is entered.
func (s *BaseStellaParserListener) EnterParenthesisedExpr(ctx *ParenthesisedExprContext) {}

// ExitParenthesisedExpr is called when production ParenthesisedExpr is exited.
func (s *BaseStellaParserListener) ExitParenthesisedExpr(ctx *ParenthesisedExprContext) {}

// EnterTail is called when production Tail is entered.
func (s *BaseStellaParserListener) EnterTail(ctx *TailContext) {}

// ExitTail is called when production Tail is exited.
func (s *BaseStellaParserListener) ExitTail(ctx *TailContext) {}

// EnterRecord is called when production Record is entered.
func (s *BaseStellaParserListener) EnterRecord(ctx *RecordContext) {}

// ExitRecord is called when production Record is exited.
func (s *BaseStellaParserListener) ExitRecord(ctx *RecordContext) {}

// EnterLogicAnd is called when production LogicAnd is entered.
func (s *BaseStellaParserListener) EnterLogicAnd(ctx *LogicAndContext) {}

// ExitLogicAnd is called when production LogicAnd is exited.
func (s *BaseStellaParserListener) ExitLogicAnd(ctx *LogicAndContext) {}

// EnterTypeApplication is called when production TypeApplication is entered.
func (s *BaseStellaParserListener) EnterTypeApplication(ctx *TypeApplicationContext) {}

// ExitTypeApplication is called when production TypeApplication is exited.
func (s *BaseStellaParserListener) ExitTypeApplication(ctx *TypeApplicationContext) {}

// EnterLetRec is called when production LetRec is entered.
func (s *BaseStellaParserListener) EnterLetRec(ctx *LetRecContext) {}

// ExitLetRec is called when production LetRec is exited.
func (s *BaseStellaParserListener) ExitLetRec(ctx *LetRecContext) {}

// EnterLogicOr is called when production LogicOr is entered.
func (s *BaseStellaParserListener) EnterLogicOr(ctx *LogicOrContext) {}

// ExitLogicOr is called when production LogicOr is exited.
func (s *BaseStellaParserListener) ExitLogicOr(ctx *LogicOrContext) {}

// EnterTryWith is called when production TryWith is entered.
func (s *BaseStellaParserListener) EnterTryWith(ctx *TryWithContext) {}

// ExitTryWith is called when production TryWith is exited.
func (s *BaseStellaParserListener) ExitTryWith(ctx *TryWithContext) {}

// EnterPred is called when production Pred is entered.
func (s *BaseStellaParserListener) EnterPred(ctx *PredContext) {}

// ExitPred is called when production Pred is exited.
func (s *BaseStellaParserListener) ExitPred(ctx *PredContext) {}

// EnterTypeAsc is called when production TypeAsc is entered.
func (s *BaseStellaParserListener) EnterTypeAsc(ctx *TypeAscContext) {}

// ExitTypeAsc is called when production TypeAsc is exited.
func (s *BaseStellaParserListener) ExitTypeAsc(ctx *TypeAscContext) {}

// EnterNatRec is called when production NatRec is entered.
func (s *BaseStellaParserListener) EnterNatRec(ctx *NatRecContext) {}

// ExitNatRec is called when production NatRec is exited.
func (s *BaseStellaParserListener) ExitNatRec(ctx *NatRecContext) {}

// EnterUnfold is called when production Unfold is entered.
func (s *BaseStellaParserListener) EnterUnfold(ctx *UnfoldContext) {}

// ExitUnfold is called when production Unfold is exited.
func (s *BaseStellaParserListener) ExitUnfold(ctx *UnfoldContext) {}

// EnterRef is called when production Ref is entered.
func (s *BaseStellaParserListener) EnterRef(ctx *RefContext) {}

// ExitRef is called when production Ref is exited.
func (s *BaseStellaParserListener) ExitRef(ctx *RefContext) {}

// EnterDotTuple is called when production DotTuple is entered.
func (s *BaseStellaParserListener) EnterDotTuple(ctx *DotTupleContext) {}

// ExitDotTuple is called when production DotTuple is exited.
func (s *BaseStellaParserListener) ExitDotTuple(ctx *DotTupleContext) {}

// EnterFix is called when production Fix is entered.
func (s *BaseStellaParserListener) EnterFix(ctx *FixContext) {}

// ExitFix is called when production Fix is exited.
func (s *BaseStellaParserListener) ExitFix(ctx *FixContext) {}

// EnterLet is called when production Let is entered.
func (s *BaseStellaParserListener) EnterLet(ctx *LetContext) {}

// ExitLet is called when production Let is exited.
func (s *BaseStellaParserListener) ExitLet(ctx *LetContext) {}

// EnterAssign is called when production Assign is entered.
func (s *BaseStellaParserListener) EnterAssign(ctx *AssignContext) {}

// ExitAssign is called when production Assign is exited.
func (s *BaseStellaParserListener) ExitAssign(ctx *AssignContext) {}

// EnterTuple is called when production Tuple is entered.
func (s *BaseStellaParserListener) EnterTuple(ctx *TupleContext) {}

// ExitTuple is called when production Tuple is exited.
func (s *BaseStellaParserListener) ExitTuple(ctx *TupleContext) {}

// EnterConsList is called when production ConsList is entered.
func (s *BaseStellaParserListener) EnterConsList(ctx *ConsListContext) {}

// ExitConsList is called when production ConsList is exited.
func (s *BaseStellaParserListener) ExitConsList(ctx *ConsListContext) {}

// EnterPatternBinding is called when production patternBinding is entered.
func (s *BaseStellaParserListener) EnterPatternBinding(ctx *PatternBindingContext) {}

// ExitPatternBinding is called when production patternBinding is exited.
func (s *BaseStellaParserListener) ExitPatternBinding(ctx *PatternBindingContext) {}

// EnterBinding is called when production binding is entered.
func (s *BaseStellaParserListener) EnterBinding(ctx *BindingContext) {}

// ExitBinding is called when production binding is exited.
func (s *BaseStellaParserListener) ExitBinding(ctx *BindingContext) {}

// EnterMatchCase is called when production matchCase is entered.
func (s *BaseStellaParserListener) EnterMatchCase(ctx *MatchCaseContext) {}

// ExitMatchCase is called when production matchCase is exited.
func (s *BaseStellaParserListener) ExitMatchCase(ctx *MatchCaseContext) {}

// EnterPatternCons is called when production PatternCons is entered.
func (s *BaseStellaParserListener) EnterPatternCons(ctx *PatternConsContext) {}

// ExitPatternCons is called when production PatternCons is exited.
func (s *BaseStellaParserListener) ExitPatternCons(ctx *PatternConsContext) {}

// EnterPatternTuple is called when production PatternTuple is entered.
func (s *BaseStellaParserListener) EnterPatternTuple(ctx *PatternTupleContext) {}

// ExitPatternTuple is called when production PatternTuple is exited.
func (s *BaseStellaParserListener) ExitPatternTuple(ctx *PatternTupleContext) {}

// EnterPatternList is called when production PatternList is entered.
func (s *BaseStellaParserListener) EnterPatternList(ctx *PatternListContext) {}

// ExitPatternList is called when production PatternList is exited.
func (s *BaseStellaParserListener) ExitPatternList(ctx *PatternListContext) {}

// EnterPatternRecord is called when production PatternRecord is entered.
func (s *BaseStellaParserListener) EnterPatternRecord(ctx *PatternRecordContext) {}

// ExitPatternRecord is called when production PatternRecord is exited.
func (s *BaseStellaParserListener) ExitPatternRecord(ctx *PatternRecordContext) {}

// EnterPatternVariant is called when production PatternVariant is entered.
func (s *BaseStellaParserListener) EnterPatternVariant(ctx *PatternVariantContext) {}

// ExitPatternVariant is called when production PatternVariant is exited.
func (s *BaseStellaParserListener) ExitPatternVariant(ctx *PatternVariantContext) {}

// EnterPatternAsc is called when production PatternAsc is entered.
func (s *BaseStellaParserListener) EnterPatternAsc(ctx *PatternAscContext) {}

// ExitPatternAsc is called when production PatternAsc is exited.
func (s *BaseStellaParserListener) ExitPatternAsc(ctx *PatternAscContext) {}

// EnterPatternInt is called when production PatternInt is entered.
func (s *BaseStellaParserListener) EnterPatternInt(ctx *PatternIntContext) {}

// ExitPatternInt is called when production PatternInt is exited.
func (s *BaseStellaParserListener) ExitPatternInt(ctx *PatternIntContext) {}

// EnterPatternInr is called when production PatternInr is entered.
func (s *BaseStellaParserListener) EnterPatternInr(ctx *PatternInrContext) {}

// ExitPatternInr is called when production PatternInr is exited.
func (s *BaseStellaParserListener) ExitPatternInr(ctx *PatternInrContext) {}

// EnterPatternTrue is called when production PatternTrue is entered.
func (s *BaseStellaParserListener) EnterPatternTrue(ctx *PatternTrueContext) {}

// ExitPatternTrue is called when production PatternTrue is exited.
func (s *BaseStellaParserListener) ExitPatternTrue(ctx *PatternTrueContext) {}

// EnterPatternInl is called when production PatternInl is entered.
func (s *BaseStellaParserListener) EnterPatternInl(ctx *PatternInlContext) {}

// ExitPatternInl is called when production PatternInl is exited.
func (s *BaseStellaParserListener) ExitPatternInl(ctx *PatternInlContext) {}

// EnterPatternVar is called when production PatternVar is entered.
func (s *BaseStellaParserListener) EnterPatternVar(ctx *PatternVarContext) {}

// ExitPatternVar is called when production PatternVar is exited.
func (s *BaseStellaParserListener) ExitPatternVar(ctx *PatternVarContext) {}

// EnterParenthesisedPattern is called when production ParenthesisedPattern is entered.
func (s *BaseStellaParserListener) EnterParenthesisedPattern(ctx *ParenthesisedPatternContext) {}

// ExitParenthesisedPattern is called when production ParenthesisedPattern is exited.
func (s *BaseStellaParserListener) ExitParenthesisedPattern(ctx *ParenthesisedPatternContext) {}

// EnterPatternSucc is called when production PatternSucc is entered.
func (s *BaseStellaParserListener) EnterPatternSucc(ctx *PatternSuccContext) {}

// ExitPatternSucc is called when production PatternSucc is exited.
func (s *BaseStellaParserListener) ExitPatternSucc(ctx *PatternSuccContext) {}

// EnterPatternFalse is called when production PatternFalse is entered.
func (s *BaseStellaParserListener) EnterPatternFalse(ctx *PatternFalseContext) {}

// ExitPatternFalse is called when production PatternFalse is exited.
func (s *BaseStellaParserListener) ExitPatternFalse(ctx *PatternFalseContext) {}

// EnterPatternUnit is called when production PatternUnit is entered.
func (s *BaseStellaParserListener) EnterPatternUnit(ctx *PatternUnitContext) {}

// ExitPatternUnit is called when production PatternUnit is exited.
func (s *BaseStellaParserListener) ExitPatternUnit(ctx *PatternUnitContext) {}

// EnterPatternCastAs is called when production PatternCastAs is entered.
func (s *BaseStellaParserListener) EnterPatternCastAs(ctx *PatternCastAsContext) {}

// ExitPatternCastAs is called when production PatternCastAs is exited.
func (s *BaseStellaParserListener) ExitPatternCastAs(ctx *PatternCastAsContext) {}

// EnterLabelledPattern is called when production labelledPattern is entered.
func (s *BaseStellaParserListener) EnterLabelledPattern(ctx *LabelledPatternContext) {}

// ExitLabelledPattern is called when production labelledPattern is exited.
func (s *BaseStellaParserListener) ExitLabelledPattern(ctx *LabelledPatternContext) {}

// EnterTypeTuple is called when production TypeTuple is entered.
func (s *BaseStellaParserListener) EnterTypeTuple(ctx *TypeTupleContext) {}

// ExitTypeTuple is called when production TypeTuple is exited.
func (s *BaseStellaParserListener) ExitTypeTuple(ctx *TypeTupleContext) {}

// EnterTypeTop is called when production TypeTop is entered.
func (s *BaseStellaParserListener) EnterTypeTop(ctx *TypeTopContext) {}

// ExitTypeTop is called when production TypeTop is exited.
func (s *BaseStellaParserListener) ExitTypeTop(ctx *TypeTopContext) {}

// EnterTypeBool is called when production TypeBool is entered.
func (s *BaseStellaParserListener) EnterTypeBool(ctx *TypeBoolContext) {}

// ExitTypeBool is called when production TypeBool is exited.
func (s *BaseStellaParserListener) ExitTypeBool(ctx *TypeBoolContext) {}

// EnterTypeRef is called when production TypeRef is entered.
func (s *BaseStellaParserListener) EnterTypeRef(ctx *TypeRefContext) {}

// ExitTypeRef is called when production TypeRef is exited.
func (s *BaseStellaParserListener) ExitTypeRef(ctx *TypeRefContext) {}

// EnterTypeRec is called when production TypeRec is entered.
func (s *BaseStellaParserListener) EnterTypeRec(ctx *TypeRecContext) {}

// ExitTypeRec is called when production TypeRec is exited.
func (s *BaseStellaParserListener) ExitTypeRec(ctx *TypeRecContext) {}

// EnterTypeSum is called when production TypeSum is entered.
func (s *BaseStellaParserListener) EnterTypeSum(ctx *TypeSumContext) {}

// ExitTypeSum is called when production TypeSum is exited.
func (s *BaseStellaParserListener) ExitTypeSum(ctx *TypeSumContext) {}

// EnterTypeVar is called when production TypeVar is entered.
func (s *BaseStellaParserListener) EnterTypeVar(ctx *TypeVarContext) {}

// ExitTypeVar is called when production TypeVar is exited.
func (s *BaseStellaParserListener) ExitTypeVar(ctx *TypeVarContext) {}

// EnterTypeVariant is called when production TypeVariant is entered.
func (s *BaseStellaParserListener) EnterTypeVariant(ctx *TypeVariantContext) {}

// ExitTypeVariant is called when production TypeVariant is exited.
func (s *BaseStellaParserListener) ExitTypeVariant(ctx *TypeVariantContext) {}

// EnterTypeUnit is called when production TypeUnit is entered.
func (s *BaseStellaParserListener) EnterTypeUnit(ctx *TypeUnitContext) {}

// ExitTypeUnit is called when production TypeUnit is exited.
func (s *BaseStellaParserListener) ExitTypeUnit(ctx *TypeUnitContext) {}

// EnterTypeNat is called when production TypeNat is entered.
func (s *BaseStellaParserListener) EnterTypeNat(ctx *TypeNatContext) {}

// ExitTypeNat is called when production TypeNat is exited.
func (s *BaseStellaParserListener) ExitTypeNat(ctx *TypeNatContext) {}

// EnterTypeBottom is called when production TypeBottom is entered.
func (s *BaseStellaParserListener) EnterTypeBottom(ctx *TypeBottomContext) {}

// ExitTypeBottom is called when production TypeBottom is exited.
func (s *BaseStellaParserListener) ExitTypeBottom(ctx *TypeBottomContext) {}

// EnterTypeParens is called when production TypeParens is entered.
func (s *BaseStellaParserListener) EnterTypeParens(ctx *TypeParensContext) {}

// ExitTypeParens is called when production TypeParens is exited.
func (s *BaseStellaParserListener) ExitTypeParens(ctx *TypeParensContext) {}

// EnterTypeFun is called when production TypeFun is entered.
func (s *BaseStellaParserListener) EnterTypeFun(ctx *TypeFunContext) {}

// ExitTypeFun is called when production TypeFun is exited.
func (s *BaseStellaParserListener) ExitTypeFun(ctx *TypeFunContext) {}

// EnterTypeForAll is called when production TypeForAll is entered.
func (s *BaseStellaParserListener) EnterTypeForAll(ctx *TypeForAllContext) {}

// ExitTypeForAll is called when production TypeForAll is exited.
func (s *BaseStellaParserListener) ExitTypeForAll(ctx *TypeForAllContext) {}

// EnterTypeRecord is called when production TypeRecord is entered.
func (s *BaseStellaParserListener) EnterTypeRecord(ctx *TypeRecordContext) {}

// ExitTypeRecord is called when production TypeRecord is exited.
func (s *BaseStellaParserListener) ExitTypeRecord(ctx *TypeRecordContext) {}

// EnterTypeList is called when production TypeList is entered.
func (s *BaseStellaParserListener) EnterTypeList(ctx *TypeListContext) {}

// ExitTypeList is called when production TypeList is exited.
func (s *BaseStellaParserListener) ExitTypeList(ctx *TypeListContext) {}

// EnterRecordFieldType is called when production recordFieldType is entered.
func (s *BaseStellaParserListener) EnterRecordFieldType(ctx *RecordFieldTypeContext) {}

// ExitRecordFieldType is called when production recordFieldType is exited.
func (s *BaseStellaParserListener) ExitRecordFieldType(ctx *RecordFieldTypeContext) {}

// EnterVariantFieldType is called when production variantFieldType is entered.
func (s *BaseStellaParserListener) EnterVariantFieldType(ctx *VariantFieldTypeContext) {}

// ExitVariantFieldType is called when production variantFieldType is exited.
func (s *BaseStellaParserListener) ExitVariantFieldType(ctx *VariantFieldTypeContext) {}
