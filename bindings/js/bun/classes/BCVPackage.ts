import BasePackage from "./BasePackage";
import type CurrencyValue from "../enums/CurrencyValue";

export default class BCVPackage extends BasePackage {
	static get_currency_value(currency: CurrencyValue) {
		return this.toString(this.library.symbols.BCVGetCurrencyValue(currency)!);
	}
	static get_currency_value_as_float(currency: CurrencyValue) {
		return Number(
			this.toString(this.library.symbols.BCVGetCurrencyValue(currency)!),
		);
	}
	static get_dollar_value() {
		return this.toString(this.library.symbols.BCVGetDollarValue()!);
	}
	static get_dollar_value_as_float() {
		return Number(this.toString(this.library.symbols.BCVGetDollarValue()!));
	}
}
