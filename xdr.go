//
// The common part of XDR should be in its own package eventually.
// package xdr
package main

import (
	"fmt"	// P3
	"reflect"
)

// P3 temporary
//
// Running through introspection of structures to be dumped as XDR stream
func Print(p interface{}) {
	var t reflect.Type = reflect.TypeOf(p)
	v := reflect.ValueOf(p)
	// s := reflect.ValueOf(p).Elem()
	switch t.Kind() {
	  case reflect.Struct:
		fmt.Printf("struct\n")
		for i := 0; i < v.NumField(); i++ {
			// vf := v.Field(i)
			fs := t.Field(i)
			fmt.Printf("%d: %s %s\n", i, fs.Name, fs.Type)
		}
	  default:
		fmt.Printf("unknown\n")
	}
}
