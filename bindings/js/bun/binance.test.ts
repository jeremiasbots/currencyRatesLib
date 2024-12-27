import { expect, describe, test } from "bun:test";
import BinancePackage from "./classes/BinancePackage";

describe("BinancePackage.get_dollar_value():", () => {
	test("get dollar value", () => {
		expect(BinancePackage.get_dollar_value()).toBeString();
	});
});

describe("BinancePackage.get_dollar_value_as_float():", () => {
	test("get dollar value", () => {
		expect(BinancePackage.get_dollar_value_as_float()).toBeNumber();
	});
});

describe("BinancePackage.get_dollar_data()):", () => {
	test("get dollar value", () => {
		expect(BinancePackage.get_dollar_data()).toBeObject();
	});
});
