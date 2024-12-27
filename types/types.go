package types

type APIResponse struct {
	Source     string   `json:"fuente"`
	Name       string   `json:"nombre"`
	Buying     *float64 `json:"compra"`
	Selling    *float64 `json:"venta"`
	Average    float64  `json:"promedio"`
	UpdateDate string   `json:"fechaActualizacion"`
}
