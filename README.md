# Demonstration of GoLand bug

* Two modules, a library and application
* Code is shared using go replace directive (see [application/go.mod](application/go.mod))
* Library implements a connection pool, with a connection implementing io.Closer
* Application implements a connection as a 'resource connection', which implements io.Closer
* Go cli compiles and runs this just fine. This can be shown by running `make run`
* Yet GoLand (`Build #GO-233.11799.228, built on December 1, 2023`) shows compile errors:
```
Cannot use interfaces.ResourceConnection as the type io.CloserType does not implement 'io.Closer' as some methods are missing:Close() error
```

**Note**: if application/interfaces/interface.go were to be moved to /application/app/ then GoLand would not
log compile errors.