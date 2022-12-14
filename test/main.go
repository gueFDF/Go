package main

import (
	_ "Go/test/hello"
	"fmt"
	"math"
	"strings"
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
			fmt.Printf("(%d,%d)=%d ",k1,k2,v2)
		}
	}
	fmt.Println()
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

	// for k1, v1 := range f {
	// 	for k2, v2 := range v1 {
	// 		fmt.Printf("(%d,%d)=%d\n", k1, k2, v2)
	// 	}
	// }

	printarr(&f)
}
