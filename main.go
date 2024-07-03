package main

import (
	// 可以理解为是 项目名（go module 名） + 包的相对路径（从项目根目录到包的路径）
	"github.com/zrh099/golearn/pkg/demo1"
)

func main() {
	// 此处 demo1 是包名，而不是包的路径名
	// 但因为包所在目录和包名一致，所以可能会误认为，此处是上面 import 路径的末尾
	// 这样保持一致的好处就是，当 demo1 包函数发生错误是，可以通过相同名称在 import 找到对应的项目路径，然后排查并解决错误
	demo1.Hello1()
}
