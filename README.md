# chvt
一个使用golang实现的chvt库，逻辑实现参考了kbd中的c代码。

[chvt.c](https://kernel.googlesource.com/pub/scm/linux/kernel/git/legion/kbd/+/1.10/src/chvt.c)
[getfd.c](https://kernel.googlesource.com/pub/scm/linux/kernel/git/legion/kbd/+/1.10/src/getfd.c)

另外推荐一篇vt相关的专业帖子

[Using Virtual Terminals under Linux](http://asm.sourceforge.net/articles/vt.html)

## chvt工具
工具通过cmd/main.go生成
```bash
$ go build -o chvt cmd/main.go
```