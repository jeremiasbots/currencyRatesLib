import { expect, describe, test } from "bun:test";
import BCVPackage from "./classes/BCVPackage";
import CurrencyValue from "./enums/CurrencyValue";

describe("BCVPackage.get_currency_value(CurrencyValue):", () => {
  test("get euro value", () => {
    expect(BCVPackage.get_currency_value(CurrencyValue.Euro)).toBeString();
  });
  test("get dollar value", () => {
    expect(
      BCVPackage.get_currency_value(CurrencyValue.AmericanDollar)
    ).toBeString();
  });
  test("get yuan value", () => {
    expect(BCVPackage.get_currency_value(CurrencyValue.Yuan)).toBeString();
  });
  test("get turkish lira value", () => {
    expect(
      BCVPackage.get_currency_value(CurrencyValue.TurkishLira)
    ).toBeString();
  });
  test("get ruble value", () => {
    expect(BCVPackage.get_currency_value(CurrencyValue.Ruble)).toBeString();
  });
});

describe("BCVPackage.get_dollar_value():", () => {
  test("get dollar value", () => {
    expect(BCVPackage.get_dollar_value()).toBeString();
  });
});

describe("BCVPackage.get_dollar_value_as_float():", () => {
  test("get dollar value", () => {
    expect(BCVPackage.get_dollar_value_as_float()).toBeNumber();
  });
});

describe("BCVPackage.get_currency_value_as_float(CurrencyValue):", () => {
  test("get euro value", () => {
    expect(
      BCVPackage.get_currency_value_as_float(CurrencyValue.Euro)
    ).toBeNumber();
  });
  test("get dollar value", () => {
    expect(
      BCVPackage.get_currency_value_as_float(CurrencyValue.AmericanDollar)
    ).toBeNumber();
  });
  test("get yuan value", () => {
    expect(
      BCVPackage.get_currency_value_as_float(CurrencyValue.Yuan)
    ).toBeNumber();
  });
  test("get turkish lira value", () => {
    expect(
      BCVPackage.get_currency_value_as_float(CurrencyValue.TurkishLira)
    ).toBeNumber();
  });
  test("get ruble value", () => {
    expect(
      BCVPackage.get_currency_value_as_float(CurrencyValue.Ruble)
    ).toBeNumber();
  });
});

describe("BCVPackage.get_dollar_data():", () => {
  test("get dollar data", () => {
	console.log(BCVPackage.get_dollar_data())
    expect(BCVPackage.get_dollar_data()).toBeObject();
  });
});
