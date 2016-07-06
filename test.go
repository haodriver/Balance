package main

import (
	//"github.com/stretchr/testify/mock"
	//"os"
	"github.com/stretchr/testify/mock"
	"serial"
	"testing"
	//"Balance"
	//"github.com/golang/mock/gomock"
)

/* Test WriteMessage
 */
type MyInterface interface {
	WriteMessage(s, inputCommand)
}

func (m *MyInterface) name() {

}

func TestWriteMessageEmptyCommand(t *testing.T) string {
	/*
		basic_test.go catches emtpy/inaccessble port
	*/

	c0 := &serial.Config{Name: "Port1", Baud: 9600}
	s0, err := serial.OpenPort(c0)
	if err != nil {
		t.Fatal(err)
	}

	// test case of writing an empty message
	if WriteMessage(s0, "") != "Empty Command" {
		return "missing empty command case"
	}

	return "Empty Message Test Success"
	// test case of writing a proper message

}

/* func TestWriteMessageShortCommand(t *testing.T) string {
	// test case of writing a message without proper <cr>
	if WriteMessage(s0, "Q") != "No <cr> and <lf> at the end of the command" {
		return "missing incomplete command case "
	}
	return "Incomplet Message Test Success"
}

func TestWriteMessageComplete(t *testing.T) string {
	if WriteMessage(s0, "Q\r\n") != "Writing Successful" {
		return "Error in Connection"
	}
	return "Full Message Test Success"
}
*/

//func TestReadMessage

//func TestReadMessage(t *testing.T) {

//}
