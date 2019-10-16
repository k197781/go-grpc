### compile
After adding protocol definition to `.proto`, you need to compile.
```
protoc greeter.proto --go_out=plugins=grpc:.
```

### start server
```
go run server/server.go
```

### execute
```
go run client/client.go world
> Greeting: Hello world
```

### create gRPC document
```
protoc --doc_out=./docs --doc_opt=html,index.html *.proto
```
