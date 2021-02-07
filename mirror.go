//Package mirror implements run-time 'mirror' reflection, allowing struct fields to be passed to functions in a type-safe manner.
package mirror

import (
	"fmt"
	"reflect"
	"strconv"
)

//Field of a struct.
type Field struct {
	reflect.StructField

	Parent int
}

type key struct {
	reflect.Type

	offset int
	groups string
}

//Type is a mirrored type used to store a field-mapping.
type Type struct {
	//list stores every field in the type in a flat sequential slice.
	list []Field

	//maps from a key to the index in the list.
	maps map[key]int

	//size of the types offset range & current offset value.
	size map[reflect.Type]int
}

func (t *Type) deserialize(value reflect.Value) (k key) {
	rtype := value.Type()
	k.Type = rtype

	//Deserialize the offset out of the type.
	switch rtype.Kind() {

	case reflect.String:
		k.offset, _ = strconv.Atoi(value.String())

	case reflect.Int:
		k.offset = int(value.Int())

	case reflect.Struct:
		k.offset = 0
		for i := 0; i < rtype.NumField(); i++ {
			k.groups += "." + fmt.Sprint(t.deserialize(value.Field(i)).offset)
		}

	default:
		panic("mirror.Type.Field unsupported struct-field type: " + rtype.String())
	}

	return
}

func (t *Type) serialise(parent int, offset int, value reflect.Value) key {
	rtype := value.Type()
	group := ""

	//Serialise the offset into the type somehow.
	switch value.Type().Kind() {

	case reflect.String:
		//convert offset to string representation.
		value.SetString(fmt.Sprint(offset))

	case reflect.Int:
		//store offset directly as an int.
		value.SetInt(int64(offset))

	case reflect.Struct:
		//caclulate children.
		for i := 0; i < rtype.NumField(); i++ {
			field := rtype.Field(i)

			index := len(t.list)
			t.list = append(t.list, Field{field, parent})

			var offset = t.size[field.Type]
			t.size[field.Type] = offset + 1

			group += "." + fmt.Sprint(offset)

			t.maps[t.serialise(index, offset, value.Field(i))] = index
		}
	default:
		panic("mirror: unsupported struct-field type " + value.Type().String())
	}

	if group != "" {
		offset = 0
	}

	return key{rtype, offset, group}
}

//Reflect reflects the fields of the value onto the mirror.
func (t *Type) Reflect(value interface{}) {
	var rvalue = reflect.ValueOf(value)
	if rvalue.Type().Kind() != reflect.Ptr || rvalue.Elem().Type().Kind() != reflect.Struct {
		panic("mirror.Type.Reflect must be passed a struct pointer")
	}

	//reset the mirror.
	t.list = nil
	t.maps = make(map[key]int)
	t.size = make(map[reflect.Type]int)

	rvalue = rvalue.Elem()
	rtype := rvalue.Type()
	for i := 0; i < rtype.NumField(); i++ {
		field := rtype.Field(i)

		index := len(t.list)
		t.list = append(t.list, Field{field, -1})

		var offset = t.size[field.Type]
		t.size[field.Type] = offset + 1

		t.maps[t.serialise(index, offset, rvalue.Field(i))] = index
	}
}

//Field returns the value's field, if the value was not initialised
//with this Type's call to Reflect, Field panics or returns an undefined field.
func (t Type) Field(field interface{}) Field {
	return t.list[t.maps[t.deserialize(reflect.ValueOf(field))]]
}

//Path returns the field's path in Go syntax rooted at the mirrored type, if the value was not initialised
//with this Type's call to Reflect, Path panics or returns an undefined field.
func (t Type) Path(field interface{}) string {
	var fields = []Field{t.list[t.maps[t.deserialize(reflect.ValueOf(field))]]}

	for f := fields[len(fields)-1]; f.Parent != -1; f = t.list[f.Parent] {
		fields = append(fields, t.list[f.Parent])
	}

	var path string
	for i := len(fields) - 1; i >= 0; i-- {
		path += "." + fields[i].Name
	}

	return path
}
