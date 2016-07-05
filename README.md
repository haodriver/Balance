# Balance
## WriteMessage 
WriteMessage takes in the port and command specified by the manual (p67) and writes it to the device through the port.
## ReadMessage
ReadMessage takes in the port and return the weight of the current sample
## CreateFile 
CreateFile takes in the weight information returned by ReadMessage and the sample ID given by the user and creates a file with the name of "Weight of Sample XX" with the weight information and a time stamp.
