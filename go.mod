module github.com/learn/v1.0.0

go 1.12

require github.com/sirupsen/logrus v1.3.0

require (
	algorithms v1.0.0
	github.com/googollee/go-socket.io v1.4.1
	github.com/pkg/errors v0.8.1 // indirect
	github.com/stretchr/testify v1.3.0
	gotest.tools v2.2.0+incompatible
)

replace algorithms => ./github.com/learn/algorithms
