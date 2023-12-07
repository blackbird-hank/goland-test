package app

import (
	"goland-test/application/interfaces"
	lib_interfaces "goland-test/library/interfaces"
	"goland-test/library/lib"

	"github.com/google/uuid"
)

var _ interfaces.ResourceConnection = (*resourceConnection)(nil)
var _ lib_interfaces.ConnectionFactory[interfaces.ResourceConnection] = (*resourceConnectionFactory)(nil)

type resourceConnection struct {
	id string
}

func (*resourceConnection) Close() error {
	return nil
}

func (r *resourceConnection) Id() string {
	return r.id
}

type resourceConnectionFactory struct{}

func (r *resourceConnectionFactory) Connect() interfaces.ResourceConnection {
	return &resourceConnection{
		id: uuid.NewString(),
	}
}

func NewResourceConnectionPool() lib_interfaces.ConnectionPool[interfaces.ResourceConnection] {
	factory := resourceConnectionFactory{}
	return lib.NewConnectionPool[interfaces.ResourceConnection](&factory)
}
