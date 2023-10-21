package main

import (
	"fmt"
	"regexp"
)

func main() {
	r := regexp.MustCompile("^[a-zA-Z]+[a-zA-Z0-9._-]*@[a-zA-Z].[a-zA-Z]")
	match := r.Match([]byte("1nas1@gmail.com"))
	fmt.Println(match)
}
