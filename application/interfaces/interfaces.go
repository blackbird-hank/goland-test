package interfaces

import "io"

type ResourceConnection interface {
	io.Closer
	Id() string
}
