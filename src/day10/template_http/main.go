package main

import (
	"fmt"
	"io"
	"net/http"
	"text/template"
)

var myTemplate *template.Template

type Result struct {
	output string
}

type Person struct {
	Name  string
	Title string
	Age   int
}

func (p *Result) Write(b []byte) (n int, err error) {
	fmt.Println("called by template")
	p.output += string(b)
	return len(b), nil
}

func userInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handle hello")
	var arr []Person
	p := Person{Name: "deardeng01", Age: 22, Title: "哈哈哈"}
	p1 := Person{Name: "deardeng02", Age: 22, Title: "哈哈哈"}
	p2 := Person{Name: "deardeng03", Age: 22, Title: "哈哈哈"}
	arr = append(arr, p)
	arr = append(arr, p1)
	arr = append(arr, p2)

	resultWriter := &Result{}
	io.WriteString(resultWriter, "hello world")
	err := myTemplate.Execute(w, arr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("template render data:", resultWriter.output)
}

func initTemplate(filename string) (err error) {
	myTemplate, err = template.ParseFiles(filename)
	if err != nil {
		fmt.Println("parse file err:", err)
		return
	}
	return
}

func main() {
	initTemplate("D:/GO DEV/src/day10/template_http/index.html")
	http.HandleFunc("/user/info", userInfo)
	err := http.ListenAndServe("0.0.0.0:8888", nil)
	if err != nil {
		fmt.Println("http listen failed")
	}
}
