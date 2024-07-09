package main

import (
	"io"
	"log"
	"net"
)

// echo is handler function that simply echoes recieved data
func echo(conn net.Conn) {
	defer conn.Close()

	/*reader := bufio.NewReader(conn)
	s, err := reader.ReadString('\n')
	if err != nil{
		log.Fatalln("Unable to read data")
	}
	log.Printf("Read %d bytes: %s", len(s), s)

	log.Println("Writing data")
	writer := bufio.NewWriter(conn)
	if _, err := writer.WriteString(s); err != nil {
		log.Fatalln("Unable to write")
	}
	writer.Flush()*/
	//Create a bufffer to store recieved data
	b := make([]byte, 512)

	for {
		//Recieve data via conn.Read into a buffer
		size, err := conn.Read(b[0:])
		if err == io.EOF {
			log.Println("Client disconnected")
			break
		}
		if err != nil {
			log.Println("Unexpected error")
			break
		}
		log.Printf("Recieved %d bytes: %s\n", size, string(b))

		//Send data via conn.Write
		log.Println("Writing data")
		if _, err := conn.Write(b[0:size]); err != nil {
			log.Fatalln("Unable to write data")
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", ":20080")
	if err != nil {
		log.Fatalln("Unable to bind to port")
	}
	log.Println("Listening on 0.0.0.0:20080")
	for {
		//Wait for connection Create net.conn on connection established
		conn, err := listener.Accept()
		log.Println("Received connection")
		if err != nil {
			log.Fatalln("Unable to accpet connection")
		}
		//Handle the conection. Using goroutine for concurrency
		go echo(conn)
	}
}
