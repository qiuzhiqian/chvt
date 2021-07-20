package main

import (
	"flag"
	"log"

	"github.com/qiuzhiqian/chvt"
)

func main() {
	var num int = -1
	log.Println("chvt")
	flag.IntVar(&num, "num", 0, "change active tty to ttyN")
	flag.Parse()

	if num <= 0 {
		flag.Usage()
		return
	}
	log.Println("num:", num)
	err := chvt.Chvt(uint(num))
	if err != nil {
		log.Println(err)
	}
}
