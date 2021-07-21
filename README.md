# chvt
一个使用golang实现的chvt库，逻辑实现参考了kbd中的c代码。

[chvt.c](https://kernel.googlesource.com/pub/scm/linux/kernel/git/legion/kbd/+/1.10/src/chvt.c)
[getfd.c](https://kernel.googlesource.com/pub/scm/linux/kernel/git/legion/kbd/+/1.10/src/getfd.c)

另外推荐一篇vt相关的专业帖子

[Using Virtual Terminals under Linux](http://asm.sourceforge.net/articles/vt.html)

[Linux Programming Hints](https://www.linuxjournal.com/article/2798)

[How VT-switching works](https://dvdhrm.wordpress.com/2013/08/24/how-vt-switching-works/)

## chvt工具
### 编译
工具通过cmd/main.go生成
```bash
$ go mod tidy
$ go build -o chvt cmd/main.go
```

### 使用
```
$ sudo ./chvt --num 8
```