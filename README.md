# blog

https://medium.com/@aravindhanjay/a-todo-app-using-grpc-web-and-vue-js-4e0c18461a3e

## Prepare protoc-gen-go

```bash
GIT_TAG="v1.2.0" # change as needed
go get -d -u github.com/golang/protobuf/protoc-gen-go
git -C "$(go env GOPATH)"/src/github.com/golang/protobuf checkout $GIT_TAG
go install github.com/golang/protobuf/protoc-gen-go
```

## build golang stub

```
protoc -I todo/ todo/todo.proto --go_out=plugins=grpc:todo
```

## Prepare protoc-gen-grpc-web

- Download
  https://github.com/grpc/grpc-web/releases/tag/1.0.4

- Move
  sudo mv ~/Downloads/protoc-gen-grpc-web-1.0.4-darwin-x86_64 \
   /usr/local/bin/protoc-gen-grpc-web

- Chmod
  chmod +x /usr/local/bin/protoc-gen-grpc-web

# Client

npm install
npm install --save bootstrap-vue
npm install --save google-protobuf grpc-web
