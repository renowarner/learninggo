package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"
)

var (
	h1 = "<!DOCTYPE html><html><head><title>Title</title></head><body><p>"
	h2 = "</p></body></html>"
	e1 = `C:\Windows\System32\rundll32.exe`
	e2 = "url.dll,FileProtocolHandler"
	e3 = "file:///C:/Users/Reno/Desktop/Project/outputter.html"
)

func main() {
	var para string
	fmt.Print("What is your Paragraph?: ")
	fmt.Scanf("%v", &para)
	
	block := []string{h1, para, h2}
	
	d1 := []byte(strings.Join(block, " "))
	err := ioutil.WriteFile("outputter.html", d1, 0644)
	if err != nil {
		panic(err)
	}

	exec.Command(e1, e2, e3).Start()
}
