CGO_ENABLED=1 GOOS=linux go build -buildmode=c-shared -o linux/currencyRatesLib.so bindings/bindings.go
CGO_ENABLED=1 GOOS=windows go build -buildmode=c-shared -o windows/currencyRatesLib.dll bindings/bindings.go
CGO_ENABLED=1 GOOS=darwin go build -buildmode=c-shared -o darwin/currencyRatesLib.dylib bindings/bindings.go
