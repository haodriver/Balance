package main

import (
	//"github.com/stretchr/testify/mock"
	//"os"
	"github.com/stretchr/testify/mock"
	//"os"
	//"serial"
	//"sync"
	//"fmt"
	//"log"
	//"syscall" "Balance"
	"os"
	"testing"
	//"github.com/golang/mock/gomock"
)

/* Make a mock serial.port
 */
type MockPort struct {
	mock.Mock
	f *os.File
}

// define the argument and return messages of this port mock
// Make all the methods that is qualified to be the port interface
func (s0 MockPort) Write(b []byte) (int, error) {
	args := s0.Called(b)
	return args.Int(0), args.Error(1)
}
func (s0 MockPort) Read(b []byte) (n int, err error) {
	args := s0.Called(b)
	return args.Int(0), args.Error(1)
}
func (s0 MockPort) Flush() (err error) {
	return s0.Flush()
}
func (s0 MockPort) Close() (err error) {
	return s0.Close()
}

//------------------------------------------------

func TestWriteMessageEmptyCommand(t *testing.T) {
	/*
		basic_test.go catches emtpy/inaccessble port
	*/

	s0 := new(MockPort)
	s0.On("Write", []byte("")).Return(0, nil)
	// test case of writing an empty message
	if WriteMessage(s0, "") != "Empty Command" {
		t.Error("Failed empty command case")
	}
	// test case of writing a proper message
}

func TestWriteMessageShortCommand(t *testing.T) {
	// test case of writing a message without proper <cr>
	s0 := new(MockPort)
	s0.On("Write", []byte("Q")).Return(3, nil)
	if WriteMessage(s0, "Q") != "No <cr> and <lf> at the end of the command" {
		t.Error("missing incomplete command case ")
	}

}

func TestWriteMessageComplete(t *testing.T) {
	s0 := new(MockPort)
	s0.On("Write", []byte("Q\r\n")).Return(3, nil)
	if WriteMessage(s0, "Q\r\n") != "Writing Successful" {
		t.Error("Error in Connection")
	}

}

//func TestReadMessage

//func TestReadMessage(t *testing.T) {

//}
