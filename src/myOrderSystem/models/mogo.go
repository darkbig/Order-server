package models

import (
	"gopkg.in/mgo.v2"
	"github.com/astaxie/beego/logs"
)

type UserInfo struct {
	UserId  string `bson:"userid"`
	Phone string	`bson:"phone"`
	PassWord string `bson:"password"`
}
type Men struct {
	Persons []UserInfo
}

const URL = "127.0.0.1:27017"

func UserMesUpData(dataBaseName string,collectionName string, tmp *UserInfo) (bool ,string) {

	session, err := mgo.Dial(URL)  //连接数据库
	if err != nil {
		logs.Error(err.Error())
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	db := session.DB(dataBaseName)	 //数据库名称
	collection := db.C(collectionName) //如果该集合已经存在的话，则直接返回
	err = collection.Insert(&tmp)
	if err != nil {
		return false,err.Error()
	}
	return true,""
}
