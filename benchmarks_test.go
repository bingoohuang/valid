package valid

import (
	"bytes"
	sql "database/sql/driver"
	"testing"
	"time"
)

func BenchmarkFieldSuccess(b *testing.B) {
	valid := New()
	s := "1"

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = valid.Var(&s, "len=1")
	}
}

func BenchmarkFieldSuccessParallel(b *testing.B) {
	valid := New()
	s := "1"

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = valid.Var(&s, "len=1")
		}
	})
}

func BenchmarkFieldFailure(b *testing.B) {
	valid := New()
	s := "12"

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = valid.Var(&s, "len=1")
	}
}

func BenchmarkFieldFailureParallel(b *testing.B) {
	valid := New()
	s := "12"

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = valid.Var(&s, "len=1")
		}
	})
}

func BenchmarkFieldArrayDiveSuccess(b *testing.B) {
	valid := New()
	m := []string{"val1", "val2", "val3"}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		_ = valid.Var(m, "required,dive,required")
	}
}

func BenchmarkFieldArrayDiveSuccessParallel(b *testing.B) {
	valid := New()
	m := []string{"val1", "val2", "val3"}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = valid.Var(m, "required,dive,required")
		}
	})
}

func BenchmarkFieldArrayDiveFailure(b *testing.B) {
	valid := New()
	m := []string{"val1", "", "val3"}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = valid.Var(m, "required,dive,required")
	}
}

func BenchmarkFieldArrayDiveFailureParallel(b *testing.B) {
	valid := New()
	m := []string{"val1", "", "val3"}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = valid.Var(m, "required,dive,required")
		}
	})
}

func BenchmarkFieldMapDiveSuccess(b *testing.B) {
	valid := New()
	m := map[string]string{"val1": "val1", "val2": "val2", "val3": "val3"}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		_ = valid.Var(m, "required,dive,required")
	}
}

func BenchmarkFieldMapDiveSuccessParallel(b *testing.B) {
	valid := New()
	m := map[string]string{"val1": "val1", "val2": "val2", "val3": "val3"}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = valid.Var(m, "required,dive,required")
		}
	})
}

func BenchmarkFieldMapDiveFailure(b *testing.B) {
	valid := New()
	m := map[string]string{"": "", "val3": "val3"}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = valid.Var(m, "required,dive,required")
	}
}

func BenchmarkFieldMapDiveFailureParallel(b *testing.B) {
	valid := New()
	m := map[string]string{"": "", "val3": "val3"}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = valid.Var(m, "required,dive,required")
		}
	})
}

func BenchmarkFieldMapDiveWithKeysSuccess(b *testing.B) {
	valid := New()
	m := map[string]string{"val1": "val1", "val2": "val2", "val3": "val3"}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		_ = valid.Var(m, "required,dive,keys,required,endkeys,required")
	}
}

func BenchmarkFieldMapDiveWithKeysSuccessParallel(b *testing.B) {
	valid := New()
	m := map[string]string{"val1": "val1", "val2": "val2", "val3": "val3"}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = valid.Var(m, "required,dive,keys,required,endkeys,required")
		}
	})
}

func BenchmarkFieldMapDiveWithKeysFailure(b *testing.B) {
	valid := New()
	m := map[string]string{"": "", "val3": "val3"}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = valid.Var(m, "required,dive,keys,required,endkeys,required")
	}
}

func BenchmarkFieldMapDiveWithKeysFailureParallel(b *testing.B) {
	valid := New()
	m := map[string]string{"": "", "val3": "val3"}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = valid.Var(m, "required,dive,keys,required,endkeys,required")
		}
	})
}

func BenchmarkFieldCustomTypeSuccess(b *testing.B) {
	valid := New()
	valid.RegisterCustomTypeFunc(ValidateValuerType, (*sql.Valuer)(nil), valuer{})
	val := valuer{
		Name: "1",
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = valid.Var(val, "len=1")
	}
}

func BenchmarkFieldCustomTypeSuccessParallel(b *testing.B) {
	valid := New()
	valid.RegisterCustomTypeFunc(ValidateValuerType, (*sql.Valuer)(nil), valuer{})
	val := valuer{
		Name: "1",
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = valid.Var(val, "len=1")
		}
	})
}

func BenchmarkFieldCustomTypeFailure(b *testing.B) {
	valid := New()
	valid.RegisterCustomTypeFunc(ValidateValuerType, (*sql.Valuer)(nil), valuer{})
	val := valuer{}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = valid.Var(val, "len=1")
	}
}

func BenchmarkFieldCustomTypeFailureParallel(b *testing.B) {
	valid := New()
	valid.RegisterCustomTypeFunc(ValidateValuerType, (*sql.Valuer)(nil), valuer{})
	val := valuer{}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = valid.Var(val, "len=1")
		}
	})
}

func BenchmarkFieldOrTagSuccess(b *testing.B) {
	valid := New()
	s := "rgba(0,0,0,1)"

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = valid.Var(s, "rgb|rgba")
	}
}

func BenchmarkFieldOrTagSuccessParallel(b *testing.B) {
	valid := New()
	s := "rgba(0,0,0,1)"

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = valid.Var(s, "rgb|rgba")
		}
	})
}

func BenchmarkFieldOrTagFailure(b *testing.B) {
	valid := New()
	s := "#000"

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = valid.Var(s, "rgb|rgba")
	}
}

func BenchmarkFieldOrTagFailureParallel(b *testing.B) {
	valid := New()
	s := "#000"

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = valid.Var(s, "rgb|rgba")
		}
	})
}

func BenchmarkStructLevelValidationSuccess(b *testing.B) {
	valid := New()
	valid.RegisterStructValidation(StructValidationTestStructSuccess, TestStruct{})

	tst := TestStruct{
		String: "good value",
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = valid.Struct(tst)
	}
}

func BenchmarkStructLevelValidationSuccessParallel(b *testing.B) {
	valid := New()
	valid.RegisterStructValidation(StructValidationTestStructSuccess, TestStruct{})

	tst := TestStruct{
		String: "good value",
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = valid.Struct(tst)
		}
	})
}

func BenchmarkStructLevelValidationFailure(b *testing.B) {
	valid := New()
	valid.RegisterStructValidation(StructValidationTestStruct, TestStruct{})

	tst := TestStruct{
		String: "good value",
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = valid.Struct(tst)
	}
}

func BenchmarkStructLevelValidationFailureParallel(b *testing.B) {
	valid := New()
	valid.RegisterStructValidation(StructValidationTestStruct, TestStruct{})

	tst := TestStruct{
		String: "good value",
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = valid.Struct(tst)
		}
	})
}

func BenchmarkStructSimpleCustomTypeSuccess(b *testing.B) {
	valid := New()
	valid.RegisterCustomTypeFunc(ValidateValuerType, (*sql.Valuer)(nil), valuer{})

	val := valuer{
		Name: "1",
	}

	type Foo struct {
		Valuer   valuer `valid:"len=1"`
		IntValue int    `valid:"min=5,max=10"`
	}

	validFoo := &Foo{Valuer: val, IntValue: 7}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = valid.Struct(validFoo)
	}
}

func BenchmarkStructSimpleCustomTypeSuccessParallel(b *testing.B) {
	valid := New()
	valid.RegisterCustomTypeFunc(ValidateValuerType, (*sql.Valuer)(nil), valuer{})
	val := valuer{
		Name: "1",
	}

	type Foo struct {
		Valuer   valuer `valid:"len=1"`
		IntValue int    `valid:"min=5,max=10"`
	}
	validFoo := &Foo{Valuer: val, IntValue: 7}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = valid.Struct(validFoo)
		}
	})
}

func BenchmarkStructSimpleCustomTypeFailure(b *testing.B) {
	valid := New()
	valid.RegisterCustomTypeFunc(ValidateValuerType, (*sql.Valuer)(nil), valuer{})

	val := valuer{}

	type Foo struct {
		Valuer   valuer `valid:"len=1"`
		IntValue int    `valid:"min=5,max=10"`
	}
	validFoo := &Foo{Valuer: val, IntValue: 3}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = valid.Struct(validFoo)
	}
}

func BenchmarkStructSimpleCustomTypeFailureParallel(b *testing.B) {
	valid := New()
	valid.RegisterCustomTypeFunc(ValidateValuerType, (*sql.Valuer)(nil), valuer{})

	val := valuer{}

	type Foo struct {
		Valuer   valuer `valid:"len=1"`
		IntValue int    `valid:"min=5,max=10"`
	}
	validFoo := &Foo{Valuer: val, IntValue: 3}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = valid.Struct(valid.Struct(validFoo))
		}
	})
}

func BenchmarkStructFilteredSuccess(b *testing.B) {
	valid := New()

	type Test struct {
		Name     string `valid:"required"`
		NickName string `valid:"required"`
	}

	test := &Test{
		Name: "Joey Bloggs",
	}
	byts := []byte("Name")
	fn := func(ns []byte) bool {
		return !bytes.HasSuffix(ns, byts)
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = valid.StructFiltered(test, fn)
	}
}

func BenchmarkStructFilteredSuccessParallel(b *testing.B) {
	valid := New()

	type Test struct {
		Name     string `valid:"required"`
		NickName string `valid:"required"`
	}

	test := &Test{
		Name: "Joey Bloggs",
	}
	byts := []byte("Name")
	fn := func(ns []byte) bool {
		return !bytes.HasSuffix(ns, byts)
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = valid.StructFiltered(test, fn)
		}
	})
}

func BenchmarkStructFilteredFailure(b *testing.B) {
	valid := New()

	type Test struct {
		Name     string `valid:"required"`
		NickName string `valid:"required"`
	}

	test := &Test{
		Name: "Joey Bloggs",
	}

	byts := []byte("NickName")

	fn := func(ns []byte) bool {
		return !bytes.HasSuffix(ns, byts)
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = valid.StructFiltered(test, fn)
	}
}

func BenchmarkStructFilteredFailureParallel(b *testing.B) {
	valid := New()

	type Test struct {
		Name     string `valid:"required"`
		NickName string `valid:"required"`
	}

	test := &Test{
		Name: "Joey Bloggs",
	}
	byts := []byte("NickName")
	fn := func(ns []byte) bool {
		return !bytes.HasSuffix(ns, byts)
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = valid.StructFiltered(test, fn)
		}
	})
}

func BenchmarkStructPartialSuccess(b *testing.B) {
	valid := New()

	type Test struct {
		Name     string `valid:"required"`
		NickName string `valid:"required"`
	}

	test := &Test{
		Name: "Joey Bloggs",
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = valid.StructPartial(test, "Name")
	}
}

func BenchmarkStructPartialSuccessParallel(b *testing.B) {
	valid := New()

	type Test struct {
		Name     string `valid:"required"`
		NickName string `valid:"required"`
	}

	test := &Test{
		Name: "Joey Bloggs",
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = valid.StructPartial(test, "Name")
		}
	})
}

func BenchmarkStructPartialFailure(b *testing.B) {
	valid := New()

	type Test struct {
		Name     string `valid:"required"`
		NickName string `valid:"required"`
	}

	test := &Test{
		Name: "Joey Bloggs",
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = valid.StructPartial(test, "NickName")
	}
}

func BenchmarkStructPartialFailureParallel(b *testing.B) {
	valid := New()

	type Test struct {
		Name     string `valid:"required"`
		NickName string `valid:"required"`
	}

	test := &Test{
		Name: "Joey Bloggs",
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = valid.StructPartial(test, "NickName")
		}
	})
}

func BenchmarkStructExceptSuccess(b *testing.B) {
	valid := New()

	type Test struct {
		Name     string `valid:"required"`
		NickName string `valid:"required"`
	}

	test := &Test{
		Name: "Joey Bloggs",
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = valid.StructExcept(test, "Nickname")
	}
}

func BenchmarkStructExceptSuccessParallel(b *testing.B) {
	valid := New()

	type Test struct {
		Name     string `valid:"required"`
		NickName string `valid:"required"`
	}

	test := &Test{
		Name: "Joey Bloggs",
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = valid.StructExcept(test, "NickName")
		}
	})
}

func BenchmarkStructExceptFailure(b *testing.B) {
	valid := New()

	type Test struct {
		Name     string `valid:"required"`
		NickName string `valid:"required"`
	}

	test := &Test{
		Name: "Joey Bloggs",
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = valid.StructExcept(test, "Name")
	}
}

func BenchmarkStructExceptFailureParallel(b *testing.B) {
	valid := New()

	type Test struct {
		Name     string `valid:"required"`
		NickName string `valid:"required"`
	}

	test := &Test{
		Name: "Joey Bloggs",
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = valid.StructExcept(test, "Name")
		}
	})
}

func BenchmarkStructSimpleCrossFieldSuccess(b *testing.B) {
	valid := New()

	type Test struct {
		Start time.Time
		End   time.Time `valid:"gtfield=Start"`
	}

	now := time.Now().UTC()
	then := now.Add(time.Hour * 5)
	test := &Test{
		Start: now,
		End:   then,
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = valid.Struct(test)
	}
}

func BenchmarkStructSimpleCrossFieldSuccessParallel(b *testing.B) {
	valid := New()

	type Test struct {
		Start time.Time
		End   time.Time `valid:"gtfield=Start"`
	}

	now := time.Now().UTC()
	then := now.Add(time.Hour * 5)
	test := &Test{
		Start: now,
		End:   then,
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = valid.Struct(test)
		}
	})
}

func BenchmarkStructSimpleCrossFieldFailure(b *testing.B) {
	valid := New()

	type Test struct {
		Start time.Time
		End   time.Time `valid:"gtfield=Start"`
	}

	now := time.Now().UTC()
	then := now.Add(time.Hour * -5)

	test := &Test{
		Start: now,
		End:   then,
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = valid.Struct(test)
	}
}

func BenchmarkStructSimpleCrossFieldFailureParallel(b *testing.B) {
	valid := New()

	type Test struct {
		Start time.Time
		End   time.Time `valid:"gtfield=Start"`
	}

	now := time.Now().UTC()
	then := now.Add(time.Hour * -5)
	test := &Test{
		Start: now,
		End:   then,
	}
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = valid.Struct(test)
		}
	})
}

func BenchmarkStructSimpleCrossStructCrossFieldSuccess(b *testing.B) {
	valid := New()

	type Inner struct {
		Start time.Time
	}

	type Outer struct {
		Inner     *Inner
		CreatedAt time.Time `valid:"eqcsfield=Inner.Start"`
	}

	now := time.Now().UTC()
	inner := &Inner{
		Start: now,
	}
	outer := &Outer{
		Inner:     inner,
		CreatedAt: now,
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = valid.Struct(outer)
	}
}

func BenchmarkStructSimpleCrossStructCrossFieldSuccessParallel(b *testing.B) {
	valid := New()

	type Inner struct {
		Start time.Time
	}

	type Outer struct {
		Inner     *Inner
		CreatedAt time.Time `valid:"eqcsfield=Inner.Start"`
	}

	now := time.Now().UTC()
	inner := &Inner{
		Start: now,
	}
	outer := &Outer{
		Inner:     inner,
		CreatedAt: now,
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = valid.Struct(outer)
		}
	})
}

func BenchmarkStructSimpleCrossStructCrossFieldFailure(b *testing.B) {
	valid := New()
	type Inner struct {
		Start time.Time
	}

	type Outer struct {
		Inner     *Inner
		CreatedAt time.Time `valid:"eqcsfield=Inner.Start"`
	}

	now := time.Now().UTC()
	then := now.Add(time.Hour * 5)

	inner := &Inner{
		Start: then,
	}

	outer := &Outer{
		Inner:     inner,
		CreatedAt: now,
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = valid.Struct(outer)
	}
}

func BenchmarkStructSimpleCrossStructCrossFieldFailureParallel(b *testing.B) {
	valid := New()

	type Inner struct {
		Start time.Time
	}

	type Outer struct {
		Inner     *Inner
		CreatedAt time.Time `valid:"eqcsfield=Inner.Start"`
	}

	now := time.Now().UTC()
	then := now.Add(time.Hour * 5)

	inner := &Inner{
		Start: then,
	}

	outer := &Outer{
		Inner:     inner,
		CreatedAt: now,
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = valid.Struct(outer)
		}
	})
}

func BenchmarkStructSimpleSuccess(b *testing.B) {
	valid := New()
	type Foo struct {
		StringValue string `valid:"min=5,max=10"`
		IntValue    int    `valid:"min=5,max=10"`
	}

	validFoo := &Foo{StringValue: "Foobar", IntValue: 7}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = valid.Struct(validFoo)
	}
}

func BenchmarkStructSimpleSuccessParallel(b *testing.B) {
	valid := New()
	type Foo struct {
		StringValue string `valid:"min=5,max=10"`
		IntValue    int    `valid:"min=5,max=10"`
	}
	validFoo := &Foo{StringValue: "Foobar", IntValue: 7}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = valid.Struct(validFoo)
		}
	})
}

func BenchmarkStructSimpleFailure(b *testing.B) {
	valid := New()
	type Foo struct {
		StringValue string `valid:"min=5,max=10"`
		IntValue    int    `valid:"min=5,max=10"`
	}

	invalidFoo := &Foo{StringValue: "Fo", IntValue: 3}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = valid.Struct(invalidFoo)
	}
}

func BenchmarkStructSimpleFailureParallel(b *testing.B) {
	valid := New()
	type Foo struct {
		StringValue string `valid:"min=5,max=10"`
		IntValue    int    `valid:"min=5,max=10"`
	}

	invalidFoo := &Foo{StringValue: "Fo", IntValue: 3}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = valid.Struct(invalidFoo)
		}
	})
}

func BenchmarkStructComplexSuccess(b *testing.B) {
	valid := New()
	tSuccess := &TestString{
		Required:  "Required",
		Len:       "length==10",
		Min:       "min=1",
		Max:       "1234567890",
		MinMax:    "12345",
		Lt:        "012345678",
		Lte:       "0123456789",
		Gt:        "01234567890",
		Gte:       "0123456789",
		OmitEmpty: "",
		Sub: &SubTest{
			Test: "1",
		},
		SubIgnore: &SubTest{
			Test: "",
		},
		Anonymous: struct {
			A string `valid:"required"`
		}{
			A: "1",
		},
		Iface: &Impl{
			F: "123",
		},
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = valid.Struct(tSuccess)
	}
}

func BenchmarkStructComplexSuccessParallel(b *testing.B) {
	valid := New()
	tSuccess := &TestString{
		Required:  "Required",
		Len:       "length==10",
		Min:       "min=1",
		Max:       "1234567890",
		MinMax:    "12345",
		Lt:        "012345678",
		Lte:       "0123456789",
		Gt:        "01234567890",
		Gte:       "0123456789",
		OmitEmpty: "",
		Sub: &SubTest{
			Test: "1",
		},
		SubIgnore: &SubTest{
			Test: "",
		},
		Anonymous: struct {
			A string `valid:"required"`
		}{
			A: "1",
		},
		Iface: &Impl{
			F: "123",
		},
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = valid.Struct(tSuccess)
		}
	})
}

func BenchmarkStructComplexFailure(b *testing.B) {
	valid := New()
	tFail := &TestString{
		Required:  "",
		Len:       "",
		Min:       "",
		Max:       "12345678901",
		MinMax:    "",
		Lt:        "0123456789",
		Lte:       "01234567890",
		Gt:        "1",
		Gte:       "1",
		OmitEmpty: "12345678901",
		Sub: &SubTest{
			Test: "",
		},
		Anonymous: struct {
			A string `valid:"required"`
		}{
			A: "",
		},
		Iface: &Impl{
			F: "12",
		},
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = valid.Struct(tFail)
	}
}

func BenchmarkStructComplexFailureParallel(b *testing.B) {
	valid := New()
	tFail := &TestString{
		Required:  "",
		Len:       "",
		Min:       "",
		Max:       "12345678901",
		MinMax:    "",
		Lt:        "0123456789",
		Lte:       "01234567890",
		Gt:        "1",
		Gte:       "1",
		OmitEmpty: "12345678901",
		Sub: &SubTest{
			Test: "",
		},
		Anonymous: struct {
			A string `valid:"required"`
		}{
			A: "",
		},
		Iface: &Impl{
			F: "12",
		},
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = valid.Struct(tFail)
		}
	})
}

type TestOneof struct {
	Color string `valid:"oneof=red green"`
}

func BenchmarkOneof(b *testing.B) {
	w := &TestOneof{Color: "green"}
	val := New()
	for i := 0; i < b.N; i++ {
		_ = val.Struct(w)
	}
}

func BenchmarkOneofParallel(b *testing.B) {
	w := &TestOneof{Color: "green"}
	val := New()

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = val.Struct(w)
		}
	})
}
