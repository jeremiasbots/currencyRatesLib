// @bun
var{fileURLToPath:i}=globalThis.Bun;import{CString as n,dlopen as u,FFIType as r,suffix as c}from"bun:ffi";class t{static library=u(i(new URL(`currencyRatesLib.${c}`,import.meta.url)),{BCVGetCurrencyValue:{args:[r.int8_t],returns:r.ptr},BCVGetDollarValue:{args:[],returns:r.ptr},BinanceGetDollarValue:{args:[],returns:r.ptr},BinanceGetDollarData:{args:[],returns:r.ptr},MonitorDollarGetDollarValue:{args:[],returns:r.ptr},MonitorDollarGetDollarData:{args:[],returns:r.ptr},FreeString:{args:[r.ptr],returns:r.void}});static toString(a){let s=new n(a);return this.library.symbols.FreeString(s.ptr),s.toString()}}class e extends t{static get_currency_value(a){return this.toString(this.library.symbols.BCVGetCurrencyValue(a))}static get_currency_value_as_float(a){return Number(this.toString(this.library.symbols.BCVGetCurrencyValue(a)))}static get_dollar_value(){return this.toString(this.library.symbols.BCVGetDollarValue())}static get_dollar_value_as_float(){return Number(this.toString(this.library.symbols.BCVGetDollarValue()))}}class l extends t{static get_dollar_value(){return this.toString(this.library.symbols.BinanceGetDollarValue())}static get_dollar_value_as_float(){return Number(this.toString(this.library.symbols.BinanceGetDollarValue()))}static get_dollar_data(){return JSON.parse(this.toString(this.library.symbols.BinanceGetDollarData()))}}class o extends t{static get_dollar_value(){return this.toString(this.library.symbols.MonitorDollarGetDollarValue())}static get_dollar_value_as_float(){return Number(this.toString(this.library.symbols.MonitorDollarGetDollarValue()))}static get_dollar_data(){return JSON.parse(this.toString(this.library.symbols.MonitorDollarGetDollarData()))}}var g={bcv:e,binance:l,monitor_dollar:o},C=g;export{C as default};
