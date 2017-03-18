package main

import (
	"fmt"
	"os"
	"net/url"
)

func main() {
	fmt.Println("URL encoder")
	msg := os.Args[1] 
	fmt.Println("original:" + msg)
	fmt.Println("encoded :" + url.QueryEscape(msg))
}