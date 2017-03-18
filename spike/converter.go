package main

import (
    "fmt"
    "io"
    "net/http"
    "os"
    "image"
    "image/jpeg"
    "image/png"
    "log"
)

func main() {
    var url string = "https://c.s-microsoft.com/ja-jp/CMSImages/spk-ushio.jpg?version=f6328834-736f-4b59-2f85-7b198346ef4d"

    response, err := http.Get(url)
    if err != nil {
        panic(err)
    }
    defer response.Body.Close()
    dataofImage := ImageRead(response.Body)
    fmt.Println("Reading image...")
    fmt.Println("Converting image to png ...")
    FormatPng(dataofImage)
}

func ImageRead(inputImage io.ReadCloser) (image image.Image) {
    img, err := jpeg.Decode(inputImage)
    if err != nil {
        log.Fatal(err)
    }
    inputImage.Close()
    return img
}

func FormatPng(img image.Image) {
    out, err := os.Create("converted.png")
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    err = png.Encode(out, img)
    if err!= nil {
        fmt.Println(err)
        os.Exit(1)
    }
    fmt.Println("\n success...\n")
}