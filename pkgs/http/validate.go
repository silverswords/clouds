package core

import (
	"reflect"
	"sync"

	"gopkg.in/go-playground/validator.v8"
)

var (
	// Validator -
	Validator = &defaultValidator{}
)

type defaultValidator struct {
	once     sync.Once
	validate *validator.Validate
}

// Validate validate data
func Validate(obj interface{}) error {
	return Validator.ValidateStruct(obj)
}

// ValidateStruct receives any kind of type, but only performed struct or pointer to struct type.
func (v *defaultValidator) ValidateStruct(obj interface{}) error {
	value := reflect.ValueOf(obj)
	valueType := value.Kind()
	if valueType == reflect.Ptr {
		valueType = value.Elem().Kind()
	}

	if valueType == reflect.Struct {
		v.lazyinit()
		if err := v.validate.Struct(obj); err != nil {
			return err
		}
	}

	return nil
}

func (v *defaultValidator) lazyinit() {
	v.once.Do(func() {
		config := &validator.Config{TagName: "zeit"}
		v.validate = validator.New(config)
	})
}
