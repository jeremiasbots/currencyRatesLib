import { expect, describe, test } from "bun:test";
import MonitorDollarPackage from "./classes/MonitorDollarPackage";

describe("MonitorDollarPackage.get_dollar_value():", () => {
	test("get dollar value", () => {
		expect(MonitorDollarPackage.get_dollar_value()).toBeString();
	});
});

describe("MonitorDollarPackage.get_dollar_value_as_float():", () => {
	test("get dollar value", () => {
		expect(MonitorDollarPackage.get_dollar_value_as_float()).toBeNumber();
	});
});

describe("MonitorDollarPackage.get_dollar_data()):", () => {
	test("get dollar value", () => {
		expect(MonitorDollarPackage.get_dollar_data()).toBeObject();
	});
});
