package notblank

import (
	"testing"

	"github.com/bingoohuang/valid"
	"github.com/go-playground/assert/v2"
)

type test struct {
	String    string      `valid:"notblank"`
	Array     []int       `valid:"notblank"`
	Pointer   *int        `valid:"notblank"`
	Number    int         `valid:"notblank"`
	Interface interface{} `valid:"notblank"`
	Func      func()      `valid:"notblank"`
}

func TestNotBlank(t *testing.T) {
	v := valid.New()
	err := v.RegisterValidation("notblank", NotBlank)
	assert.Equal(t, nil, err)

	// Errors
	var x *int
	invalid := test{
		String:    " ",
		Array:     []int{},
		Pointer:   x,
		Number:    0,
		Interface: nil,
		Func:      nil,
	}
	fieldsWithError := []string{
		"String",
		"Array",
		"Pointer",
		"Number",
		"Interface",
		"Func",
	}

	errors := v.Struct(invalid).(valid.ValidationErrors)
	var fields []string
	for _, err := range errors {
		fields = append(fields, err.Field())
	}

	assert.Equal(t, fieldsWithError, fields)

	// No errors
	y := 1
	x = &y
	va := test{
		String:    "str",
		Array:     []int{1},
		Pointer:   x,
		Number:    1,
		Interface: "value",
		Func:      func() {},
	}

	err = v.Struct(va)
	assert.Equal(t, nil, err)
}
