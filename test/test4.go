package test

import (
	"publicChain/client"
)

/**
 * @author: linfeifei
 * @email: 2778368047@qq.com
 * @phone: 18170618733
 * @DateTime: 2022/3/29 11:13
 **/
/*
	实现命令行接口
*/
func main() {
	//客户端交互接口
	cli := client.Cli{}
	//执行  --> 命令行输入命令
	cli.Run()

}