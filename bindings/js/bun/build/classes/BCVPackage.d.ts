import BasePackage from "./BasePackage";
import type CurrencyValue from "../enums/CurrencyValue";
export default class BCVPackage extends BasePackage {
    static get_currency_value(currency: CurrencyValue): string;
    static get_currency_value_as_float(currency: CurrencyValue): number;
    static get_dollar_value(): string;
    static get_dollar_value_as_float(): number;
    static get_dollar_data(): any;
}
