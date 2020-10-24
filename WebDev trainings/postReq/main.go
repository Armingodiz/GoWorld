package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", ":8585")
	if err != nil {
		log.Fatalln(err)
	}
	defer listener.Close()
	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go serve85(connection)
	}
}
func serve85(connection net.Conn) {
	scanner := bufio.NewScanner(connection)
	i := 0
	var method, path string
	for scanner.Scan() {
		txt := scanner.Text()
		fmt.Println(txt)
		if i == 0 {
			ln := strings.Fields(txt)
			method = ln[0]
			path = ln[1]
		}
		if txt == "" {
			fmt.Println("DONEEEEEEE !")
			break
		}
		i++
	}
	recHandler(method , path , connection)
}
func recHandler(m, p string, connection net.Conn) {
	switch {
	case m == "GET" && p == "/":
		handleIndex(connection)
	case m == "GET" && p == "/submit":
		handleSubmit(connection)
	case m == "POST" && p == "/submit":
		handleApplyPost(connection)
	default:
		handleIndex(connection)
	}
}
func handleIndex(connection net.Conn)  {

	body := `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title> INDEX </title>
</head>
<body>
	<h1> INDEX PAGE IS OPEN </h1>
    <a href="/">  index Page </a><br>
	<a href="/submit">  submit Page   </a><br>
</body>
</html>
	`
	io.WriteString(connection, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(connection, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(connection, "Content-Type: text/html\r\n")
	io.WriteString(connection, "\r\n")
	io.WriteString(connection, body)
	connection.Close()
}
func handleSubmit(connection net.Conn)  {
	body := `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title> INDEX </title>
</head>
<body>
	<h1> SUBMIT PAGE IS OPEN </h1>
    <a href="/">  index Page </a><br>
	<a href="/submit">  submit Page   </a><br>
    <form action="/submit" method="POST">
    <input type="text" placeholder=" my name ">
    <input type="submit" value="submit">
    </form>
</body>
</html>
	`
	io.WriteString(connection, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(connection, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(connection, "Content-Type: text/html\r\n")
	io.WriteString(connection, "\r\n")
	io.WriteString(connection, body)
	connection.Close()
}
func handleApplyPost(connection net.Conn)  {
	body := `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title> INDEX </title>
</head>
<body>
	<h1> SUBMIT IS DONE  </h1>
    <a href="/">  index Page </a><br>
	<a href="/submit">  submit Page   </a><br>
</body>
</html>
	`
	io.WriteString(connection, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(connection, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(connection, "Content-Type: text/html\r\n")
	io.WriteString(connection, "\r\n")
	io.WriteString(connection, body)
	connection.Close()
}

