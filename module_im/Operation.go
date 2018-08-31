package main


type Operator interface {
	connect(msg *Msg) (key string, err error)
}

type DefaultOperator struct {
}

func (o *DefaultOperator) connect(msg *Msg) (key string, err error) {
	key, err = Connect(msg)
	return
}
