package main

import (
	"fmt"
	"github.com/bingoohuang/valid"

	ten "github.com/bingoohuang/valid/translations/en"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
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

// use a single instance , it caches struct info
var (
	uni *ut.UniversalTranslator
	va  *valid.Validate
)

func main() {
	// NOTE: ommitting allot of error checking for brevity

	en := en.New()
	uni = ut.New(en, en)

	// this is usually know or extracted from http 'Accept-Language' header
	// also see uni.FindTranslator(...)
	trans, _ := uni.GetTranslator("en")

	va = valid.New()
	ten.RegisterDefaultTranslations(va, trans)

	translateAll(trans)
	translateIndividual(trans)
	translateOverride(trans) // yep you can specify your own in whatever locale you want!
}

func translateAll(trans ut.Translator) {
	type User struct {
		Username string `valid:"required"`
		Tagline  string `valid:"required,lt=10"`
		Tagline2 string `valid:"required,gt=1"`
	}

	user := User{
		Username: "Joeybloggs",
		Tagline:  "This tagline is way too long.",
		Tagline2: "1",
	}

	err := va.Struct(user)
	if err != nil {

		// translate all error at once
		errs := err.(valid.ValidationErrors)

		// returns a map with key = namespace & value = translated error
		// NOTICE: 2 errors are returned and you'll see something surprising
		// translations are i18n aware!!!!
		// eg. '10 characters' vs '1 character'
		fmt.Println(errs.Translate(trans))
	}
}

func translateIndividual(trans ut.Translator) {
	type User struct {
		Username string `valid:"required"`
	}

	var user User

	err := va.Struct(user)
	if err != nil {

		errs := err.(valid.ValidationErrors)

		for _, e := range errs {
			// can translate each error one at a time.
			fmt.Println(e.Translate(trans))
		}
	}
}

func translateOverride(trans ut.Translator) {
	va.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0} must have a value!", true) // see universal-translator for details
	}, func(ut ut.Translator, fe valid.FieldError) string {
		t, _ := ut.T("required", fe.Field())

		return t
	})

	type User struct {
		Username string `valid:"required"`
	}

	var user User

	err := va.Struct(user)
	if err != nil {

		errs := err.(valid.ValidationErrors)

		for _, e := range errs {
			// can translate each error one at a time.
			fmt.Println(e.Translate(trans))
		}
	}
}
