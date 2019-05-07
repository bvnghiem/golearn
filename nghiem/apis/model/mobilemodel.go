package model

import (
	"golearn/nghiem/apis/entities"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MobileModel struct {
	Db         *mgo.Database
	Collection string
}

func (mobileModel MobileModel) FindAll() (mobiles []entities.Mobile, err error) {
	err = mobileModel.Db.C(mobileModel.Collection).Find(bson.M{}).All(&mobiles)
	return mobiles, err
}
