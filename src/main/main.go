package main

import (
	"net"
	"bufio"
	"strconv"
	"log"
	"strings"
	"bytes"
	"time"
	"net/url"
)


func main() {
	port := 8080;
	ln, err := net.Listen("tcp", ":" + strconv.Itoa(port))
	if err != nil {
		panic("Failed start server: " + err.Error());
	}
	log.Print("Server started at " + strconv.Itoa(port) + " port")
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	bufr := bufio.NewReader(conn)
	buf := make([]byte, 1024)

	var input bytes.Buffer

	defer conn.Close();
	for {
		readBytes, err := bufr.Read(buf)
		if err != nil {
			log.Printf("handle connection error, err=%s", err)
			return
		}

		input.Write(buf[:readBytes]) // Сохраняем полученные даты


		httpRequestEnd := "\r\n\r\n"
		readString := string(buf[:readBytes])
		if (strings.Contains(readString, httpRequestEnd)) {
			break	// конец запроса
		}
	}

	response := parseInputData(input.String())
	conn.Write(response)
}

type request struct {
	method string
	url string
	protocol string
}

func isMethodAllowed(method string) (bool)  {
	return strings.Compare(method, "GET") == 0 || strings.Compare(method, "HEAD") == 0;
}

func parseInputData(input string) ([]byte) {
	var infoLine= strings.Split(input, "\r\n")[0]

	var splitLine= strings.Split(infoLine, " ")

	decoded_url, _ := url.QueryUnescape(splitLine[1])
	userRequest := request{
		method:   splitLine[0],
		url:      decoded_url,
		protocol: splitLine[2],
	}

	var response bytes.Buffer
	response.WriteString(userRequest.protocol)
	response.WriteString(" ")

	var file File

	if !isMethodAllowed(userRequest.method) {
		response.WriteString("405 Method Not Allowed")
		response.WriteString("\r\n")
		response.WriteString("Allow: GET, HEAD")
		response.WriteString("\r\n")

	} else {
		head := strings.Compare(userRequest.method, "HEAD") == 0;
		file = GetFile(userRequest.url, head)
		switch file.status {
		case 200:
			{
				response.WriteString("200 OK\r\n")
				response.WriteString("Content-Type: " +
					file.content_type + "\r\n")
				response.WriteString("Content-Length: " +
					strconv.Itoa(file.length) + "\r\n")
				break
			}

		case 403:
			{
				response.WriteString("403 Forbidden\r\n")
				break
			}

		case 404:
			{
				response.WriteString("404 File Not Found\r\n")
				break
			}
		default:
			break
		}
	}

	// Дописываем хедеры
	response.WriteString("Date: " + time.Now().String())
	response.WriteString("\r\n")
	response.WriteString("Server: Golang HTTP Server")
	response.WriteString("\r\n")
	response.WriteString("Connection: Close")
	response.WriteString("\r\n")
	response.WriteString("\r\n")

	if (file.status == 200) {
		response.Write(file.content)
	}

	return response.Bytes()
}