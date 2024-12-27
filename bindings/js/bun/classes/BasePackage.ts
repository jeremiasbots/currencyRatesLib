import { fileURLToPath } from "bun";
import { CString, dlopen, FFIType, suffix, type Pointer } from "bun:ffi";

export default class BasePackage {
	protected static library = dlopen(
		fileURLToPath(new URL(`currencyRatesLib.${suffix}`, import.meta.url)),
		{
			BCVGetCurrencyValue: {
				args: [FFIType.int8_t],
				returns: FFIType.ptr,
			},
			BCVGetDollarValue: {
				args: [],
				returns: FFIType.ptr,
			},
			BinanceGetDollarValue: {
				args: [],
				returns: FFIType.ptr,
			},
			BinanceGetDollarData: {
				args: [],
				returns: FFIType.ptr,
			},
			MonitorDollarGetDollarValue: {
				args: [],
				returns: FFIType.ptr,
			},
			MonitorDollarGetDollarData: {
				args: [],
				returns: FFIType.ptr,
			},
			FreeString: {
				args: [FFIType.ptr],
				returns: FFIType.void,
			},
		},
	);

	protected static toString(array: Pointer) {
		const cStr = new CString(array);
		this.library.symbols.FreeString(cStr.ptr);
		return cStr.toString();
	}
}
