package roawa

// TransactionRequest for sending a request for a new transactions
type TransactionRequest struct {
	CompanyNumber string `json:"company_number"`
	Reference     string `json:"reference"`
	Description   string `json:"description"`
}

// TransactionResponse for result of a new transaction
type TransactionResponse struct {
	ID string `json:"id"`
}
