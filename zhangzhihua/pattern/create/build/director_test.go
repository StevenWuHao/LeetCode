package build

import (
	txt2 "LeetCode/zhangzhihua/pattern/create/build/txt"
	"testing"
)

//建造者模式
//具体组合交给Director去处理
func TestIndex(t *testing.T) {
	//html := &html.Html{}
	txt := &txt2.Text{}
	d := &Director{}
	d.SetBuild(txt)
	d.GetBodyFoot()
	txt.PrintStr()
}
