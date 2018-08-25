# Linux下的tty终端修改显示中文乱码

2016年11月28日 00:59:56

​		

虚拟机下tty终端显示的中文全是“菱形”符号，所以重新下载了个字体fbterm

下载命令：sudo apt-get install fbterm

会自动下载完成，下载完成后运行

sudo fbterm 即可

更改字体和字体大小

vi ~/.fbtermrc

```
font-names=Ubuntu Mono
font-size=14
```

