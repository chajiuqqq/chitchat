package gen

//go:generate protoc -I../proto --go_out=../pb --go_opt=paths=source_relative --go-grpc_out=../pb --go-grpc_opt=paths=source_relative ../proto/threadService.proto ../proto/authService.proto
