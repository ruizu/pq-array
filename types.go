package pqarray

import (
	"strconv"
)

type Int64Slice []int64

func (is *Int64Slice) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	val := string(value.([]byte))
	parsed := parse(val)
	converted := []int64{}
	for _, v := range parsed {
		pv, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return nil
		}
		converted = append(converted, pv)
	}
	(*is) = Int64Slice(converted)
	return nil
}

type Float64Slice []float64

func (fs *Float64Slice) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	val := string(value.([]byte))
	parsed := parse(val)
	converted := []float64{}
	for _, v := range parsed {
		pv, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return nil
		}
		converted = append(converted, pv)
	}
	(*fs) = Float64Slice(converted)
	return nil
}

type StringSlice []string

func (ss *StringSlice) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	val := string(value.([]byte))
	(*ss) = StringSlice(parse(val))
	return nil
}
