package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}

/*
要执行 Go 语言代码可以使用 go run 命令。
$ go run hello.go
*/

/*
此外我们还可以使用 go build 命令来生成二进制文件：
$ go build hello.go 
$ ls
hello    hello.go
$ ./hello 
*/