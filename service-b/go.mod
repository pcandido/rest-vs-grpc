module github.com/pcandido/rest-vs-grpc/service-b

go 1.23.2

replace github.com/pcandido/rest-vs-grpc/proto => ../proto

require (
  google.golang.org/grpc v1.70.0
	golang.org/x/net v0.32.0 // indirect
	golang.org/x/sys v0.28.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20241202173237-19429a94021a // indirect
	google.golang.org/protobuf v1.35.2 // indirect
	github.com/pcandido/rest-vs-grpc/proto v0.0.0
)

