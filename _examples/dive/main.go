package main

import (
	"fmt"

	"github.com/bingoohuang/valid"
)

// Test ...
type Test struct {
	Array []string          `valid:"required,gt=0,dive,required"`
	Map   map[string]string `valid:"required,gt=0,dive,keys,keymax,endkeys,required,max=1000"`
}

// use a single instance of Validate, it caches struct info
var va *valid.Validate

func main() {
	va = valid.New()

	// registering alias so we can see the differences between
	// map key, value validation errors
	va.RegisterAlias("keymax", "max=10")

	var test Test

	val(test)

	test.Array = []string{""}
	test.Map = map[string]string{"test > than 10": ""}
	val(test)
}

func val(test Test) {
	fmt.Println("testing")
	err := va.Struct(test)
	fmt.Println(err)
}
