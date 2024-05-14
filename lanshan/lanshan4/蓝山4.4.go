package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

//I/O库
//I/O是指输入/输出（Input/Output）
//读出
//type Reader interface{
//    Read(p []byte)(n int,err error)
//}

//写入
//tyoe Writer interface{
//    Write(p []byte)(n int,err error)
//}
//还有两个接口分别是

//type Closer interface {
//Close() error
//}

//type Seeker interface {
//Seek(offset int64, whence int) (int64, error)
//}
//Closer 表示可以关闭的对象，Close 方法用于关闭对象，释放资源
//Seeker 表示可以随机读取的对象，Seeker 方法将当前读写位置设置为偏移量 offset 字节之后的位置，whence 可以使 0、1、2
//0：表示相对于文件起始位置，offset 必须为非负数
//1：表示相对于当前位置，offset 可以为负数
//2：表示相对于文件结尾，offset 可以为负数
//type ReaderAt interface {
//    ReadAt(p []byte, off int64) (n int, err error)     // 从指定位置读取
//}
//
//type ReaderFrom interface {
//    ReadFrom(r Reader) (n int64, err error)     // 从指定的 Reader 中读取
//}

//func Copy(dst Writer,src Reader)(written int64, err error)
//func CopyN(dst Writer, src Reader, n int64) (written int64, err error)
//func ReadAll(r Reader) ([]byte, error)
//func ReadAtLeast(r Reader, buf []byte, min int) (n int, err error)
//func ReadFull(r Reader, buf []byte) (n int, err error)

//bufio库
//Writer结构体
//type Writer struct{
//    err error
//    buf []byte
//    n int
//    wr io.Writer}

//func NewWriter(w io.Writer) *Writer
//获取一个以 w 作为底层 io.Writer 的 bufio.Writer
//func NewWriterSize(w io.Writer, size int) *Writer
//获取一个以 w 作为底层 io.Writer 缓冲区大小为 size 的 bufio.Writer

//写入缓存区
//func (b *Writer) Write(p []byte) (nn int, err error)
//将字节切片 p 的内容写入缓存中
//func (b *Writer) WriteByte(c byte) error
//将一个字符串写入缓存中
//func (b *Writer) WriteByte(c byte) error
//写入单个字节
//func (b *Writer) WriteRune(r rune) (size int, err error)
//写入一个 unicode 码值

//操作缓存区
//func (b *Writer) Flush() error
//缓存中的所有数据写入到底层的 io.Writer 对象中
//func (b *Writer) Available() int
//返回底层缓冲区的字节数
//func (b *Writer) Reset(w io.Writer)
//清除缓存，并将底层 io.Writer 对象设置为 w
//func (b *Writer) Size() int
//返回底层缓冲区的大小

//Reader
//bufio.Reader 是一个带有缓冲区的 io.Reader 接口的实现，
//它会在内存中存储从底层 io.Reader 中读取到的数据，然后先从内存缓冲区中读取数据，
//这样可以减少访问底层 io.Reader 对象的次数以及减轻操作系统的压力。结构体定义：
//
//type Reader struct {
//    buf          []byte
//    rd           io.Reader // reader provided by the client
//    r, w         int       // buf read and write positions
//    err          error
//    lastByte     int // last byte read for UnreadByte -1 means invalid
//    lastRuneSize int //size of last rune read for UnreadRune; -1 means invalid
//}

//初始化缓冲区
//func NewReader(rd io.Reader) *Reader
//获取一个以 rd 作为底层 io.Reaer 的 bufio.Reader
//func NewReaderSize(rd io.Reader, size int) *Reader
//获取一个以 rd 作为底层 io.Reaer 缓冲区大小为 size 的 bufio.Reaer

//读取缓存区
//r，w 是两个偏移量表示缓冲区中读写的位置。当从缓冲区中读取数据时，r 增加，
//当调用底层 io.Reader 的 Read 方法读取数据到缓冲区时，w 增加。
//func (b *Reader) Read(p []byte) (n int, err error)
//从缓冲区中读取数据到 p 中
//func (b *Reader) ReadByte() (byte, error)
//从缓冲区中读取一个字节
//func (b *Reader) ReadRune() (r rune, size int, err error)
//从缓冲区中读取一个 UTF-8 编码的字符
//func (b *Reader) ReadLine() (line []byte, isPrefix bool, err error)
//从缓冲区中读取一行

//操作
//func (b *Reader) Peek(n int) ([]byte, error)
//读取 n 个字节，但不改变偏移量
//func (b *Reader) Reset(r io.Reader)
// 清除缓存，并将底层的 io.Reader 重新设置为传入的 r

//strings & bytes
//由于 Go 中的字符串是只读的，所以 strings 包中只实现了读相关的接口，结构体定义如下：
//
//type Reader struct {
//    s        string
//    i        int64 // current reading index
//    prevRune int   // index of previous rune; or < 0
//}
//初始化
//func NewReader(s string) *Reader       // 创建一个从 s 读取数据的 Reader
//读取
//func (r *Reader) Read(b []byte) (n int, err error)   // 读取数据到 p
//func (r *Reader) ReadByte() (b byte, err error)  // 读取一个字节

//bytes
//跟 strings.Reader 差不多 bytes.Reader 也实现了读相关的接口，结构体定义如下：
//
//type Reader struct {
//    s        []byte
//    i        int64 // current reading index
//    prevRune int   // index of previous rune; or < 0
//}
//func NewReader(b []byte) *Reader   // 创建一个从 b 读取数据的 Reader

//除此之外，bytes 包中还有一个 bytes.Buffer 结构体实现了读写相关的接口，结构体定义如下：
//
//type Buffer struct {
//    buf      []byte // contents are the bytes buf[off : len(buf)]
//    off      int    // read at &buf[off], write at &buf[len(buf)]
//    lastRead readOp // last read operation, so that Unread* can work correctly.
//}

//func NewBuffer(buf []byte)
//将 buf 作为出初始内容并创建一个 Buffer，
//func NewBufferString(s string) *Buffer
//将 []byte(s) 作为初始内容创建一个 Buffer

//写入和读取
//func (b *Buffer) Write(p []byte) (n int, err error)    // 将 p 写入缓冲区
//func (b *Buffer) WriteString(s string) (n int, err error)   // 将 s 写入缓冲区
//func (b *Buffer) Read(p []byte) (n int, err error) // 从缓冲区读取数据到 p

//操作缓冲区
//func (b *Buffer) Reset()   // 清空缓冲区的内容
//func (b *Buffer) Len() int     // 返回缓冲区中未读取的字节数
//func (b *Buffer) Cap() int      // 返回缓冲区的容量
//func (b *Buffer) Bytes() []byte   // 以字节切片返回所有未读取的数据
//func (b *Buffer) String() string    // 以字符串返回所有未读取的数据

func main() {
	//io打开文件，采用缓冲访问写入
	//再额外加入os.O_APPEND标志可以把指针控制到文末，而不是会覆盖数据
	myfile, err := os.OpenFile("wenzi.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("打开文件出错:", err)
		return
	}
	defer myfile.Close() //关闭文件
	//初始化缓存区与文件关联
	writer := bufio.NewWriter(myfile)
	//写入字符切片到缓存区
	data := []byte("123456789")
	_, err = writer.Write(data)
	if err != nil {
		fmt.Println("写入数据到缓冲区出错:", err)
		return
	}
	//获取缓冲区可用字节数
	canuse_byte := writer.Available()
	fmt.Println("缓存区还可用字节数：%d\n", canuse_byte)
	//获取缓冲区的大小
	canuse_size := writer.Size()
	fmt.Println("底层缓冲区的总大小：%d\n", canuse_size)

	//刷新缓冲区，即将缓冲区的文件写入底层关联文件
	err = writer.Flush()
	if err != nil {
		fmt.Println("刷新缓冲区出错:", err)
		return
	}
	fmt.Println("刷新缓冲区成功")

	//初始化一个缓冲读取器与文件关联
	myfile, _ = os.Open("wenzi.txt")
	defer myfile.Close()
	reader := bufio.NewReader(myfile)

	j := 1
	reader.Discard(j - 1)
	k := 6
	peekdata, err := reader.Peek(k - j + 1)
	if err != nil {
		if err != io.EOF {
			fmt.Println("Peek 出错:", err)
			return
		}
	}
	fmt.Printf("Peek 到的数据: %s\n", peekdata)
	peekdata2, err := reader.Peek(2)
	if err != nil {
		if err != io.EOF {
			fmt.Println("Peek 出错:", err)
			return
		}
	}
	fmt.Printf("Peek 到的数据: %s\n", peekdata2)
	//从缓冲区读取数据
	read_data := make([]byte, 4)
	n, err := reader.Read(read_data)
	if err != nil {
		if err != io.EOF {
			fmt.Println("从缓冲区读取数据出错:", err)
			return
		}
	}
	fmt.Printf("从缓冲区读取 %d 字节的数据: %s\n", n, read_data)

	read2 := make([]byte, 6)
	x, err := reader.Read(read2)
	fmt.Printf("从缓冲区读取 %d 字节的数据: %s\n", x, read2)
	peekdata3, err := reader.Peek(k - j + 1)
	if err != nil {
		if err != io.EOF {
			fmt.Println("Peek 出错:", err)
			return
		}
	}
	fmt.Printf("Peek 到的数据: %s\n", peekdata3)
	peekdata4, err := reader.Peek(1)
	if err != nil {
		if err != io.EOF {
			fmt.Println("Peek 出错:", err)
			return
		}
	}
	fmt.Printf("Peek 到的数据: %s\n", peekdata4)

	read3 := make([]byte, 2)
	y, err := reader.Read(read3)
	fmt.Printf("从缓冲区读取 %d 字节的数据: %s\n", y, read3)
	//在执行 Peek() 方法后，文件读取指针并没有发生偏移，所以后续的读取操作可以继续从原来的位置开始。
	//对于指定的(j,k)字节切片的无损失读取我们可以先跳过j-1个字节再读取j到k字节的数据
	//而在缓冲区结束之后再次采用.Read()会改变指针地从原文本提取数据，这个指针改变是和peek是共享的！！！！！！！！！！！！！！！！！！！！！！！！！
	// 使用 Reset 方法
	newBuffer := bytes.NewBufferString("Reset Buffer")
	reader.Reset(newBuffer)

	resetData, err := reader.Read(read_data)
	//reader.Read() 方法返回两个值：
	//一个整数，表示实际读取的字节数。这个值告诉你在这次读取操作中到底读取了多少数据。
	//一个错误对象，如果在读取过程中出现了错误的话。
	if err != nil {
		fmt.Println("从重置后的缓冲区读取数据出错:", err)
		return
	}
	fmt.Printf("从重置后的缓冲区读取 %d 字节的数据: %s\n", resetData, read_data[:resetData])
}
