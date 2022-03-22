package main

import "fmt"

/**
 * @author: linfeifei
 * @email: 2778368047@qq.com
 * @phone: 18170618733
 * @DateTime: 2022/3/15 9:43
 **/
/*
	实现区块链的交互：
		1、可选择功能：遍历、添加区块、查看区块信息
		2、命令行接口、提供功能菜单、灵活选择
		3、封装功能（创建区块链、添加区块、遍历区块、查看区块...）
		4、
*/
/*
	思路：

*/
func main() {
	Menu()
}
func Menu() {

	fmt.Println("**********功能菜单*********")
	fmt.Println("1、创建区块链")
	fmt.Println("2、添加区块")
	fmt.Println("3、遍历区块")
	fmt.Println("4、查看区块")
	fmt.Println("5、查看帮助菜单")
	fmt.Println("请输入你的功能...")
	var num int
	fmt.Scanln(&num)
	fmt.Println(num)

	switch num {
	case 1:
		A()
	case 2:
		B()
	case 3:
		C()
	case 4:
		D()
	default:
		fmt.Println("输入有误，请检查")

	}

}
func A() {

}
func B() {

}
func C() {

}
func D() {

}
