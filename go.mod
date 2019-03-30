module github.com/learn/v1.0.0

go 1.12

require github.com/sirupsen/logrus v1.3.0

require (
	algorithms v1.0.0
	github.com/Arafatk/glot v0.0.0-20180312013246-79d5219000f0
	github.com/googollee/go-socket.io v1.4.1
	github.com/pkg/errors v0.8.1 // indirect
	github.com/stretchr/testify v1.3.0
	golang.org/x/crypto v0.0.0-20190313024323-a1f597ede03a // indirect
	golang.org/x/lint v0.0.0-20190313153728-d0100b6bd8b3 // indirect
	golang.org/x/net v0.0.0-20190313220215-9f648a60d977 // indirect
	golang.org/x/sys v0.0.0-20190316082340-a2f829d7f35f // indirect
	golang.org/x/tools v0.0.0-20190315214010-f0bfdbff1f9c // indirect
	gonum.org/v1/plot v0.0.0-20190312081609-cd8a2043e414
	gotest.tools v2.2.0+incompatible
	graphs v1.0.0
	server v1.0.0
)

replace algorithms => ./github.com/learn/algorithms

replace server => ./github.com/learn/server

replace graphs => ./github.com/learn/graphs
