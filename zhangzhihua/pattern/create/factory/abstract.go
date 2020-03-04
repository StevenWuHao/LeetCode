package factory

import "fmt"

//抽象工厂
//案例：windows和mac分别都有button和input,该用例展示如跟据抽象工厂创建出相同系的windows组件

//抽象产品
type Button interface {
	draw()
}

type Input interface {
	pan()
}

type WinButton struct {
}

type WinInput struct {
}

type MacButton struct {
}

type MacInput struct {
}

func (w WinButton) draw() {
	fmt.Println("Win button")
}

func (w WinInput) pan() {
	fmt.Println("Win input")
}

func (w MacButton) draw() {
	fmt.Println("Mac button")
}

func (w MacInput) pan() {
	fmt.Println("Mac input")
}

type Factory interface {
	newButton() Button
	newInput() Input
}

type WinFactory struct{}

func (w WinFactory) newButton() Button {
	fmt.Println("win new button factory")
	return &WinButton{}
}

func (w WinFactory) newInput() Input {
	fmt.Println("win new input factory")
	return &WinInput{}
}

type MacFactory struct {
}

func (m MacFactory) newButton() Button {
	fmt.Println("mac new button factory")
	return &MacButton{}
}

func (m MacFactory) newInput() Input {
	fmt.Println("mac new input factory")
	return &MacInput{}
}
