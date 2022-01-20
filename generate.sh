#protoc --go_out=. --go_opt=paths=source_relative greet/greetpb/greet.proto

# go - server, client, dart - client
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative  --dart_out=grpc:greet/greetpb/dart  greet/greetpb/greet.proto
