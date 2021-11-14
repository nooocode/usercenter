protoc -I . --go_out=. --go-triple_out=. user.proto common.proto api.proto menu.proto role.proto tenant.proto identity.proto
protoc-go-inject-tag -input=api.pb.go
protoc-go-inject-tag -input=common.pb.go
protoc-go-inject-tag -input=identity.pb.go
protoc-go-inject-tag -input=menu.pb.go
protoc-go-inject-tag -input=role.pb.go
protoc-go-inject-tag -input=tenant.pb.go
protoc-go-inject-tag -input=user.pb.go