package main

import (
	"fmt"
	"goland-test/application/app"
)

func main() {
	pool := app.NewResourceConnectionPool()
	conn := pool.GetConnection()
	fmt.Printf("Connection id: %s\n", conn.Id())
}
