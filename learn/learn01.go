package learn

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
)

const (
	a = iota
	b
	c
	d
	e = 9
	f
)

//首字母大写->可以被外部包使用
//首字母小写->只能在包内使用
func Learn01() {
	fmt.Printf("%d - %d - %d - %d - %d - %d \n\n", a, b, c, d, e, f)
	//数值交换
	a, b := 1, 2
	fmt.Printf("%d - %d\n", a, b)
	a, b = b, a
	fmt.Printf("%d - %d", a, b)
	//创建sin函数图片
	createImage()
}

func createImage() {
	const size = 300
	//创建灰图
	pic := image.NewGray(image.Rect(0, 0, size, size))
	//遍历像素
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			//填充为白色
			pic.SetGray(x, y, color.Gray{255})
		}
	}
	// 0-size生成最大坐标
	for x := 0; x < size; x++ {
		// 让sin的值的范围在0~2Pi之间
		s := float64(x) * 2 * math.Pi / size
		// sin的幅度为一半的像素。向下偏移一半像素并翻转
		y := size/2 - math.Sin(s)*size/2
		// 用黑色绘制sin轨迹
		pic.SetGray(x, int(y), color.Gray{0})
	}
	//创建文件
	file, err := os.Create("sin.png")
	if err != nil {
		log.Fatal(err)
	}
	//写入
	pErr := png.Encode(file, pic)
	if pErr != nil {
		log.Fatal(err)
	}
	file.Close()
}
