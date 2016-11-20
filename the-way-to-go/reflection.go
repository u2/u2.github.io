package main

import (
	"fmt"
	"reflect"
)

func main() {
	e1()
	e2()
	e3()
  e4()
}

func e1() {
	type MyInt int
	var x MyInt = 7
	v := reflect.ValueOf(x)

	y := v.Interface().(MyInt) // y will have type float64.
	fmt.Println(y)

	fmt.Println(v.Interface())
}

func e2() {
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	fmt.Println("settability of v:", v.CanSet())
}

func e3() {
	var x float64 = 3.4
	p := reflect.ValueOf(&x) // Note: take the address of x.
	fmt.Println("type of p:", p.Type())
	fmt.Println("settability of p:", p.CanSet())
}

func e4() {
	var x float64 = 3.4
	p := reflect.ValueOf(&x) // Note: take the address of x.
	v := p.Elem()
	fmt.Println("settability of v:", v.CanSet())
}


