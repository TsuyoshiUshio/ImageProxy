package main

import (
    "net/http"
    "net/url"
    "github.com/labstack/echo"
    "io"
    "image"
    "image/jpeg"
    "image/png"
    "log"
)

func main() {
    e := echo.New()
    e.GET("/convertedimage/:url", func(c echo.Context) error {
        source, err := url.QueryUnescape(c.Param("url"))
        response, err := http.Get(source)
        if err != nil {
            panic(err)
        }
        defer response.Body.Close()
        log.Printf(response.Header.Get("Content-type"));

        if response.Header.Get("Content-type") == "image/jpeg" {
            jpegImage := ImageRead(response.Body)

            pr, pw := io.Pipe()
            go func() {
                err = png.Encode(pw, jpegImage)
                pw.Close()
            }()
            return c.Stream(http.StatusOK, "image/png", pr)
        } else {
            return c.Stream(http.StatusOK, "image/png", response.Body)
        }
    })
    e.Logger.Fatal(e.Start(":8090"))
}

func ImageRead(inputImage io.ReadCloser) (image image.Image) {
    img, err := jpeg.Decode(inputImage)
    if err != nil {
        log.Fatal(err)
    }
    inputImage.Close()
    return img
}

