//
// The common part of XDR should be in its own package eventually.
// package xdr
package main

import (
	"reflect"
)

type XDR struct {
	buf []byte
}

func (x *XDR) Encode(p XDRe) {
	var t reflect.Type = reflect.TypeOf(p)
	var v reflect.Value = reflect.ValueOf(p)
	switch t.Kind() {
	  case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			fs := t.Field(i)
			vf := v.Field(i)

			// The documentation for the package reflect tells us
			// "Inteface(v Value) panics if the Value was obtained
			// by accessing unexported struct fields."
			//   var vi interface{} = vf.Interface()
			//   x.Encode(vi)
			// Instead, we pluck concrete values with a switch.

			switch fs.Type.Kind() {
			  case reflect.Int:	// we treat this as Int32
				vi := vf.Int()
				x.EncodeInt(int(vi))
			  case reflect.Int64:
				vi := vf.Int()
				x.EncodeInt64(vi)
			  case reflect.String:
				vs := vf.String()
				x.EncodeString(vs)
			// case reflect.Struct:
			// Same provlem as with taking Interface() above
			//	x.Encode(.....)
			  default:
				;
			}
		}
	  default:
		// As it happens, currently we call Encode() on structs only
		// XXX then panic() here
		;
	}
}

func (x *XDR) EncodeInt(v int) {
	var a4 *[4]byte = new([4]byte)
	for j := 0; j < 4; j++ {
		a4[j] = byte(v >> uint((3-j)*8))
	}
	x.append(a4[0:4])
}

func (x *XDR) EncodeInt64(v int64) {
	var a8 *[8]byte = new([8]byte)
	for j := 0; j < 8; j++ {
		a8[j] = byte(v >> uint((7-j)*8))
	}
	x.append(a8[0:8])
}

// "In Go, a string is in effect a read-only slice of bytes."
func (x *XDR) EncodeString(v string) {
	var l int = len(v)
	var bv []byte = make([]byte, 4 + ((l + 3) & ^3))
	bv[0] = byte(l >> 24)
	bv[1] = byte(l >> 16)
	bv[2] = byte(l >> 8)
	bv[3] = byte(l)
	copy(bv[4:4+l], v)
	x.append(bv)
}

func (x *XDR) append(bv []byte) {
	// XXX Temporary while we apply expanding slice
	// with array reallocation. See
	// http://blog.golang.org/go-slices-usage-and-internals
	//	buflen := len(x.buf)
	//	newbuf := new([buflen + len(bv)]byte)
	//	if buflen != 0 {
	//		copy(newbuf, x.buf)
	//	}
	//	copy(newbuf[buflen:], bv)
	//	x.buf = newbuf
	// Or, even better
	// https://golang.org/doc/effective_go.html#append
	x.buf = append(x.buf, bv...)
}

func (x *XDR) Fetch() ([]byte) {
	return x.buf
}

type XDRe interface {
	XDRencode(x *XDR)
}
