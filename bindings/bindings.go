package main

/*
#include <stdlib.h>
#include <stdint.h>
*/
import "C"
import (
	"encoding/json"
	"unsafe"

	"github.com/jeremiasbots/currencyRatesLib/go/bcv"
	"github.com/jeremiasbots/currencyRatesLib/go/binance"
	"github.com/jeremiasbots/currencyRatesLib/go/monitordollar"
)

//export BCVGetCurrencyValue
func BCVGetCurrencyValue(currency C.int8_t) *C.char {
	currencyRealValue := int8(currency)
	value, err := bcv.GetCurrencyValue(int8(currencyRealValue))
	if err != nil {
		panic(err)
	}
	finalValueC := C.CString(value)
	return finalValueC
}

//export BCVGetCurrencyValueAsDouble
func BCVGetCurrencyValueAsDouble(currency C.int8_t) C.double {
	currencyRealValue := int8(currency)
	value, err := bcv.GetCurrencyValueAsFloat64(currencyRealValue)
	if err != nil {
		panic(err)
	}
	finalValueC := C.double(value)
	return finalValueC
}

//export BCVGetCurrencyValueAsFloat
func BCVGetCurrencyValueAsFloat(currency C.int8_t) C.float {
	currencyRealValue := int8(currency)
	value, err := bcv.GetCurrencyValueAsFloat32(currencyRealValue)
	if err != nil {
		panic(err)
	}
	finalValueC := C.float(value)
	return finalValueC
}

//export BCVGetDollarValue
func BCVGetDollarValue() *C.char {
	value, err := bcv.GetDollarValue()
	if err != nil {
		panic(err)
	}
	valueCString := C.CString(value)
	return valueCString
}

//export BCVGetDollarValueAsDouble
func BCVGetDollarValueAsDouble() C.double {
	value, err := bcv.GetDollarValueAsFloat64()
	if err != nil {
		panic(err)
	}
	valueCDouble := C.double(value)
	return valueCDouble
}

//export BCVGetDollarValueAsFloat
func BCVGetDollarValueAsFloat() C.float {
	value, err := bcv.GetDollarValueAsFloat32()
	if err != nil {
		panic(err)
	}
	valueCDouble := C.float(value)
	return valueCDouble
}

//export BinanceGetDollarValue
func BinanceGetDollarValue() *C.char {
	value, err := binance.GetDollarValue()
	if err != nil {
		panic(err)
	}
	valueCString := C.CString(value)
	return valueCString
}

//export BinanceGetDollarValueAsDouble
func BinanceGetDollarValueAsDouble() C.double {
	value, err := binance.GetDollarValueAsFloat64()
	if err != nil {
		panic(err)
	}
	valueCDouble := C.double(value)
	return valueCDouble
}

//export BinanceGetDollarValueAsFloat
func BinanceGetDollarValueAsFloat() C.float {
	value, err := binance.GetDollarValueAsFloat32()
	if err != nil {
		panic(err)
	}
	valueCFloat := C.float(value)
	return valueCFloat
}

//export BinanceGetDollarData
func BinanceGetDollarData() *C.char {
	value, err := binance.GetDollarData()
	if err != nil {
		panic(err)
	}
	jsonValue, jsonErr := json.Marshal(value)
	if jsonErr != nil {
		panic(jsonErr)
	}
	valueCString := C.CString(string(jsonValue))
	return valueCString
}

//export MonitorDollarGetDollarValue
func MonitorDollarGetDollarValue() *C.char {
	value, err := monitordollar.GetDollarValue()
	if err != nil {
		panic(err)
	}
	valueCString := C.CString(value)
	return valueCString
}

//export MonitorDollarGetDollarValueAsFloat
func MonitorDollarGetDollarValueAsFloat() C.float {
	value, err := monitordollar.GetDollarValueAsFloat32()
	if err != nil {
		panic(err)
	}
	valueCString := C.float(value)
	return valueCString
}

//export MonitorDollarGetDollarValueAsDouble
func MonitorDollarGetDollarValueAsDouble() C.double {
	value, err := monitordollar.GetDollarValueAsFloat64()
	if err != nil {
		panic(err)
	}
	valueCString := C.double(value)
	return valueCString
}

//export MonitorDollarGetDollarData
func MonitorDollarGetDollarData() *C.char {
	value, err := monitordollar.GetDollarData()
	if err != nil {
		panic(err)
	}
	json, jsonErr := json.Marshal(value)
	if jsonErr != nil {
		panic(jsonErr)
	}
	valueCString := C.CString(string(json))
	return valueCString
}

//export FreeString
func FreeString(str *C.char) {
	defer C.free(unsafe.Pointer(str))
}

func main() {}
