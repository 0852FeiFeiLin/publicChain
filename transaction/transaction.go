package transaction

import (
	"bytes"
	"encoding/gob"
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

//序列化
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
func (txs *Transaction) DeSerialize() ([]byte, error) {
	return nil, nil
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
