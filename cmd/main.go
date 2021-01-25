package main

import (
	"bufio"
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
	reader := bufio.NewReader(p)

	go func() {
		fmt.Println("start reading using a channel")
		for v := range p.Listen('\n') {
			if err == nil {
				fmt.Print("value from channel:" + string(v))
			}
		}
	}()

	fmt.Println("start reading using a file")
	for {
		line, err := reader.ReadBytes('\n')
		if err == nil {
			fmt.Print("value from file:" + string(line))
		}
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
