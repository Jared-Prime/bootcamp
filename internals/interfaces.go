package internals

type I interface {
	DoSomeWork()
}

type T struct {
	a int
}

func (t *T) DoSomeWork() {
}

func RunExample() {
	t := &T{}
	i := I(t)
	print(i)
}
