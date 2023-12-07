package lib

import (
	lib_interfaces "goland-test/library/interfaces"
	"io"
)

var _ lib_interfaces.ConnectionPool[io.Closer] = (*connectionPool[io.Closer])(nil)

type connectionPool[C io.Closer] struct {
	factory lib_interfaces.ConnectionFactory[C]
}

func NewConnectionPool[C io.Closer](factory lib_interfaces.ConnectionFactory[C]) lib_interfaces.ConnectionPool[C] {
	return &connectionPool[C]{
		factory: factory,
	}
}

func (c *connectionPool[C]) GetConnection() C {
	return c.factory.Connect()
}
