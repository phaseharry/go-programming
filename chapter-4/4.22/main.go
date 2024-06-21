package main

import (
	"errors"
	"fmt"
)

/*
Type switch - switch statements that are based on the value's type assertion.
- type switch cannot use fallthrough. only one case will be hit and then type switch is done processing
*/
func doubler(v interface{}) (string, error) {
	// t will be the value of v but with its associated type enabled
	switch t := v.(type) {
	// for string & bool, we only handle one type per switch case so we can just process it
	// in their respective case statements
	case string:
		return t + t, nil
	case bool:
		if t == true {
			return "truetrue", nil
		}
		return "falsefalse", nil
	// since we're handling both float32 & float64 in this case, we have to check t again to see whether it is a float32 or a float64.
	case float32, float64:
		if f, ok := v.(float64); ok {
			return fmt.Sprint(f * 2), nil
		}
		// if above type assertion fails (not float64, but float32) then will use below case
		return fmt.Sprint(t.(float32) * 2), nil
	case int:
		return fmt.Sprint(t * 2), nil
	case int8:
		return fmt.Sprint(t * 2), nil
	case int16:
		return fmt.Sprint(t * 2), nil
	case int32:
		return fmt.Sprint(t * 2), nil
	case int64:
		return fmt.Sprint(t * 2), nil
	case uint:
		return fmt.Sprint(t * 2), nil
	case uint8:
		return fmt.Sprint(t * 2), nil
	case uint16:
		return fmt.Sprint(t * 2), nil
	case uint32:
		return fmt.Sprint(t * 2), nil
	case uint64:
		return fmt.Sprint(t * 2), nil
	default:
		return "", errors.New("unsupported type passed")
	}
}

func main() {
	res, _ := doubler(-5)
	fmt.Println("-5: ", res)
	res, _ = doubler(5)
	fmt.Println("5: ", res)
	res, _ = doubler("yum")
	fmt.Println("yum: ", res)
	res, _ = doubler(true)
	fmt.Println("true: ", res)
	res, _ = doubler(float32(3.14))
	fmt.Println("3.14: ", res)
}
