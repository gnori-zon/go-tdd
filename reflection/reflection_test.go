package reflection

import (
	"slices"
	"testing"
)

type StructWithOneStringField struct {
	StrField string
}

type StructWithTwoStringField struct {
	FirstStrField  string
	SecondStrField string
}

type StructWithTwoStringFieldAndOneIntField struct {
	FirstStrField  string
	SecondStrField string
	IntField       int
}

type StructWithTwoNestedStructWithTwoStringFields struct {
	FirstStruct  StructWithTwoStringField
	SecondStruct StructWithTwoStringField
}

type StructWithPointer struct {
	StructWithTwoStringField *StructWithTwoStringField
}

type StructWithDoublePointer struct {
	StructWithTwoStringField **StructWithTwoStringField
}

type StructWithSliceWithStructWithTwoStrings struct {
	StructSlice []StructWithTwoStringField
}

type StructWithSliceWithStrings struct {
	Slice []string
}

type StructWithTwoStringsArray struct {
	Array [2]string
}

type StructWithPointerStringSlice struct {
	SlicePtr *[]string
}

type StructWithPointerPointerStructsSlice struct {
	SlicePtr *[]*StructWithTwoStringField
}

var pointerStruct = &StructWithTwoStringField{FirstStrField: "first", SecondStrField: "second"}

func TestWalk(t *testing.T) {
	cases := []struct {
		name          string
		input         interface{}
		expectedCalls []string
	}{
		{
			name:          "struct with one string field",
			input:         StructWithOneStringField{StrField: "str"},
			expectedCalls: []string{"str"},
		},
		{
			name:          "struct with two string fields",
			input:         StructWithTwoStringField{FirstStrField: "first", SecondStrField: "second"},
			expectedCalls: []string{"first", "second"},
		},
		{
			name:          "struct with two string fields and one int field",
			input:         StructWithTwoStringFieldAndOneIntField{FirstStrField: "first", SecondStrField: "second", IntField: 42},
			expectedCalls: []string{"first", "second"},
		},
		{
			name: "struct with two nest struct fields with 2 string fields",
			input: StructWithTwoNestedStructWithTwoStringFields{
				FirstStruct:  StructWithTwoStringField{FirstStrField: "1first", SecondStrField: "1second"},
				SecondStruct: StructWithTwoStringField{FirstStrField: "2first", SecondStrField: "2second"},
			},
			expectedCalls: []string{"1first", "1second", "2first", "2second"},
		},
		{
			name:          "pinter struct with two string fields",
			input:         &StructWithTwoStringField{FirstStrField: "first", SecondStrField: "second"},
			expectedCalls: []string{"first", "second"},
		},
		{
			name:          "pointer struct with nested pointer struct with two string fields",
			input:         &StructWithPointer{&StructWithTwoStringField{FirstStrField: "first", SecondStrField: "second"}},
			expectedCalls: []string{"first", "second"},
		},
		{
			name:          "struct with nested double pointer struct with two string fields",
			input:         &StructWithDoublePointer{&pointerStruct},
			expectedCalls: []string{"first", "second"},
		},
		{
			name:          "struct with slice with structs with two string fields",
			input:         StructWithSliceWithStructWithTwoStrings{StructSlice: []StructWithTwoStringField{{FirstStrField: "first", SecondStrField: "second"}}},
			expectedCalls: []string{"first", "second"},
		},
		{
			name:          "struct with slice with strings",
			input:         StructWithSliceWithStrings{[]string{"first", "second"}},
			expectedCalls: []string{"first", "second"},
		},
		{
			name:          "struct with array with strings",
			input:         StructWithTwoStringsArray{[2]string{"first", "second"}},
			expectedCalls: []string{"first", "second"},
		},
		{
			name:          "slice struct with one struct with two string fields",
			input:         []StructWithTwoStringField{{FirstStrField: "first", SecondStrField: "second"}},
			expectedCalls: []string{"first", "second"},
		},
		{
			name:          "struct with pointer string slice",
			input:         StructWithPointerStringSlice{&[]string{"first4", "second4"}},
			expectedCalls: []string{"first4", "second4"},
		},
		{
			name:          "struct with pointer pointer structs slice",
			input:         StructWithPointerPointerStructsSlice{&[]*StructWithTwoStringField{{FirstStrField: "first6", SecondStrField: "second6"}}},
			expectedCalls: []string{"first6", "second6"},
		},
		{
			name:          "map with strings values",
			input:         map[string]string{"1": "first", "2": "second"},
			expectedCalls: []string{"first", "second"},
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			var got []string
			Walk(testCase.input, func(input string) {
				got = append(got, input)
			})

			assertContains(t, got, testCase.expectedCalls)
		})
	}

	t.Run("walk chan with struct with two string fields", func(t *testing.T) {
		aChannel := make(chan StructWithTwoStringField)

		go func() {
			aChannel <- StructWithTwoStringField{FirstStrField: "first1", SecondStrField: "second1"}
			aChannel <- StructWithTwoStringField{FirstStrField: "first2", SecondStrField: "second2"}
			close(aChannel)
		}()

		want := []string{"first1", "second1", "first2", "second2"}
		var got []string
		Walk(aChannel, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, want)
	})

	t.Run("walk chan with ints", func(t *testing.T) {
		aChannel := make(chan int)

		go func() {
			aChannel <- 1
			aChannel <- 2
			close(aChannel)
		}()

		var want []string
		var got []string
		Walk(aChannel, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, want)
	})

	t.Run("walk func structs with two string fields supplier", func(t *testing.T) {
		supplier := func() (StructWithTwoStringField, StructWithTwoStringField) {
			return StructWithTwoStringField{FirstStrField: "first1", SecondStrField: "second1"}, StructWithTwoStringField{FirstStrField: "first2", SecondStrField: "second2"}
		}
		want := []string{"first1", "second1", "first2", "second2"}
		var got []string
		Walk(supplier, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, want)
	})

	t.Run("walk func int supplier", func(t *testing.T) {
		supplier := func() int {
			return 1
		}
		var want []string
		var got []string
		Walk(supplier, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, want)
	})

	t.Run("walk func with args strings supplier", func(t *testing.T) {
		supplier := func(argument string) (string, string) {
			return "1", "2"
		}
		var want []string
		var got []string
		Walk(supplier, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, want)
	})
}

func assertContains(t testing.TB, got, want []string) {
	t.Helper()
	if len(got) != len(want) {
		t.Fatalf("got length: %d; want length: %d", len(got), len(want))
	}
	for _, item := range want {
		if !slices.Contains(got, item) {
			t.Errorf("got: %v; want contains: %v", got, item)
		}
	}
}
