import type APIDollarData from "../types/APIDollarData";
import BasePackage from "./BasePackage";
export default class MonitorDollarPackage extends BasePackage {
    static get_dollar_value(): string;
    static get_dollar_value_as_float(): number;
    static get_dollar_data(): APIDollarData;
}
