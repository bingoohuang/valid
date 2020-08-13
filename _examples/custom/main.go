package main

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"reflect"

	"github.com/bingoohuang/valid"
)

// DbBackedUser User struct
type DbBackedUser struct {
	Name sql.NullString `valid:"required"`
	Age  sql.NullInt64  `valid:"required"`
}

// use a single instance of Validate, it caches struct info
var va *valid.Validate

func main() {
	va = valid.New()

	// register all sql.Null* types to use the ValidateValuer CustomTypeFunc
	va.RegisterCustomTypeFunc(ValidateValuer, sql.NullString{}, sql.NullInt64{}, sql.NullBool{}, sql.NullFloat64{})

	// build object for validation
	x := DbBackedUser{Name: sql.NullString{String: "", Valid: true}, Age: sql.NullInt64{Int64: 0, Valid: false}}

	err := va.Struct(x)
	if err != nil {
		fmt.Printf("Err(s):\n%+v\n", err)
	}
}

// ValidateValuer implements valid.CustomTypeFunc
func ValidateValuer(field reflect.Value) interface{} {
	if valuer, ok := field.Interface().(driver.Valuer); ok {

		val, err := valuer.Value()
		if err == nil {
			return val
		}
		// handle the error how you want
	}

	return nil
}
