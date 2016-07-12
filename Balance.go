package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"serial"
	"time"
)

type port interface {
	Write(b []byte) (n int, err error)
	Read(b []byte) (n int, err error)
}

/* WriteMessage
Takes in an argument of a connected port
*/

// check if s needs to be a pointer
func WriteMessage(s port, inputCommand string) string {
	fmt.Println("Writing Messages")
	//Check if the command is empty
	if inputCommand == "" {
		return "Empty Command"
	}
	// Check if the command has proper carrier return
	length := len(inputCommand)
	if length >= 3 {
		if inputCommand[(length-2):length] != "\r\n" {
			return "No <cr> and <lf> at the end of the command"
		}
	}
	if length < 3 {
		return "No <cr> and <lf> at the end of the command"
	}
	// Command with correct format gets turned into bytes
	command := []byte(inputCommand)
	_, err := s.Write(command)
	if err != nil {
		log.Fatal(err)
	}
	return "Writing Successful"

}

/* ReadMessage
Takes in an argument of a connected port
Returns a string of the message that the devices reponds
*/
func ReadMessage(s port) string {
	buf := make([]byte, 17)
	n, err := s.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The Weight is %q", buf[:n])
	//show the weight result on the terminal screen
	return string(buf[:])
}

/* CreateFile
takes in a filename (string) and data (string)
creates a file with filename and writes data
*/
func CreateFile(text string, weight string) string {
	//Check if the filename input (text) is empty
	if text == "" {
		return "Re-enter the sample ID"
	}
	// Append proper filename
	filename := "Weight of Sample "
	filename += text
	filename += ".txt"
	// Preppend a time stamp to the weight data
	data := time.Now().String()
	data += "\r\n"
	data += weight
	data += "\r"
	os.Create(filename)
	ioutil.WriteFile(filename, []byte(data), 0644)
	return "Writing File Successful"
}

func main() {
	c := &serial.Config{Name: "COM3", Baud: 9600, Parity: 'N', StopBits: 1, ReadTimeout: 5000}
	s, err := serial.OpenPort(c)
	fmt.Println("Accessing the balance")
	if err != nil {
		log.Fatal(err)
	}
	//According to the device manual
	//Q\r\n is used as a command for immediate weight data
	WriteMessage(s, "Q\r\n")
	time.Sleep(time.Second * 1)
	weight := ReadMessage(s)
	//Ask user input for the Sample ID
	fmt.Println("\r\nEnter the Sample ID")
	var text string
	fmt.Scanln(&text)

	CreateFile(text, weight)

}
