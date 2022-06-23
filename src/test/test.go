package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	var a int = 0
	aa, _ := os.Open("peaga.txt")

	ss := make([]byte, 1020)
	filea, _ := aa.Stat()
	fsize := filea.Size()
	fsizeS := strconv.FormatInt(fsize, 10)
	fmt.Println(fsizeS)
	newFile, _ := os.Create("lol.txt")
	for {
		st, err := aa.Read(ss)
		a += st
		a := int64(a)
		newFile.Write(ss)
		if err == io.EOF || a == fsize {
			break
		}
	}
	defer newFile.Close()
}
