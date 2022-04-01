package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"syscall"
	"time"
)

var pipeFile = "/tmp/pipe.ipc"

func main() {
	_ = os.Remove(pipeFile)

	err := syscall.Mkfifo(pipeFile, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	go write()
	go read()
	select {}

}

func read() {
	fmt.Println("open a named pipe file for read.")
	file, _ := os.OpenFile(pipeFile, os.O_RDWR, os.ModeNamedPipe)
	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadBytes('\n')
		fmt.Println("read...")
		if err == nil {
			fmt.Print("load string: " + string(line))
		}
	}
}

func write() {
	fmt.Println("start schedule writing.")
	f, err := os.OpenFile(pipeFile, os.O_RDWR, 0777)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	i := 0
	for {
		fmt.Println("write string to named pipe file.")
		f.WriteString(fmt.Sprintf("test write times:%d\n", i))
		i++
		time.Sleep(time.Second)
		if i == 10 {
			break
		}
	}
}
