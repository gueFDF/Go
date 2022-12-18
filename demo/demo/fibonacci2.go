package demo

import (
	"fmt"
	"log"
	"math"
	"math/big"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

func Fib(n int) (value int) {
	if n <= 1 {
		value = 1
	} else {
		value = Fib(n-1) + Fib(n-2)
	}
	return
}

// 闭包写法
func Fib_close() func() int {
	a, b := 1, 1

	return func() int {
		a, b = b, a+b
		return b
	}
}

// 缓存优化写法
var fibs [1000]uint64

func Fib_catch(a int) uint64 {
	if a == 0 || a == 1 {
		fibs[a] = 1
	} else if fibs[a] == 0 {
		fibs[a] = Fib_catch(a-1) + Fib_catch(a-2)
	} else {
		return fibs[a]
	}
	return fibs[a]
}

func Print(n int) {
	if n <= 0 {
		fmt.Println(n)
		return
	} else {
		fmt.Println(n)
		Print(n - 1)
	}
}

func Factoral(n uint64) (value uint64) {
	value = 1
	if n > 0 {
		value = n * Factoral(n-1)
	}
	return
}

func Test_map() {
	asciiOnly := func(c rune) rune {
		if c > 127 {
			return ' '
		}
		return c
	}
	fmt.Println(strings.Map(asciiOnly, "Jérôme Österreich"))
}

func Adder(a int) func(b int) int {
	return func(b int) int {
		return a + b
	}
}

func MakeAddSuffix(sufixx string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, sufixx) {
			return name + sufixx
		} else {
			return name
		}
	}
}

// 使用闭包进行debug
func Debug() func() {
	return func() {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("%s:%d", file, line)
	}
}

// 扩容切片,增长因子为factor
func Enlarge(s []int, factor int) []int {
	news := make([]int, len(s)*factor)
	copy(news, s)
	s = news
	return s
}

func Test_regexp() {
	searchIn := "John: 2578.34 William: 4567.23 Steve: 5632.18"
	pat := "[0-9]+.[0-9]+" //正则

	f := func(s string) string {
		v, _ := strconv.ParseFloat(s, 32)
		fmt.Println("v: ", v)
		return strconv.FormatFloat(v*2, 'f', 2, 32)
	}

	if ok, _ := regexp.Match(pat, []byte(searchIn)); ok {
		fmt.Println("Match Found!")
	}

	//解析一个正则表达式，如果成功则返回一个可用于匹配文本的Regexp对象
	re, _ := regexp.Compile(pat)
	//将匹配到的部分替换为"##.#"
	str := re.ReplaceAllString(searchIn, "##.#")
	fmt.Println(str)
	//将匹配到的部分用函数f处理
	str2 := re.ReplaceAllStringFunc(searchIn, f)
	fmt.Println(str2)
}

// 测试big包
func Test_big() {
	im := big.NewInt(math.MaxInt64)
	in := im
	io := big.NewInt(1956)
	ip := big.NewInt(1)
	ip.Mul(im, in).Add(ip, im).Div(ip, io)
	fmt.Printf("Big Int: %v\n", ip)
	// Here are some calculations with bigInts:
	rm := big.NewRat(math.MaxInt64, 1956)
	rn := big.NewRat(-1956, math.MaxInt64)
	ro := big.NewRat(19, 56)
	rp := big.NewRat(1111, 2222)
	rq := big.NewRat(1, 1)
	rq.Mul(rm, rn).Add(rq, ro).Mul(rq, rp)
	fmt.Printf("Big Rat: %v\n", rq)
}

type test_car interface {
	start()
}


type Car struct {
	test_car
}