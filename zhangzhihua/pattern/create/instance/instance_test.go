package instance

import (
	"runtime"
	"strconv"
	"testing"
	time2 "time"
)

func TestNewDBSafe(t *testing.T) {
	runtime.GOMAXPROCS(3)

	time := 10
	for time > 0 {
		go NewDBSafe(strconv.Itoa(time))
		time--
	}
	time2.Sleep(time2.Second * 3)
}

func TestGetInstance(t *testing.T) {
	runtime.GOMAXPROCS(3)
	time := 10
	for time > 0 {
		go GetInstanceSafe(strconv.Itoa(time))
		time--
	}
	time2.Sleep(time2.Second * 3)
}

func TestIdx(t *testing.T) {

	name1 := "mysql"
	db1 := NewDBUnsafe(name1)
	t.Logf("First new %s \n", db1.name)

	name2 := "sql"
	db1 = NewDBUnsafe(name2)
	t.Logf("First new %s \n", db1.name)

	name3 := "sql"
	db1 = NewDBUnsafe(name3)
	t.Logf("First new %s \n", db1.name)

}

/*
测试了线程安全问题
*/
func TestUnsafe(t *testing.T) {
	time := 10
	for time > 0 {
		go NewDBUnsafe(strconv.Itoa(time))
		time--
	}
	time2.Sleep(time2.Second * 3)
}
