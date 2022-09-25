package reflects

// 反射的限制：
// 1. 到 Go 1.16 为止，我们无法通过反射动态创建一个接口类型。这是Go反射目前的一个限制
// 2. 使用反射动态创建结构体类型的时候可能会有各种不完美的情况出现
// 3. 我们无法通过反射来声明一个新的类型

import (
	"fmt"
	"reflect"
)

type T []interface{ m() }

func (T) m() {}

type F func(string, int) bool

func (f F) m(s string) bool {
	return f(s, 32)
}

func (f F) M() {}

type I interface {
	m(s string) bool
	M()
}

func Reflect1() {
	tp := reflect.TypeOf(new(interface{}))
	tt := reflect.TypeOf(T{})
	fmt.Println(tp.Kind(), tt.Kind()) // ptr slice

	// 使用间接的方法得到表示两个接口类型的 reflect.Type 值
	ti, tim := tp.Elem(), tt.Elem()
	fmt.Println(ti.Kind(), tim.Kind()) // interface interface

	fmt.Println(tt.Implements(tim))  // true
	fmt.Println(tp.Implements(tim))  // false
	fmt.Println(tim.Implements(tim)) // true

	// 所有的类型都实现了任何空接口类型
	fmt.Println(tp.Implements(ti))  // true
	fmt.Println(tt.Implements(ti))  // true
	fmt.Println(tim.Implements(ti)) // true
	fmt.Println(ti.Implements(ti))  // true
}

func Reflect2() {
	var x struct {
		F F
		i I
	}

	tx := reflect.TypeOf(x)
	fmt.Println(tx.Kind())        // struct
	fmt.Println(tx.NumField())    // 2
	fmt.Println(tx.Field(1).Name) // i
	// 包路径（PkgPath）是非导出字段（或方法）的内在属性
	fmt.Println(tx.Field(0).PkgPath)
	fmt.Println(tx.Field(1).PkgPath) // main

	tf, ti := tx.Field(0).Type, tx.Field(1).Type
	fmt.Println(tf.Kind())               // func
	fmt.Println(tf.IsVariadic())         // false
	fmt.Println(tf.NumIn(), tf.NumOut()) // 2 1
	t0, t1, t2 := tf.In(0), tf.In(1), tf.Out(0)
	// 下一行打印出 string int bool
	fmt.Println(t0.Kind(), t1.Kind(), t2.Kind())

	fmt.Println(tf.NumMethod(), ti.NumMethod()) // 1 2
	fmt.Println(tf.Method(0).Name)              // M
	fmt.Println(ti.Method(1).Name)              // m
	_, ok1 := tf.MethodByName("m")
	_, ok2 := ti.MethodByName("m")
	fmt.Println(ok1, ok2) // false true
}

// 1. 对于非接口类型，reflect.Type.NumMethod方法只返回一个类型的所有导出的方法（包括通过
//		内嵌得来的隐式方法）的个数，并且 方法reflect.Type.MethodByName不能用来获取一个类型
//		的非导出方法； 而对于接口类型，则并无这些限制
// 2. 虽然reflect.Type.NumField方法返回一个结构体类型的所有字段（包括非导出字段）的数
//		目，但是不推荐 Ң 使用方法reflect.Type.FieldByName来获取非导出字段

type Tii struct {
	X    int  `max:"99" min:"0" default:"0"`
	Y, Z bool `optional:"yes"`
}

func Reflect3() {
	tii := reflect.TypeOf(Tii{})
	x := tii.Field(0).Tag
	y := tii.Field(1).Tag
	z := tii.Field(2).Tag
	fmt.Println(reflect.TypeOf(x)) // reflect.StructTag

	// v 的类型为 string
	v, present := x.Lookup("max")
	fmt.Println(len(v), present)      // 2 true
	fmt.Println(x.Get("max"))         // 99
	fmt.Println(x.Lookup("optional")) // false
	fmt.Println(y.Lookup("optional")) // yes true
	fmt.Println(z.Lookup("optional")) // yes true
}

func DynamicReflect() {
	ta := reflect.ArrayOf(5, reflect.TypeOf(123))
	fmt.Println(ta) // [5]int
	tc := reflect.ChanOf(reflect.SendDir, ta)
	fmt.Println(tc) // chan<= [5]int
	tp := reflect.PtrTo(ta)
	fmt.Println(tp) // *[5]int
	ts := reflect.SliceOf(tp)
	fmt.Println(ts) // []*[5]int
	tm := reflect.MapOf(ta, tc)
	fmt.Println(tm) // map[[5]int]chan<- [5]int
	tf := reflect.FuncOf([]reflect.Type{ta}, []reflect.Type{tp, tc}, false)
	fmt.Println(tf) // func([5]int) (*[5]int, chan<- [5]int)
	tt := reflect.StructOf([]reflect.StructField{
		{Name: "Age", Type: reflect.TypeOf("abc")},
	})
	fmt.Println(tt)            // struct { Age string }
	fmt.Println(tt.NumField()) // 1
}
