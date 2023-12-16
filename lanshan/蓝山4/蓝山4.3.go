package main

import (
	"fmt"
	"io"
	"os"
)

//os包
//文件
//1.打开一个文件只用于读取
//func Open(name string)(file*File,err error)
//name string：表示要打开的文件的名称（包括路径）。
//函数返回两个值：
//file *File：表示打开的文件。它是一个指向文件的指针。
//err error：在打开文件过程中出现的错误（如果有）。
//如果打开文件成功，则err为nil。

//2.指定方式打开文件或者创建文件
//func OpenFile(name string, flag int, perm FileMode) (*File, error)
//使用指定的标志 flag（如 O_RDONLY 等）打开指定名称的文件，
//或使用指定的权限模式 perm（如 0666）创建文件
//flag 可以选下列的值:
//
//os.O_RDONLY：只读
//os.O_WRONLY：只写
//os.O_RDWR：读写
//os.O_APPEND：追加写
//os.O_CREATE：如果文件不存在则创建
//os.O_TRUNC：打开时清空文件
//os.O_EXCL: 和 os.O_CREATE 配合使用，文件必须不存在
//os.O_TRUNC: 打开时清空文件
//perm 表示文件的权限，
//在 Unix 和类 Unix 系统（如 Linux）中，
//文件的权限通常用一个三位数的八进制数来表示，这个数被称为权限模式。
//其中，第一位表示文件类型，
//接下来的三位表示所有者（owner）的权限，
//再接下来的三位表示所属组（group）的权限，
//最后三位表示其他用户的权限，所以这个参数在 Windows 下是没有意义的

//3.创建文件
//func Create(name string) (file *File, err error)

//4.删除文件或目录
//func Remove(path string) error
//删除 name 指定的文件或目录
//func RemoveAll(path string) error
//删除 path 指定的文件，
//如果 path 是目录，则删除它包含的所有下级对象，
//它会尝试删除所有对象，除非遇到错误

//5.写入文件
//func(f*File)Write(b []byte)(n int,err error)
//向文件中写入 len(b) 字节数据。它返回写入的字节数和可能遇到的任何错误。

//6.读取文件
//func (f *File) Read(b []byte) (n int, err error)
//从f中读取最多 len(b) 字节数据并写入 b。它返回读取的字节数和可能遇到的任何错误。

//7.重命名文件或目录
//func Rename(oldpath, newpath string) error
//修改 oldpath 指定的文件或目录，或移动一个文件

//8.获取当前工作目录
//func Getwd()(dir string,err error)

//9.改变当前工作目录
//func Chdir(dir string)error
//将工作目录修改为 dir 指定的目录

//10.判断错误
//func IsNotExist(err error) bool // 文件不存在
//func IsExist(err error) bool // 文件已经存在
//func IsPermission(err error) bool // 无权限
//func IsTimeout(err error) bool // 超时

//环境变量
//1.获取所有环境变量
//func Environ()[]string
//获取所有的环境变量，并返回格式为 key=value 的字符串切片

//2.获取指定的环境变量
//func Getenv(key string)string
//获取名为key的环境变量，若不存在会返回空字符串

//3.设置环境变量
//func Setenv(key,value,string)error
//设置名为key的环境变量

// 命令行参数
// os 包中使用 Args 字符串切片表示命令行参数，
// 其中第一个参数 os.Args[0] 是程序名
func main() {
	//创建文件
	newfile, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("出错：", err)
	} else {
		defer newfile.Close()
		fmt.Println("成功创建文件：", newfile.Name())
	}
	//defer newfile.Close()语句的作用是在当前函数执行完毕后
	// （无论是正常返回还是发生错误导致的异常返回），
	//自动调用newfile.Close()来关闭文件。
	//这样做的好处是确保文件在函数执行完毕时被正确关闭，
	//从而释放文件资源并避免资源泄漏。
	//打开文件
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("出错：", err)
	} else {
		defer file.Close()
		fmt.Println("打开文件：", file.Name())
	}
	// 打开文件（使用 OpenFile）
	anotherFile, err := os.OpenFile("anotherfile.txt", os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("打开文件出错:", err)
	} else {
		defer anotherFile.Close()
		fmt.Println("成功打开文件:", anotherFile.Name())
	}
	// 写入文件
	dataToWrite := []byte("Hello, Golang!")
	bytesWritten, err := anotherFile.Write(dataToWrite)
	if err != nil {
		fmt.Println("写入文件出错:", err)
		return
	}
	fmt.Printf("成功写入 %d 字节到文件\n", bytesWritten)
	// 读取文件
	readBuffer := make([]byte, 100)
	bytesRead, err := file.Read(readBuffer)
	if err != nil {
		if err != io.EOF {
			fmt.Println("读取文件出错:", err)
			return
		}
	}
	fmt.Printf("从文件中读取 %d 字节的数据: %s\n", bytesRead, readBuffer[:bytesRead])

	// 移除文件
	_, _ = os.Create("toRemove.txt")
	err = os.Remove("toRemove.txt")
	if err != nil {
		fmt.Println("删除文件出错:", err)
	} else {
		fmt.Println("成功删除文件")
	}

	// 重命名文件
	err = os.Rename("anotherfile.txt", "renamedfile.txt")
	if err != nil {
		fmt.Println("重命名文件出错:", err)
	} else {
		fmt.Println("成功重命名文件")
	}

	// 创建目录
	err = os.Mkdir("example_dir", 0755) // 0755是目录权限，表示所有者具有读写执行权限，其他人具有读和执行权限
	if err != nil {
		fmt.Println("创建目录出错:", err)
		return
	}
	// 改变工作目录
	err = os.Chdir("example_dir")
	if err != nil {
		fmt.Println("改变工作目录出错:", err)
	} else {
		newCurrentDir, _ := os.Getwd()
		fmt.Println("新的工作目录:", newCurrentDir)
	}

	// 获取当前工作目录
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("获取当前工作目录出错:", err)
	} else {
		fmt.Println("当前工作目录:", currentDir)
	}
	_, _ = os.Create("dir_file")
	_ = os.Chdir("..")

	// 递归删除目录及其内容
	err = os.RemoveAll("example_dir")
	if err != nil {
		fmt.Println("删除目录出错:", err)
	} else {
		fmt.Println("成功删除目录及其内容")
	}

	// 获取环境变量
	goPath := os.Getenv("GOPATH")
	fmt.Println("GOPATH环境变量:", goPath)

	// 设置环境变量
	err = os.Setenv("MY_VARIABLE", "my_value")
	if err != nil {
		fmt.Println("设置环境变量出错:", err)
	} else {
		fmt.Print("成功设置环境变量：")
		fmt.Println(os.Getenv("MY_VARIABLE"))
	}

	// 获取命令行参数
	fmt.Println("命令行参数:", os.Args)

}
