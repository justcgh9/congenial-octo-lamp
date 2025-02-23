package env

import "fmt"

type Env struct {
	elements Stack[map[string]Type]
}

func (e *Env) Push() {
	e.elements.Push(make(map[string]Type, 64))
}

func (e *Env) Put(n string, t Type) {
	mp, _ := e.elements.Peek()
	mp[n] = t
}

func (e *Env) Check(n string) Type {
	if e.elements.IsEmpty() {return Erroneos{}}

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

func (e *Env) Print () {
	// e.elements.Pop()
	mp, _ := e.elements.Peek()
	for k, v := range mp {
		fmt.Println(k ,v.Type())
	}
}