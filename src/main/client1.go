package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:9000")
	if err != nil {
		panic(err)
	}
	name := "DDD"
	writing(conn, name)
	for {
		//Read data

		// Write data to buffer
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter text: ")
		text, _ := reader.ReadString('\r')
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
