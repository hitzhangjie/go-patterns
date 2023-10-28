package bridge

type Color interface {
	String() string
}

type Red struct{}

func (r Red) String() string {
	return "red"
}

type Blue struct{}

func (b Blue) String() string {
	return "blue"
}
