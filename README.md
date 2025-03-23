# rest-vs-grpc

https://www.youtube.com/watch?v=UVdY-EK3c3o


```
protoc \
  -I proto/proto \
  --go_out=proto/pb \
  --go-grpc_out=proto/pb \
  --go_opt=paths=source_relative \
  --go-grpc_opt=paths=source_relative \
  complex_data.proto
```