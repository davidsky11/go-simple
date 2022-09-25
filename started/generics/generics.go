package generics

import "fmt"

// Go 1.18 版本开始，官方正式支持泛型

// 泛型核心就三个概念：
// 		1. Type parameters for functions and types 类型参数，可以用于泛型函数以及泛型类型
//		2. Type sets defined by interfaces
//				Go 1.18之前，interface用来定义方法集( a set of methods)。
//				Go 1.18开始，还可以使用interface来定义类型集(a set of types)，作为类型参数的Type constraint(类型限制)
//		3. Type inference
//				类型推导，可以帮助我们在写代码的时候不用传递类型实参，由编译器自行推导。

// 通过引入 类型形参 和 类型实参 这两个概念，我们让一个函数获得了处理多种不同类型数据的能力，这种编程方式被称为 泛型编程。
// 通过Go的 接口+反射 不也能实现这样的动态数据处理吗？是的，泛型能实现的功能通过接口+反射也基本能实现。
//		使用过反射的人都知道反射机制有很多问题：
//		1. 用起来发麻
//		2. 失去了编译时的类型检查，容易出错
//		3. 性能不太理想

// ===> 经验之谈：
//		如果你经常要分别为不同的类型写完全相同逻辑的代码，那么使用泛型将是最合适的选择

// 泛型类型不能直接拿来使用，必须传入 类型实参 (Type argument)将其确定为具体的类型之后才可使用。
//		而传入类型实参确定具体类型的操作被称为 实例化 (Instantiations)

func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

var ints = map[string]int64{
	"first":  34,
	"second": 12,
}

var floats = map[string]float64{
	"first":  35.98,
	"second": 26.99,
}

func NonGenericSum() {
	fmt.Printf("Non-Generic Sums: %v and %v\n",
		SumInts(ints), SumFloats(floats))
}

// 通用函数处理

// 为K类型参数指定类型约束comparable。专门针对此类情况，comparable在 Go 中预先声明了约束。
//		它允许任何类型的值可以用作比较运算符==和的操作数!=。Go 要求映射键具有可比性。所以声明K as comparable是必要的，这样您就可以K在 map 变量中用作键。
//		它还确保调用代码对映射键使用允许的类型。

// 为V类型参数指定一个约束，它是两种类型的联合：int64和float64。
//		Using|指定两种类型的联合，这意味着此约束允许任何一种类型。
//		编译器将允许任一类型作为调用代码中的参数。

// 指定m参数是 type map[K]V，其中K和V 是已经为类型参数指定的类型。
//		请注意，我们知道map[K]V是有效的地图类型，因为K它是可比较的类型。
//		如果我们没有声明K可比较，编译器将拒绝对map[K]V.

func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func GenericSum() {
	fmt.Printf("Generic Sums: %v and %v\n",
		SumIntsOrFloats[string, int64](ints),
		SumIntsOrFloats[string, float64](floats))

	// 当 Go 编译器可以推断您要使用的类型时，您可以在调用代码中省略类型参数。
	//		编译器从函数参数的类型推断类型参数。
	fmt.Printf("Generic Sums, type parameters inferred: %v and %v\n",
		SumIntsOrFloats(ints),
		SumIntsOrFloats(floats))
}

// 声明类型约束

type Number interface {
	int64 | float64
}

func SumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func GenericSumWithConstraint() {
	fmt.Printf("Generic Sums with Constraint: %v and %v\n",
		SumNumbers(ints),
		SumNumbers(floats))
}

type MyMap[KEY int | string, VALUE float32 | float64] map[KEY]VALUE

// 泛型类型generic type： MyMap[KEY, VALUE]
// 类型约束type constraint： key部分 int | string, value部分 float32 | float64
// 类型形参列表type parameter list： KEY int | string, VALUE float32 | float64

var a MyMap[string, float64] = map[string]float64{
	"af_score":  9.6,
	"bob_score": 8.8,
}

// 类型实参type argument： MyMap[string, float64]        即 泛型类型实例化

// WowStruct 类型形参的互相套用
type WowStruct[T int | float32, S []T] struct {
	Data     S
	MaxValue T
	MinValue T
}

var ws = WowStruct[float32, []float32]{
	Data:     []float32{1.1, 1.8, 9.9, 9.8},
	MaxValue: 10,
	MinValue: 0,
}

// 泛型的套娃
type Slice[T int | string | float32 | float64] []T

// ✗ 错误。泛型类型Slice[T]的类型约束中不包含uint, uint8
//type UintSlice[T uint | uint8] Slice[T]

// ✓ 正确。基于泛型类型Slice[T]定义了新的泛型类型 FloatSlice[T] 。FloatSlice[T]只接受float32和float64两种类型
type FloatSlice[T float32 | float64] Slice[T]

// ✓ 正确。基于泛型类型Slice[T]定义的新泛型类型 IntAndStringSlice[T]
type IntAndStringSlice[T int | string] Slice[T]

// ✓ 正确 基于IntAndStringSlice[T]套娃定义出的新泛型类型
type IntSlice[T int] IntAndStringSlice[T]

// 在map中套一个泛型类型Slice[T]
type WowMap[T int | string] map[string]Slice[T]

// 在map中套Slice[T]的另一种写法
type WowMap2[T Slice[int] | Slice[string]] map[string]T

type MySlice[T int | float32 | string] []T

func (s MySlice[T]) Sum() T {
	var sum T
	for _, value := range s {
		sum += value
	}
	return sum
}

func GenericForReceiver() {
	var s MySlice[int] = []int{1, 2, 3, 4}
	var t MySlice[float32] = []float32{1.0, 2, 3, 4}

	var y MySlice[string] = []string{"1", "2", "3", "4"}

	fmt.Println(s.Sum())
	fmt.Println(t.Sum())
	fmt.Println(y.Sum())

}
