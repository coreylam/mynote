// 第一行代码 package main 定义了包名。你必须在源文件中非注释的第一行指明这个文件属于哪个包，如：package main。package main表示一个可独立执行的程序，每个 Go 应用程序都包含一个名为 main 的包。
package main

// import "fmt" 告诉 Go 编译器这个程序需要使用 fmt 包（的函数，或其他元素），fmt 包实现了格式化 IO（输入/输出）的函数。
import "fmt"

func print_hello(){
    // var result int
    println("hello")
}

func test_closure() func() int{
    i := 0
    return func() int {
        i++
        return i
    }
}
// func main() 是程序开始执行的函数。main 函数是每一个可执行程序所必须包含的，一般来说都是在启动后第一个执行的函数（如果有 init() 函数则会先执行该函数）。
func main(){
    // fmt.Println(...) 可以将字符串输出到控制台，并在最后自动增加换行字符
    fmt.Printf("hello, world\n")
    print_hello()
    func_count := test_closure()
    for i:=0; i<5; i++ {
        println(func_count())
    }

    // 当标识符（包括常量、变量、类型、函数名、结构字段等等）以一个大写字母开头，如：Group1，那么使用这种形式的标识符的对象就可以被外部包的代码所使用（客户端程序需要先导入这个包），这被称为导出（像面向对象语言中的 public）；标识符如果以小写字母开头，则对包外是不可见的，但是他们在整个包的内部是可见并且可用的（像面向对象语言中的 private ）。

}



// $ go run hello.go
// hello, world


// 关键字
// 下面列举了 Go 代码中会使用到的 25 个关键字或保留字：
// break	    default	        func	interface	select
// case	    defer	        go	    map	        struct
// chan	    else	        goto	package	    switch
// const	    fallthrough	    if	    range	    type
// continue	for	            import	return	    var


// 除了以上介绍的这些关键字，Go 语言还有 36 个预定义标识符：
// uint, uint8, uint16, uint32, uint64
// int, int8, int16, int32, int64
// float32, float64
// complex64, complex128
// byte, uintptr

// bool

// string

// false, true, iota, nil


// append, cap, close, complex
// copy, imag, len, make, new, panic
// print, println, real, recover

//类型不同多个变量, 全局变量, 局部变量不能使用这种方式
var (
    vname1 int
    vname2 string
)

// go语言的数据类型:
// 数字类型 
//     int, int8, int16, int32, int64
//     uint, uint8, uint16, uint32, uint64
//     float32, float64
//     complex64, complex128
//     byte, rune, uintptr

// 布尔型
//     bool


// 字符串类型:
//     string


// 派生类型:
//     (a) 指针类型（Pointer）
//     (b) 数组类型
//     (c) 结构化类型(struct)
//     (d) 联合体类型 (union)
//     (e) 函数类型
//     (f) 切片类型
//     (g) 接口类型（interface）
//     (h) Map 类型
//     (i) Channel 类型
