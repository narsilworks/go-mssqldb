package mssql

import (
	"strconv"
)

type ExtParam struct {
	Value interface{}
	Size  int
	Scale uint8
	Prec  uint8
}

// NewExtParam creates a new ExtParam struct with the following arguments:
//
//	value - value for the query
//	opts - an options list that should be in the following sequence:
//		size - should be a string or integer
//		scale - should ba a string or integer
//		precision - should be a string or integer
func NewExtParam(value interface{}, opts ...interface{}) ExtParam {

	var (
		size  int
		scale uint8
		prec  uint8
	)

	// Process options
	if len(opts) == 1 {
		switch t := opts[0].(type) {
		case string:
			if size64, err := strconv.ParseInt(t, 10, 32); err == nil {
				size = int(size64)
			}
		case int:
			size = t
		default:
			size = -1
		}
	}

	if len(opts) == 2 {
		switch t := opts[1].(type) {
		case string:
			if size64, err := strconv.ParseInt(t, 10, 32); err == nil {
				scale = uint8(size64)
			}
		case int:
			scale = uint8(t)
		case uint8:
			scale = t
		default:
			scale = 0
		}
	}

	if len(opts) == 3 {
		switch t := opts[2].(type) {
		case string:
			if size64, err := strconv.ParseInt(t, 10, 32); err == nil {
				prec = uint8(size64)
			}
		case int:
			prec = uint8(t)
		case uint8:
			prec = t
		default:
			scale = 0
		}
	}

	return ExtParam{
		Value: value,
		Size:  size,
		Scale: scale,
		Prec:  prec,
	}
}

func makeExtParam(val ExtParam) (res param) {
	param, err := makeParam(val.Value, nil)
	if err != nil {
		return
	}

	res = param

	if val.Size != -1 {
		res.declSize = val.Size
	}

	if val.Scale != 0 {
		res.ti.Scale = val.Scale
	}

	if val.Prec != 0 {
		res.ti.Prec = val.Prec
	}

	return
}
