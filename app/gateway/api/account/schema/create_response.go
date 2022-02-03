package schema

type CreateAccountResponse struct {
	// AccountID is the id generated that represents this account
	AccountID string `json:"account_id"`

	// Token is a JWT token containing claims about the logged in account owner
	Token string `json:"token"`
}
