git init

grpc boiler template

## 1 Setting

1. golang 
```
go env 로 GOPATH 확인필요
```
2. gland 또는 vscode
3. protoc
``` brew install protobuf  
brew install clang-format  
go get -u github.com/golang/protobuf/protoc-gen-go
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
flutter pub global activate protoc_plugin   //flutter grpc 사용시 
```


## 2 시작하기
1. proto 작성
> /greet/greetpb/greet.proto 
./generate.sh


