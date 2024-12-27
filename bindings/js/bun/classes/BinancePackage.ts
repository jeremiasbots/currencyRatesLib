import type APIDollarData from "../types/APIDollarData";
import BasePackage from "./BasePackage";

export default class BinancePackage extends BasePackage {
	static get_dollar_value() {
		return this.toString(this.library.symbols.BinanceGetDollarValue()!);
	}
	static get_dollar_value_as_float() {
		return Number(this.toString(this.library.symbols.BinanceGetDollarValue()!));
	}
	static get_dollar_data(): APIDollarData {
		return JSON.parse(
			this.toString(this.library.symbols.BinanceGetDollarData()!),
		);
	}
}
