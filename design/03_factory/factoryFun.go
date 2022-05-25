// 工厂方法
package factory

import "fmt"


type BaseDBV2 struct {
	address string
	port    int64
	account string
	passwd  string
}

func (this *BaseDB) InitV2(address string, port int64, account string, passwd string) {
	this.address = address
	this.port = port
	this.account = account
	this.passwd = passwd
}

type MysqlDBV2 struct {
	BaseDB
}

func (mdb *MysqlDBV2) ConnectV2() string {
	return fmt.Sprintf("连接了mysql,地址是:%s", fmt.Sprintf("%s:%d", mdb.address, mdb.port))
}

type MongoDBV2 struct {
	BaseDB
}

func (mdb *MongoDBV2) ConnectV2() string {
	return fmt.Sprintf("mongo,地址是:%s", fmt.Sprintf("%s:%d", mdb.address, mdb.port))
}

// 定义创建Mysql工厂类 声明方法 获取Mysql对象
type CreateMysqlDBFactoryV2 struct {
}

func (cdb *CreateMysqlDBFactoryV2) CreateDBConnectV2() (bdb *MysqlDBV2, err error) {
	bdb = &MysqlDBV2{}
	return bdb, nil
}

// 定义创建Mongodb工厂类 声明方法 获取Mongodb对象
type CreateMongodbDBFactoryV2 struct {
}

func (cdb *CreateMongodbDBFactoryV2) CreateDBConnectV2() (bdb *MongoDBV2, err error) {
	bdb = &MongoDBV2{}
	return bdb, nil
}
