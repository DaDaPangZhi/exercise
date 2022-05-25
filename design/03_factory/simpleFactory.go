// 简单工厂
package factory

import "fmt"

type DBconnect interface {
	Init(address string, port int64, account string, passwd string)
	Connect() string
}

type BaseDB struct {
	address string
	port    int64
	account string
	passwd  string
}

func (this *BaseDB) Init(address string, port int64, account string, passwd string) {
	this.address = address
	this.port = port
	this.account = account
	this.passwd = passwd
}

type MysqlDB struct {
	BaseDB
}

func (mdb *MysqlDB) Connect() string {
	return fmt.Sprintf("连接了mysql,地址是:%s", fmt.Sprintf("%s:%d", mdb.address, mdb.port))
}

type MongoDB struct {
	BaseDB
}

func (mdb *MongoDB) Connect() string {
	return fmt.Sprintf("mongo,地址是:%s", fmt.Sprintf("%s:%d", mdb.address, mdb.port))
}

// 定义工厂类 声明方法 根据参数返回不同对象
type CreateDBFactory struct {
}

func (cdb *CreateDBFactory) CreateDBConnect(dbname string) (bdb DBconnect, err error) {
	switch dbname {
	case "mysql":
		bdb = &MysqlDB{}
	case "mongodb":
		bdb = &MongoDB{}
	default:
		bdb = &MysqlDB{}
	}
	return
}
