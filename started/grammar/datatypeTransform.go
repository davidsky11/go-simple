package grammar

import "fmt"

var a001 int = 100

// 类型转换时，需要考虑两种类型之间的关系和范围，是否会发生数值截断
// 布尔型无法与其他类型进行转换
func transform() {
	float001 := float64(a001)
	string001 := string(a001)
	fmt.Printf("%f, %s", float001, string001)
}

func cnEnTrans() {
	chinese := 100
	english := 90.9
	avg := (chinese + int(english)) / 2
	avg2 := (float64(chinese) + english) / 2
	fmt.Printf("%T, %d\n", avg, avg)
	fmt.Printf("%T, %f\n", avg2, avg2)
}

// 整型转字符串类型
// 这种类型的转换，其实相当于byte或rune转string。
// int数值是ASCII码的编号或unicode字符集的编号，转成string就是根据字符集，将对应编号的字符查找出来。
// 当该数值超出unicode编号范围，则转成的字符串显示为乱码。
// 例如，19968转string，就是“一”。

// • ASCII字符集中数字的十进制范围是48～57；
// • ASCII字符集中大写字母的十进制范围是65～90；
// • ASCII字符集中小写字母的十进制范围是97～122；
// • unicode字符集中汉字的范围是4e00～9fa5，十进制范围是19968～40869。

func intTrans() {
	int002 := 97
	int003 := 19968
	result := string(int002)
	fmt.Println(result)
	result = string(int003)
	fmt.Println(result)
}

// 在Go语言中，不允许字符串转int，会产生如下错误
// 		cannot convert str (type string) to type int
