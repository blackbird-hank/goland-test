package lib_interfaces

import "io"

type ConnectionFactory[C io.Closer] interface {
	Connect() C
}

type ConnectionPool[C io.Closer] interface {
	GetConnection() C
}
