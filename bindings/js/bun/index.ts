import BCVPackage from "./classes/BCVPackage";
import BinancePackage from "./classes/BinancePackage";
import MonitorDollarPackage from "./classes/MonitorDollarPackage";

const packagesObject = {
	bcv: BCVPackage,
	binance: BinancePackage,
	monitor_dollar: MonitorDollarPackage,
};

export default packagesObject;
