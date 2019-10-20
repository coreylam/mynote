# vscode远程开发环境搭建

## 1. 准备远程服务器
[购买链接](https://buy.cloud.tencent.com/)

## 2. 本地按照vscode
[vscode下载链接](https://code.visualstudio.com/)

## 3. 安装sshclient
git 中包含 ssh.exe，通过安装git 配置环境变量，使用ssh。 [点击下载安装git](https://git-scm.com/download/win)
将该路径（ C:\Program Files\Git\usr\bin）添加到 PATH环境变量，如果是安装到其它路径的对应修改路径。

## 4. vscode中安装远程插件
在扩展中按照 remote-ssh
安装后勾选 Remote.SSH: Show Login Terminal 下的 `Always reveal the SSH login terminal`

## 5. 添加ssh远程配置
配置 ~/.ssh/config 文件， 内容示例如下：
```shell
Host <配置名称，自定义>
    HostName <远端服务器IP>
    User <ssh登录的用户名>
	Port <ssh登录的端口>
    IdentityFile <私钥路径>
```

其中 私钥 的生成如下：

```shell
#在本地机器生成密钥对
ssh-keygen -t rsa

#将公钥拷贝到服务器上 
ssh-copy-id -p <ssh登录的端口>  <ssh登录的用户名>@<远端服务器IP>
```

