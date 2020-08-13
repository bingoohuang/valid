package main

import (
	"fmt"
	"github.com/bingoohuang/valid"
)

// User contains user information
type User struct {
	FirstName      string     `valid:"required"`
	LastName       string     `valid:"required"`
	Age            uint8      `valid:"gte=0,lte=130"`
	Email          string     `valid:"required,email"`
	FavouriteColor string     `valid:"iscolor"`                // alias for 'hexcolor|rgb|rgba|hsl|hsla'
	Addresses      []*Address `valid:"required,dive,required"` // a person can have a home and cottage...
}

// Address houses a users address information
type Address struct {
	Street string `valid:"required"`
	City   string `valid:"required"`
	Planet string `valid:"required"`
	Phone  string `valid:"required"`
}

// use a single instance of Validate, it caches struct info
var va *valid.Validate

func main() {
	va = valid.New()

	validateStruct()
	validateVariable()
}

func validateStruct() {
	address := &Address{
		Street: "Eavesdown Docks",
		Planet: "Persphone",
		Phone:  "none",
	}

	user := &User{
		FirstName:      "Badger",
		LastName:       "Smith",
		Age:            135,
		Email:          "Badger.Smith@gmail.com",
		FavouriteColor: "#000-",
		Addresses:      []*Address{address},
	}

	// returns nil or ValidationErrors ( []FieldError )
	err := va.Struct(user)
	if err != nil {

		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		if _, ok := err.(*valid.InvalidValidationError); ok {
			fmt.Println(err)
			return
		}

		for _, err := range err.(valid.ValidationErrors) {

			fmt.Println(err.Namespace())
			fmt.Println(err.Field())
			fmt.Println(err.StructNamespace())
			fmt.Println(err.StructField())
			fmt.Println(err.Tag())
			fmt.Println(err.ActualTag())
			fmt.Println(err.Kind())
			fmt.Println(err.Type())
			fmt.Println(err.Value())
			fmt.Println(err.Param())
			fmt.Println()
		}

		// from here you can create your own error messages in whatever language you wish
		return
	}

	// save user to database
}

func validateVariable() {
	myEmail := "joeybloggs.gmail.com"

	errs := va.Var(myEmail, "required,email")

	if errs != nil {
		fmt.Println(errs) // output: Key: "" Error:Field validation for "" failed on the "email" tag
		return
	}

	// email ok, move on
}
