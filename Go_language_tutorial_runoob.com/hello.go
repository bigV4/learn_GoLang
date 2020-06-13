// helloworld.go
package main //定义了包名。你必须在源文件中非注释的第一行指明这个文件属于哪个包，如：package main。package main表示一个可独立执行的程序，每个 Go 应用程序都包含一个名为 main 的包。

import (
"fmt" //告诉 Go 编译器这个程序需要使用 fmt 包（的函数，或其他元素），fmt 包实现了格式化 IO（输入/输出）的函数。
"./myMath"
)
/*
关于包，根据本地测试得出以下几点：

 文件名与包名没有直接关系，不一定要将文件名与包名定成同一个。
 文件夹名与包名没有直接关系，并非需要一致。
 同一个文件夹下的文件只能有一个包名，否则编译报错。
*/

func main() { //注意⚠️，{ 不能在单独的行上
	/* 这是我的第一个简单的程序 */
	fmt.Println("Hello, World!")  //可以将字符串输出到控制台，并在最后自动增加换行字符 \n。
	fmt.Println("Google" + "Runoob")  //Go 语言的字符串可以通过 + 实现
	fmt.Println(mathClass.Add(1,1))
	fmt.Println(mathClass.Sub(1,1))
	printvar()
	_,numb,strs := numbers() //只获取函数返回值的后两个
	fmt.Println(numb,strs)
}

func printvar(){

    // 声明一个变量并初始化
    var a = "RUNOOB"
    fmt.Println(a)

    // 没有初始化就为零值
    var b int
    fmt.Println(b)

    // bool 零值为 false
    var c bool
    fmt.Println(c)
}

//一个可以返回多个值的函数
func numbers()(int,int,string){
	a , b , c := 1 , 2 , "str"
	return a,b,c
}

//要执行 Go 语言代码可以使用 go run 命令。 go run hello.go
//此外我们还可以使用 go build 命令来生成二进制文件：$ go build hello.go 
