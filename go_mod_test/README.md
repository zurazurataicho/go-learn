# Go self package module test

## Commands

```shell
mkdir greetings
cd greetings
go mod init zura.org/greetings
vi greetings.go

cd ..
mkdir hello
cd hello
go mod init zura.org/hello
vi hello.go

go mod edit -replace zura.org/greetings=../greetings
go mod tidy
go run .
```
