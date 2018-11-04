package main

import "fmt"

//eng
func processEng(str string) bool {
	for i := 0; i < len(str); i++ {
		if i == len(str)/2 {
			break
		}
		last := len(str) - i - 1
		if str[i] != str[last] {
			return false
		}
	}
	return true
}

//ch
func process(str string) bool {
	t := []rune(str)
	length := len(t)

	for i, v := range t {
		fmt.Printf("%v-%v-%d\n", i, v, len(string(v)))
	}

	for i := range t {
		if i == length/2 {
			break
		}
		last := length - i - 1
		if t[i] != t[last] {
			return false
		}
	}
	return true
}

func main() {
	var str string
	fmt.Scanf("%s", &str)
	if process(str) {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}
