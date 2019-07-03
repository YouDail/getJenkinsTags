rm -f projTags.micro.go	projTags.pb.go
protoc --proto_path=. --micro_out=. --go_out=.  projTags.proto
