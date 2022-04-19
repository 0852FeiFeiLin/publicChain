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
	chain, _ := entity.NewBlockChain("")
	defer chain.DB.Close()
	//并赋值给client客户端，不然使用的时候会报空指针异常
	cl.bc = chain
	//判断数据库操作对象是否存在，写在这里就不需要每次在方法里面写了
	if cl.bc == nil {
		fmt.Println("区块链db不存在！")
		return
	}
	//判断输入的长度
	if len(os.Args) < 2 {
		return
	}
	switch os.Args[1] {
	//功能1
	case "createblockchain":
		cl.createChain()
	//功能2
	case "addblock":
		cl.addBlock()
	//功能3
	case "printchain":
		cl.printChain()
	//功能4
	case "getblockcount":
		cl.getBlockCount()
	//功能5  (没有实现)
	case "getblockinfo":
		cl.getBlockInfo()
	//功能6
	case "getlastblock":
		cl.getLastBlock()
	//功能7
	case "getfirstblock":
		cl.getFirstBlock()
	case "getallblock":
		cl.getAllBlock()
	//功能n
	case "help":
		cl.help()

	default:
		fmt.Println("please check it！not have this function~")
		//退出
		os.Exit(1)
	}
}

/*
	对应上面的所有功能
*/
func (cl *Cli) createChain() {
	createBlockChain := flag.NewFlagSet("createBlockchain", flag.ExitOnError)
	//获取参数
	address := createBlockChain.String("address", "", "创世区块的交易信息")
	//解析
	createBlockChain.Parse(os.Args[2:])
	//调用方法
	//判断区块链.db()是否存在，存在就不创建，不存在NewBlockChain
	exits := tools.FileExits("blockchain.db")
	if exits {
		fmt.Println("区块链已经存在，不能再创建了...")
		return
	}
	_, err := entity.NewBlockChain(*address)
	if err != nil {
		fmt.Println("创建区块链失败")
		return
	}
	fmt.Println("区块链创建成功")
}
func (cl *Cli) addBlock() {
	addBlock := flag.NewFlagSet("addblock", flag.ExitOnError)
	//先判断区块链是否存在，
	exits := tools.FileExits("blockchain.db")
	if !exits {
		fmt.Println("区块链不存在，请创建区块链后添加区块！")
		return
	}
	//获得参数
	data := addBlock.String("data", "", "区块的交易信息")
	//解析
	addBlock.Parse(os.Args[2:])
	//调用方法
	err := cl.bc.AddBlockToChain([]byte(*data))
	//判断返回值
	if err != nil {
		fmt.Println("添加区块失败！")
		return
	}
	fmt.Println("添加区块成功")
}

func (cl *Cli) printChain() {
	//先判断区块链是否存在，
	exits := tools.FileExits("blockchain.db")
	if !exits {
		fmt.Println("区块链不存在！")
		return
	}
	//获取所有区块对象切片
	blocks, err := cl.bc.GetAllBlock()
	if err != nil {
		fmt.Println("获取AllBlocks对象失败")
		return
	}
	//迭代切片输出
	for _, b := range blocks {
		fmt.Printf(":Prev. hash: %x\n", b.PrevHash)
		fmt.Printf("Data: %s\n", b.Data)
		fmt.Printf("Hash: %x\n", b.NowHash)
		fmt.Println()
	}
	fmt.Println("遍历完成！！")
}

func (cl *Cli) getBlockCount() {
	//先判断区块链是否存在，
	exits := tools.FileExits("blockchain.db")
	if !exits {
		fmt.Println("区块链不存在！")
		return
	}
	blocks, err := cl.bc.GetAllBlock()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("区块总数量：", len(blocks))
}

func (cl *Cli) getLastBlock() {
	//先判断区块链是否存在，
	exits := tools.FileExits("blockchain.db")
	if !exits {
		fmt.Println("区块链不存在！")
		return
	}
	blocks, err := cl.bc.GetAllBlock()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//最后一个切片对象也就是最后一个区块
	fmt.Printf("Prev. hash: %x\n", blocks[0].PrevHash)
	fmt.Printf("Data: %s\n", blocks[0].Data)
	fmt.Printf("Hash: %x\n", blocks[0].NowHash)
}
/*
	获取第一个区块信息
*/
func (cl *Cli) getFirstBlock() {
	//先判断区块链是否存在，
	exits := tools.FileExits("blockchain.db")
	if !exits {
		fmt.Println("区块链不存在！")
		return
	}
	blocks, err := cl.bc.GetAllBlock()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//最后一个切片对象也就是最后一个区块
	fmt.Printf("Prev. hash: %x\n", blocks[len(blocks)-1].PrevHash)
	fmt.Printf("Data: %s\n", blocks[len(blocks)-1].Data)
	fmt.Printf("Hash: %x\n", blocks[len(blocks)-1].NowHash)
}
//获取单个区块的信息
func (cl *Cli) getBlockInfo() {
	//先判断区块链是否存在，
	exits := tools.FileExits("blockchain.db")
	if !exits {
		fmt.Println("区块链不存在！")
		return
	}
	/*
		问题：我们的blotDB只能桶key获取到值，能不能根据具体的vlaue获取到整个信息呢？
	*/
	getBlockInfo := flag.NewFlagSet("getblockinfo", flag.ExitOnError)
	hash := getBlockInfo.String("hash", "", "区块hash")
	cl.bc.GetBlockInfo([]byte(*hash))
}

//获取所有区块,(到时候获取所有区块个数，获取单个区块，都很方便)
func (cl *Cli) getAllBlock() {
	//先判断区块链是否存在，
	exits := tools.FileExits("blockchain.db")
	if !exits {
		fmt.Println("区块链不存在！")
		return
	}
	//获取到了所有区块对象
	blocks, err := cl.bc.GetAllBlock()
	if err != nil {
		fmt.Println("所有区块获取失败...")
		return
	}
	fmt.Println("获取到的区块对象个数：", len(blocks))
}

func (cl *Cli) help() {
	fmt.Println("main.exe Command --data ?")
	fmt.Println("Has the following Command:")
	fmt.Println("\t \t createBlockchain --data Transaction information of Genesis block")
	fmt.Println("\t \t addblock --data Transaction information of this block")
	fmt.Println("\t \t getblockinfo --hash The hash of this block")
	fmt.Println("\t \t printchain")
	fmt.Println("\t \t getblockconut")
	fmt.Println("\t \t getfirstblock")
	fmt.Println("\t \t getlastblock")
	fmt.Println()
}

