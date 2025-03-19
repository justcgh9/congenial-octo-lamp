package checker

import "stella-implementation-in-go/env"

type Env struct {
	elements env.Stack[map[string]Type]
}

func (e *Env) Push() {
	e.elements.Push(make(map[string]Type, 64))
}

func (e *Env) Put(n string, t Type) {
	mp, _ := e.elements.Peek()
	mp[n] = t
}

func (e *Env) Check(n string) Type {
	if e.elements.IsEmpty() {return Type{
		Kind: "",
	}}

	mp, _ := e.elements.Pop()

	val, ok := mp[n]
	if !ok {
		val = e.Check(n)
	}

	e.elements.Push(mp)
	return val
}

func (e *Env) Pop() {
	_, _ = e.elements.Pop()
}