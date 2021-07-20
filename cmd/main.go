package main

import (
	"log"

	"github.com/qiuzhiqian/chvt"
)

func main() {
	log.Println("chvt")
	err := chvt.Chvt(9)
	if err != nil {
		log.Println(err)
	}
}
