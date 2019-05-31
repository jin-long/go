package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err:= ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "file error : %v", err)
			continue
		}
		for _, line:=range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	
	for strline, line:=range counts {
		fmt.Printf("line:%d, strline:%s \n", line, strline)
	}
}
