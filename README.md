``Golang GRPC POC``
```
4 types:
unary
server streaming
client streaming
bi directional streaming
```
```
Steps to setup grpc and protocol buffers ->

https://grpc.io/docs/languages/go/quickstart/
https://go-zero.dev/docs/prepare/protoc-install/

If you make any changes in .proto file, then you need to generate the code again.
Below is the command to do the same:
$protoc --go_out=. --go-grpc_out=. proto/greet.proto
$protoc --go_out=. --go-grpc_out=. proto/employee.proto

```

```
cd go-grpc && go mod tidy
cd server && go run *.go
cd client && go run *.go
```