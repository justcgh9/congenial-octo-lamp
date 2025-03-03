package visitor

import (
	"fmt"
	"os"
	"stella-implementation-in-go/env"
)

func isAny(arg interface{}) bool {
	_, ok := arg.(env.Any)
	return ok
}

func (v *Visitor) err(str string) {

	if v.subtyping != 0 && str == "ERROR_UNEXPECTED_TYPE_FOR_EXPRESSION" {
		str = "ERROR_UNEXPECTED_SUBTYPE"
	}

	fmt.Println(str)
	os.Exit(1)
}

func addIfExists(mp map[string]bool, key string) {
	if _, ok := mp[key]; ok {
		mp[key] = true
	}
}

func removeIfExists(mp map[string]bool, key string) {
	if _, ok := mp[key]; ok {
		mp[key] = false
	}
}

func eraseLiterals(t interface{} ) env.Type {
	switch t := t.(type) {
	default:
		return t.(env.Type)
	case env.Func:
		t.IsAnonymous = false
		return t
	case env.List:
		t.IsLiteral = false
		return t
	case env.Tuple:
		t.IsLiteral = false
		return t
	case env.Record:
		t.IsLiteral = false
		return t
	case env.Reference:
		t.IsLiteral = false
		return t
	}
}


func contains(str string, args ...string) bool {
	if len(args) <= 0 {
		return false
	}

	for _, arg := range args {
		if arg == str {
			return true
		}
	}

	return false
}

func startsWith(s, t string) bool {
	if len(s) > len(t) {
		return false
	}

	return s == t[:len(s)]
}

func ascriptSumTypes(f, s env.Sum) env.Sum {
	
	fl, fok := f.Left.(env.Sum)
	sl, sok := s.Left.(env.Sum) 
	
	if fok && sok { 
		fl.Left = ascriptSumTypes(fl, sl)
	}

	fr, fok := f.Right.(env.Sum)
	sr, sok := s.Right.(env.Sum)

	if fok && sok { 
		fl.Right = ascriptSumTypes(fr, sr)
	}

	if f.Left.Type() == (env.Any{}).Type() {
		f.Left = s.Left
	}

	if f.Right.Type() == (env.Any{}).Type() {
		f.Right = s.Right
	}
	
	return f
}

func isAmbiguousType(t interface{}) (interface{}, bool) {
	switch t := t.(type) {
	default:
		return nil, false
	case env.Any, nil:
		return nil, true
	case env.Func:
		return isAmbiguousType(t.Return)
	case env.List:
		if t.T != nil { 
			return isAmbiguousType(t.T)
		}
		return env.List{}, true
	case env.Tuple:
		for _, v := range t.Elements {
			if val, ok := isAmbiguousType(v); ok {
				return val, true
			}
		}
		return nil, false
	case env.Record:
		for _, v := range t.Elements {
			if val, ok := isAmbiguousType(v); ok {
				return val, true
			}
		}
		return nil, false
	case env.Reference:
		val, ok := isAmbiguousType(t.UnderlyingType)
		if !ok {
			return nil, false
		}

		if val == nil { val = env.Reference{}}

		return val, true
	case *env.Throw:
		val, ok := isAmbiguousType(t.UnderlyingType)
		if !ok {
			return nil, false
		}

		if val == nil { val = env.Throw{}}

		return val, true
	case env.Sum:
		if isAny(t.Left) || isAny(t.Right) {
			return env.Sum{}, true
		}
		val, ok := isAmbiguousType(t.Left)
		if !ok {
			val, ok := isAmbiguousType(t.Right)
			if !ok {
				return nil, false
			}

			if val == nil { val = env.Sum{}}

			return val, true
		}

		if val == nil { val = env.Sum{}}

		return val, true
	}
}

func throwAmbiguousType(t interface{}) {
	if val, ok := isAmbiguousType(t); ok { // list (inl (0)) - Sum Type, true
		switch val.(type) {
		case env.Sum:
			fmt.Println("ERROR_AMBIGUOUS_SUM_TYPE")
			os.Exit(1)
		case env.List:
			fmt.Println("ERROR_AMBIGUOUS_LIST_TYPE")
			os.Exit(1)
		case env.Reference:
			fmt.Println("ERROR_AMBIGUOUS_REFERENCE_TYPE")
			os.Exit(1)
		case env.Throw:
			fmt.Println("ERROR_AMBIGUOUS_THROW_TYPE")
			os.Exit(1)
		}
	}
}

func (v Visitor) isSubtype(left, right interface{}) bool {

	if _, ok := right.(env.Top); ok {return true}

	switch left := left.(type) {
	default:
		return false
	case *env.Bottom:
		return true
	case env.Nat:
		_, ok := right.(env.Nat)
		return ok
	case env.Bool:
		_, ok := right.(env.Bool)
		return ok
	case env.Unit:
		_, ok := right.(env.Unit)
		return ok
	case env.Func:
		right, ok := right.(env.Func)
		if !ok { return false }
		
		if len(left.Args) != len(right.Args) { return false }
		
		for i := range left.Args {
			if ! v.isSubtype(right.Args[i], left.Args[i]) {return false} 
		}

		return v.isSubtype(left.Return, right.Return)
	case env.Tuple:
		right, ok := right.(env.Tuple)
		if !ok { return false }

		if len(left.Elements) != len(right.Elements) { return false }
		
		for i := range left.Elements {
			if ! v.isSubtype(left.Elements[i], right.Elements[i]) {return false} 
		}

		return true

	case env.Record:
		right, ok := right.(env.Record)
		if !ok { return false }

		if len(left.Elements) < len(right.Elements) { v.err("ERROR_MISSING_RECORD_FIELDS") }
		
		for i := range right.Elements {

			_, ok := left.Elements[i]	

			if !ok {
				v.err("ERROR_MISSING_RECORD_FIELDS")
			}

			if !ok || !v.isSubtype(left.Elements[i], right.Elements[i]) {return false} 
		}

		return true
	case env.Sum:
		right, ok := right.(env.Sum)
		if !ok {
			return false
		}

		return v.isSubtype(left.Left, right.Left) && v.isSubtype(left.Right, right.Right)
	case env.List:
		right, ok := right.(env.List)
		if !ok { return false }

		return v.isSubtype(left.T, right.T)

	case env.Reference:
		right, ok := right.(env.Reference)

		return ok && v.isSubtype(left.UnderlyingType, right.UnderlyingType) && v.isSubtype(right.UnderlyingType, left.UnderlyingType)

	case env.Variant:
		right, ok := right.(env.Variant)
		if !ok { return false }

		if len(left.Elements) > len(right.Elements) { v.err("ERROR_UNEXPECTED_VARIANT_LABEL") }

		for k := range left.Elements {
			_, ok := right.Elements[k]

			if !ok {v.err("ERROR_UNEXPECTED_VARIANT_LABEL")}

			if !ok || !v.isSubtype(left.Elements[k], right.Elements[k]) { return false }
		}

		return true
	}
}

// func isAmbiguousSumType(t interface{}) bool {
// 	switch t := t.(type) {
// 	default:
// 		return false
// 	case env.Sum:
// 		if isAny(t.Left) || isAny(t.Right) {
// 			return true
// 		}
// 		return isAmbiguousSumType(t.Left) || isAmbiguousSumType(t.Right)
// 	}
// }

// func isAmbiguousListType(t interface{}) bool {
// 	switch t := t.(type) {
// 	default:
// 		return false
// 	case env.List:
// 		if t.T == nil || isAny(t.T) {
// 			return true
// 		}
// 		return isAmbiguousListType(t.T)
// 	}
// }

// func isAmbiguousReferenceType(t interface{}) bool {
// 	switch t := t.(type) {
// 	default:
// 		return false
// 	case env.Reference:
// 		if t.UnderlyingType == nil || isAny(t.UnderlyingType) {
// 			return true
// 		}
// 		return isAmbiguousReferenceType(t.UnderlyingType)
// 	}
// }



// func getAllVariations(t interface{}, prefix string) []string {
// 	switch t := t.(type) {
// 	default:
// 	case env.Nat:
// 		return []string{"nat", "nat 0"}
// 	case env.Bool:
// 		return []string{"true", "false"}
// 	case env.Tuple:
// 		combs := make([][]string, len(t.Elements))
// 		for _, val := range t.Elements {
// 			combs[i] = getAllVariations()
// 		}
// 	}
// 	return []string{}
// }