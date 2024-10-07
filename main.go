package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	addr := os.Args[1]
	fmt.Printf("SERVER ADDRESS IS %v\n", addr)

	path := os.Args[2]
	fmt.Printf("PATH IS %v\n", path)

	stream, _ := net.Dial("tcp", fmt.Sprintf("%v:1919", addr))
	fmt.Printf("CONNECTED TO %v\n", stream.RemoteAddr())

	file, _ := os.Open(path)
	filename := filepath.Base(path)
	fmt.Printf("FILE %v OPENED\n", filename)

	stream.Write([]byte(fmt.Sprintf("%v\n", filename)))
	fmt.Printf("FILENAME SENT\n")

	br := bufio.NewReader(stream)
	msgBytes, _, _ := br.ReadLine()
	msg := strings.TrimSpace(string(msgBytes))
	fmt.Printf("%v RECEIVED FROM THE SERVER\n", msg)

	bytes, _ := io.Copy(stream, file)

	fmt.Printf("SENT %v BYTES\n", bytes)
	stream.Close()

	fmt.Printf("DONE\n")

}
