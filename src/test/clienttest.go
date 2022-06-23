package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:9000")
	if err != nil {
		panic(err)
	}
	name := "TRD"
	writing(conn, name)
	defer disconnected(conn)
	recivefile(conn)
	/*for {
		//Read data

		// Write data to buffer
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter text: ")
		text, _ := reader.ReadString('\r')
		txt := name + ": " + text
		writing(conn, txt)

	}*/

}

func reading(conn net.Conn) (a string) {
	rr := make([]byte, 1020)
	aa, err := conn.Read(rr)
	if err != nil {
		panic(err)
	} else {
		return string(rr[:aa])
	}
}
func writing(conn net.Conn, txt string) {
	data := []byte(txt)
	conn.Write(data)
}
func sendfile(conn net.Conn, fname string) {
	var a int = 0
	aa, _ := os.Open(fname)

	ss := make([]byte, 100)
	filea, _ := aa.Stat()
	fsize := filea.Size()
	fsizeS := strconv.FormatInt(fsize, 10)
	writing(conn, fsizeS)
	for {
		msg := reading(conn)
		if msg != "" {
			fmt.Println(msg)
			break
		}
	}
	for {
		fmt.Println(a)
		st, err := aa.Read(ss)
		a += st
		a := int64(a)
		fmt.Println(ss)

		conn.Write(ss[0:st])

		if err == io.EOF || a == fsize {
			break
		}
	}
}
func recivefile(conn net.Conn) {
	p := 0
	var fsize int

	rr := make([]byte, 1020)
	newFile, _ := os.Create("lol.doc")
	//add file size
	for {
		fsize, _ = strconv.Atoi(reading(conn))
		if fsize != 0 {
			fmt.Println(fsize)
			break
		}
	}
	writing(conn, "all good")
	//file
	for {
		tt, err := conn.Read(rr)
		p += tt
		fmt.Println(tt)
		newFile.Write(rr[:tt])

		if err == io.EOF || p >= fsize {
			fmt.Println("end")
			break
		}

	}
	defer newFile.Close()

}
func disconnected(conn net.Conn) {
	if err := recover(); err != nil {
		conn.Close()
		fmt.Println(err)
	}
}
