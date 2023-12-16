package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//时间函数签名
//需要调用的时候写time.???，后面是返回的数据类型，返回的数据内容与函数自然语言名称一致
//{func (t Time) Clock() (hour, min, sec int) // 返回 t 对应的时、分、秒
//func (t Time) Hour() int
//func (t Time) Minute() int
//func (t Time) Second() int
//func (t Time) Nanosecond() int
//func (t Time) Date() (year int, month Month, day int) // 返回 t 对应的年、月、日，
//创建的时候需要输入(年，月，日，时，分，秒)6个输入量，其中月可以是int类型，也可以是time。Month类型
//func (t Time) Year() int
//func (t Time) Month() Month
//func (t Time) Day() int}

//操作时间函数签名
//{func (t Time) Add(d Duration) Time  // `Add` 返回时间 t+d
//func (t Time) AddDate(years int, months int, days int) Time // `AddDate` 返回时间 t+date
//func (t Time) Sub(u Time) Duration  // `Sub` 返回时间 t-d
//func (t Time) Equal(u Time) bool    // `Equal` 判断 t 和 u 是否相同，
//注意如果精度不同，两个时间也是不一样的
//
//func (t Time) Before(u Time) bool   // `Before` 判断 t 是否在 u 之后
//func (t Time) After(u Time) bool    // `After` 判断 t 是否在 u 之前
//
//func Since(t Time) Duration         // `Since` 返回 t 到现在经过的时间
//func Until(t Time) Duration         // `Until` 返回现在到 t 的时间
//func (t Time) Truncate(d Duration) Time  // `Truncate` 将时间阶段为某一精度}

// 转化时间函数签名
// 时间转化到时间戳
// {func (t Time) Unix() int64        // 秒
// func (t Time) UnixMilli() int64   // 毫秒
// func (t Time) UnixMicro() int64   // 微秒
// func (t Time) UnixNano() int64    // 纳秒 }

// 时间戳转换到时间
// {func Unix(sec int64, nsec int64) Time   // 秒和纳秒
// func UnixMilli(msec int64) Time         // 毫秒
// func UnixMicro(usec int64) Time         // 微秒}

// 时间格式化转化为字符串
// {Go语言中格式化时间模板不是常见的 Y-m-d H:M:S 而是使用Go语言的诞生时间
//2006 年 1 月 2 号 15 点 04 分 05 秒，也就是 2006-01-02 15:04:05
//
//func (t Time) Format(layout string) string
//Format 根据 layout 指定的格式返回 t 代表的时间点的格式化文本表示}

// 字符串转化为时间
// {func Parse(layout, value string) (Time, error)
//Parse 根据 layout 指定的格式解析一个格式化的时间字符串并返回它代表的时间}

// 字符串解析为时间
// {func ParseDuration(s string) (Duration, error)
//ParseDuration 解析一个时间段字符串，如"300ms"、"-1.5h"、"2h45m"。
//合法的单位有"ns"、"us" /"µs"、"ms"、"s"、"m"、"h"}

// 时间段表示为float64
// {func (d Duration) Hours() float64
// func (d Duration) Minutes() float64
// func (d Duration) Seconds() float64
// 分别将时间段表示 float64 类型的小时、分钟、秒数}

//定时器
//{在 Go 语言中，time 包提供了 Timer 和 Ticker 两种类型，用于处理定时和周期性的时间事件。
//Timer
//func NewTimer(d Duration) *Timer
//Timer 类型表示一个单次事件。在未来的某个时间点，
//Timer 会发送一个值到它的 C 通道，使用 ticker 的 Stop 方法可以提前停止它

//Ticker
//
//func NewTicker(d Duration) *Ticker
//Ticker 表示一个周期事件，跟 Timer 差不多}

// {下面是示例运用
// 1.加载某一地区的时间
// 2.获取时分秒年月日
// 3.操作时间}
func main() {
	//1.加载某一地区的时间
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}
	now := time.Now()
	fmt.Println(now.In(loc))
	//{这段代码是一个Go语言程序的主函数（main函数），它用于获取并打印当前时间在上海时区的时间。
	//
	//首先，代码使用time.LoadLocation("Asia/Shanghai")来加载上海时区的地理位置信息。
	//如果加载失败，它会将错误消息打印出来并返回。
	//
	//然后，代码使用time.Now()获取当前的时间，并将其保存在变量now中。
	//
	//最后，代码使用fmt.Println(now.In(loc))将now变量的值转换为上海时区的时间，并打印出来。
	//
	//总的来说，这段代码的作用是获取并打印当前时间在上海时区的时间。}

	//2.获取时分秒年月日
	//(1)使用Clock()得到时分秒的三个参数
	hour, min, sec := now.Clock()
	fmt.Println("时：%d 分：%d 秒：%d\n", hour, min, sec)
	//(2)使用Hour(),Minute(),Second()单独获取每个参数
	fmt.Println("时：%d 分：%d 秒：%d\n", now.Hour(), now.Minute(), now.Second())
	//(3)使用Date()获取年月日的三个参数
	year, month, day := now.Date()
	fmt.Println("年：%d 月：%s 日：%d\n", year, month, day)
	// (4)使用 Year()、Month()、Day() 单独获取年、月、日
	fmt.Printf("年：%d 月：%s 日：%d\n", now.Year(), now.Month(), now.Day())

	//3.操作时间
	//(1)在当前时间上加上n小时/分钟/秒
	n := 2
	after_n_hour := now.Add(time.Duration(n) * time.Hour)
	//time.Hour的数据类型是time.Duiation类型，整型的n无法直接运算,
	//但是具体的数据例如2*time.Hour没问题
	//使用time.Minute/time.Second同理得到时间加法
	//在加的内容前面加负号“-”既是时间减法
	fmt.Println(after_n_hour)

	//(2)now.AddDate整体改变时间
	after_some_date := now.AddDate(1, 2, 3)
	fmt.Println(after_some_date)

	//(3)计算任意两个时间的时间差
	time1 := now.Add(2 * time.Hour)
	time2 := now.Add(-1 * time.Hour)
	shijiancha := time1.Sub(time2) //运算time1-time2
	fmt.Println(shijiancha)

	//(4)判断两个时间是否相同
	true_or_false := time1.Equal(time2)
	fmt.Println("两个时间是否相同：", true_or_false)

	//(5)判断时间的先后顺序
	true_or_false_Before := time1.Before(time2) //判断time1是否在time2之前
	true_or_false_After := time1.After(time2)   //判断time1是否在time2之后
	fmt.Println("time1是否在time2之前：", true_or_false_Before)
	fmt.Println("time1是否在time2之后", true_or_false_After)

	//(6)计算从某个时间到现在的时间差
	time1_until := time.Until(time1) //会返回时间差的绝对值
	fmt.Println(time1_until)

	//(7)截断时间到指定精度
	truncatedTime := now.Truncate(time.Hour)
	fmt.Println("截断到小时的时间：", truncatedTime)
	//{now.Truncate(time.Hour) 是 Go 语言中的一个时间截断操作。
	//它将给定时间 now 的分钟、秒和纳秒部分都设置为零，并将其舍入到最近的整点小时。
	//例如，如果 now 是 2023-11-20 14:38:12，
	//那么 truncatedTime 将会是 2023-11-20 14:00:00，分钟、秒和纳秒被设置为零，只保留了小时信息
	//这种操作通常用于处理时间精度，或者在进行时间比较时忽略较小的时间单位。}

	//(8)格式化时间输出
	formattedTime := time1.Format("2006-01-02 15:04:05")
	//这里括号内的"2006-01-02 15:04:05"时间只是示例格式，实际输出时间是time1
	fmt.Println("格式化时间：", formattedTime)

	//(9)解析字符串为时间
	timeStr := "2023-11-15 12:30:45"
	parsedTime, err := time.Parse("2006-01-02 15:04:05", timeStr)
	//这里括号内的"2006-01-02 15:04:05"时间只是示例格式，实际解析时间为timeStr
	if err != nil {
		fmt.Println("解析时间出错：", err)
	} else {
		fmt.Println("解析后的时间：", parsedTime)
	}
	durationStr := "1h30m45s"
	parsedDuration, err := time.ParseDuration(durationStr)
	if err != nil {
		fmt.Println("解析持续时间出错：", err)
	} else {
		fmt.Printf("解析后的持续时间：%v，总小时数：%f\n", parsedDuration, parsedDuration.Hours())
	}
	//{这段代码使用了Go语言中的持续时间解析操作。
	//time.ParseDuration(durationStr) 中的字符串durationStr表示待解析的持续时间字符串。
	//
	//假设 durationStr 是 "1h30m45s"，经过解析后，如果没有错误发生，
	//将会得到一个parsedDuration的持续时间对象，表示解析后的持续时间。
	//
	//如果解析过程中出现错误，那么 err 变量将不为空，并且可以通过 err.Error() 方法获取具体的错误信息。
	//
	//如果解析成功，将会输出 "解析后的持续时间："以及解析后的持续时间值，
	//并且通过 parsedDuration.Hours() 方法可以获取总小时数。
	//
	//注意，持续时间解析操作要求待解析的持续时间字符串符合一定的格式，
	//例如"1h30m45s"表示1小时30分45秒。在持续时间字符串中，"h"表示小时，"m"表示分钟，"s"表示秒。}

	//定时器
	// 创建一个通道来接收中断信号
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	// 启动一个goroutine等待中断信号
	//{在此示例中，我们使用os/signal包来创建一个接收中断信号的通道interrupt。Notify函数用于将中断信号（os.Interrupt）和系统终止信号（syscall.SIGTERM）发送到该通道。
	//
	//然后，我们启动一个goroutine来等待通道中的中断信号。一旦接收到中断信号，我们打印一条消息并调用os.Exit(0)来终止程序的执行。
	//
	//在主函数的剩余部分，你可以编写你的业务逻辑代码。定时器将继续按照设定的时间间隔触发并输出当前时间，直到接收到中断信号为止。
	//
	//当你想要中断程序时，可以通过按下键盘上的Ctrl+C（发送中断信号）来触发中断信号，并且程序将在接收到中断信号后退出。}
	go func() {
		<-interrupt
		fmt.Println("接收到中断信号，程序即将退出")
		os.Exit(0)
	}()

	k := 2
	ticker := time.NewTicker(time.Duration(k) * time.Second) // 创建一个以k秒为间隔的定时器
	defer ticker.Stop()                                      // 确保在函数结束时停止定时器

	for {
		<-ticker.C                         // 等待定时器通道触发,当定时器达到设定的时间间隔时，通道会发送一个时间值。
		fmt.Println("Tick at", time.Now()) // 输出当前时间
	}
}
