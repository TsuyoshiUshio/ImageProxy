package main

import (
	"encoding/base64"
	"os"
	"fmt"
)

func main() {
	fmt.Println("base64 encoder")
	msg := os.Args[1] 
	fmt.Println("origina:" + msg)
	encoded := base64.StdEncoding.EncodeToString([]byte(msg))
	fmt.Println("encoded:" + encoded)
}