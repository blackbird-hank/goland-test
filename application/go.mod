module goland-test/application

replace goland-test/library => ../library

go 1.21

require (
	github.com/google/uuid v1.4.0
	goland-test/library v0.0.0-00010101000000-000000000000
)
