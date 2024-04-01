// Code generated from StellaParser.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // StellaParser

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// StellaParserListener is a complete listener for a parse tree produced by StellaParser.
type StellaParserListener interface {
	antlr.ParseTreeListener

	// EnterStart_Program is called when entering the start_Program production.
	EnterStart_Program(c *Start_ProgramContext)

	// EnterStart_Expr is called when entering the start_Expr production.
	EnterStart_Expr(c *Start_ExprContext)

	// EnterStart_Type is called when entering the start_Type production.
	EnterStart_Type(c *Start_TypeContext)

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterLanguageCore is called when entering the LanguageCore production.
	EnterLanguageCore(c *LanguageCoreContext)

	// EnterAnExtension is called when entering the AnExtension production.
	EnterAnExtension(c *AnExtensionContext)

	// EnterDeclFun is called when entering the DeclFun production.
	EnterDeclFun(c *DeclFunContext)

	// EnterDeclFunGeneric is called when entering the DeclFunGeneric production.
	EnterDeclFunGeneric(c *DeclFunGenericContext)

	// EnterDeclTypeAlias is called when entering the DeclTypeAlias production.
	EnterDeclTypeAlias(c *DeclTypeAliasContext)

	// EnterDeclExceptionType is called when entering the DeclExceptionType production.
	EnterDeclExceptionType(c *DeclExceptionTypeContext)

	// EnterDeclExceptionVariant is called when entering the DeclExceptionVariant production.
	EnterDeclExceptionVariant(c *DeclExceptionVariantContext)

	// EnterInlineAnnotation is called when entering the InlineAnnotation production.
	EnterInlineAnnotation(c *InlineAnnotationContext)

	// EnterParamDecl is called when entering the paramDecl production.
	EnterParamDecl(c *ParamDeclContext)

	// EnterFold is called when entering the Fold production.
	EnterFold(c *FoldContext)

	// EnterAdd is called when entering the Add production.
	EnterAdd(c *AddContext)

	// EnterIsZero is called when entering the IsZero production.
	EnterIsZero(c *IsZeroContext)

	// EnterVar is called when entering the Var production.
	EnterVar(c *VarContext)

	// EnterTypeAbstraction is called when entering the TypeAbstraction production.
	EnterTypeAbstraction(c *TypeAbstractionContext)

	// EnterDivide is called when entering the Divide production.
	EnterDivide(c *DivideContext)

	// EnterLessThan is called when entering the LessThan production.
	EnterLessThan(c *LessThanContext)

	// EnterDotRecord is called when entering the DotRecord production.
	EnterDotRecord(c *DotRecordContext)

	// EnterGreaterThan is called when entering the GreaterThan production.
	EnterGreaterThan(c *GreaterThanContext)

	// EnterEqual is called when entering the Equal production.
	EnterEqual(c *EqualContext)

	// EnterThrow is called when entering the Throw production.
	EnterThrow(c *ThrowContext)

	// EnterMultiply is called when entering the Multiply production.
	EnterMultiply(c *MultiplyContext)

	// EnterConstMemory is called when entering the ConstMemory production.
	EnterConstMemory(c *ConstMemoryContext)

	// EnterList is called when entering the List production.
	EnterList(c *ListContext)

	// EnterTryCatch is called when entering the TryCatch production.
	EnterTryCatch(c *TryCatchContext)

	// EnterTryCastAs is called when entering the TryCastAs production.
	EnterTryCastAs(c *TryCastAsContext)

	// EnterHead is called when entering the Head production.
	EnterHead(c *HeadContext)

	// EnterTerminatingSemicolon is called when entering the TerminatingSemicolon production.
	EnterTerminatingSemicolon(c *TerminatingSemicolonContext)

	// EnterNotEqual is called when entering the NotEqual production.
	EnterNotEqual(c *NotEqualContext)

	// EnterConstUnit is called when entering the ConstUnit production.
	EnterConstUnit(c *ConstUnitContext)

	// EnterSequence is called when entering the Sequence production.
	EnterSequence(c *SequenceContext)

	// EnterConstFalse is called when entering the ConstFalse production.
	EnterConstFalse(c *ConstFalseContext)

	// EnterAbstraction is called when entering the Abstraction production.
	EnterAbstraction(c *AbstractionContext)

	// EnterConstInt is called when entering the ConstInt production.
	EnterConstInt(c *ConstIntContext)

	// EnterVariant is called when entering the Variant production.
	EnterVariant(c *VariantContext)

	// EnterConstTrue is called when entering the ConstTrue production.
	EnterConstTrue(c *ConstTrueContext)

	// EnterSubtract is called when entering the Subtract production.
	EnterSubtract(c *SubtractContext)

	// EnterTypeCast is called when entering the TypeCast production.
	EnterTypeCast(c *TypeCastContext)

	// EnterIf is called when entering the If production.
	EnterIf(c *IfContext)

	// EnterApplication is called when entering the Application production.
	EnterApplication(c *ApplicationContext)

	// EnterDeref is called when entering the Deref production.
	EnterDeref(c *DerefContext)

	// EnterIsEmpty is called when entering the IsEmpty production.
	EnterIsEmpty(c *IsEmptyContext)

	// EnterPanic is called when entering the Panic production.
	EnterPanic(c *PanicContext)

	// EnterLessThanOrEqual is called when entering the LessThanOrEqual production.
	EnterLessThanOrEqual(c *LessThanOrEqualContext)

	// EnterSucc is called when entering the Succ production.
	EnterSucc(c *SuccContext)

	// EnterInl is called when entering the Inl production.
	EnterInl(c *InlContext)

	// EnterGreaterThanOrEqual is called when entering the GreaterThanOrEqual production.
	EnterGreaterThanOrEqual(c *GreaterThanOrEqualContext)

	// EnterInr is called when entering the Inr production.
	EnterInr(c *InrContext)

	// EnterMatch is called when entering the Match production.
	EnterMatch(c *MatchContext)

	// EnterLogicNot is called when entering the LogicNot production.
	EnterLogicNot(c *LogicNotContext)

	// EnterParenthesisedExpr is called when entering the ParenthesisedExpr production.
	EnterParenthesisedExpr(c *ParenthesisedExprContext)

	// EnterTail is called when entering the Tail production.
	EnterTail(c *TailContext)

	// EnterRecord is called when entering the Record production.
	EnterRecord(c *RecordContext)

	// EnterLogicAnd is called when entering the LogicAnd production.
	EnterLogicAnd(c *LogicAndContext)

	// EnterTypeApplication is called when entering the TypeApplication production.
	EnterTypeApplication(c *TypeApplicationContext)

	// EnterLetRec is called when entering the LetRec production.
	EnterLetRec(c *LetRecContext)

	// EnterLogicOr is called when entering the LogicOr production.
	EnterLogicOr(c *LogicOrContext)

	// EnterTryWith is called when entering the TryWith production.
	EnterTryWith(c *TryWithContext)

	// EnterPred is called when entering the Pred production.
	EnterPred(c *PredContext)

	// EnterTypeAsc is called when entering the TypeAsc production.
	EnterTypeAsc(c *TypeAscContext)

	// EnterNatRec is called when entering the NatRec production.
	EnterNatRec(c *NatRecContext)

	// EnterUnfold is called when entering the Unfold production.
	EnterUnfold(c *UnfoldContext)

	// EnterRef is called when entering the Ref production.
	EnterRef(c *RefContext)

	// EnterDotTuple is called when entering the DotTuple production.
	EnterDotTuple(c *DotTupleContext)

	// EnterFix is called when entering the Fix production.
	EnterFix(c *FixContext)

	// EnterLet is called when entering the Let production.
	EnterLet(c *LetContext)

	// EnterAssign is called when entering the Assign production.
	EnterAssign(c *AssignContext)

	// EnterTuple is called when entering the Tuple production.
	EnterTuple(c *TupleContext)

	// EnterConsList is called when entering the ConsList production.
	EnterConsList(c *ConsListContext)

	// EnterPatternBinding is called when entering the patternBinding production.
	EnterPatternBinding(c *PatternBindingContext)

	// EnterBinding is called when entering the binding production.
	EnterBinding(c *BindingContext)

	// EnterMatchCase is called when entering the matchCase production.
	EnterMatchCase(c *MatchCaseContext)

	// EnterPatternCons is called when entering the PatternCons production.
	EnterPatternCons(c *PatternConsContext)

	// EnterPatternTuple is called when entering the PatternTuple production.
	EnterPatternTuple(c *PatternTupleContext)

	// EnterPatternList is called when entering the PatternList production.
	EnterPatternList(c *PatternListContext)

	// EnterPatternRecord is called when entering the PatternRecord production.
	EnterPatternRecord(c *PatternRecordContext)

	// EnterPatternVariant is called when entering the PatternVariant production.
	EnterPatternVariant(c *PatternVariantContext)

	// EnterPatternAsc is called when entering the PatternAsc production.
	EnterPatternAsc(c *PatternAscContext)

	// EnterPatternInt is called when entering the PatternInt production.
	EnterPatternInt(c *PatternIntContext)

	// EnterPatternInr is called when entering the PatternInr production.
	EnterPatternInr(c *PatternInrContext)

	// EnterPatternTrue is called when entering the PatternTrue production.
	EnterPatternTrue(c *PatternTrueContext)

	// EnterPatternInl is called when entering the PatternInl production.
	EnterPatternInl(c *PatternInlContext)

	// EnterPatternVar is called when entering the PatternVar production.
	EnterPatternVar(c *PatternVarContext)

	// EnterParenthesisedPattern is called when entering the ParenthesisedPattern production.
	EnterParenthesisedPattern(c *ParenthesisedPatternContext)

	// EnterPatternSucc is called when entering the PatternSucc production.
	EnterPatternSucc(c *PatternSuccContext)

	// EnterPatternFalse is called when entering the PatternFalse production.
	EnterPatternFalse(c *PatternFalseContext)

	// EnterPatternUnit is called when entering the PatternUnit production.
	EnterPatternUnit(c *PatternUnitContext)

	// EnterPatternCastAs is called when entering the PatternCastAs production.
	EnterPatternCastAs(c *PatternCastAsContext)

	// EnterLabelledPattern is called when entering the labelledPattern production.
	EnterLabelledPattern(c *LabelledPatternContext)

	// EnterTypeTuple is called when entering the TypeTuple production.
	EnterTypeTuple(c *TypeTupleContext)

	// EnterTypeTop is called when entering the TypeTop production.
	EnterTypeTop(c *TypeTopContext)

	// EnterTypeBool is called when entering the TypeBool production.
	EnterTypeBool(c *TypeBoolContext)

	// EnterTypeRef is called when entering the TypeRef production.
	EnterTypeRef(c *TypeRefContext)

	// EnterTypeRec is called when entering the TypeRec production.
	EnterTypeRec(c *TypeRecContext)

	// EnterTypeAuto is called when entering the TypeAuto production.
	EnterTypeAuto(c *TypeAutoContext)

	// EnterTypeSum is called when entering the TypeSum production.
	EnterTypeSum(c *TypeSumContext)

	// EnterTypeVar is called when entering the TypeVar production.
	EnterTypeVar(c *TypeVarContext)

	// EnterTypeVariant is called when entering the TypeVariant production.
	EnterTypeVariant(c *TypeVariantContext)

	// EnterTypeUnit is called when entering the TypeUnit production.
	EnterTypeUnit(c *TypeUnitContext)

	// EnterTypeNat is called when entering the TypeNat production.
	EnterTypeNat(c *TypeNatContext)

	// EnterTypeBottom is called when entering the TypeBottom production.
	EnterTypeBottom(c *TypeBottomContext)

	// EnterTypeParens is called when entering the TypeParens production.
	EnterTypeParens(c *TypeParensContext)

	// EnterTypeFun is called when entering the TypeFun production.
	EnterTypeFun(c *TypeFunContext)

	// EnterTypeForAll is called when entering the TypeForAll production.
	EnterTypeForAll(c *TypeForAllContext)

	// EnterTypeRecord is called when entering the TypeRecord production.
	EnterTypeRecord(c *TypeRecordContext)

	// EnterTypeList is called when entering the TypeList production.
	EnterTypeList(c *TypeListContext)

	// EnterRecordFieldType is called when entering the recordFieldType production.
	EnterRecordFieldType(c *RecordFieldTypeContext)

	// EnterVariantFieldType is called when entering the variantFieldType production.
	EnterVariantFieldType(c *VariantFieldTypeContext)

	// ExitStart_Program is called when exiting the start_Program production.
	ExitStart_Program(c *Start_ProgramContext)

	// ExitStart_Expr is called when exiting the start_Expr production.
	ExitStart_Expr(c *Start_ExprContext)

	// ExitStart_Type is called when exiting the start_Type production.
	ExitStart_Type(c *Start_TypeContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitLanguageCore is called when exiting the LanguageCore production.
	ExitLanguageCore(c *LanguageCoreContext)

	// ExitAnExtension is called when exiting the AnExtension production.
	ExitAnExtension(c *AnExtensionContext)

	// ExitDeclFun is called when exiting the DeclFun production.
	ExitDeclFun(c *DeclFunContext)

	// ExitDeclFunGeneric is called when exiting the DeclFunGeneric production.
	ExitDeclFunGeneric(c *DeclFunGenericContext)

	// ExitDeclTypeAlias is called when exiting the DeclTypeAlias production.
	ExitDeclTypeAlias(c *DeclTypeAliasContext)

	// ExitDeclExceptionType is called when exiting the DeclExceptionType production.
	ExitDeclExceptionType(c *DeclExceptionTypeContext)

	// ExitDeclExceptionVariant is called when exiting the DeclExceptionVariant production.
	ExitDeclExceptionVariant(c *DeclExceptionVariantContext)

	// ExitInlineAnnotation is called when exiting the InlineAnnotation production.
	ExitInlineAnnotation(c *InlineAnnotationContext)

	// ExitParamDecl is called when exiting the paramDecl production.
	ExitParamDecl(c *ParamDeclContext)

	// ExitFold is called when exiting the Fold production.
	ExitFold(c *FoldContext)

	// ExitAdd is called when exiting the Add production.
	ExitAdd(c *AddContext)

	// ExitIsZero is called when exiting the IsZero production.
	ExitIsZero(c *IsZeroContext)

	// ExitVar is called when exiting the Var production.
	ExitVar(c *VarContext)

	// ExitTypeAbstraction is called when exiting the TypeAbstraction production.
	ExitTypeAbstraction(c *TypeAbstractionContext)

	// ExitDivide is called when exiting the Divide production.
	ExitDivide(c *DivideContext)

	// ExitLessThan is called when exiting the LessThan production.
	ExitLessThan(c *LessThanContext)

	// ExitDotRecord is called when exiting the DotRecord production.
	ExitDotRecord(c *DotRecordContext)

	// ExitGreaterThan is called when exiting the GreaterThan production.
	ExitGreaterThan(c *GreaterThanContext)

	// ExitEqual is called when exiting the Equal production.
	ExitEqual(c *EqualContext)

	// ExitThrow is called when exiting the Throw production.
	ExitThrow(c *ThrowContext)

	// ExitMultiply is called when exiting the Multiply production.
	ExitMultiply(c *MultiplyContext)

	// ExitConstMemory is called when exiting the ConstMemory production.
	ExitConstMemory(c *ConstMemoryContext)

	// ExitList is called when exiting the List production.
	ExitList(c *ListContext)

	// ExitTryCatch is called when exiting the TryCatch production.
	ExitTryCatch(c *TryCatchContext)

	// ExitTryCastAs is called when exiting the TryCastAs production.
	ExitTryCastAs(c *TryCastAsContext)

	// ExitHead is called when exiting the Head production.
	ExitHead(c *HeadContext)

	// ExitTerminatingSemicolon is called when exiting the TerminatingSemicolon production.
	ExitTerminatingSemicolon(c *TerminatingSemicolonContext)

	// ExitNotEqual is called when exiting the NotEqual production.
	ExitNotEqual(c *NotEqualContext)

	// ExitConstUnit is called when exiting the ConstUnit production.
	ExitConstUnit(c *ConstUnitContext)

	// ExitSequence is called when exiting the Sequence production.
	ExitSequence(c *SequenceContext)

	// ExitConstFalse is called when exiting the ConstFalse production.
	ExitConstFalse(c *ConstFalseContext)

	// ExitAbstraction is called when exiting the Abstraction production.
	ExitAbstraction(c *AbstractionContext)

	// ExitConstInt is called when exiting the ConstInt production.
	ExitConstInt(c *ConstIntContext)

	// ExitVariant is called when exiting the Variant production.
	ExitVariant(c *VariantContext)

	// ExitConstTrue is called when exiting the ConstTrue production.
	ExitConstTrue(c *ConstTrueContext)

	// ExitSubtract is called when exiting the Subtract production.
	ExitSubtract(c *SubtractContext)

	// ExitTypeCast is called when exiting the TypeCast production.
	ExitTypeCast(c *TypeCastContext)

	// ExitIf is called when exiting the If production.
	ExitIf(c *IfContext)

	// ExitApplication is called when exiting the Application production.
	ExitApplication(c *ApplicationContext)

	// ExitDeref is called when exiting the Deref production.
	ExitDeref(c *DerefContext)

	// ExitIsEmpty is called when exiting the IsEmpty production.
	ExitIsEmpty(c *IsEmptyContext)

	// ExitPanic is called when exiting the Panic production.
	ExitPanic(c *PanicContext)

	// ExitLessThanOrEqual is called when exiting the LessThanOrEqual production.
	ExitLessThanOrEqual(c *LessThanOrEqualContext)

	// ExitSucc is called when exiting the Succ production.
	ExitSucc(c *SuccContext)

	// ExitInl is called when exiting the Inl production.
	ExitInl(c *InlContext)

	// ExitGreaterThanOrEqual is called when exiting the GreaterThanOrEqual production.
	ExitGreaterThanOrEqual(c *GreaterThanOrEqualContext)

	// ExitInr is called when exiting the Inr production.
	ExitInr(c *InrContext)

	// ExitMatch is called when exiting the Match production.
	ExitMatch(c *MatchContext)

	// ExitLogicNot is called when exiting the LogicNot production.
	ExitLogicNot(c *LogicNotContext)

	// ExitParenthesisedExpr is called when exiting the ParenthesisedExpr production.
	ExitParenthesisedExpr(c *ParenthesisedExprContext)

	// ExitTail is called when exiting the Tail production.
	ExitTail(c *TailContext)

	// ExitRecord is called when exiting the Record production.
	ExitRecord(c *RecordContext)

	// ExitLogicAnd is called when exiting the LogicAnd production.
	ExitLogicAnd(c *LogicAndContext)

	// ExitTypeApplication is called when exiting the TypeApplication production.
	ExitTypeApplication(c *TypeApplicationContext)

	// ExitLetRec is called when exiting the LetRec production.
	ExitLetRec(c *LetRecContext)

	// ExitLogicOr is called when exiting the LogicOr production.
	ExitLogicOr(c *LogicOrContext)

	// ExitTryWith is called when exiting the TryWith production.
	ExitTryWith(c *TryWithContext)

	// ExitPred is called when exiting the Pred production.
	ExitPred(c *PredContext)

	// ExitTypeAsc is called when exiting the TypeAsc production.
	ExitTypeAsc(c *TypeAscContext)

	// ExitNatRec is called when exiting the NatRec production.
	ExitNatRec(c *NatRecContext)

	// ExitUnfold is called when exiting the Unfold production.
	ExitUnfold(c *UnfoldContext)

	// ExitRef is called when exiting the Ref production.
	ExitRef(c *RefContext)

	// ExitDotTuple is called when exiting the DotTuple production.
	ExitDotTuple(c *DotTupleContext)

	// ExitFix is called when exiting the Fix production.
	ExitFix(c *FixContext)

	// ExitLet is called when exiting the Let production.
	ExitLet(c *LetContext)

	// ExitAssign is called when exiting the Assign production.
	ExitAssign(c *AssignContext)

	// ExitTuple is called when exiting the Tuple production.
	ExitTuple(c *TupleContext)

	// ExitConsList is called when exiting the ConsList production.
	ExitConsList(c *ConsListContext)

	// ExitPatternBinding is called when exiting the patternBinding production.
	ExitPatternBinding(c *PatternBindingContext)

	// ExitBinding is called when exiting the binding production.
	ExitBinding(c *BindingContext)

	// ExitMatchCase is called when exiting the matchCase production.
	ExitMatchCase(c *MatchCaseContext)

	// ExitPatternCons is called when exiting the PatternCons production.
	ExitPatternCons(c *PatternConsContext)

	// ExitPatternTuple is called when exiting the PatternTuple production.
	ExitPatternTuple(c *PatternTupleContext)

	// ExitPatternList is called when exiting the PatternList production.
	ExitPatternList(c *PatternListContext)

	// ExitPatternRecord is called when exiting the PatternRecord production.
	ExitPatternRecord(c *PatternRecordContext)

	// ExitPatternVariant is called when exiting the PatternVariant production.
	ExitPatternVariant(c *PatternVariantContext)

	// ExitPatternAsc is called when exiting the PatternAsc production.
	ExitPatternAsc(c *PatternAscContext)

	// ExitPatternInt is called when exiting the PatternInt production.
	ExitPatternInt(c *PatternIntContext)

	// ExitPatternInr is called when exiting the PatternInr production.
	ExitPatternInr(c *PatternInrContext)

	// ExitPatternTrue is called when exiting the PatternTrue production.
	ExitPatternTrue(c *PatternTrueContext)

	// ExitPatternInl is called when exiting the PatternInl production.
	ExitPatternInl(c *PatternInlContext)

	// ExitPatternVar is called when exiting the PatternVar production.
	ExitPatternVar(c *PatternVarContext)

	// ExitParenthesisedPattern is called when exiting the ParenthesisedPattern production.
	ExitParenthesisedPattern(c *ParenthesisedPatternContext)

	// ExitPatternSucc is called when exiting the PatternSucc production.
	ExitPatternSucc(c *PatternSuccContext)

	// ExitPatternFalse is called when exiting the PatternFalse production.
	ExitPatternFalse(c *PatternFalseContext)

	// ExitPatternUnit is called when exiting the PatternUnit production.
	ExitPatternUnit(c *PatternUnitContext)

	// ExitPatternCastAs is called when exiting the PatternCastAs production.
	ExitPatternCastAs(c *PatternCastAsContext)

	// ExitLabelledPattern is called when exiting the labelledPattern production.
	ExitLabelledPattern(c *LabelledPatternContext)

	// ExitTypeTuple is called when exiting the TypeTuple production.
	ExitTypeTuple(c *TypeTupleContext)

	// ExitTypeTop is called when exiting the TypeTop production.
	ExitTypeTop(c *TypeTopContext)

	// ExitTypeBool is called when exiting the TypeBool production.
	ExitTypeBool(c *TypeBoolContext)

	// ExitTypeRef is called when exiting the TypeRef production.
	ExitTypeRef(c *TypeRefContext)

	// ExitTypeRec is called when exiting the TypeRec production.
	ExitTypeRec(c *TypeRecContext)

	// ExitTypeAuto is called when exiting the TypeAuto production.
	ExitTypeAuto(c *TypeAutoContext)

	// ExitTypeSum is called when exiting the TypeSum production.
	ExitTypeSum(c *TypeSumContext)

	// ExitTypeVar is called when exiting the TypeVar production.
	ExitTypeVar(c *TypeVarContext)

	// ExitTypeVariant is called when exiting the TypeVariant production.
	ExitTypeVariant(c *TypeVariantContext)

	// ExitTypeUnit is called when exiting the TypeUnit production.
	ExitTypeUnit(c *TypeUnitContext)

	// ExitTypeNat is called when exiting the TypeNat production.
	ExitTypeNat(c *TypeNatContext)

	// ExitTypeBottom is called when exiting the TypeBottom production.
	ExitTypeBottom(c *TypeBottomContext)

	// ExitTypeParens is called when exiting the TypeParens production.
	ExitTypeParens(c *TypeParensContext)

	// ExitTypeFun is called when exiting the TypeFun production.
	ExitTypeFun(c *TypeFunContext)

	// ExitTypeForAll is called when exiting the TypeForAll production.
	ExitTypeForAll(c *TypeForAllContext)

	// ExitTypeRecord is called when exiting the TypeRecord production.
	ExitTypeRecord(c *TypeRecordContext)

	// ExitTypeList is called when exiting the TypeList production.
	ExitTypeList(c *TypeListContext)

	// ExitRecordFieldType is called when exiting the recordFieldType production.
	ExitRecordFieldType(c *RecordFieldTypeContext)

	// ExitVariantFieldType is called when exiting the variantFieldType production.
	ExitVariantFieldType(c *VariantFieldTypeContext)
}
