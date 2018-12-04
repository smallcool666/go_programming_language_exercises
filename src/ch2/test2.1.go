package main
/*
	向tempconv包添加类型、常量和函数用来处理Kelvin绝对温度的转换，Kelvin 绝对零度是−273.15°C，Kelvin绝对温度1K和摄氏度1°C的单位间隔是一样的。
*/
import (
	"fmt"
	"ch2/tempconv"
)

type Kelvin float64

const (
	AbsoluteZeroC Kelvin = 0
	FreezingC     Kelvin = 273.15
	BoilingC      Kelvin = 373.15
)

func (k Kelvin) String() string{
	return fmt.Sprintf("%g°T", k)
}

func (k Kelvin) KToC() tempconv.Celsius {
	return tempconv.Celsius(k - 273.15)
}

func  main()  {
	fmt.Println(Kelvin.KToC(AbsoluteZeroC))
	fmt.Println(Kelvin.KToC(FreezingC))
	fmt.Println(Kelvin.KToC(BoilingC))
}
