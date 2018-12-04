package main
/*
	写一个通用的单位转换程序，用类似cf程序的方式从命令行读取参数，如果缺省的话则是从标准输入读取参数，
然后做类似Celsius和Fahrenheit的单位转换，长度单位可以对应英尺和米，重量单位可以对应磅和公斤等。
*/
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Kilogram float64
type Pound float64

func (k Kilogram) String() string {
	return fmt.Sprintf("%g kg", k)
}

func (p Pound) String() string {
	return fmt.Sprintf("%g lb", p)
}

func (k Kilogram) KToP() Pound {
	return Pound(k * 2.20462)
}

func (p Pound) PToK() Kilogram {
	return Kilogram(p * 0.453592)
}

func main() {
	var res string
	params := os.Args[1:]
	if len(params) == 0 {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			param := input.Text()
			res = getNum(param)
			fmt.Println(param + " = " + res)
		}
	} else {
		for _, v := range params {
			res = getNum(v)
			fmt.Println(v + " = " + res)
		}
	}
}

func getNum(str string) string {
	var unit string
	//判断后缀，非法直接返回
	param := strings.Trim(str, " ")
	if strings.HasSuffix(param, "kg") {
		unit = "kg"
	} else if strings.HasSuffix(param, "lb") {
		unit = "lb"
	} else {
		return fmt.Sprint("must has kg or lb \n")
	}

	//获取数字并转换成float
	numStr := strings.Trim(param, unit)
	num, err := strconv.ParseFloat(numStr, 64)
	if err != nil {
		return fmt.Sprint("please input a int or float type data\n")
	}
	//kg转换成lb
	if unit == "kg" {
		return Pound.String(Kilogram.KToP(Kilogram(num)))
	}
	return  Kilogram.String(Pound.PToK(Pound(num)))
}
