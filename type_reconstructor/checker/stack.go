package checker

import "stella-implementation-in-go/env"

type Scope struct {
    Vars map[string]int
}

type Stack = env.Stack[Scope]