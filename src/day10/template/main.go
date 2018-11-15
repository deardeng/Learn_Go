package main

import (
	"fmt"
	"html/template"
	"os"
)

type Person struct {
	Name  string
	Title string
	Age   string
}

func main() {
	t, err := template.ParseFiles("D:/GO DEV/src/day10/template/index.html")
	if err != nil {
		fmt.Println("parse file err:", err)
		return
	}
	p := Person{Name: "deardeng", Age: "22", Title: "啊傻瓜"}
	if err := t.Execute(os.Stdout, p); err != nil {
		fmt.Println("There was an error:", err.Error())
	}
}
