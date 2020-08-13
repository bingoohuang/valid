package valid

import "reflect"

// FieldLevel contains all the information and helper functions
// to valid a field.
type FieldLevel interface {
	// returns the top level struct, if any.
	Top() reflect.Value

	// returns the current fields parent struct, if any or
	// the comparison value if called 'VarWithValue'
	Parent() reflect.Value

	// returns current field for validation.
	Field() reflect.Value

	// returns the field's name with the tag
	// name taking precedence over the fields actual name.
	FieldName() string

	// returns the struct field's name.
	StructFieldName() string

	// returns param for validation against current field.
	Param() string

	// GetTag returns the current validations tag name.
	GetTag() string

	// ExtractType gets the actual underlying type of field value.
	// It will dive into pointers, customTypes and return you the
	// underlying value and it's kind.
	ExtractType(field reflect.Value) (value reflect.Value, kind reflect.Kind, nullable bool)

	// GetStructField traverses the parent struct to retrieve a specific field denoted by the provided namespace
	// in the param and returns the field, field kind, if it's a nullable type and whether is was successful in retrieving
	// the field at all.
	//
	// NOTE: when not successful ok will be false, this can happen when a nested struct is nil and so the field
	// could not be retrieved because it didn't exist.
	GetStructField() (field reflect.Value, fieldKind reflect.Kind, nullable bool, found bool)

	// GetStructFieldOKAdvanced is the same as GetStructFieldOK except that it accepts the parent struct to start looking for
	// the field and namespace allowing more extensibility for validators.
	GetStructFieldV2(val reflect.Value, namespace string) (
		field reflect.Value, fieldKind reflect.Kind, nullable bool, found bool)
}

var _ FieldLevel = new(validator)

// Field returns current field for validation.
func (v *validator) Field() reflect.Value {
	return v.flField
}

// FieldName returns the field's name with the tag.
// name taking precedence over the fields actual name.
func (v *validator) FieldName() string {
	return v.cf.altName
}

// GetTag returns the current validations tag name.
func (v *validator) GetTag() string {
	return v.ct.tag
}

// StructFieldName returns the struct field's name.
func (v *validator) StructFieldName() string {
	return v.cf.name
}

// Param returns param for validation against current field.
func (v *validator) Param() string {
	return v.ct.param
}

// GetStructFieldOK returns Param returns param for validation against current field.
func (v *validator) GetStructField() (reflect.Value, reflect.Kind, bool, bool) {
	return v.getStructFieldOKInternal(v.slflParent, v.ct.param)
}

// GetStructFieldOKAdvanced is the same as GetStructFieldOK except that it accepts the parent struct to start looking for
// the field and namespace allowing more extensibility for validators.
func (v *validator) GetStructFieldV2(val reflect.Value, namespace string) (
	field reflect.Value, fieldKind reflect.Kind, nullable bool, found bool) {
	return v.getStructFieldOKInternal(val, namespace)
}
