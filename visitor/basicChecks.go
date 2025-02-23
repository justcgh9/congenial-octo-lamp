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

func err(str string) {
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