package otto

import (
	"strconv"
)

// Number

func numberValueFromNumberArgumentList(argumentList []Value) Value {
	if len(argumentList) > 0 {
		return toValue(toNumber(argumentList[0]))
	}
	return toValue(0)
}

func builtinNumber(call FunctionCall) Value {
	return numberValueFromNumberArgumentList(call.ArgumentList)
}

func builtinNewNumber(self *_object, _ Value, argumentList []Value) Value {
	return toValue(self.runtime.newNumber(numberValueFromNumberArgumentList(argumentList)))
}

func builtinNumber_toFixed(call FunctionCall) Value {
	precision := toIntegerFloat(call.Argument(0))
	if call.This.IsNaN() {
		return toValue("NaN")
	}
	if 0 > precision {
		panic(newRangeError("RangeError: toFixed() precision must be greater than 0"))
	}
	return toValue(strconv.FormatFloat(toFloat(call.This), 'f', int(precision), 64))
}

func builtinNumber_toExponential(call FunctionCall) Value {
	precision := float64(-1)
	if value := call.Argument(0); value.IsDefined() {
		precision = toIntegerFloat(value)
		if 0 > precision {
			panic(newRangeError("RangeError: toExponential() precision must be greater than 0"))
		}
	}
	if call.This.IsNaN() {
		return toValue("NaN")
	}
	return toValue(strconv.FormatFloat(toFloat(call.This), 'e', int(precision), 64))
}
