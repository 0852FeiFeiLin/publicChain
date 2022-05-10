package test

/**
 * @author: linfeifei
 * @email: 2778368047@qq.com
 * @phone: 18170618733
 * @DateTime: 2022/2/22 9:09
 **/
/*
实现功能：
	1、创建区块
	2、创建区块链
	3、添加区块到区块链里面
	4、遍历查看区块链结构
*//*
func main() {
	//创建区块
	block1 := block.NewBlock([]byte("第一个区块"), nil) //交易信息，上一个区块
	fmt.Println(string(block1.Data))

	//创建区块链（必须要有第一个创世区块）
	bc := block.NewBlockChain([]byte("创世区块"))  //交易信息
	fmt.Println(string(bc.Blocks[0].Data))  //通过索引获取具体的区块

	//持续添加区块
	bc.AddBlockToChain([]byte("hello")) //这就是第二个区块添加到区块链
	bc.AddBlockToChain([]byte("feifei")) //第三个区块添加到区块链

	fmt.Println("——————————————————————遍历区块链---------------------")
	//遍历区块链查看结构
	for _,block := range bc.Blocks {
		fmt.Printf("prevHash： %x\n",block.PrevHash)
		fmt.Printf("data： %s\n",block.Data)
		fmt.Printf("hash： %x\n",block.NowHash)
	}

}*/
