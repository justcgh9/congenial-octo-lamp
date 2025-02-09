package visitor

import "stella-implementation-in-go/env"

func isAny(arg interface{}) bool {
	_, ok := arg.(env.Any)
	return ok
}