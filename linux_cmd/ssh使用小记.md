# ssh使用小记



- 免密登录

```shell
ssh -i <指定私钥> root@<远程主机IP>
```



- 指定快捷登录

```shell
#路径: ~/.ssh/config
Host <快捷登录名称>
        HostName <远程主机IP>
        Port 22
        User root
        IdentityFile <私钥文件路径>


#登录
ssh <快捷登录名称>
```



