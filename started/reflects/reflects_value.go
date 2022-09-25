package reflects

// 有两种方法获取一个代表着一个指针所引用着的值的reflect.Value值
//		1. 通过调用代表着此指针值的reflect.Value值的Elem方法。
// 		2. 将代表着此指针值的reflect.Value值的传递给一个reflect.Indirect函数调用。（如果传
//			递给一个reflect.Indirect函数调用的实参不代表着一个指针值，则此调用返回此实参的一个复制。）

import (
	"fmt"
	"reflect"
	"time"
)

func ReflectValue1() {
	n := 123
	p := &n
	vp := reflect.ValueOf(p)
	fmt.Println(vp.CanSet(), vp.CanAddr()) // false false
	vn := vp.Elem()                        // 取得 vp 的底层指针值引用的值的代表值
	fmt.Println(vn.CanSet(), vn.CanAddr()) // true true
	vn.Set(reflect.ValueOf(789))           // <=> vn.SetInt(789)
	fmt.Println(n)                         // 789
}

func Reflectvalue2() {
	var s struct {
		X interface{} // 一个导出字段
		y interface{} // 一个非导出字段（小写）
	}
	vp := reflect.ValueOf(&s)
	// 如果 vp 代表着一个指针，下一行等价于 vs := vp.Elem()
	vs := reflect.Indirect(vp)
	// vx 和 vy 都各自代表着一个接口值
	vx, vy := vs.Field(0), vs.Field(1)
	fmt.Println(vx.CanSet(), vx.CanAddr()) // true true
	// vy is addressable but not modifiable.
	fmt.Println(vy.CanSet(), vy.CanAddr()) // false true

	vb := reflect.ValueOf(123)
	vx.Set(vb) // ok，因为 vx 带包的值是可修改的

	//vy.Set(vb) // 会造成恐慌，因为 vy 代表的值是不可修改的
	fmt.Println(s)                      // {123 <nil>}
	fmt.Println(vx.IsNil(), vy.IsNil()) // false true

}

// 将一个自定义泛型函数绑定到不同的类型的函数值上

func InvertSlice(args []reflect.Value) (result []reflect.Value) {
	inSlice, n := args[0], args[0].Len()
	outSlice := reflect.MakeSlice(inSlice.Type(), 0, n)
	for i := n - 1; i >= 0; i-- {
		element := inSlice.Index(i)
		outSlice = reflect.Append(outSlice, element)
	}
	return []reflect.Value{outSlice}
}

func Bind(p interface{}, f func([]reflect.Value) []reflect.Value) {
	// invert 代表一个函数值
	invert := reflect.ValueOf(p).Elem()
	invert.Set(reflect.MakeFunc(invert.Type(), f))
}

func CustomReflectFunc2DuckingTypeValue() {
	var invertInts func([]int) []int
	Bind(&invertInts, InvertSlice)
	fmt.Println(invertInts([]int{2, 3, 5})) // [5 3 2]

	var invertStrs func([]string) []string
	Bind(&invertStrs, InvertSlice)
	fmt.Println(invertStrs([]string{"Go", "C"})) // [C Go]
}

type Tm struct {
	A, b int
}

func (t Tm) AddSubThenScale(n int) (int, int) {
	return n * (t.A + t.b), n * (t.A - t.b)
}

func MethodByNameCall() {
	t := Tm{5, 2}
	vt := reflect.ValueOf(t)
	vm := vt.MethodByName("AddSubThenScale")

	results := vm.Call([]reflect.Value{reflect.ValueOf(3)})
	fmt.Println(results[0].Int(), results[1].Int()) // 21 9

	neg := func(x int) int {
		return -x
	}

	vf := reflect.ValueOf(neg)
	fmt.Println(vf.Call(results[:1])[0].Int()) // -21
	fmt.Println(vf.Call([]reflect.Value{
		vt.FieldByName("A"), // 如果是字段b，则造成恐慌
	})[0].Int()) // -5
}

func ReflectSetMapIndex() {
	valueOf := reflect.ValueOf
	m := map[string]int{"unix": 1973, "windows": 1985}
	v := valueOf(m)

	// 第二个实参为Value零值时，表示删除一个映射条目
	v.SetMapIndex(valueOf("windows"), reflect.Value{})
	v.SetMapIndex(valueOf("linux"), valueOf(1991))
	for i := v.MapRange(); i.Next(); {
		fmt.Println(i.Key(), "\t:", i.Value())
	}
}

func ReflectChannelValue() {
	c := make(chan string, 2)
	vc := reflect.ValueOf(c)
	vc.Send(reflect.ValueOf("C"))
	succeeded := vc.TrySend(reflect.ValueOf("Go"))
	fmt.Println(succeeded) // true
	succeeded = vc.TrySend(reflect.ValueOf("C++"))
	fmt.Println(succeeded)          // false
	fmt.Println(vc.Len(), vc.Cap()) // 2 2
	vs, succeeded := vc.TryRecv()
	fmt.Println(vs.String(), succeeded) // C true
	vs, sentBeforeClosed := vc.Recv()
	fmt.Println(vs.String(), sentBeforeClosed) // Go false
	vs, succeeded1 := vc.TryRecv()
	//vs, succeeded = vc.TryRecv()   等价于上面一行
	fmt.Println(vs.String()) // <invalid Value>
	fmt.Println(succeeded1)  // false
}

func ReflectSelect() {
	c := make(chan int, 1)
	vc := reflect.ValueOf(c)
	succeeded := vc.TrySend(reflect.ValueOf(123))
	fmt.Println(succeeded, vc.Len(), vc.Cap()) // true 1 1

	vSend, vZero := reflect.ValueOf(789), reflect.Value{}
	branches := []reflect.SelectCase{
		{Dir: reflect.SelectDefault, Chan: vZero, Send: vZero},
		{Dir: reflect.SelectRecv, Chan: vc, Send: vZero},
		{Dir: reflect.SelectSend, Chan: vc, Send: vSend},
	}
	selIndex, vRecv, sentBeforeClosed := reflect.Select(branches)
	fmt.Println(selIndex)         // 1
	fmt.Println(sentBeforeClosed) // true
	fmt.Println(vRecv.Int())      // 123
	vc.Close()

	// 再模拟一次 select 流程控制代码块。因为 vc 已经关闭了，所以需将最后一个 cae 分支去除，
	// 否则它可能会造成一个恐慌。
	selIndex, _, sentBeforeClosed = reflect.Select(branches[:2])
	fmt.Println(selIndex, sentBeforeClosed) // 1 false
}

func ReflectZeroValue() {
	var z reflect.Value
	fmt.Println(z) // <invalid reflect.Value>
	v := reflect.ValueOf((*int)(nil)).Elem()
	fmt.Println(v)      // <invalid reflect.Value>
	fmt.Println(v == z) // true
	var i = reflect.ValueOf([]interface{}{nil}).Index(0)
	fmt.Println(i)             // <nil>
	fmt.Println(i.Elem())      // <invalid reflect.Value>
	fmt.Println(i.Elem() == z) // true
}

func ReflectInterfaceValue() {
	vx := reflect.ValueOf(123)
	vy := reflect.ValueOf("abc")
	vz := reflect.ValueOf([]bool{false, true})
	vt := reflect.ValueOf(time.Time{})

	x := vx.Interface().(int)
	y := vy.Interface().(string)
	z := vz.Interface().([]bool)
	m := vt.MethodByName("IsZero").Interface().(func() bool)
	fmt.Println(x, y, z, m()) // 123 abc [false true] true

	type T struct{ x int }
	t := &T{3}
	v := reflect.ValueOf(t).Elem().Field(0)
	fmt.Println(v) // 3
	// 调用一个代表着非导出字段的reflect.Value值的Interface方法将导致一个恐慌。
	fmt.Println(v.Interface()) // panic
}
