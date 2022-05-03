package transaction

import (
	"bytes"
	"encoding/gob"
	"errors"
	"publicChain/entity"
	"publicChain/tools"
)

/**
 * @author: linfeifei
 * @email: 2778368047@qq.com
 * @phone: 18170618733
 * @DateTime: 2022/4/19 9:34
 * @Description:
		交易 = 交易Id + 交易输入 + 交易输出
		1、需要先有交易输入和交易输出才可以创建交易，
		2、一笔交易里面可以有多个交易输入和多个交易输出，所以是[]Input切片类型
		3、交易结构体TXid和和交易输入结构体里面的TXid虽然还是一样的，但是代表的意思是不一样的，
 **/
type Transaction struct {
	//交易的唯一标识，是一个流水号
	TXid []byte
	//交易输入（可以有多个交易输入）
	Input []Input
	//交易输出（可以有多个交易输出）
	OutPut []OutPut
}

//交易序列化方法
func (txs *Transaction) Serialize() ([]byte, error) {
	var result bytes.Buffer

	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(txs)
	if err != nil {
		return nil, err
	}
	//返回序列化结果
	return result.Bytes(), nil
}

//反序列化
func (txs *Transaction) DeSerialize(txsByte []byte) (*Transaction, error) {
	//把交易信息传入，声明缓冲流对象，
	buffer := bytes.NewBuffer(txsByte)
	decoder := gob.NewDecoder(buffer)
	//要转换的类型
	var tx Transaction
	err := decoder.Decode(&tx)
	if err != nil {
		return nil, err
	}
	//返回反序列化后的结构体对象
	return &tx, nil
}

/*
	coinbase交易
*/
func NewCoinBase(address string) (*Transaction, error) { //address 是矿工的账户
	//实例化交易对象
	cb := Transaction{
		Input: nil,
		OutPut: []OutPut{
			//coinbase只有一个交易输出，所以就写一个大括号
			{
				Value: 50,
				//锁定脚本里面的是一个账户，公钥
				ScriptPubKey: []byte(address),
				//[]byte("zhang")
			},
			//正常的交易，是有两个交易输出，也就是两个大括号{}，{}代表两个交易输出
		},
	}
	txsByte, err := cb.Serialize()
	if err != nil {
		return nil, err
	}
	//把交易对象进行hash计算，然后当作txid
	cb.TXid = txsByte
	//返回的是coinBase交易对象
	return &cb, nil
}

/*
	创建交易,返回交易
		参数:(交易发送者，接受者，金额)
*/
func NewTransaction(from, to string, amount uint) (*Transaction, error) {
	/*
		1、创建Input
			a、在已经有的交易中，去寻找可用的交易输出，
				怎么找？
					思路：
						1、先找到区块链中的所有区块，
						2、然后从区块中找到所有的交易，
						3、然后找到所有的Output，
						4、然后筛选出所有和from有关的Output。（交易输入同上)
				余额 = 所有的收入（交易输出） - 所有的支出（交易输入）

			b、从所有的可用的交易输出中，取出一部分，判断是否足够（够用就行）
			c、构建Input
		2、创建Output
		3、给txid赋值
		4、返回交易对象，
	*/
	//创建区块链对象
	bc, err2 := entity.NewBlockChain("")
	if err2 != nil {
		return nil, err2
	}
	//创建Input的准备工作
	//a、
	//余额  = 交易输出 - 交易输入  方法 *************还没写
	output := bc.FindAllOutput(from) //txid  下标
	input, err2 := bc.FindAllInput(from)
	if err2 != nil {
		return nil, err2
	}

	//相减方法（抹除）
	//寻找余额spendOutputs  = 所有的交易输出  - 所有的交易输入
	/*spendOutputs := bc.FindSpendOutputs(output, input)

	//判断余额是否够用
	if spendOutputs == nil {
		return nil,errors.New("没有可用的余额~")
	}*/
	//我们需要使用结构体来存储：txid  vout  面额  (UTXO结构体)
	//寻找余额未消费的UTXO  （不妥）
	spendOutputs, totalAmount := bc.FindSpendOutputs(output, input, amount) //返回值1：需要用到的所有的钱，返回值2：所有钱的金额（对应关系）
	if spendOutputs == nil {
		return nil, errors.New("没有可用的余额~")
	}
	//b、从所有的可用的交易输出中，取出一部分，判断是否足够（够用就行）
	/*
		//纪录余额
		var totalAmount uint = 0
		var totalNums int
		for index, utxo := range utxos {
			totalAmount += utxo.Value//(value 修改为uint类型)
			if totalAmount >= amount { //如果余额大于要转的钱，说明足够，
				totalNums = index +1
				break //如果够了，那就不搜口袋看钱了
			}
		}*/
	if totalAmount < amount { //如果余额小于要转的钱，说明不够，
		return nil, errors.New("余额不足！！！")
	}

	//c、构建input （因为一笔交易可能会有多个input，[10,10,20,30]）
	allInput := make([]Input, 0)          //这次交易要用到的所有交易输入 [10,10,20,30]
	for _, output := range spendOutputs { //遍历XXX
		input := NewInput(output.Txid, output.Index, []byte(from)) //参数三是from，是因为from要用这笔钱，所以是from的script
		allInput = append(allInput, input)
	}

	//2、创建OutPut  。。。。

	tx := Transaction{ //实例化交易对象
		OutPut: nil,
		Input:  nil,
	}
	//序列化
	byteTx, err := tx.Serialize()
	if err != nil {
		return nil, err
	}
	hahs := tools.GetSha256Hash(byteTx)
	//3、给txid赋值
	tx.TXid = hahs
	//4、返回交易对象
	return &tx, nil

}
