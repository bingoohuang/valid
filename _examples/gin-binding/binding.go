package main

import (
	"reflect"
	"sync"

	"github.com/bingoohuang/valid"
	"github.com/gin-gonic/gin/binding"
)

type defaultValidator struct {
	once  sync.Once
	valid *valid.Validate
}

var _ binding.StructValidator = &defaultValidator{}

func (v *defaultValidator) ValidateStruct(obj interface{}) error {
	if kindOfData(obj) == reflect.Struct {

		v.lazyinit()

		if err := v.valid.Struct(obj); err != nil {
			return err
		}
	}

	return nil
}

func (v *defaultValidator) Engine() interface{} {
	v.lazyinit()
	return v.valid
}

func (v *defaultValidator) lazyinit() {
	v.once.Do(func() {
		v.valid = valid.New()
		v.valid.SetTagName("binding")

		// add any custom validations etc. here
	})
}

func kindOfData(data interface{}) reflect.Kind {
	value := reflect.ValueOf(data)
	valueType := value.Kind()

	if valueType == reflect.Ptr {
		valueType = value.Elem().Kind()
	}
	return valueType
}
