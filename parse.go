package gpapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"reflect"
	"strings"
)

func Interface2Values(arg interface{}) (r url.Values, err error) {
	err = nil
	typ := reflect.TypeOf(arg)
	r = url.Values{}
	for i := 0; i < typ.NumField(); i++ {
		fld := typ.Field(i)
		tagstr := fld.Tag.Get("gpapi")
		tagstrA := strings.Split(tagstr, ",")
		key := fld.Name
		val := "optional"
		for _, t := range tagstrA {
			if t == "required" {
				val = "required"
			} else if t != "" && t != "optional" {
				key = t
			}
		}
		r.Set(key, val)
	}
	r.Set("Action", strings.TrimSuffix(typ.Name(), "Arg"))
	return
}

func ValidateRequest(arg interface{}) (r url.Values, err error) {
	typ := reflect.TypeOf(arg)
	val := reflect.ValueOf(arg)
	r = url.Values{}
	for i := 0; i < typ.NumField(); i++ {
		fld := typ.Field(i)
		var valdata string
		var valdataA []string
		switch val.Field(i).Kind() {
		case reflect.Array, reflect.Slice:
			valdataA = []string{}
			for ii := 0; ii < val.Field(i).Len(); ii++ {
				valdataA = append(valdataA, val.Field(i).Index(i).String())
			}
		case reflect.String:
			valdata = val.Field(i).String()
		}
		tagstr := fld.Tag.Get("gpapi")
		tagstrA := strings.Split(tagstr, ",")
		for _, t := range tagstrA {
			if t == "required" && valdata == "" && len(valdataA) == 0 {
				err = fmt.Errorf("%s is required: %+v", fld.Name, arg)
				return
			}
		}
		if valdata != "" {
			key := fld.Name
			if len(tagstrA) != 0 && tagstrA[0] != "" && tagstrA[0] != "required" {
				key = tagstrA[0]
			}
			r.Set(key, valdata)
		}
		for ii, val := range valdataA {
			r.Set(fmt.Sprintf("%s.%d", fld.Name, ii+1), val)
		}
	}
	r.Set("Action", strings.TrimSuffix(typ.Name(), "Arg"))
	return
}

func ParseResponse(rd io.Reader, result interface{}) (err error) {
	var vv map[string]interface{}
	dec := json.NewDecoder(rd)
	err = dec.Decode(&vv)
	for k, v := range vv {
		if strings.HasSuffix(k, "Response") {
			rd2, wr := io.Pipe()
			enc := json.NewEncoder(wr)
			go enc.Encode(v)
			dec2 := json.NewDecoder(rd2)
			return dec2.Decode(result)
		}
	}
	rd2, wr := io.Pipe()
	enc := json.NewEncoder(wr)
	go enc.Encode(vv)
	dec2 := json.NewDecoder(rd2)
	return dec2.Decode(result)
}
