package utils

import (
	"fmt"
	"log"
	"os"
	//"path/filepath"
)

// GL exported
var GL *log.Logger

// EL exported
var EL *log.Logger

// R exported
var R *log.Logger

func init() {
	//gpath := os.Getenv("GOPATH")
	//path := filepath.Join(gpath, "/src/quoteqdl/files/sys.log")

	f, err := os.OpenFile("file/sys.log", os.O_RDWR|os.O_CREATE|os.O_APPEND|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("Error opening file:", err)
		//os.Exit(1)
	}
	f2, err := os.OpenFile("file/resp.log", os.O_RDWR|os.O_CREATE|os.O_APPEND|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("Error opening file:", err)
		//os.Exit(1)
	}
	R = log.New(f2, "RESP:\t", log.Ldate|log.Ltime|log.Lshortfile)
	GL = log.New(f, "LOG:\t", log.Ldate|log.Ltime|log.Lshortfile)
	EL = log.New(f, "ERROR:\t", log.Ldate|log.Ltime|log.Lshortfile)
}
