package txt

import (
	"fmt"
)

type Text struct {
	echoStr string
}

func (text *Text) MakeTitle(str string) {
	text.echoStr += "title :" + str + "//r//n"
}

func (text *Text) MakeBody(str string) {
	text.echoStr += "body :" + str
}

func (text *Text) MakeFoot(str string) {
	text.echoStr += "foot :" + str + "//r//n"
}

func (text *Text) PrintStr() {
	fmt.Println(text.echoStr)
}
