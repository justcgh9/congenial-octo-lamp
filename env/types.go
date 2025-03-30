package env

import (
	"fmt"
	"sort"
	"stella-implementation-in-go/parser"
)

type Type interface {
	Type() string
}

type Nat struct {}

func (n Nat) Type() string {
	return "Nat"
}

type Bool struct {
	Val string
}

func (b Bool) Type() string {
	return "Bool"
}

type Func struct {
	Args []Type
	Return Type
	IsAnonymous bool
}

func (f Func) Type() string {
	ans := "fn("
	for i, arg := range f.Args {
		ans += arg.Type()
		if i < len(f.Args) - 1 {
			ans += ","
		}
	} 
	ans += ")" + "->" + f.Return.Type()

	return ans
}

type Unit struct {}

func (u Unit) Type() string {
	return "Unit"
}

type Tuple struct {
	Elements []Type
	IsLiteral bool
}

func (t Tuple) Type() string {
	ans := "("

	for i := range t.Elements {
		ans += t.Elements[i].Type()
		if i < len(t.Elements) - 1 {
			ans += ","
		}
	}

	ans += ")"
	return ans
}

type Binding struct {
	Name string
	T Type
}

func (b Binding) Type() string {
	return b.Name + ":" + b.T.Type()
}

type Erroneos struct {}

func (e Erroneos) Type() string {
	return ""
}

type Any struct {}

func (a Any) Type() string {
	return "any"
}

type Sum struct {
	Left  Type
	Right Type
}

func (s Sum) Type() string {
	return s.Left.Type() + "+" + s.Right.Type()
}

type List struct {
	T Type
	IsLiteral bool
	// Elements []Type
}

func (l List) Type() string {
	if l.T == nil {
		return "[]"
	}
	return "[" + l.T.Type() + "]"
}

type Record struct {
	Elements map[string]Type
	IsLiteral bool
}

func (r Record) Type () string {
	ans := "{"

	keys := make([]string, 0, len(r.Elements))
	for key := range r.Elements {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, key := range keys {
		ans += key + ":" + r.Elements[key].Type() + ","
	}

	if ans[len(ans) - 1] == ',' {
		ans = ans[:(len(ans) - 1)]
	}

	ans += "}"
	return ans
}

type Reference struct {
	UnderlyingType Type
	IsLiteral bool
}

func (r Reference) Type() string {
	return "&" + r.UnderlyingType.Type()
}

type Throw struct {
	UnderlyingType Type
	// IsLiteral bool
}

func (t Throw) Type() string {
	return "throw(" + t.UnderlyingType.Type() + ")"
}

type Variant struct {
	Elements map[string]Type
	IsLiteral bool
}

func (r Variant) Type () string {
	ans := "<|"

	keys := make([]string, 0, len(r.Elements))
	for key := range r.Elements {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, key := range keys {
		if r.Elements[key] == nil {
			ans += key + ","
			continue
		}
		ans += key + ":" + r.Elements[key].Type() + ","
	}

	if ans[len(ans) - 1] == ',' {
		ans = ans[:(len(ans) - 1)]
	}

	ans += "|>"
	return ans
}

type Top struct {}

func (t Top) Type () string {
	return "Top"
}

type Bottom struct {}

func (t Bottom) Type () string {
	return "Bottom"
}

type GenericAbstraction struct {
	Generics map[string]Type
	T 		 Type
	Bindings []Binding
	Ctx 	 parser.IExprContext
}

func (g GenericAbstraction) Type () string {
	
	ans := "["

	keys := make([]string, 0, len(g.Generics))
	for key := range g.Generics {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, key := range keys {
		ans += key + ","
	}

	if ans[len(ans) - 1] == ',' {
		ans = ans[:(len(ans) - 1)]
	}

	ans += "] " + g.T.Type()


	return ans
}

type Var struct {
	ID 	int
}

func (v Var) Type() string {
	return fmt.Sprintf("var %d", v.ID)
}