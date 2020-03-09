package html

import (
	"fmt"
)

type Html struct {
	echoStr string
}

func (html *Html) MakeTitle(str string) {
	html.echoStr += "<html> <title>" + str + "</title>"
}

func (html *Html) MakeBody(str string) {
	html.echoStr += "<body><div>" + str + "</div>"
}

func (html *Html) MakeFoot(str string) {
	html.echoStr += "<div> foot :" + str + "</div></body></html>"
}

func (html *Html) PrintStr() {
	fmt.Println(html.echoStr)
}
