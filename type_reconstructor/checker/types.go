package checker

type Type struct {
	Kind 	string
	Data 	any
}

type TypeVar struct {
    ID int
}

type Substitution map[int]Type

type Constraint struct {
	Left 	Type
	Right	Type
}

func NewConstraint(left, right Type) Constraint {
	return Constraint{
		Left: left,
		Right: right,
	}
}

const (
	Int 			= "int"
	Bool 			= "bool"
	Unit 			= "unit"
	Func 			= "func"
	List 			= "list"
	Reference 		= "ref"
	Tuple 			= "tuple"
	Record 			= "record"
	Variant 		= "variant"
	Sum 			= "sum"
	UnificationVar 	= "var"
)

func NewInt() Type {
	return Type{Kind: Int}
}


func NewUnit() Type {
	return Type{Kind: Unit}
}


func NewBool() Type {
	return Type{Kind: Bool}
}


type FuncType struct {
	Param 	Type
	Return 	Type
}

func NewFunc(param, result Type) Type {
	return Type{
		Kind: Func,
		Data: FuncType{
			Param: param,
			Return: result,
		},
	}
}

type ListType struct {
	Content	Type
}

func NewList(content Type) Type {
	return Type{
		Kind: List,
		Data: ListType{
			Content: content,
		},
	}
}

type ReferenceType struct {
	Content	Type
}

func NewReference(content Type) Type {
	return Type{
		Kind: Reference,
		Data: ReferenceType{
			Content: content,
		},
	}
}

type TupleType struct {
	Content []Type
}

func NewTuple(content []Type) Type {
	return Type{
		Kind: Tuple,
		Data: TupleType{
			Content: content,
		},
	}
}

type RecordType struct {
	Content map[string]Type
}

func NewRecord(content map[string]Type) Type {
	return Type{
		Kind: Record,
		Data: RecordType{
			Content: content,
		},
	}
}

type SumType struct {
	Left	Type
	Right	Type
}

func NewSum(left, right Type) Type {
	return Type{
		Kind: Sum,
		Data: SumType{
			Left: left,
			Right: right,
		},
	}
}

type VariantType struct {
	Content map[string]Type
}

func NewVariant(content map[string]Type) Type {
	return Type{
		Kind: Variant,
		Data: VariantType{
			Content: content,
		},
	}
}