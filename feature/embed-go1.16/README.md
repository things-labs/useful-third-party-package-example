# go 1.16 feature

[原文](https://mp.weixin.qq.com/s?__biz=MzAxMTA4Njc0OQ==&mid=2651442728&idx=1&sn=c73548dc85c22a651cb3282b6de24709&chksm=80bb10dab7cc99cccfe2a49b13825a8b8d33a9b42dfaf4630fef0eb2eafae6f8d43b6cc16ec3&mpshare=1&scene=24&srcid=11198bNTZthirgq2QaVfYwTP&sharer_sharetime=1605769359482&sharer_shareid=fbafc624aa53cd09857fb0861ac2a16d&exportkey=AXlBYBmFf3Ffhp8wJcVOvmo%3D&pass_ticket=yXtG%2FwKZlgpcLNsJaTEezqzFsT%2Bf464n09HsOnW%2Frpkzt3%2BJL4XH99RrIPqUcJR3&wx_header=0#rd)

## `//go:embed` 指令

之前第三方的库,基本是基于 `go generate`,将静态资源文件生成go源文件,最后编译进二进制文件中.
官方的实现,通过 `//go:embed` 指令,在编译时将静态资源嵌入二进制文件中.然后,Go 通过标准库,让用户能够访问这些内嵌的资源.

### `//go:embed` 指令的用法

#### 规则
在变量声明上方,通过 `//go:embed` 指令指定一个或多个符合 `path.Match` 模式的要嵌入的文件或目录.
相关规则或使用注意如下:

- 跟其他指令一样,`//` 和 `go:embed` 之间不能有空格.（不会报错,但该指令会被编译器忽略）
- 指令和变量声明之间可以有空行或普通注释,但不能有其他语句；

```go
//go:embed message.txt

var message string
```
以上代码是允许的,不过建议紧挨着,而且建议变量声明和指令之间也别加注释,注释应该放在指令上方.

- 变量的类型只能是 `string`、`[]byte` 或 `embed.FS`,即使是这三个类型的别名也不行；
- 允许有多个 `//go:embed` 指令.多个文件或目录可以通过空格分隔,也可以写多个指令.
```go
//go:embed image template
//go:embed html/index.html
var content embed.FS
```
- 文件或目录使用的是相对路径,相对于指令所在 Go 源文件所在的目录,路径分隔符永远使用 `/`；当文件或目录名包含空格时,可以使用双引号或反引号括起来
- 对于目录,会以该目录为根,递归的方式嵌入所有文件和子目录
- 变量的声明可以是导出或非导出的；可以是全局也可以在函数内部；但只能是声明,不能给初始化值
```go
//go:embed message.txt
var message string = "" // 编译不通过：go:embed cannot apply to var with initializer
```
- 只能内嵌模块内的文件,比如 .git/* 或软链接文件无法匹配；空目录会被忽略；
- 模式不能包含 `.` 或 `..`,也不能以 `/` 开始,如果要匹配当前目录所有文件,应该使用 `*` 而不是 `.`；
- 以`.`或`_`开头的文件或目录将被忽略
## 标准库

和 embed 相关的标准库有 5 个,其中 2 个是新增的:embed 和 io/fs；net/http,text/template 和 html/template 包依赖 io/fs 包,而 embed.FS 类型实现了 io/fs 包的 FS 接口,因此这 3 个包可以使用 embed.FS.（Go1.16 发布时可能还会增加其他包或修改一些包的内容

