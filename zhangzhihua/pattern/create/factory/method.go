package factory

import "fmt"

//工厂方法
//将实例推迟到子类中去创建
type Logger interface {
	echo()
}

type File struct{}
type Stdout struct{}

func (f File) echo() {
	fmt.Println("File")
}

func (f Stdout) echo() {
	fmt.Println("Stdout")
}

type NewProduct interface {
	new() Logger
}

//这里表示抽象工厂
type OutPutFactory struct {}

func (o OutPutFactory) new() Logger {
	fmt.Println("i am Abstract OutPutFactory")
	return File{}
}

type FileFactory struct {
	OutPutFactory
}

type StdoutFactory struct {
	OutPutFactory
}

func (f FileFactory) new(name string) Logger {
	fmt.Printf("name: %s i am file FileFactory \n", name)
	return File{}
}

func (f StdoutFactory) new() Logger {
	fmt.Println("i am file StdoutFactory")
	return Stdout{}
}
