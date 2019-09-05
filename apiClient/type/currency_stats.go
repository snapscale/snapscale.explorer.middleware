package _type

type Currency struct {
	Supply    string `json:"supply"`
	MaxSupply string `json:"max_supply"`
	Issuer    string `json:"issuer"`
}

type CurrencyStatsRequest struct {
	Symbol string `json:"symbol"`
	Code   string `json:"code"`
}

type CurrencyStatsResponse struct {
	XST Currency `json:"XST"`
}
