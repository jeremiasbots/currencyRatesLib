from cffi import FFI
import os

ffi = FFI()
with open("./currencyRatesLib.h", "r") as f:
    contenido = f.read()
ffi.cdef(contenido)
lib = ffi.dlopen("./currencyRatesLib.dll")
