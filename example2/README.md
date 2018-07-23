## Command to generate pb
```shell
protoc --go_out=plugins=grpc:. workout.proto
```

```shell
protoc -I . --go_out=plugins=grpc:. ./*.proto
```