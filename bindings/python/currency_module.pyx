cdef extern from "currencyRatesLib.h":
    ctypedef struct GetCurrencyValue_return:
        char* r0
        char* r1
    GetCurrencyValue_return GetCurrencyValue(char currency)

def get_currency_value(currency: int):
    struct = GetCurrencyValue(currency)
    return struct