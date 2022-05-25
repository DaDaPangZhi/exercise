// 单例模式 一个类只有一个对象
package singleton

type Singleton struct{}

var singleton *Singleton

// 饿汉模式 系统加载就完成初始化
func init() {
	singleton = &Singleton{}
}

func GetInstance() *Singleton {
	return singleton
}
