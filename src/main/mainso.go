package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
)

func main() {
	i := make([]string, 0, 5)
	ln, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		} else {
			i = append(i, "user"+strconv.Itoa(len(i)+1))
		}
		go handle(conn)

	}
}
func handle(conn net.Conn) {
	ll := reading(conn)
	if ll == "TRD" || ll == "DDD" {
		tt := fmt.Sprintf("%s connected", ll)
		fmt.Println(tt)
		adress := conn.RemoteAddr()
		as := adress.String()
		fmt.Println(as[6:])

	} else {
		fmt.Println("disconnected")
		conn.Close()
	}
	for {
		/*reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter text: ")
		text, _ := reader.ReadString('\n')*/

		data := []byte("text")
		conn.Write(data)

		rr := make([]byte, (125 * 4))
		aa, err := conn.Read(rr)
		if err != nil {
			continue
		} else {
			fmt.Println(string(rr[0:aa]))
		}
		defer fmt.Println("disconnected")
	}

	/*scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		fmt.Println(string(scanner.Bytes()))

	}*/
}

func reading(conn net.Conn) (a string) {
	rr := make([]byte, (124 * 4))
	aa, err := conn.Read(rr)
	if err != nil {
		panic(err)
	}
	return string(rr[:aa])

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
