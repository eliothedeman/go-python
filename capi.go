package python

/*
#include "Python.h"
*/
import (
	"C"
	"fmt"
	"encoding/binary"
)
import "reflect"

// format units
var (
	Format_Units = map[string]valueBuilder{
		"s":   vb_String,
		"s#":  vb_String,
		"s*":  vb_String,
		"z":   vb_String,
		"z#":  vb_String,
		"z*":  vb_String,
		"u":   vb_String,
		"u#":  vb_String,
		"es":  vb_String,
		"et":  vb_String,
		"es#": vb_String,
		"et#": vb_String,
	}
)

// valueBuilder creates a python object from an interface
type valueBuilder func(interface{}) *PyObject

// vb_String return the string value of i
func vb_String(i interface{}) *PyObject {
	// stop short on nil
	if i == nil {
		return &PyString{}
	}

	var s string
	switch t := i.(type) {
	// if we are looking at a byte slice, treat it as a string
	case []byte:
		s = string(t)
	// everything else should be converted by fmt correctly
	default:
		s = fmt.Sprintf("%v", i)
	}
	return PyString_FromString(s)
}

// vb_TinyInt return a pyobject containing a TinyInt
func vb_Int(i interface{}) *PyObject {
	var v int64
	switch t := i.(type) {
	case []byte:
		
		// not sure exaclty what to do here. I am guessing just treat the first 8 bytes a a 64 bit int
		if len(t) < 8 {
			t = append(t,make([]byte, 8 - len(t))
		}
		v,_ = binary.Varint(t[:8])
	case string:
		b := []byte(t)

		// not sure exaclty what to do here. I am guessing just treat the first 8 bytes a a 64 bit int
		if len(t) < 8 {
			b = append(b,make([]byte, 8 - len(t))
		}
		v,_ = binary.Varint(b[:8])

	// get the value as an int see if there is a way to do this without using reflection
	case uint8, int8, int16, uint16, int32, uint32, int, uint, int64, uint64:
		v = reflect.ValueOf(t).Int()

	// all other types that cant be put into an int type, assume we are ok with a zero value
	default:
		v = 0
	}

	return PyLong_FromLong(v)
}

// PyObject* Py_BuildValue(const char *format, ...)
// Return value: New reference.
func Py_BuildValue(format string, args ...interface{}) *PyObject {
	return nil
}

// PyMethodDef
// ml_name	char *	name of the method
// ml_meth	PyCFunction	pointer to the C implementation
// ml_flags	int	flag bits indicating how the call should be constructed
// ml_doc	char *	points to the contents of the docstring
type PyMethodDef struct {
	Name  string // name of the method
	Meth  func(self, args *PyObject) *PyObject
	Flags MethodDefFlags
	Doc   string
}

type MethodDefFlags int

const (
	MethVarArgs  MethodDefFlags = C.METH_VARARGS
	MethKeyWords                = C.METH_KEYWORDS
	MethNoArgs                  = C.METH_NOARGS
	MethO                       = C.METH_O
	MethOldArgs                 = C.METH_OLDARGS
	MethClass                   = C.METH_CLASS
	MethStatic                  = C.METH_STATIC
	MethCoexist                 = C.METH_COEXIST
)
