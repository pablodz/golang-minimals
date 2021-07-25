

1. go get google.golang.org/grpc
2. go get -u github.com/golang/protobuf/protoc-gen-go
3. protoc -I datafiles/ datafiles/transaction.proto --go_out=plugins=grpc:datafiles

