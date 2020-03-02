package instance

import (
	"fmt"
	"sync"
)

/**
单例

解决问题：某一个类在全局应用中只需要存在一个

场景应用案例：全局配置，全局资源等

说明：非线程安全，可查看测试用例 TestUnsafe
*/

type DB struct {
	name string
}

var instance *DB

//普通单例
func NewDBUnsafe(nm string) *DB {

	if instance == nil {
		instance = &DB{nm}
	}
	fmt.Println(instance.name)
	return instance
}

/**
线程安全 单例
*/
var DBLock sync.Mutex

func NewDBSafe(nm string) *DB {
	if instance == nil {

		DBLock.Lock()
		defer DBLock.Unlock()
		if instance == nil {
			instance = &DB{nm}
		}
	}
	fmt.Println(instance.name)
	return instance
}

/**
sync.Once 是go内置的一个单例方法
*/
var once sync.Once

func GetInstanceSafe(nm string) *DB {
	once.Do(func() {
		instance = &DB{nm}
	})
	fmt.Println(instance.name)
	return instance

}
