package visitor

import "stella-implementation-in-go/env"

func isAny(arg interface{}) bool {
	_, ok := arg.(env.Any)
	return ok
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

func isAmbiguousSumType(t interface{}) bool {
	switch t := t.(type) {
	default:
		return false
	case env.Sum:
		if isAny(t.Left) || isAny(t.Right) {
			return true
		}
		return isAmbiguousSumType(t.Left) || isAmbiguousSumType(t.Right)
	}
}

func isAmbiguousListType(t interface{}) bool {
	switch t := t.(type) {
	default:
		return false
	case env.List:
		if t.T == nil || isAny(t.T) {
			return true
		}
		return isAmbiguousListType(t.T)
	}
}

func isAmbiguousReferenceType(t interface{}) bool {
	switch t := t.(type) {
	default:
		return false
	case env.Reference:
		if t.UnderlyingType == nil || isAny(t.UnderlyingType) {
			return true
		}
		return isAmbiguousReferenceType(t.UnderlyingType)
	}
}


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