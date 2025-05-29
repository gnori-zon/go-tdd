package reflection

import (
	"reflect"
)

func Walk(x interface{}, fn func(input string)) {
	var values []reflect.Value
	values = append(values, reflect.ValueOf(x))

	for len(values) > 0 {
		value := removeFirst(&values)
		value = dereferenceWhileNeeded(value)

		switch value.Kind() {
		case reflect.String:
			fn(value.String())
		case reflect.Struct:
			appendAllStructFields(value, &values)
		case reflect.Slice, reflect.Array:
			appendAllCollectionItems(value, &values)
		case reflect.Map:
			appendAllMapValues(value, &values)
		case reflect.Chan:
			appendAllChanValues(value, &values)
		case reflect.Func:
			appendFuncResultIfNotRequiredArgs(value, &values)
		}
	}
}

func removeFirst(values *[]reflect.Value) reflect.Value {
	value := (*values)[0]
	*values = (*values)[1:]
	return value
}

func dereferenceWhileNeeded(value reflect.Value) reflect.Value {
	for value.Kind() == reflect.Pointer {
		value = value.Elem()
	}
	return value
}

func appendAllStructFields(value reflect.Value, values *[]reflect.Value) {
	numFields := value.NumField()
	for i := 0; i < numFields; i++ {
		field := value.Field(i)
		*values = append(*values, field)
	}
}

func appendAllCollectionItems(value reflect.Value, values *[]reflect.Value) {
	for i := 0; i < value.Len(); i++ {
		*values = append(*values, value.Index(i))
	}
}

func appendAllMapValues(value reflect.Value, values *[]reflect.Value) {
	for _, key := range value.MapKeys() {
		*values = append(*values, value.MapIndex(key))
	}
}

func appendAllChanValues(value reflect.Value, values *[]reflect.Value) {
	for {
		if received, ok := value.Recv(); ok {
			*values = append(*values, received)
		} else {
			break
		}
	}
}

func appendFuncResultIfNotRequiredArgs(value reflect.Value, values *[]reflect.Value) {
	if value.Type().NumIn() == 0 {
		funcResult := value.Call(nil)
		*values = append(*values, funcResult...)
	}
}
