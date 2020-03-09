package build

type Director struct {
	build Builder
}

func (d *Director) SetBuild(build Builder) {
	d.build = build
}

func (d *Director) GetTitleBody() {
	d.build.MakeTitle("getTitleBody")
	d.build.MakeBody("makeBody")
}

func (d *Director) GetBodyFoot() {
	d.build.MakeBody("makeBody")
	d.build.MakeFoot("MakeFoot")
}

func (d *Director) PrintStr() {
}
