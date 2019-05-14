package models

import (
	"golearn/nghiem/middlewares/entities"

	"gopkg.in/mgo.v2/bson"

	"golang.org/x/crypto/bcrypt"

	"gopkg.in/mgo.v2"
)

type AccountModel struct {
	Db         *mgo.Database
	Collection string
}

func (aModel AccountModel) Create(account *entities.Account) error {
	encryptedPass, _ := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	account.Password = string(encryptedPass)
	return aModel.Db.C(aModel.Collection).Insert(&account)
}
func (aModel AccountModel) Check(username, password string) bool {
	var account entities.Account
	err := aModel.Db.C(aModel.Collection).Find(bson.M{
		"username": username,
	}).One(&account)
	if err != nil {
		return false
	} else {
		return bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password)) == nil
	}
}
