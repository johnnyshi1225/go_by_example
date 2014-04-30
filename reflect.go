package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age int
}

func main() {
	var x float64 = 3.4
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	fmt.Println(t)
	fmt.Println(v)

	fmt.Println(v.Type())
	fmt.Println(v.Kind())
	fmt.Println(v.Float())

	p :=reflect.ValueOf(&x)
	fmt.Println(p)
	v = p.Elem()
	fmt.Println(v)
	v.SetFloat(7.1)

	fmt.Println(x)

	john := Person{"john", 25}
	fmt.Println(john)
	s := reflect.ValueOf(&john).Elem()
	typeOfPersion := s.Type()

	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Println(f.Type())
		fmt.Printf("%d: %s %s = %v\n", i, typeOfPersion.Field(i).Name, f.Type(), f.Interface())
	}

}

