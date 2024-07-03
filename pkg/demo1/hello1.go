// - 此处 package 关键字就是指定包名
// - 子目录内的函数文件，可以随意命名，如 hello1.go，自己识别就好
// - 注意：子目录内可以创建多个函数文件，但同一个子目录的所有文件只能属于同一个包名，就是必须设置相同的 package
// - 如 demo1 子目录，后续创建的文件，必须添加 package demo1
package demo1

import "fmt"

func Hello1() {
	fmt.Println("this is Hello-1 from demo1")
}
