package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func readAllFile() {
	fileName := "/Users/zhanglu/Documents/doing/tmp/test/文件读写/test.txt"
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))
}

func readAllLinne() {
	fileName := "/Users/zhanglu/Documents/doing/tmp/test/文件读写/test.txt"
	fileObj, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer fileObj.Close()

	reader := bufio.NewReader(fileObj)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println(line)
			return
		}
		if err != nil {
			break
		}
		fmt.Print(line)
	}
}

func write() {
	fileName := "./test3.txt"
	fileObj, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer fileObj.Close()
	str := "who's your dady\n"
	fileObj.Write([]byte(str))
	fileObj.WriteString("hello world\n")
}

func main() {
	//readAllLinne()
	write()
}
