// 单例模式 一个类只有一个对象
package singleton

import "sync"

type LazySingleton struct{}

var (
	lazySingleton *LazySingleton
	once          = &sync.Once{}
)

// 懒汉模式 对象被使用才完成初始化
func GetLazyInstance() *LazySingleton {
	if lazySingleton == nil {
		once.Do(func() {
			lazySingleton = &LazySingleton{}
		})
	}
	return lazySingleton
}
