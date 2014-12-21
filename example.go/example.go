package main

import (
	"fmt"
	"github.com/prestonTao/hashcash"
	"time"
)

func main() {
	message := "taopopoo@126.com"
	bits := 20
	t1 := time.Now()
	nonce := hashcash.Work(message, bits)
	t2 := time.Now()
	fmt.Println("工作时间：", t2.Sub(t1).Seconds())
	ok := hashcash.Check(message, bits, nonce)
	fmt.Println(ok, nonce)
}
