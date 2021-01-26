package main

import (
	"fmt"
	"github.com/erezlevip/named-pipes/pkg"
	"log"
	"time"
)

func main() {
	p, err := pkg.NewPipe("main.p")
	if err != nil {
		log.Fatal(err)
	}
	go writeTests(p)
	for v := range p.Listen('\n') {
		fmt.Print("value from channel:" + string(v))
	}
}

func writeTests(p pkg.Pipe) {
	fmt.Println("start writing")
	defer p.Close()
	for i := 0; i < 10; i++ {
		p.Write([]byte(fmt.Sprintf("test write times:%d\n", i)))
		i++
		time.Sleep(time.Second)
	}
}
