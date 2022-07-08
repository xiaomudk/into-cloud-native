package main

import (
    "fmt"
    "io/ioutil"
)

func main() {

    content, err := ioutil.ReadFile("test.txt")

    if err != nil {
	    fmt.Errorf("open err:%s", err.Error())
    }
    fmt.Println(string(content))
}