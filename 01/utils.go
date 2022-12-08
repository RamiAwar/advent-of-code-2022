package main

import "fmt"

func ParseInt(val string) (int64, error) {
	var intVal int64
	factor := int64(1)
	for i := len(val) - 1; i >= 0; i-- {
		c := val[i]
		if c >= '0' && c <= '9' {
			intVal += int64(c-'0') * factor
		} else {
			return intVal, fmt.Errorf("bad int: %q", val)
		}
		factor *= 10
	}
	return intVal, nil
}
