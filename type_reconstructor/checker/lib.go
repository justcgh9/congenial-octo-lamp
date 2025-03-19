package checker

import "fmt"

func AddConstraint(constraint Constraint, constraints *[]Constraint) {
	*constraints = append(*constraints, constraint)
}

func applySubst(t Type, subst Substitution) Type {
    if t.Kind == UnificationVar {
        if replacement, found := subst[t.Data.(TypeVar).ID]; found {
            return applySubst(replacement, subst)
        }
        return t
    }

    switch t.Kind {
    case Func:
        f := t.Data.(FuncType)
        return NewFunc(applySubst(f.Param, subst), applySubst(f.Return, subst))
    case List:
        l := t.Data.(ListType)
        return NewList(applySubst(l.Content, subst))
    case Reference:
        r := t.Data.(ReferenceType)
        return NewReference(applySubst(r.Content, subst))
    case Tuple:
        tuple := t.Data.(TupleType)
        newContent := make([]Type, len(tuple.Content))
        for i, v := range tuple.Content {
            newContent[i] = applySubst(v, subst)
        }
        return NewTuple(newContent)
    case Record:
        record := t.Data.(RecordType)
        newContent := make(map[string]Type, len(record.Content))
        for i, v := range record.Content {
            newContent[i] = applySubst(v, subst)
        }
        return NewRecord(newContent)
    case Sum:
        s := t.Data.(SumType)
        return NewSum(applySubst(s.Left, subst), applySubst(s.Right, subst))
    case Variant:
        v := t.Data.(VariantType)
        newContent := make(map[string]Type)
        for k, v := range v.Content {
            newContent[k] = applySubst(v, subst)
        }
        return NewVariant(newContent)
    default:
        return t
    }
}


func occursCheck(v TypeVar, t Type, subst Substitution) bool {
    t = applySubst(t, subst)

    if t.Kind == UnificationVar {
        return t.Data.(TypeVar).ID == v.ID
    }

    switch t.Kind {
    case Func:
        f := t.Data.(FuncType)
        return occursCheck(v, f.Param, subst) || occursCheck(v, f.Return, subst)
    case List:
        return occursCheck(v, t.Data.(ListType).Content, subst)
    case Reference:
        return occursCheck(v, t.Data.(ReferenceType).Content, subst)
    case Tuple:
        for _, el := range t.Data.(TupleType).Content {
            if occursCheck(v, el, subst) {
                return true
            }
        }
    case Record:
        for _, el := range t.Data.(RecordType).Content {
            if occursCheck(v, el, subst) {
                return true
            }
        }
    case Sum:
        s := t.Data.(SumType)
        return occursCheck(v, s.Left, subst) || occursCheck(v, s.Right, subst)
    case Variant:
        for _, el := range t.Data.(VariantType).Content {
            if occursCheck(v, el, subst) {
                return true
            }
        }
    }
    return false
}

func unify(t1, t2 Type, subst Substitution) error {
    t1 = applySubst(t1, subst)
    t2 = applySubst(t2, subst)

    if equals(t1, t2) {
        return nil
    }

    if t1.Kind == UnificationVar {
        return unifyVar(t1.Data.(TypeVar), t2, subst)
    }
    if t2.Kind == UnificationVar {
        return unifyVar(t2.Data.(TypeVar), t1, subst)
    }

    if t1.Kind != t2.Kind {
        return fmt.Errorf("type mismatch: %v vs %v", t1.Kind, t2.Kind)
    }

    switch t1.Kind {
    case Func:
        f1, f2 := t1.Data.(FuncType), t2.Data.(FuncType)
        if err := unify(f1.Param, f2.Param, subst); err != nil {
            return err
        }
        return unify(f1.Return, f2.Return, subst)

    case List:
        return unify(t1.Data.(ListType).Content, t2.Data.(ListType).Content, subst)

    case Reference:
        return unify(t1.Data.(ReferenceType).Content, t2.Data.(ReferenceType).Content, subst)

    case Tuple:
        tup1, tup2 := t1.Data.(TupleType), t2.Data.(TupleType)
        if len(tup1.Content) != len(tup2.Content) {
            return fmt.Errorf("tuple length mismatch: %v vs %v", len(tup1.Content), len(tup2.Content))
        }
        for i := range tup1.Content {
            if err := unify(tup1.Content[i], tup2.Content[i], subst); err != nil {
                return err
            }
        }

    case Record:
        rec1, rec2 := t1.Data.(RecordType), t2.Data.(RecordType)
        if len(rec1.Content) != len(rec2.Content) {
            return fmt.Errorf("record length mismatch")
        }
        for i := range rec1.Content {
            if err := unify(rec1.Content[i], rec2.Content[i], subst); err != nil {
                return err
            }
        }

    case Sum:
        sum1, sum2 := t1.Data.(SumType), t2.Data.(SumType)
        if err := unify(sum1.Left, sum2.Left, subst); err != nil {
            return err
        }
        return unify(sum1.Right, sum2.Right, subst)

    case Variant:
        var1, var2 := t1.Data.(VariantType), t2.Data.(VariantType)
        if len(var1.Content) != len(var2.Content) {
            return fmt.Errorf("variant length mismatch")
        }
        for key, t1 := range var1.Content {
            t2, ok := var2.Content[key]
            if !ok {
                return fmt.Errorf("variant key mismatch: %v", key)
            }
            if err := unify(t1, t2, subst); err != nil {
                return err
            }
        }
    default:
        return fmt.Errorf("unknown type: %v", t1.Kind)
    }
    return nil
}


func unifyVar(v TypeVar, t Type, subst Substitution) error {
    if existing, found := subst[v.ID]; found {
        return unify(existing, t, subst)
    }
    if occursCheck(v, t, subst) {
        // InfiniteType()
        return fmt.Errorf("recursive type: %v in %v", v, t)
    }
    subst[v.ID] = t
    return nil
}

func unifyConstraints(constraints []Constraint) (Substitution, error) {
    subst := make(Substitution)
    worklist := constraints

    for len(worklist) > 0 {
        c := worklist[len(worklist)-1]
        worklist = worklist[:len(worklist)-1]

        // Check if the same variable appears in both sides before unifying
        if _, ok := commonVar(c.Left, c.Right); ok && !equals(c.Left, c.Right) {
            InfiniteType()
            return nil, fmt.Errorf("infinite type detected")
        }

        if err := unify(c.Left, c.Right, subst); err != nil {
            return nil, err
        }
    }

    return subst, nil
}

func commonVar(t1, t2 Type) (TypeVar, bool) {
    vars1 := collectVars(t1)
    vars2 := collectVars(t2)

    for _, v := range vars1 {
        if _, found := vars2[v.ID]; found {
            return v, true
        }
    }
    return TypeVar{}, false
}

func collectVars(t Type) map[int]TypeVar {
    vars := make(map[int]TypeVar)

    var collect func(Type)
    collect = func(t Type) {
        t = applySubst(t, make(Substitution)) // Normalize type

        if t.Kind == UnificationVar {
            v := t.Data.(TypeVar)
            vars[v.ID] = v
            return
        }

        switch t.Kind {
        case Func:
            f := t.Data.(FuncType)
            collect(f.Param)
            collect(f.Return)
        case List:
            collect(t.Data.(ListType).Content)
        case Reference:
            collect(t.Data.(ReferenceType).Content)
        case Tuple:
            for _, el := range t.Data.(TupleType).Content {
                collect(el)
            }
        case Record:
            for _, el := range t.Data.(RecordType).Content {
                collect(el)
            }
        case Sum:
            s := t.Data.(SumType)
            collect(s.Left)
            collect(s.Right)
        case Variant:
            for _, el := range t.Data.(VariantType).Content {
                collect(el)
            }
        }
    }

    collect(t)
    return vars
}


func equals(t1, t2 Type) bool {
    if t1.Kind != t2.Kind {
        return false
    }

    switch t1.Kind {
    case Int, Bool, Unit, UnificationVar:
        return true // Simple types have no internal data to compare.

    case Func:
        f1, ok1 := t1.Data.(FuncType)
        f2, ok2 := t2.Data.(FuncType)
        return ok1 && ok2 && equals(f1.Param, f2.Param) && equals(f1.Return, f2.Return)

    case List:
        l1, ok1 := t1.Data.(ListType)
        l2, ok2 := t2.Data.(ListType)
        return ok1 && ok2 && equals(l1.Content, l2.Content)

    case Reference:
        r1, ok1 := t1.Data.(ReferenceType)
        r2, ok2 := t2.Data.(ReferenceType)
        return ok1 && ok2 && equals(r1.Content, r2.Content)

    case Tuple:
        t1Data, ok1 := t1.Data.(TupleType)
        t2Data, ok2 := t2.Data.(TupleType)
        if !ok1 || !ok2 || len(t1Data.Content) != len(t2Data.Content) {
            return false
        }
        for i := range t1Data.Content {
            if !equals(t1Data.Content[i], t2Data.Content[i]) {
                return false
            }
        }
        return true

    case Record:
        r1, ok1 := t1.Data.(RecordType)
        r2, ok2 := t2.Data.(RecordType)
        if !ok1 || !ok2 || len(r1.Content) != len(r2.Content) {
            return false
        }
        for key, v1 := range r1.Content {
            v2, exists := r2.Content[key]
            if !exists || !equals(v1, v2) {
                return false
            }
        }
        return true

    case Variant:
        v1, ok1 := t1.Data.(VariantType)
        v2, ok2 := t2.Data.(VariantType)
        if !ok1 || !ok2 || len(v1.Content) != len(v2.Content) {
            return false
        }
        for key, v1Type := range v1.Content {
            v2Type, exists := v2.Content[key]
            if !exists || !equals(v1Type, v2Type) {
                return false
            }
        }
        return true

    case Sum:
        s1, ok1 := t1.Data.(SumType)
        s2, ok2 := t2.Data.(SumType)
        return ok1 && ok2 && equals(s1.Left, s2.Left) && equals(s1.Right, s2.Right)

    default:
        return false
    }
}
