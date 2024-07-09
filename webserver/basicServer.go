package main

import(
"fmt"
"log"
"os"
)

//fooReader defines an io.reader to read from stdin.
type FooReader struct{}

//Read reads data from stdin.
func(fooreader *FooReader) Read(b []byte)(int, error){
	fmt.Print("in > ")
	return os.Stdin.Read(b)
}

//writer defines an io.writer to write to stdout
type FooWriter struct{}

func(FooWriter *FooWriter) Write(b []byte) (int,error){
	fmt.Print("out > ")
	return os.Stdout.Write(b)
}

func main(){
	//instatiate reader and writer
	var (
		reader FooReader
		writer FooWriter
	)

//create buffer to hold input/output
input := make([]byte, 4096)

//use reader to read input
s, err := reader.Read(input)
if err != nil {
	log.Fatalln("Unable to read data")
}
fmt.Printf("Read %d bytes from stdin\n", s)

//use writer to write output
s, err = writer.Write(input)
if err != nil{
	log.Fatalln("UNable to write data")
}
fmt.Printf("Write %d bytes to stdout\n", s)
}