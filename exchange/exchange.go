package exchange

type Exchange struct {
	List []interface{}
}

func New(exchanges ...interface{}) *Exchange {
	return &Exchange{exchanges}
}
