/// <reference types="bun-types" />
import { FFIType, type Pointer } from "bun:ffi";
export default class BasePackage {
    protected static library: import("bun:ffi").Library<{
        BCVGetCurrencyValue: {
            args: FFIType.int8_t[];
            returns: FFIType.ptr;
        };
        BCVGetDollarValue: {
            args: never[];
            returns: FFIType.ptr;
        };
        BCVGetDollarData: {
            args: never[];
            returns: FFIType.ptr;
        };
        BinanceGetDollarValue: {
            args: never[];
            returns: FFIType.ptr;
        };
        BinanceGetDollarData: {
            args: never[];
            returns: FFIType.ptr;
        };
        MonitorDollarGetDollarValue: {
            args: never[];
            returns: FFIType.ptr;
        };
        MonitorDollarGetDollarData: {
            args: never[];
            returns: FFIType.ptr;
        };
        FreeString: {
            args: FFIType.ptr[];
            returns: FFIType.void;
        };
    }>;
    protected static toString(array: Pointer): string;
}
