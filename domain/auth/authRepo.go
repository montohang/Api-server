package auth

import "api_server/model"

type IAuthRepo interface {
	ReadAccountByEmail(string)(*model.Account,error)
	CreateAccount(account *model.Account,id string) error
}