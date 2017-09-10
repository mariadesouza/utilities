package structutil

import (
	"reflect"
)

func StructToMap(v interface{}) *map[string]interface{} {

	//modelRefType := v.Type()
	structmap := make(map[string]interface{})

	s := reflect.ValueOf(v).Elem() //element values
	typeofT := s.Type()            // type of Element - name of element
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		elementName := typeofT.Field(i).Name
		//typeOfStructElement := f.Type()
		valueOfElement := f.Interface()
		structmap[elementName] = valueOfElement
	}
	return &structmap
}
