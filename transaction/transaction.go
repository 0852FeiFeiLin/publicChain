package transaction

import (
	"bytes"
	"encoding/gob"
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
		return nil,err
	}
	//返回反序列化后的结构体对象
	return &tx,nil
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
func NewTransaction(from ,to string,amount uint)(*Transaction,error){
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
/*	bc, err2 := entity.NewBlockChain("")
	if err2 != nil {
		return nil, err2
	}
	//余额  = 交易输出 - 交易输入  方法 *************还没写
	output := bc.FindAllOutput(from)
	input, err2 := bc.FindAllInput(from)
	if err2 != nil {
		return nil,err2
	}
	//相减方法 *************还没写******************************
*/
	tx := Transaction{//实例化交易对象
		OutPut: nil,
		Input: nil,
	}
	//序列化
	byteTx,err := tx.Serialize()
	if err != nil {
		return nil, err
	}
	hahs := tools.GetSha256Hash(byteTx)
	//3、
	tx.TXid = hahs
	//4、
	return &tx,nil

}