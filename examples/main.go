package main

import (
	"fmt"
	"github.com/energye/go-bindata/examples/pkg/assets"
)

//参数说明： //go:generate go-bindata -fs 固定格式, -o:二进制资源go文件输出存放目录 -pkg:包名 ./assets 要生成的二进制资源目录
//go:generate go-bindata -fs -o=pkg/assets/assets.go -pkg=assets ./assets

func main() {
	// 需要先使用 go generate 命令生成静态资源字节go文件
	// 如未执行 go generate 此处有错误
	data, err := assets.AssetFile().ReadFile("assets/assets.txt")
	fmt.Println(len(data), string(data), err)
}
