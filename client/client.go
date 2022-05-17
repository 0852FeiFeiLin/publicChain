package client

import (
	"flag"
	"fmt"
	"os"
	"publicChain/block"
	"publicChain/tools"
	"publicChain/transaction"
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
	bc *block.BlockChain
}

func (cl *Cli) Run() {
	//使用区块链对象
	chain, _ := block.NewBlockChain("zhang")
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
	case "send":
		cl.send()
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
	//功能8
	case "getallblock":
		cl.getAllBlock()
	//功能9
	case "getbalance":
		cl.getBalance()
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
	创建区块链：
			main.exe createChain --address  "矿工账户"
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
	_, err := block.NewBlockChain(*address)
	if err != nil {
		fmt.Println("创建区块链失败")
		return
	}
	fmt.Println("区块链创建成功")
}

/*
	添加区块  --->  发起交易
		想添加区块到区块链中，首先需要有交易，那我们就需要先发起交易，产生一笔交易(收钱人 给钱人 给钱的金额)
		main.exe send  --from "zhang" --to  “li” --amount 50
*/
func (cl *Cli) send() {
	/*
		1、创建一笔交易transaction
		2、把这笔交易存储到区块中，并保存到区块链中，
	*/
	sendflag := flag.NewFlagSet("send", flag.ExitOnError)
	from := sendflag.String("from", "", "交易发起者的地址")
	to := sendflag.String("to", "", "交易接受者的地址")
	//无符号，正整数，不能是负数
	amount := sendflag.Uint("amount", 0, "交易的金额")
	//解析
	sendflag.Parse(os.Args[2:])
	//1、创建普通交易
	newTransaction, err := cl.bc.NewTransaction(*from, *to, *amount)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//2、把交易存储到区块中，并保存到区块链中
	/*chain, err := block.NewBlockChain("")
	if err != nil {
		fmt.Println(err.Error())
		return
	}*/
	/*
		因为发起一笔交易，就会产生一笔coinBase交易，也就是记账人的奖励
		注意：我们这边谁产生了这笔交易，谁就是记账人，就得到coinBase奖励
	*/
	base, err := cl.bc.NewCoinBase(*from)

	err = cl.bc.AddBlockToChain([]transaction.Transaction{*newTransaction, *base})
	if err != nil {
		fmt.Println("区块添加失败！")
		return
	}
	fmt.Println("区块添加成功！")
}

//打印区块链
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
	for _, b := range blocks {
		fmt.Printf("PrevHash: %x\n", b.PrevHash)
		fmt.Printf("区块Hash: %x\n", b.NowHash)
		fmt.Printf("DataSize: %d\n", len(b.Txs)) //交易数目
		//遍历切片集合
		for _, tx := range b.Txs {
			//交易hash值
			fmt.Printf("\t交易hash:%x\n", tx.TXid)

			//有几个交易输入
			fmt.Printf("\t\t有%d个交易输入:\n", len(tx.Input))
			for i, input := range tx.Input {
				fmt.Printf("\t\t\t消费%d,来自%x,下标%d\n", i, input.TXid, input.VOut)
			}

			//有几个交易输出
			fmt.Printf("\t\t有%d个交易输出:\n", len(tx.OutPut))
			for i, output := range tx.OutPut {
				//
				fmt.Printf("\t\t\t收入%d,金额%d,属于%s\n", i, output.Value, string(output.ScriptPubKey))
			}

			//fmt.Printf("\t交易输出:%s\n",string(tx.OutPut[0].ScriptPubKey))
		}
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
	fmt.Printf("PrevHash: %x\n", blocks[0].PrevHash)
	fmt.Printf("区块Hash: %x\n", blocks[0].NowHash)
	fmt.Printf("DataSize: %d\n", len(blocks[0].Txs))

	//遍历切片集合
	for _, tx := range blocks[0].Txs {
		//交易hash值
		fmt.Printf("\t交易hash:%x\n", tx.TXid)

		//有几个交易输入
		fmt.Printf("\t\t有%d个交易输入:\n", len(tx.Input))
		for i, input := range tx.Input {
			fmt.Printf("\t\t\t消费%d,来自%x,下标%d\n", i, input.TXid, input.VOut)
		}

		//有几个交易输出
		fmt.Printf("\t\t有%d个交易输出:\n", len(tx.OutPut))
		for i, output := range tx.OutPut {
			//
			fmt.Printf("\t\t\t收入%d,金额%d,属于%s\n", i, output.Value, string(output.ScriptPubKey))
		}
	}

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
	fmt.Printf("PrevHash: %x\n", blocks[len(blocks)-1].PrevHash)
	fmt.Printf("Hash: %x\n", blocks[len(blocks)-1].NowHash)
	fmt.Printf("DataSize: %d\n", len(blocks[len(blocks)-1].Txs))
	for _, tx := range blocks[len(blocks)-1].Txs {
		//交易hash值
		fmt.Printf("\t交易hash:%x\n", tx.TXid)
		//有几个交易输入
		fmt.Printf("\t\t有%d个交易输入:\n", len(tx.Input))
		for i, input := range tx.Input {
			fmt.Printf("\t\t\t消费%d,来自%x,下标%d\n", i, input.TXid, input.VOut)
		}

		//有几个交易输出
		fmt.Printf("\t\t有%d个交易输出:\n", len(tx.OutPut))
		for i, output := range tx.OutPut {
			//
			fmt.Printf("\t\t\t收入%d,金额%d,属于%s\n", i, output.Value, string(output.ScriptPubKey))
		}
	}

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

//查询余额 getbalance --address "zhang"
func (cl *Cli) getBalance() {
	//先判断区块链是否存在，
	exits := tools.FileExits("blockchain.db")
	if !exits {
		fmt.Println("区块链不存在！")
		return
	}
	getbalance := flag.NewFlagSet("getbalance", flag.ExitOnError)
	from := getbalance.String("address", "", "需要查询余额的地址")
	getbalance.Parse(os.Args[2:])
	balance, err := cl.bc.GetUTXO(*from)
	if err != nil {
		return
	}
	fmt.Printf("%s的余额为：%d\n", *from, balance)
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
