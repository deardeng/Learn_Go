package main

import (
	"testing"
	"time"
)

func TestSave(t *testing.T) {
	stu := &student{
		Name: "sut01",
		Sex:  "man",
		Age:  18,
	}

	err := stu.Save()
	if err != nil {
		t.Fatalf("save student error")
	}
}

func TestLoad(t *testing.T) {
	stu := &student{
		Name: "sut01",
		Sex:  "man",
		Age:  18,
	}

	err := stu.Save()
	if err != nil {
		t.Fatalf("save student error")
	}

	stu2 := &student{}
	time.Sleep(10 * time.Second)
	err = stu2.Load()
	if err != nil {
		t.Fatalf("load student failed, err :%v", err)
	}

	if stu.Name != stu2.Name {
		t.Fatalf("load student failed, name not equal.")
	}

	if stu.Age != stu2.Age {
		t.Fatalf("load student failed, age not equal.")
	}

	if stu.Sex != stu2.Sex {
		t.Fatalf("load student failed, sex not equal.")
	}
}
