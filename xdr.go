//
// The common part of XDR should be in its own package eventually.
// package xdr
package main

import (
	"fmt"	// P3
	"reflect"
)

type XDR struct {
	buf []byte
}
// XXX try to create a special interface XDRe, use here instead of interface{}
func (x *XDR) Encode(p interface{}) {
	var t reflect.Type = reflect.TypeOf(p)
	var v reflect.Value = reflect.ValueOf(p)
	switch t.Kind() {
	  case reflect.Struct:
		fmt.Printf("Encode struct\n")
		for i := 0; i < v.NumField(); i++ {
			fs := t.Field(i)
			fmt.Printf("%d: %s %s\n", i, fs.Name, fs.Type)
			vf := v.Field(i)

			// The documentation for the package reflect tells us
			// "Inteface(v Value) panics if the Value was obtained
			// by accessing unexported struct fields."
			//   var vi interface{} = vf.Interface()
			//   x.Encode(vi)
			// Instead, we pluck concrete values with a switch.

			var bv []byte
			switch fs.Type.Kind() {
			  case reflect.Int:	// we treat this as Int32
				vi := vf.Int()
				fmt.Printf("Encode field %s int %d\n", fs.Name, int(vi))
				var a4 *[4]byte = new([4]byte)
				a4[0] = byte(vi >> 24)
				a4[1] = byte(vi >> 16)
				a4[2] = byte(vi >> 8)
				a4[3] = byte(vi)
				bv = a4[0:4]
			  case reflect.String:
				fmt.Printf("Encode field %s string\n", fs.Name)
				vs := vf.String()
				// bv = []byte(vs)
				var l int = len(vs)
				bv = make([]byte, 4 + ((l + 3) & ^3))
				bv[0] = byte(l >> 24)
				bv[1] = byte(l >> 16)
				bv[2] = byte(l >> 8)
				bv[3] = byte(l)
				copy(bv[4:4+l], vs)
			  default:
				fmt.Printf("Encode field %s unknown\n", fs.Name)
				bv = nil
			}

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
	  default:
		// As it happens, currently we call Encode() on structs only
		// XXX then panic() here
		fmt.Printf("Encode unknown\n")
	}

	// XXX
	// var a *[10]byte = new([10]byte)
	// x.buf = a[0:10]
}
func (x *XDR) Fetch() ([]byte) {
	return x.buf
}
