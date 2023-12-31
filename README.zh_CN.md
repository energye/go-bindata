### 为Go执行文件内置二进制静态资源

### 适用在Go中低版本不支持 embed.FS 包

### 使用方式

> 1. 克隆该项目
> 
> `git clone https://github.com/energye/go-bindata.git`
> 
> 2. 进入 go-bindata/go-bindata 目录 
> 
> 执行命令: `go install` 编译并安装到go/bin目录, 不出意外将得到 go-bindata 执行文件在go/bin目录中
> 
> 3. 内置资源到exe需先准备好静态资源放到你自己的目录, 参考示例 [examples](examples) [assets](examples%2Fassets)
> 
> 开始使用:
> 1. 使用go的 go:generate注解命令, 写到main函数go文件中开头位置
> 2. 在main函数目录执行 `go generate` 命令生成存放静态资源的字节go文件
> 3. 代码示例
```go
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
	fmt.Println(len(data), string(data), err) // 18 静态资源文件 <nil>
}

```