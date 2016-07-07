package main

import (
	//"github.com/stretchr/testify/mock"
	//"os"
	"github.com/stretchr/testify/mock"
	"os"
	//"serial"
	//"sync"
	//"fmt"
	//"log"
	"syscall"
	"testing"
	//"Balance"
	//"github.com/golang/mock/gomock"
)

/* Make a mock serial.port
 */
type MockPort struct {
	mock.Mock
	f *os.File
}

// define the argument and return messages of this port mock
func (p *MockPort) Close() error {
	return p.f.Close()
}
func (p *MockPort) Write(command []byte) (n int, err error) {
	return p.f.Write(command)
}
func (p *MockPort) Read(buf []byte) (n int, err error) {
	return p.f.Read(buf)
}
func (p *MockPort) Flush() error {
	const TCFLSH = 0x540B
	_, _, err := syscall.Syscall(
		syscall.SYS_IOCTL,
		uintptr(p.f.Fd()),
		uintptr(TCFLSH),
		uintptr(syscall.TCIOFLUSH),
	)
	return err
}

//------------------------------------------------

func TestWriteMessageEmptyCommand(t *testing.T) string {
	/*
		basic_test.go catches emtpy/inaccessble port
	*/

	s0 := new(MockPort)
	s0.On("Write", []byte("")).Return(1, nil)

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
