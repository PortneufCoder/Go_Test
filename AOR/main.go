package main

import (
	"fmt"

	"github.com/gobuffalo/packr" 
)



func main() {
	box := packr.NewBox("../AOR/serverContent")
	data := box.String("SIP_Stack.json")
	fmt.Println("Contents of file:", data)

	handleRequests()
}
