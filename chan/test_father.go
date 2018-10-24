package main

import (
	"fmt"
)

type Conf struct {
	Name string
}

func Test(conf *Conf) {
	fmt.Println(conf.Name)
}

type CConf struct {
	Conf
	ID int
}

func main() {

	cc := &CConf{Conf{"11111"}, 1}

	Test(&cc.Conf)
	fmt.Println(cc.Conf.Name)
}
