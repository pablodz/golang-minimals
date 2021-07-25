# Create proto files 

1. go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
2. protoc --go_out=. *.proto