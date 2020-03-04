package factory

import "testing"

func TestAb(t *testing.T) {
	var f1 Factory = WinFactory{}
	f1.newButton().draw()
	f1.newInput().pan()

	var f2 Factory = MacFactory{}
	f2.newButton().draw()
	f2.newInput().pan()
}
