package main

import (
	"fmt"
	"github.com/bingoohuang/valid"
)

// MyStruct ..
type MyStruct struct {
	String string `valid:"is-awesome"`
}

// use a single instance of Validate, it caches struct info
var va *valid.Validate

func main() {
	va = valid.New()
	va.RegisterValidation("is-awesome", ValidateMyVal)

	s := MyStruct{String: "awesome"}

	err := va.Struct(s)
	if err != nil {
		fmt.Printf("Err(s):%+v\n%+v\n", s, err)
	}

	s.String = "not awesome"
	err = va.Struct(s)
	if err != nil {
		fmt.Printf("Err(s):%+v\n%+v\n", s, err)
	}
}

// ValidateMyVal implements valid.Func
func ValidateMyVal(fl valid.FieldLevel) bool {
	return fl.Field().String() == "awesome"
}
