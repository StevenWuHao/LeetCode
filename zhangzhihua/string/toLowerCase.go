package string

import (
	"bytes"
	"fmt"
)

func main() {

	str := "GOLANG"
	echo := toLowerCase(str)
	fmt.Println(echo)
}

func toLowerCase(str string) string {
	if len(str) == 0 {
		return str
	}

	var buff bytes.Buffer
	for _, v := range str {
		if v >= 65 && v <= 90 {
			buff.WriteRune(v + 32)
		} else {
			buff.WriteRune(v)
		}
	}
	return buff.String()
}
