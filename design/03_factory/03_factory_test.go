package factory_test

import (
	factory "design/03_factory"
	"fmt"
	"testing"
)

func TestSimpleFactory(t *testing.T) {
	Simplefactory := &factory.CreateDBFactory{}
	mysqlObject, err := Simplefactory.CreateDBConnect("mysql")
	if err != nil {
		panic("mysql error")
	}
	mysqlObject.Init("mysqladdress", 3306, "kengerukong", "kengerukong")
	mysqlConnect := mysqlObject.Connect()
	fmt.Println("mysql connect is " + mysqlConnect)
}

func TestFactoryFunc(t *testing.T) {
	mysqlFactory := &factory.CreateMysqlDBFactoryV2{}
	mysqlObject, err := mysqlFactory.CreateDBConnectV2()
	if err != nil {
		panic("mysql error")
	}
	mysqlObject.InitV2("mysqladdress", 3306, "kengerukong", "kengerukong")
	mysqlConnect := mysqlObject.ConnectV2()
	fmt.Println("mysql connect is " + mysqlConnect)

	mongodbFactory := &factory.CreateMongodbDBFactoryV2{}
	mongodbObject, err := mongodbFactory.CreateDBConnectV2()
	if err != nil {
		panic("mongodb error")
	}
	mongodbObject.InitV2("mongodbaddress", 3307, "kengerukong", "kengerukong")
	mongodbConnect := mongodbObject.ConnectV2()
	fmt.Println("mongodb connect is " + mongodbConnect)
}
