package main

import (

)

/**
 * @author: linfeifei
 * @email: 2778368047@qq.com
 * @phone: 18170618733
 * @DateTime: 2022/3/1 11:14
 **/
/*
实现功能：
	1、创建区块链，
	2、进行添加区块（NewBlock实现了pow算法），从而提升了难度
	3、查看区块的随机数
*/
/*func main() {
	//创建区块链
	bc := entity.NewBlockChain([]byte("创世区块"))
	fmt.Println(string(bc.Blocks[0].Data))

	//添加区块到区块链（在newBlock里面实现pow算法）
	bc.AddBlockToChain([]byte("222"))
	fmt.Println(string(bc.Blocks[1].Data))

	//现在有两个区块了，然后我们打印随机数看一下
	fmt.Println("创世区块的随机数：",bc.Blocks[0].Nonce)
	fmt.Println("222区块的随机数：",bc.Blocks[1].Nonce)
}*/