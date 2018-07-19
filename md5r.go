package md5r

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"reflect"
	"strings"
)

// String generates and returns a md5 hash of the given string
func String(ins ...string) string {
	h := md5.New()
	for _, in := range ins {
		h.Write([]byte(in))
	}
	return hex.EncodeToString(h.Sum(nil))
}

// Struct generates and returns a md5 hash of the given struct's values excluding any fields with the ignore:"md5r"
func Struct(ins ...interface{}) string {
	h := md5.New()
	// fmt.Println(ins)
	for _, in := range ins {
		// fmt.Println(reflect.TypeOf(in))
		t := reflect.TypeOf(in)
		v := reflect.ValueOf(in)

		if reflect.ValueOf(in).Kind() == reflect.Ptr {
			t = t.Elem()
			v = v.Elem()
		}

		for i := 0; i < t.NumField(); i++ {
			if !strings.Contains(t.Field(i).Tag.Get("ignore"), "md5") {
				h.Write([]byte(fmt.Sprintf("%v", v.Field(i))))
			}
		}
	}
	return hex.EncodeToString(h.Sum(nil))
}

func Array(input []interface{}) string {
	h := md5.New()
	// fmt.Println(ins)
	for _, in := range input {
		fmt.Println(reflect.TypeOf(in))
		t := reflect.TypeOf(in)
		v := reflect.ValueOf(in)

		if reflect.ValueOf(in).Kind() == reflect.Ptr {
			t = t.Elem()
			v = v.Elem()
		}

		for i := 0; i < t.NumField(); i++ {
			if !strings.Contains(t.Field(i).Tag.Get("ignore"), "md5") {
				h.Write([]byte(fmt.Sprintf("%v", v.Field(i))))
			}
		}
	}
	return hex.EncodeToString(h.Sum(nil))
}
