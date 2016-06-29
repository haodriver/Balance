package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"serial"
	"strings"
	"time"
)

func read(s *serial.Port) string {
	buf := make([]byte, 17)

	n, err := s.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	return string(buf[:])
	fmt.Printf("%q", buf[:])
}

func write(s *serial.Port) {
	fmt.Println("Writing Messages")
	n, err := s.Write([]byte("Q\r\n"))

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(n)
}

func main() {
	fmt.Println("Sending Request for Weight Data")
	c := &serial.Config{Name: "COM3", Baud: 9600, Parity: 'N', StopBits: 1, ReadTimeout: 5000}
	s, err := serial.OpenPort(c)

	//fmt.Println(c)
	if false {
		fmt.Println(s)
		log.Fatal(err)
	}

	fmt.Println("Opening the Port is successful")
	if err != nil {
		log.Fatal(err)
	}

	write(s)
	time.Sleep(time.Second * 1)
	weight := read(s)
	//Opening file and creating a writer
	logFile, err := os.Open("C:/Desktop/BalanceLog.txt")
	check(err)
	dataSlice := ""
	dataSlice += time.Now().String()
	dataSlice += "\n"
	dataSlice += weight
	err := ioutil.WriteFile(logFile, dataSlice, perm)
	check(err)

}
