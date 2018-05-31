package schema

import "testing"

func TestValidateFields(t *testing.T) {
	type testStruct struct {
		Name string `json:"name" validate:"string,min=5,max=10"`
	}
	x := testStruct{Name: "dd"}
	err := ValidateFields("devices", x)
	if len(err) < 1 {
		t.Error("should have 1 atleast error")
	}
}

func TestValidateFields1(t *testing.T) {
	type testSubStruct struct {
		SomeField string `validate:"string,min=1,max=100"`
	}
	type testStruct struct {
		T testSubStruct
	}
	x := testStruct{testSubStruct{}}
	err := ValidateFields("devices", x)
	if len(err) < 1 {
		t.Error("Sub struct is not validated")
	}
}

func BenchmarkValidateFields(b *testing.B) {
	type testStruct struct {
		Name string `json:"name" validate:"string,min=1,max=10"`
		Age  int
		Addr string `validate:"string,min=1,max=20"`
	}
	for i := 0; i < b.N; i++ {
		_ = ValidateFields("devices", testStruct{})
	}
}
