package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:9000")
	if err != nil {
		panic(err)
	}
	name := "TRD"
	writing(conn, "TRD")
	for {
		//Read data
		reading(conn)
		// Write data to buffer
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter text: ")
		text, _ := reader.ReadString('\n')
		txt := name + ": " + text
		writing(conn, txt)
	}
}

func reading(conn net.Conn) {
	rr := make([]byte, (124 * 4))
	aa, err := conn.Read(rr)
	if err != nil {
		panic(err)
	} else {
		fmt.Println(string(rr[:aa]))
	}
}
func writing(conn net.Conn, txt string) {
	data := []byte(txt)
	conn.Write(data)
}
func sendfile(conn net.Conn, fname string) {
	var a int = 0
	aa, _ := os.Open(fname)

	ss := make([]byte, 1020)
	filea, _ := aa.Stat()
	fsize := filea.Size()

	writing(conn, string(fsize))
	for {
		st, err := aa.Read(ss)
		a += st
		a := int64(a)
		fmt.Println(string(ss))
		conn.Write(ss)
		if err == io.EOF || a == fsize {
			break
		}
	}
}
func recivefile(conn net.Conn) {
	var p int = 0
	lm := make([]byte, 1020)
	rr := make([]byte, 1020)
	newFile, _ := os.Create("lol.txt")
	//add file size
	km, _ := conn.Read(lm)
	fsize := string(lm[:km])
	//file
	for {
		tt, err := conn.Read(rr)
		p += tt
		p := string(p)
		newFile.Write(rr)
		if err == io.EOF || p == fsize {
			break
		}
	}
	defer newFile.Close()

}
