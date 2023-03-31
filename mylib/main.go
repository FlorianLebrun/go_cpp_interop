package main

import (
	"C"
	"unsafe"
)

type MyData struct {
	a C.int
	b C.int
}

func (o *MyData) Add() C.int {
	return o.a + o.b
}

//export Add
func Add(a, b C.int) C.int {
	return a + b
}

func MakePtr() unsafe.Pointer {
	p := new(MyData)
	return unsafe.Pointer(p)
}

//export RetPtr
func RetPtr() unsafe.Pointer {
	return MakePtr()
}

func main() {
}
