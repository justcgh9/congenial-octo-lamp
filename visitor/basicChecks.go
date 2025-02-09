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