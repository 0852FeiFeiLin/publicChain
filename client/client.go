package client

import (
	"flag"
	"fmt"
	"os"
	"publicChain/entity"
	"publicChain/tools"
)

/**
 * @author: linfeifei
 * @email: 2778368047@qq.com
 * @phone: 18170618733
 * @DateTime: 2022/3/29 10:52
 **/
/*
	用户的交互入口：
		只用于负责读取用户传递的命令和参数
		并进行解析
		传递解析参数，调用对应的功能
*/
type Cli struct {
	//区块链
	bc *entity.BlockChain
}

func (cl *Cli) Run() {
	//使用区块链对象
	chain, _ := entity.NewBlockChain([]byte("创世区块"))
	defer chain.DB.Close()
	client := Cli{
		bc: chain,
	}
	//获取用户输入
	args := os.Args
	//确定功能对象
	createBlockChain := flag.NewFlagSet("createBlockchain", flag.ExitOnError)
	addBlock := flag.NewFlagSet("addblock", flag.ExitOnError)
	//printChain := flag.NewFlagSet("printchain", flag.ExitOnError)
	//getBlockCount := flag.NewFlagSet("getblockcount",flag.ExitOnError)
	getBlockInfo :=  flag.NewFlagSet("getblockinfo",flag.ExitOnError)
	//help :=  flag.NewFlagSet("help",flag.ExitOnError)

	//判断输入的具体功能
	switch args[1] {
	//功能1
	case "createBlockchain":
		//获取参数
		data := createBlockChain.String("data", "", "创世区块的交易信息")
		//解析
		createBlockChain.Parse(args[2:])
		//调用方法
		//判断区块链.db()是否存在，存在就不创建，不存在NewBlockChain
		exits := tools.FileExits("blockchain.db")
		if exits {
			fmt.Println("区块链已经存在，不能再创建了...")
			return
		}
		_, err := entity.NewBlockChain([]byte(*data))
		if err != nil {
			fmt.Println("创建区块链失败")
			return
		}
		fmt.Println("区块链创建成功")
	//功能2
	case "addblock":
		//获得参数
		data := addBlock.String("data","","区块的交易信息")
		//解析
		addBlock.Parse(args[2:])
		//调用方法
		err := client.bc.AddBlockToChain([]byte(*data))
		//判断返回值
		if err != nil {
			fmt.Println("添加区块失败！")
			return
		}
		fmt.Println("添加区块成功")
	case "printchain":
		//获得迭代对象
		iterator := client.bc.Iterator()
		for  {
			if iterator.HahNext(){
				//还有区块，获取block,打印
				block, err := iterator.Next()
				if err != nil {
					break
				}
				fmt.Printf("Prev. hash: %x\n", block.PrevHash)
				fmt.Printf("Data: %s\n", block.Data)
				fmt.Printf("Hash: %x\n", block.NowHash)
				fmt.Println()
			}else {
				fmt.Println("遍历完成！！")
				break
			}
		}

	case "getblockcount" :
		count, err := client.bc.GetBlockCount()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("区块总数量：",count)

	case "getblockinfo":
		hash := getBlockInfo.String("hash","","区块hash")
		client.bc.GetBlockInfo([]byte(*hash))
		/*
			问题：我们的blotDB只能桶key获取到值，能不能根据具体的vlaue获取到整个信息呢？
		*/

	case "getlastblock":
		block,err := client.bc.GetLastBlock()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Printf("Prev. hash: %x\n", block.PrevHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.NowHash)

	case "help":
		fmt.Println("main.exe Command --data ?")
		fmt.Println("Has the following Command:")
		fmt.Println("\t \t createBlockchain --data Transaction information of Genesis block")
		fmt.Println("\t \t addblock --data Transaction information of this block")
		fmt.Println("\t \t getblockinfo --hash The hash of this block")
		fmt.Println("\t \t printchain")
		fmt.Println("\t \t getblockconut")
		fmt.Println("\t \t getlastblock")
		fmt.Println()
	default:
		fmt.Println("please check it！not have this function~")
		//退出
		os.Exit(1)
	}
}
