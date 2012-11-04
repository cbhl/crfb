package main

import (
    "code.google.com/p/goweb/goweb"
    "fmt"
    "os"
)

func main() {
    goweb.MapFunc("/", hello)

    err := goweb.ListenAndServe(":" + os.Getenv("PORT"))
    if err != nil {
      panic(err)
    }
}

func hello(c *goweb.Context) {
    fmt.Fprintf(c.ResponseWriter, "hello, world")
}
