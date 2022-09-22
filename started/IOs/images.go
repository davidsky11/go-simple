package IOs

// image 包定义了 Image 接口：
// 		type Image interface {
//    		ColorModel() color.Model
//    		Bounds() Rectangle
//    		At(x, y int) color.Color
//		}

// color.Color 和 color.Model 类型也是接口，但是通常因为直接使用预定义的实现 image.RGBA 和 image.RGBAModel 而被忽视了。
// 这些接口和类型由 image/color 包定义。

import (
	"fmt"
	"image"
)

func RGB() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}
