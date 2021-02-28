package model

import "strings"

type Account struct {
	AccountID string            `json:"accountID"`
	Email     string            `json:"email"`
	Password  string            `json:"password"`
	Token     string            `json:"token"`
	Errors    map[string]string `json:"errors"`
}

func (account *Account) Validate() bool {
	account.Errors = make(map[string]string)
	if strings.TrimSpace(account.Email) == "" {
		account.Errors["Email"] = "enter your email"
	}
	if strings.TrimSpace(account.Password) == "" {
		account.Errors["Password"] = "enter your password"
	}

	return len(account.Errors) == 0
}