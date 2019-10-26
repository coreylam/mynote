## 安装golang

[官方文档](https://golang.google.cn/doc/)

[W3CSchool--go入门]](https://www.w3cschool.cn/go/go-environment.html)

中文指引， https://go-zh.org/doc/
https://cloud.tencent.com/developer/article/1200431

[菜鸟教程--Go语言编程](https://www.runoob.com/go/go-tutorial.html)

安装Go工具
若你要从旧版本的Go升级，那么首先必须卸载已存在的版本。

Linux、Mac OS X 和 FreeBSD 的安装包
下载此压缩包并提取到 /usr/local 目录，在 /usr/local/go 中创建Go目录树。例如：

tar -C /usr/local -xzf go$VERSION.$OS-$ARCH.tar.gz
该压缩包的名称可能不同，这取决于你安装的Go版本和你的操作系统以及处理器架构。

（此命令必须作为root或通过 sudo 运行。）

要将 /usr/local/go/bin 添加到 PATH 环境变量， 你需要将此行添加到你的 /etc/profile（全系统安装）或 $HOME/.profile 文件中：

export PATH=$PATH:/usr/local/go/bin
安装到指定位置
Go二进制发行版假定它们会被安装到 /usr/local/go （或Windows下的 c:\Go）中，但也可将Go工具安装到不同的位置。 此时你必须设置 GOROOT 环境变量来指出它所安装的位置。

例如，若你将Go安装到你的home目录下，你应当将以下命令添加到 $HOME/.profile 文件中：

export PATH=$PATH:/usr/local/go/bin
安装到指定位置
Go二进制发行版假定它们会被安装到 /usr/local/go （或Windows下的 c:\Go）中，但也可将Go工具安装到不同的位置。 此时你必须设置 GOROOT 环境变量来指出它所安装的位置。

例如，若你将Go安装到你的home目录下，你应当将以下命令添加到 $HOME/.profile 文件中：

export GOROOT=$HOME/go
export PATH=$PATH:$GOROOT/bin
注：GOROOT 仅在安装到指定位置时才需要设置。

Mac OS X安装包
打开此包文件 并跟随提示来安装Go工具。该包会将Go发行版安装到 /usr/local/go 中。

此包应该会将 /usr/local/go/bin 目录放到你的 PATH 环境变量中。 要使此更改生效，你需要重启所有打开的终端回话。

Windows
对于Windows用户，Go项目提供两种安装选项（从源码安装除外）： zip压缩包需要你设置一些环境变量，而实验性MSI安装程序则会自动配置你的安装。

MSI安装程序
打开此MSI文件 并跟随提示来安装Go工具。默认情况下，该安装程序会将Go发行版放到 c:\Go 中。

此安装程序应该会将 c:\Go\bin 目录放到你的 PATH 环境变量中。 要使此更改生效，你需要重启所有打开的命令行。

Zip压缩包
下载此zip文件 并提取到你的自选目录（我们的建议是c:\Go）：

若你选择了 c:\Go 之外的目录，你必须为你所选的路径设置 GOROOT 环境变量。

将你的Go根目录中的 bin 子目录（例如 c:\Go\bin）添加到你的 PATH 环境变量中。

在Windows下设置环境变量
在Windows下，你可以通过在系统“控制面板”中，“高级”标签上的“环境变量”按钮来设置环境变量。 Windows的一些版本通过系统“控制面板”中的“高级系统设置”选项提供此控制板。


## hello world
测试你的安装
通过构建一个简单的程序来检查Go的安装是否正确，具体操作如下：

首先创建一个名为 hello.go 的文件，并将以下代码保存在其中：

package main

import "fmt"

func main() {
    fmt.Printf("hello, world\n")
}
接着通过 go 工具运行它：

$ go run hello.go
hello, world
若你看到了“hello, world”信息，那么你的Go已被正确安装。

设置你的工作环境变量
差不多了，你只要再设置一下就好。

请阅读如何使用Go编程，它提供了使用Go工具的基本设置说明。

卸载 Go
要从你的系统中移除既有的Go安装，需删除 go 目录。 在 Linux、Mac OS X、和 FreeBSD 系统下通常为 /usr/local/go， 在 Windows 下则为 c:\Go。

你也应当从你的 PATH 环境变量中移除 Go 的 bin 目录。 在 Linux 和 FreeBSD 下你应当编辑 /etc/profile 或 $HOME/.profile。 若你是通过Mac OS X 包安装的 Go，那么你应当移除 /etc/paths.d/go 文件。 Windows 用户请阅读在 Windows 下设置环境变量一节。



```
wget https://golang.google.cn/doc/install?download=go1.13.3.linux-amd64.tar.gz
tar -C /usr/local -xzf go1.13.3.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
```

```
ln -s  /usr/local/go/bin/go go
ln -s  /usr/local/go/bin/gofmt gofmt
ln -s <目标文件> <链接文件>
```