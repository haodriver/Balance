package main

import (
	"fmt"

	"time"
)

func main() {
	weight := "0.02"
	dataSlice := ""
	dataSlice += time.Now().String()
	dataSlice += "\n"
	dataSlice += weight
	fmt.Println(dataSlice)
	/*err := ioutil.WriteFile(logFile, time.Now() weight, perm)
	check(err)
	*/
}
