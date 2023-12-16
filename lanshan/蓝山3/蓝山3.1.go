package main

import (
	"fmt"
)

type Shangpin struct {
	Name     string
	jiage    int
	shuliang int
	Id       int
}
type guanli interface {
	make() *Shangpin
	find(string, []Shangpin)
	pay(string, []Shangpin)
}

func (s Shangpin) make() *Shangpin {
	var x Shangpin
	fmt.Println("输入商品名称")
	fmt.Scan(&x.Name)

	fmt.Println("输入商品价格")
	fmt.Scan(&x.jiage)

	fmt.Println("输入商品数量")
	fmt.Scan(&x.shuliang)

	//fmt.Println("输入商品编号")
	//fmt.Scan(&x.Id)
	x.Id = 0
	fmt.Println("你输入的商品信息为", x)
	return &x
}
func (s Shangpin) find(findname string, a []Shangpin) {
	fmt.Println("输入你想查找的商品")
	for _, look := range a {
		if look.Name == findname {
			fmt.Println("匹配的信息:", look)
			return
		}
	}
	fmt.Println("未找到商品信息")
}
func (s Shangpin) pay(findname string, a []Shangpin) {
	fmt.Println("输入你想购买的商品")
	for i, look := range a {
		if look.Name == findname {
			fmt.Println("匹配的信息:", look)
			fmt.Println("输入你想购买的数量")
			var n int
			fmt.Scan(&n)
			a[i].shuliang -= n
			fmt.Println("购买成功，剩余库存：", a[i].shuliang)
			return

		}
	}
	fmt.Println("未找到商品")

}

func main() {
	all := make([]Shangpin, 0)
	var g guanli = &Shangpin{}

	for {
		fmt.Println("欢迎光临此商品平台")
		fmt.Println("请输入命令：1-输入商品 2-查询商品 3-购买商品 end-结束程序")
		var command string
		fmt.Scan(&command)
		switch command {
		case "1":
			x := g.make()
			found := false
			for i, item := range all {
				if item.Name == x.Name {
					all[i].shuliang += x.shuliang
					found = true
					x.Id = len(all) + 1
					break

				}
			}
			if !found {
				x.Id = len(all) + 1
				all = append(all, *x)
			}
		case "2":
			var name string
			fmt.Println("输入要查询的商品名称")
			fmt.Scan(&name)
			g.find(name, all)
		case "3":
			var name string
			fmt.Println("输入要购买的商品名称")
			fmt.Scan(&name)
			g.pay(name, all)
		case "end":
			fmt.Println("程序结束")
			return
		default:
			fmt.Println("无效的命令")
		}
	}
}
