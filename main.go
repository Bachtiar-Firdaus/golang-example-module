package main

import (
	"fmt"
	// disini kita menambahkan module yang telah kita buat tadi
	go_say_hello "github.com/Bachtiar-Firdaus/belajar-golang-modules"
)

func main() {
	// setelah kita menambahkan module kita bisa langsung menggunakan module yang telah kita buat
	fmt.Println(go_say_hello.SayHello())
	// kemudian kita bisa langsung menggunakan depedency injection yang baru
	fmt.Println(go_say_hello.SayName())
}
