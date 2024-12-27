package main

/*
#include <stdlib.h>
#include <stdint.h>
*/
import "C"
import (
	"unsafe"

	"github.com/jeremiasbots/currencyRatesLib/go/bcv"
	"github.com/jeremiasbots/currencyRatesLib/go/utils"
)

//export GetCurrencyValue
func GetCurrencyValue(currency C.int8_t) (*C.char, *C.char) {
	currencyRealValue := int8(currency)
	value, err := bcv.GetCurrencyValue(int8(currencyRealValue))
	valueC := C.CString(value)
	errC := C.CString(utils.ErrAsNil(err))
	defer C.free(unsafe.Pointer(valueC))
	defer C.free(unsafe.Pointer(errC))
	return valueC, errC
}

func main() {}
