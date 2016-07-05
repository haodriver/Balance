# Balance
## WriteMessage(s *serial.Port, inputCommand string) string{}
WriteMessage takes in a port pointer and command specified by the manual (p67), writes it to the device through the port and returns the status of the writing.
## ReadMessage(s *serial.Port) string {}
ReadMessage takes in the port and return the weight of the current sample
## CreateFile(filename string, data string) {}
CreateFile takes in the weight information returned by ReadMessage and the sample ID given by the user and creates a file with the name of "Weight of Sample XX" with the weight information and a time stamp.
