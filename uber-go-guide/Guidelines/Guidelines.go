package uberGoGuide

type F interface {
	f()
}

type S1 struct {
	Value int
}

// f() can't change the value.
// 值拷贝
func (s S1) f() {
	s.Value = 10
}

type S2 struct {
	Value int
}

// f can change the value
// 指针引用
func (s *S2) f() {
	s.Value = 10
}
