package main

import (
	"fmt"
	"publicChain/entity"
)

/**
 * @author: linfeifei
 * @email: 2778368047@qq.com
 * @phone: 18170618733
 * @DateTime: 2022/3/8 11:15
 **/
/*
	实现持久化存储：
			利用boltdb实现存储区块信息，
			桶1存储区块信息
			桶2存储上一个区块的hash

*/
/*
	迭代区块链：
			查看区块链信息
*/
func main() {
	chain, err := entity.NewBlockChain([]byte("创世区块"))
	defer chain.DB.Close()
	if err != nil {
		fmt.Println("失败", err.Error())
		return
	}/*
	fmt.Println("创世区块", chain.LastHash)
	err = chain.AddBlockToChain([]byte("1111"))
	err = chain.AddBlockToChain([]byte("2222"))
	err = chain.AddBlockToChain([]byte("3333"))
	err = chain.AddBlockToChain([]byte("4444"))
	if err != nil {
		fmt.Println(err.Error(), "添加失败")
		return
	}
	fmt.Println("添加成功")*/

	//迭代
	fmt.Println("***********迭代**********")
	iterator := chain.Iterator()
	for  {
		if iterator.HahNext(){
			//还有区块，获取block,打印
			next, err := iterator.Next()
			if err != nil {
				break
			}
			fmt.Println(string(next.Data))
		}else {
			fmt.Println("遍历完成！！")
			break
		}
	}



}
