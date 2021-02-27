package csv

import (
	"encoding/csv"
	"io"
	"reflect"
)

type Assigner interface {
	Assign(keys []string, values []string)
}

var AssignerType = reflect.TypeOf((*Assigner)(nil)).Elem()

func readAll(reader *csv.Reader, pointer interface{}) error {
	var t reflect.Type

	if t = reflect.TypeOf(pointer); t.Kind() != reflect.Ptr {
		panic("Not a pointer")
	}

	if t = t.Elem(); t.Kind() != reflect.Slice {
		panic("Not a pointer to slice")
	}

	if t = t.Elem(); !reflect.PtrTo(t).Implements(AssignerType) {
		panic("AssignerType not implemented")
	}

	var keys []string
	var out = reflect.ValueOf(pointer).Elem()

	for {
		values, err := reader.Read()

		if err == io.EOF {
			return nil
		} else if err != nil {
			return err
		} else if keys == nil {
			keys = values
			continue
		}

		ptr := reflect.New(t)
		ptr.Interface().(Assigner).Assign(keys, values)

		out.Set(reflect.Append(out, ptr.Elem()))
	}

	return nil
}

func ReadAll(reader io.Reader, pointer interface{}) error {
	return readAll(csv.NewReader(reader), pointer)
}
