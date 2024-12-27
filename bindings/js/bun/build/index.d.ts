import BCVPackage from "./classes/BCVPackage";
import BinancePackage from "./classes/BinancePackage";
import MonitorDollarPackage from "./classes/MonitorDollarPackage";
declare const packagesObject: {
    bcv: typeof BCVPackage;
    binance: typeof BinancePackage;
    monitor_dollar: typeof MonitorDollarPackage;
};
export default packagesObject;
