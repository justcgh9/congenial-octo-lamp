package env

type Stack[T any] struct {
    elements []T
}


func (s *Stack[T]) Push(value T) {
    s.elements = append(s.elements, value)
}


func (s *Stack[T]) Pop() (T, bool) {
    if len(s.elements) == 0 {
        var zeroValue T 
        return zeroValue, false
    }
    top := s.elements[len(s.elements)-1]
    s.elements = s.elements[:len(s.elements)-1]
    return top, true
}


func (s *Stack[T]) Peek() (T, bool) {
    if len(s.elements) == 0 {
        var zeroValue T 
        return zeroValue, false
    }
    return s.elements[len(s.elements)-1], true
}


func (s *Stack[T]) IsEmpty() bool {
    return len(s.elements) == 0
}