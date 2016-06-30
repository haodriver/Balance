package main

import (
	//"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"serial"
	"time"
)

func read(s *serial.Port) string {
	buf := make([]byte, 17)

	n, err := s.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The Weight is %q", buf[:n])
	return string(buf[:])
}

func write(s *serial.Port) {
	fmt.Println("Writing Messages")
	_, err := s.Write([]byte("Q\r\n"))

	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(n)
}

func main() {
	//fmt.Println("Sending Request")
	c := &serial.Config{Name: "COM3", Baud: 9600, Parity: 'N', StopBits: 1, ReadTimeout: 5000}
	s, err := serial.OpenPort(c)

	//fmt.Println(c)
	if false {
		fmt.Println(s)
		log.Fatal(err)
	}

	fmt.Println("Accessing the balance")
	if err != nil {
		log.Fatal(err)
	}

	write(s)
	time.Sleep(time.Second * 1)
	weight := read(s)
	//Opening file and creating a writer
	//reader := bufio.NewReader(os.Stdin)
	fmt.Println("\r\nEnter the Sample ID")
	var text string
	fmt.Scanln(&text)
	filename := ""
	filename += "Weight of Sample "
	filename += text
	filename += ".txt"

	data := ""
	data += time.Now().String()
	data += "\r\n"
	data += weight
	data += "\r"
	os.Create(filename)
	ioutil.WriteFile(filename, []byte(data), 0644)

}
