package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
   dat, err := ioutil.ReadFile("./rtk6-list.txt")
   check(err)
   kanjiList := strings.Split(string(dat), " ")
   fmt.Print(len(kanjiList))
}