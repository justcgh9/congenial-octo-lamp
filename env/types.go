package env

type Type interface {
	Type() string
}

type Nat struct {}

func (n Nat) Type() string {
	return "Nat"
}

type Bool struct {}

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
