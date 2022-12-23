package demo

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"math"
	"math/big"
	"os"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
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

type Test_car interface {
	Start()
}

type Car struct {
	Name string
}

func (p *Car) Start() {
	fmt.Println(p.Name)
}

func (p *Car) End() {
	fmt.Println(p.Name)
}

type Person struct {
	Name string
	Age  int
}

// 类型的String方法,用来自定义打印格式，调用打印函数时会自动调用
func (p *Person) String() string {
	return p.Name + strconv.Itoa(p.Age)
}

/*
Go 开发者不需要写代码来释放程序中不再使用的变量和结构占用的内存，在 Go 运行时中有一个独立的进程，即垃圾收集器 (GC)，
会处理这些事情，它搜索不再使用的变量然后释放它们的内存。可以通过 runtime 包访问 GC 进程。

通过调用 runtime.GC() 函数可以显式的触发 GC，但这只在某些罕见的场景下才有用，
比如当内存资源不足时调用 runtime.GC()，它会在此函数执行的点上立即释放一大片内存，
此时程序可能会有短时的性能下降（因为 GC 进程在执行）
*/

/*如果需要在一个对象 obj 被从内存移除前执行一些特殊操作，比如写到日志文件中，可以通过如下方式调用函数来实现：

runtime.SetFinalizer(obj, func(obj *typeObj))
*/

//利用空接口实现一个stack

type Stack []interface{}

func (p Stack) Len() int {
	return len(p)
}

func (p Stack) Cap() int {
	return cap(p)
}

func (p Stack) IsEmpty() bool {
	return len(p) == 0
}

func (p *Stack) Push(e interface{}) {
	*p = append(*p, e)
}
func (p Stack) Top() (interface{}, error) {
	if len(p) == 0 {
		return nil, errors.New("stack is empty")
	}
	return p[len(p)-1], nil
}
func (p *Stack) Pop() (interface{}, error) {
	stk := *p
	if len(stk) == 0 {
		return nil, errors.New("stack is enpty")
	}

	top := stk[len(stk)-1]
	*p = stk[:len(stk)-1]
	return top, nil
}

func Test_Reader() {
	nchars, nwords, nlines := 0, 0, 0
	inputPeader := bufio.NewReader(os.Stdin)

	for {
		input, err := inputPeader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}

		if input == "S\n" {
			fmt.Printf("%d %d %d \n", nchars, nwords, nlines)
			os.Exit(0)
		} else {
			nchars += len(input) - 1
			nwords = len(strings.Fields(input))
			nlines++
		}
	}

}

//读写文件练习

type Page struct {
	Title string
	Body  []byte
}

func (w *Page) Save() {
	f, err := os.OpenFile(w.Title, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
		fmt.Println("save")
		os.Exit(0)
	}

	defer f.Close()
	inputWriter := bufio.NewWriter(f)

	inputWriter.WriteString(string(w.Body))
	fmt.Println(string(w.Body))

	inputWriter.Flush()
}

func (r *Page) Load(filename string) {
	f, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
		fmt.Println("load")
		os.Exit(0)
	}
	defer f.Close()
	outputReader := bufio.NewReader(f)
	str, _ := outputReader.ReadString('\n')
	r.Body = []byte(str)
}

// 解析命令行参数
var NewLine = flag.Bool("n", false, "print newline") // echo -n flag, of type *bool

const (
	Space   = " "
	Newline = "\n"
)

func Test_echo() {
	flag.PrintDefaults()
	flag.Parse() // Scans the arg list and sets up flags
	var s string = ""
	for i := 0; i < flag.NArg(); i++ {
		if i > 0 {
			s += " "
			if *NewLine { // -n is parsed, flag becomes true
				s += Newline
			}
		}
		s += flag.Arg(i)
	}
	os.Stdout.WriteString(s)
}

var wg sync.WaitGroup

func Test_gre() {

	for i := 0; i < 10; i++ {
		wg.Add(1) // 启动一个goroutine就登记+1
		go func(j int) {
			defer wg.Done()
			fmt.Println("Hello Goroutine!", j)
			time.Sleep(time.Second)
		}(i)
	}
	wg.Wait() // 等待所有登记的goroutine都结束
}

// channel 练习
func Test_chan() {
	sig := make(chan int)
	ch1 := make(chan int)
	ch2 := make(chan int)
	// 开启goroutine将0~100的数发送到ch1中
	go func() {

		fmt.Println("进入阻塞")
		sig <- 1
		fmt.Println("被唤醒")
		x := 100
		for i := 0; i < x; i++ {
			ch1 <- i
		}
		close(ch1)
	}()
	// 开启goroutine从ch1中接收值，并将该值的平方发送到ch2中
	go func() {
		time.Sleep(time.Second)
		fmt.Println("唤醒生产者")
		<-sig
		for {
			i, ok := <-ch1 // 通道关闭后再取值ok=false
			if !ok {
				break
			}
			ch2 <- i * i
		}
		close(ch2)
	}()
	// 在主goroutine中从ch2中接收值打印

	for i := range ch2 { // 通道关闭后会退出for range循环

		fmt.Println(i)
	}
}




// 实现素数筛子
func filter(p1 chan int, p2 chan int, pp int) {
	for v:= range p1 {
		if v%2!=0 {
			p2<-v
		}
	}
	close(p2)
}

func Test_prime() {
	ch := make(chan int)

	//将所有数写入管道
	go func(Ch chan int) {
		for i := 2; i <= 100; i++ {
			Ch <- i
		}
		//写入-1表示写入结束
		close(Ch)
	}(ch)

	//进行筛选
	for {
		prime ,ok:= <-ch
		if !ok {
			break
		}
		fmt.Printf("%d ", prime)
		ch2 := make(chan int)
		go filter(ch, ch2, prime)
		ch = ch2
	}
}
