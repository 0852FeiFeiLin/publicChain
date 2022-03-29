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
思路:
	1.让程序运行起来
	2.使用终端窗凹来接受用户的输入
	3.读取用户的输入
	4.解析用户的输入
		(1)如果输入的是已有的功能，根据不同的输入来决定使用不同的功能
		(2）如果输入的是没有的功能,给用户进行提示
	5.功能调用结束后，要继续监听用户在终端的键盘输入
	注意：如果还有功能，那么就需要继续进行重复上面步骤，也就是功能5
*/
func main() {
	Menu()
}
func Menu() {
	for  {
		fmt.Println("**********功能菜单*********")
		fmt.Println("1、创建区块链")
		fmt.Println("2、添加区块")
		fmt.Println("3、遍历区块")
		fmt.Println("4、查看区块")
		fmt.Println("5、查看帮助菜单")
		fmt.Println("0、退出")
		fmt.Println("请输入你的功能...")
		var num int
		fmt.Scanln(&num)
		fmt.Println(num)
		if num == 0 {
			fmt.Println("退出系统")
			break
		}
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
			fmt.Println("输入有误，请检查!!!")
		}
	}


}
func A() {
	fmt.Println("功能A")
}
func B() {
	fmt.Println("功能B")

}
func C() {
	fmt.Println("功能C")

}
func D() {
	fmt.Println("功能D")

}
