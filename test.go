package main

import (
	//"github.com/stretchr/testify/mock"
	"os"
	"serial"
	"testing"
)

/* Test WriteMessage
 */

func TestWriteMessage(t *testing.T) string {
	/*
		basic_test.go catches emtpy/inaccessble port
	*/

	port0 := os.Getenv("PORT0")
	c0 := &serial.Config{Name: port0, Baud: 9600}
	s0, err := serial.OpenPort(c0)
	if err != nil {
		t.Fatal(err)
	}

	// test case of writing an empty message
	if WriteMessage(s0, "") != "Empty Command" {
		return "missing empty command case"
	}
	// test case of writing a message without proper <cr>
	if WriteMessage(s0, "Q") != "No <cr> and <lf> at the end of the command" {
		return "missing incomplete command case "
	}
	return "YAYYYY"
	// test case of writing a proper message

}

//func TestReadMessage(t *testing.T) {

//}
