package core

import (
	"embed"
	"fyne.io/fyne/v2"
)

type _resources struct {
	blank     *fyne.StaticResource
	mousedown *fyne.StaticResource

	face0 *fyne.StaticResource
	face1 *fyne.StaticResource
	face2 *fyne.StaticResource
	face3 *fyne.StaticResource
	face4 *fyne.StaticResource

	num0 *fyne.StaticResource
	num1 *fyne.StaticResource
	num2 *fyne.StaticResource
	num3 *fyne.StaticResource
	num4 *fyne.StaticResource
	num5 *fyne.StaticResource
	num6 *fyne.StaticResource
	num7 *fyne.StaticResource
	num8 *fyne.StaticResource
	num9 *fyne.StaticResource

	mineNum0 *fyne.StaticResource
	mineNum1 *fyne.StaticResource
	mineNum2 *fyne.StaticResource
	mineNum3 *fyne.StaticResource
	mineNum4 *fyne.StaticResource
	mineNum5 *fyne.StaticResource
	mineNum6 *fyne.StaticResource
	mineNum7 *fyne.StaticResource
	mineNum8 *fyne.StaticResource

	mine0 *fyne.StaticResource
	mine1 *fyne.StaticResource
	mine2 *fyne.StaticResource

	flag *fyne.StaticResource
}

var resources = new(_resources)

func Init(images embed.FS) {
	var (
		data []byte
	)

	data, _ = images.ReadFile("images/blank.gif")
	resources.blank = fyne.NewStaticResource("blank.gif", data)

	data, _ = images.ReadFile("images/mousedown.gif")
	resources.mousedown = fyne.NewStaticResource("mousedown.gif", data)

	data, _ = images.ReadFile("images/face0.gif")
	resources.face0 = fyne.NewStaticResource("face0.gif", data)

	data, _ = images.ReadFile("images/face1.gif")
	resources.face1 = fyne.NewStaticResource("face1.gif", data)

	data, _ = images.ReadFile("images/face2.gif")
	resources.face2 = fyne.NewStaticResource("face2.gif", data)

	data, _ = images.ReadFile("images/face3.gif")
	resources.face3 = fyne.NewStaticResource("face3.gif", data)

	data, _ = images.ReadFile("images/face4.gif")
	resources.face4 = fyne.NewStaticResource("face4.gif", data)

	data, _ = images.ReadFile("images/d0.gif")
	resources.num0 = fyne.NewStaticResource("d0.gif", data)

	data, _ = images.ReadFile("images/d1.gif")
	resources.num1 = fyne.NewStaticResource("d1.gif", data)

	data, _ = images.ReadFile("images/d2.gif")
	resources.num2 = fyne.NewStaticResource("d2.gif", data)

	data, _ = images.ReadFile("images/d3.gif")
	resources.num3 = fyne.NewStaticResource("d3.gif", data)

	data, _ = images.ReadFile("images/d4.gif")
	resources.num4 = fyne.NewStaticResource("d4.gif", data)

	data, _ = images.ReadFile("images/d5.gif")
	resources.num5 = fyne.NewStaticResource("d5.gif", data)

	data, _ = images.ReadFile("images/d6.gif")
	resources.num6 = fyne.NewStaticResource("d6.gif", data)

	data, _ = images.ReadFile("images/d7.gif")
	resources.num7 = fyne.NewStaticResource("d7.gif", data)

	data, _ = images.ReadFile("images/d8.gif")
	resources.num8 = fyne.NewStaticResource("d8.gif", data)

	data, _ = images.ReadFile("images/d9.gif")
	resources.num9 = fyne.NewStaticResource("d9.gif", data)

	data, _ = images.ReadFile("images/0.gif")
	resources.mineNum0 = fyne.NewStaticResource("0.gif", data)

	data, _ = images.ReadFile("images/1.gif")
	resources.mineNum1 = fyne.NewStaticResource("1.gif", data)

	data, _ = images.ReadFile("images/2.gif")
	resources.mineNum2 = fyne.NewStaticResource("2.gif", data)

	data, _ = images.ReadFile("images/3.gif")
	resources.mineNum3 = fyne.NewStaticResource("3.gif", data)

	data, _ = images.ReadFile("images/4.gif")
	resources.mineNum4 = fyne.NewStaticResource("4.gif", data)

	data, _ = images.ReadFile("images/5.gif")
	resources.mineNum5 = fyne.NewStaticResource("5.gif", data)

	data, _ = images.ReadFile("images/6.gif")
	resources.mineNum6 = fyne.NewStaticResource("6.gif", data)

	data, _ = images.ReadFile("images/7.gif")
	resources.mineNum7 = fyne.NewStaticResource("7.gif", data)

	data, _ = images.ReadFile("images/8.gif")
	resources.mineNum8 = fyne.NewStaticResource("8.gif", data)

	data, _ = images.ReadFile("images/mine0.gif")
	resources.mine0 = fyne.NewStaticResource("mine0.gif", data)

	data, _ = images.ReadFile("images/mine1.gif")
	resources.mine1 = fyne.NewStaticResource("mine1.gif", data)

	data, _ = images.ReadFile("images/mine2.gif")
	resources.mine2 = fyne.NewStaticResource("mine2.gif", data)

	data, _ = images.ReadFile("images/flag.gif")
	resources.flag = fyne.NewStaticResource("flag.gif", data)
}
