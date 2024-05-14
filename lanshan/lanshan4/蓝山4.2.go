package main

import (
	"fmt"
	"strconv"
)

//strconv包
// 字符串到基本数据类型
// 1.字符到int
// func Atoi(s string)(int error)

// 2.字符到int64或者uint64
// func ParseInt(s string,base int,bitSize int)(i int64,err error)
//func ParseUint(s string, base int, bitSize int) (uint64, error)
// base int：表示要使用的进制。通常使用的是10进制，也可以是2、8、16进制等。
// bitSize int：表示目标整数的位数大小。通常使用的是32位或64位。
//i int64：转换后的整数值（以int64类型表示）。
//err error：在转换过程中出现的错误（如果有）。如果转换成功，则err为nil。

//3.字符串到浮点数
//func ParseFloat(s string, bitSize int) (float64, error)

// 4.判断字符到bool值
//func ParseBool(str string) (value bool, err error)
//ParseBool 返回字符串表示的bool值。
//它接受1、0、t、f、T、F、true、false、True、False、TRUE、FALSE；否则返回错误

//基本数据类型到字符串
//1.int到字符串
//func Itoa(i int)string

// 2.int或者uint到字符串
//func FormatInt(i int64, base int) string
//func FormatUint(i uint64, base int) string

// 3.浮点数到字符串
// func FormatFloat(f float64, fmt byte, prec, bitSize int) string
//f float64：表示要格式化的浮点数。
//
//fmt byte：表示格式化的类型。常见的取值有：
//
//'b'：科学计数法（例如1.23e+10）。
//'e'：科学计数法（例如1.23e+10）。
//'E'：科学计数法，使用大写E（例如1.23E+10）。
//'f'：十进制表示，不带指数（例如123456.789）。
//'g'：根据情况选择科学计数法或十进制表示。
//'G'：根据情况选择科学计数法或十进制表示，使用大写E。
//'x'：十六进制表示，使用小写字母a-f。
//'X'：十六进制表示，使用大写字母A-F。
//prec int：表示要保留的小数位数（对于'f'、'e'、'E'和'g'格式）或总位数（对于'b'和'x'格式）。
//bitSize int：表示浮点数的位数大小。通常使用的是32位或64位。

// 4.bool到字符串
// func FormatBool(b bool) string
//根据 b 的值返回"true"或"false"

func main() {
	//字符串转化为整数
	inta := "123"
	int_a, err := strconv.Atoi(inta)
	if err != nil {
		fmt.Println("出错:", err)
	} else {
		fmt.Println("值:", int_a)
	}
	//指定进制，字符串转化为整数
	hexa := "1a"
	hex_a, err := strconv.ParseInt(hexa, 16, 0)
	if err != nil {
		fmt.Println("出错:", err)
	} else {
		fmt.Println("值:", hex_a)
	}
	//指定进制，字符串转化为无符号数
	uinta := "255"
	uint_a, err := strconv.ParseUint(uinta, 10, 0)
	if err != nil {
		fmt.Println("出错:", err)
	} else {
		fmt.Println("值:", uint_a)
	}
	//字符串转化为浮点数
	floata := "3.14"
	float_a, err := strconv.ParseFloat(floata, 64)
	if err != nil {
		fmt.Println("出错:", err)
	} else {
		fmt.Println("值:", float_a)
	}
	//字符串转化为布尔值
	boola := "true"
	bool_a, err := strconv.ParseBool(boola)
	if err != nil {
		fmt.Println("出错:", err)
	} else {
		fmt.Println("值:", bool_a)
	}
	// 整数转换为字符串
	intValueToStr := 42
	intStrFromInt := strconv.Itoa(intValueToStr)
	fmt.Println("整数转字符串:", intStrFromInt)
	// 格式化整数为字符串
	otherIntValueToStr := 12345
	otherIntStrFromInt := strconv.FormatInt(int64(otherIntValueToStr), 10)
	fmt.Println("格式化整数为字符串:", otherIntStrFromInt)
	// 格式化无符号整数为字符串
	uintValueToStr := uint64(98765)
	uintStrFromUint := strconv.FormatUint(uintValueToStr, 10)
	fmt.Println("格式化无符号整数为字符串:", uintStrFromUint)
	// 格式化浮点数为字符串
	floatValueToStr := 2.71828
	floatStrFromFloat := strconv.FormatFloat(floatValueToStr, 'f', -1, 64)
	fmt.Println("格式化浮点数为字符串:", floatStrFromFloat)
	// 布尔值转换为字符串
	boolValueToStr := true
	boolStrFromBool := strconv.FormatBool(boolValueToStr)
	fmt.Println("布尔值转字符串:", boolStrFromBool)
}
