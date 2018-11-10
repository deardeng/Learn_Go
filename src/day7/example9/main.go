package main

import (
	"fmt"
	"os"
	"time"
)

type PathError struct {
	path       string
	op         string
	createtime string
	message    string
}

func (p *PathError) Error() string {
	return fmt.Sprintf("path=%s, op=%s createTime=%s, message=%s\n", p.path, p.op, p.createtime, p.message)
}

func Open(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return &PathError{
			path:       filename,
			op:         "read",
			createtime: time.Now().Format("Mon Jan 2 15:04:05 -0700 MST 2006"),
			message:    err.Error(),
		}
	}
	defer file.Close()
	return nil
}

func main() {
	err := Open("c:/xxx")
	//one
	v, ok := err.(*PathError)
	if ok {
		fmt.Println("get path error, ", v)
	}
	//two
	if err != nil {
		fmt.Println(err)
	}
	//three
	switch v := err.(type) {
	case *PathError:
		fmt.Println("get path error,", v)
	default:
	}
}
