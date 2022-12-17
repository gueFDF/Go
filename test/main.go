package main

import (
	_ "Go/test/hello"
	"fmt"
	"math"
	"math/rand"
	"sort"
	"strings"
	"time"
)

// 遍历字符串
func traversalString() {
	//无论哪种转换，都会重新分配内存，并复制字节数组。
	s := "pprof.cn博客"
	for i := 0; i < len(s); i++ { //byte
		fmt.Printf("%v(%c) ", s[i], s[i])
	}
	fmt.Println()
	for _, r := range s { //rune
		fmt.Printf("%v(%c) ", r, r)
	}
	fmt.Println()
}

// 修改字符串
func changeString() {
	s1 := "hello"
	// 强制类型转换
	byteS1 := []byte(s1)
	byteS1[0] = 'H'
	fmt.Println(string(byteS1))

	s2 := "博客"
	runeS2 := []rune(s2)
	runeS2[0] = '狗'
	fmt.Println(string(runeS2))
}

func sqrtDemo() {
	var a, b = 3, 4

	// math.Sqrt()接收的参数是float64类型，需要强制转换
	c := int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

//数组

/*
   1. 数组：是同一种数据类型的固定长度的序列。
   2. 数组定义：var a [len]int，比如：var a [5]int，数组长度必须是常量，且是类型的组成部分。一旦定义，长度不能变。
   3. 长度是数组类型的一部分，因此，var a[5] int和var a[10]int是不同的类型。
   4. 数组可以通过下标进行访问，下标是从0开始，最后一个元素下标是：len-1
   for i := 0; i < len(a); i++ {
   }
   for index, v := range a {
   }
   5. 访问越界，如果下标在数组合法范围之外，则触发访问越界，会panic
   6. 数组是值类型，赋值和传参会复制整个数组，而不是指针。因此改变副本的值，不会改变本身的值。
   7.支持 "=="、"!=" 操作符，因为内存总是被初始化过的。
   8.指针数组 [n]*T，数组指针 *[n]T。

*/
//数组在全局的声名
var arr0 [5]int = [5]int{1, 2, 3}
var arr1 = [5]int{1, 2, 3, 4, 5}
var arr2 = [...]int{1, 2, 3, 4, 5, 6}
var str = [5]string{3: "hello world", 4: "tom"}

func printarr(arr *[4][2]int) {
	for k1, v1 := range arr {
		for k2, v2 := range v1 {
			fmt.Printf("(%d,%d)=%d ", k1, k2, v2)
		}
	}
	fmt.Println()
}

func mytest(arr [5]int, key int) {

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == key {
				fmt.Println(arr[i], arr[j])
			}
		}
	}
}

func plusOne(digits []int) []int {

	arr := []int{}
	arr1 := []int{}
	temp := 1
	for i := len(digits) - 1; i >= 0; i-- {
		arr = append(arr, (temp+digits[i])%10)
		temp = (digits[i] + temp) / 10
	}
	if temp == 1 {
		arr = append(arr, 1)
	}
	for j := len(arr) - 1; j >= 0; j-- {
		arr1 = append(arr1, arr[j])
	}

	return arr1
}

//切片

/*
1. 切片：切片是数组的一个引用，因此切片是引用类型。但自身是结构体，值拷贝传递。
2. 切片的长度可以改变，因此，切片是一个可变的数组。
3. 切片遍历方式和数组一样，可以用len()求长度。表示可用元素数量，读写操作不能超过该限制。
4. cap可以求出slice最大扩张容量，不能超出数组限制。0 <= len(slice) <= len(array)，其中array是slice引用的数组。
5. 切片的定义：var 变量名 []类型，比如 var str []string  var arr []int。
6. 如果 slice == nil，那么 len、cap 结果都等于 0。
*/

func test() {

	s1 := []int{1, 2, 3, 4, 5}
	fmt.Printf("slice s1 : %v\n", s1)
	s2 := make([]int, 10)
	fmt.Printf("slice s2 : %v\n", s2)
	copy(s2, s1)
	fmt.Printf("copied slice s1 : %v\n", s1)
	fmt.Printf("copied slice s2 : %v\n", s2)
	s3 := []int{1, 2, 3}
	fmt.Printf("slice s3 : %v\n", s3)
	s3 = append(s3, s2...)
	fmt.Printf("appended slice s3 : %v\n", s3)
	s3 = append(s3, 4, 5, 6)
	fmt.Printf("last slice s3 : %v\n", s3)

}

// 结构体
type person struct {
	name string
	city string
	age  int8
}

type student struct {
    name string
    age  int
}

func (p student) printname() {
	fmt.Println(p.name)
}

func  init() {
	fmt.Printf("sefsdfsdfsdf")
}

func main() {
	//hello.Print()
	//编译报错：./main.go:6:5: undefined: hello
	const (
		n1 = iota
		n2
		n3 = iota
		n4
	)
	const n5 = iota
	s1 := "hello"
	s2 := "world"
	S1 := []byte(s1)
	S1[0] = 'H'
	var s3 string = " "
	s3 = s1 + s2
	fmt.Println(n1, n2, n3, n4, n5)
	fmt.Println(`str := \"c:\\pprof\\main.exe\"`)
	fmt.Println(s3)
	fmt.Println(`"12345">"2"`, s1 == s2)
	fmt.Println(len(s1))
	count := strings.Count("asdasdasdadad", "as")
	fmt.Println(count)

	traversalString()
	changeString()
	sqrtDemo()

	//数组在局部的声明
	a := [3]int{1, 2}           // 未初始化元素值为 0。
	b := [...]int{1, 2, 3, 4}   // 通过初始化值确定数组长度。
	c := [5]int{2: 100, 4: 200} // 使用索引号初始化元素。
	d := [...]struct {
		name string
		age  uint8
	}{
		{"user1", 10}, // 可省略元素类型。
		{"user2", 20}, // 别忘了最后一行的逗号。
	}

	f := [4][2]int{1: {1, 2}, 3: {1, 3}}
	fmt.Println(a, b, c, d)
	fmt.Println(f)

	arrtest := [5]int{1, 2, 3, 4, 5}
	mytest(arrtest, 5)
	// for k1, v1 := range f {
	// 	for k2, v2 := range v1 {
	// 		fmt.Printf("(%d,%d)=%d\n", k1, k2, v2)
	// 	}
	// }

	printarr(&f)

	//切片
	//1.声明切片
	var a1 []int
	if a1 == nil {
		fmt.Println("是空")
	} else {
		fmt.Println("不是空")
	}
	// 2.:=
	a2 := []int{}
	// 3.make()
	var a3 []int = make([]int, 0)
	fmt.Println(a1, a2, a3)
	// 4.初始化赋值
	var a4 []int = make([]int, 0, 10)
	fmt.Println(a4)
	a5 := []int{1, 2, 3}
	fmt.Println(a5)
	// 5.从数组切片
	arr := [5]int{1, 2, 3, 4, 5}
	arr[2] = 100
	//var a6 []int
	// 前包后不包
	a6 := arr[1:3]
	fmt.Println(cap(a6), len(a6))
	fmt.Println(a6)
	a10 := arr[:]
	pp := [5]struct {
		x int
	}{}

	slice := pp[:]
	fmt.Println(slice)
	fmt.Printf("%p  %v\n", slice, pp)

	fmt.Println(plusOne(a10))
	test()

	//map的使用
	scoreMap := make(map[string]int)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	// 如果key存在ok为true,v为对应的值；不存在ok为false,v为值类型的零值
	v, ok := scoreMap["张三"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("查无此人")
	}

	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	scoreMap["王五"] = 60
	for k, v := range scoreMap {
		fmt.Println(k, v)
	}

	rand.Seed(time.Now().UnixNano()) //初始化随机种子
	var mymap = make(map[string]int, 200)
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i)
		vlaue := rand.Intn(100)
		mymap[key] = vlaue
	}

	//取出map中的key进行排序
	var keys = make([]string, 0, 200)
	for key := range mymap {
		keys = append(keys, key)
	}

	//对切片进行排序
	sort.StringSlice.Sort(keys)

	for _, key := range keys {
		fmt.Println(key, mymap[key])
	}

	//map切片
	var mapSlice = make([]map[string]string, 3)
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
	fmt.Println("after init")
	// 对切片中的map元素进行初始化
	mapSlice[0] = make(map[string]string, 10)
	mapSlice[0]["name"] = "王五"
	mapSlice[0]["password"] = "123456"
	mapSlice[0]["address"] = "红旗大街"
	mapSlice = append(mapSlice, make(map[string]string, 10))
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}

	var sliceMap = make(map[string][]string, 3)
	fmt.Println(sliceMap)
	fmt.Println("after init")
	key := "中国"
	value, ok := sliceMap[key]
	if !ok {
		value = make([]string, 0, 2)
	}

	value = append(value, "北京", "上海")
	sliceMap[key] = value
	fmt.Println(sliceMap)

	p3 := &person{}
	p3.age = 18
	p3.name = "xuaioguo"
	p3.city = "hanzhong"
	fmt.Printf("p3: %v\n", *p3)

	m := make(map[string]*student)
	stus := []student{
		{name: "pprof.cn", age: 18},
		{name: "测试", age: 23},
		{name: "博客", age: 28},
	}

	for _, stu := range stus {
		m[stu.name] = &stu
	}
	for k, v := range m {
		fmt.Println(k, "=>", v.name)
		v.printname()
	}
}
