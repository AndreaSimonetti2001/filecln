package main

import (
	"bufio"
	"filecln/logger"
	"filecln/try"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	if len(os.Args) != 3 {
		logger.Error("Usage: EXECUTABLE ADDRESS FILE.")
	}

	addr := os.Args[1]
	logger.Info("Server address is %v.", addr)

	path := os.Args[2]
	logger.Info("Path is %v.", path)

	stream, e := net.Dial("tcp", fmt.Sprintf("%v:1919", addr))
	try.Catch(e)
	logger.Info("Connected to %v.", stream.RemoteAddr())

	file, e := os.Open(path)
	try.Catch(e)
	filename := filepath.Base(path)
	logger.Info("File %v opened.", filename)

	_, e = stream.Write([]byte(fmt.Sprintf("%v\n", filename)))
	try.Catch(e)
	logger.Info("Filename sent.")

	br := bufio.NewReader(stream)
	msgBytes, _, e := br.ReadLine()
	try.Catch(e)
	msg := strings.TrimSpace(string(msgBytes))
	logger.Info("%v Received from the server.", msg)

	bytes, e := io.Copy(stream, file)
	try.Catch(e)
	logger.Info("Sent %v bytes.", bytes)

	e = stream.Close()
	try.Catch(e)
	logger.Info("Stream closed.")

	e = file.Close()
	try.Catch(e)
	logger.Info("File closed.")

}
