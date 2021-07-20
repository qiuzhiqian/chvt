# chvt
一个使用golang实现的chvt库，逻辑实现参考了kbd中的c代码。

[chvt.c](https://kernel.googlesource.com/pub/scm/linux/kernel/git/legion/kbd/+/1.10/src/chvt.c)
[getfd.c](https://kernel.googlesource.com/pub/scm/linux/kernel/git/legion/kbd/+/1.10/src/getfd.c)

## chvt工具
工具通过cmd/main.go生成
```bash
$ go build -o chvt cmd/main.go
```