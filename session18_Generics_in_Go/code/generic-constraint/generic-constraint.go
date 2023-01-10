package main

import "fmt"

func main() {
	val := MyInt(777)
	fmt.Println(GetIntWithCustomFormat[MyInt](val))

	b := Box[MyInt]{}
	b = b

	v := GetIntWithCustomFormat[Point]
	v = v
}

type StringifiedInt interface {
	~int | ~uint
	String() string // same as fmt.Stringer
	fmt.Stringer
}

type Point struct {
	x int
	y int
}

func GetIntWithCustomFormat[T StringifiedInt](val T) string {
	return val.String()
}

type ID int

type MyInt int

func (mi MyInt) String() string {
	return fmt.Sprintf("MyInt{%d}", mi)
}

// 1. func

func findFirst[T StringifiedInt](s []T, el T) {
	el.String()
}

// 2. Types - struct

type Box[T any] struct {
	val  T
	vals []T
}

func (b *Box[whatever]) Hello() {
	b.vals = make([]whatever, 10)
}
