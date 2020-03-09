package build

type Builder interface {
	MakeTitle(str string)
	MakeBody(str string)
	MakeFoot(str string)
}
