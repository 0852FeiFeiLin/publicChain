package entity

import (
	"bytes"
	"encoding/gob"
	"publicChain/pow"
	"publicChain/tools"
	"publicChain/transaction"
	"strconv"
	"time"
)

/**
 * @author: linfeifei
 * @email: 2778368047@qq.com
 * @phone: 18170618733
 * @DateTime: 2022/2/22 9:06
 **/

/*
	一个区块的结构体
*/
type Block struct {
	TimeStamp int64  //时间戳
	PrevHash  []byte //上一个区块hash值
	NowHash   []byte //当前区块hash
	//Data      []byte //交易信息
	Txs   []transaction.Transaction //交易信息(可能有多个交易)
	Nonce int64                     //随机数
}

//go语言实现接口，只需要让结构体实现结构体里面的所有方法，那么结构体就会变成一个实现了这个接口的实例，
func (block *Block) GetTimeStamp() int64 {
	return block.TimeStamp
}
func (block *Block) GetTxs() []transaction.Transaction {
	return block.Txs
}
func (block *Block) GetPrevHash() []byte {
	return block.PrevHash
}

/*
	创建区块:
		1、传递交易信息 + 上一个区块hash
		2、调用pow工人、找到区块hash值和随机数
		3、返回区块
*/
func NewBlock(data []transaction.Transaction, prevHash []byte) *Block { //交易信息，。上一个区块hash
	//实例化结构体，创建区块
	block := Block{
		TimeStamp: time.Now().Unix(),
		PrevHash:  prevHash,
		Txs:       data,
	}

	//找到pow结构，也就是工人
	//pow := pow.NewPow(block.Data,block.PrevHash, block.TimeStamp )
	//传入一个实现BlockInterface的实例
	pow := pow.NewPow(&block)
	//通过pow工人，找到随机数和hash值，并返回
	hash, nonce := pow.Run()

	//赋值到区块里面
	block.NowHash = hash
	block.Nonce = nonce

	/*//这样我们就不需要调用setHash方法了，因为在run方法里面就调用了GetHash
	//调用结构体方法，计算当前区块hash值
	block.NowHash = block.SetHash()*/

	return &block
}

/*
	序列化：将结构体数据 ---> 字节[]byte切片数据
*/
func (block *Block) Serialize() ([]byte, error) {
	//输入流对象
	var result bytes.Buffer
	//创建序列化对象
	en := gob.NewEncoder(&result)
	//进行序列化
	err := en.Encode(block)
	if err != nil {
		return nil, err
	}
	//返回序列化结果
	return result.Bytes(), nil
}

/*
	反序列化：将[]byte数据 ---> 结构体数据，
*/
func (bl *Block) DeSerialize(data []byte) (*Block, error) {
	//将字节切片转为io.Reader
	reader := bytes.NewReader(data)
	de := gob.NewDecoder(reader)
	//转为结构数据
	var block *Block
	err := de.Decode(&block) //(要转成什么类型)
	if err != nil {
		return nil, err
	}
	return block, nil
}

/*
	计算hash值 :不用了
		组成部分：时间戳 + 上一个区块hash + 交易信息 拼接组成字符串进行hash计算。
*/
func (block *Block) SetHash() ([]byte, error) {
	time := []byte(strconv.FormatInt(block.TimeStamp, 10)) //要转的数字,是什么进制
	//把随机数加入
	nonce := []byte(strconv.FormatInt(block.Nonce, 10))

	//交易切片序列化[]byte类型
	//循环遍历交易集合，然后进行序列化，然后加入交易集合里面,类型是[]byte
	txsBytes := []byte{}
	for _, value := range block.Txs {
		txsByet, err := value.Serialize()
		if err != nil {
			return nil, err
		}
		txsBytes = append(txsBytes, txsByet...)
	}
	//bytes.Join拼接字符串：拼接内容,以什么方式进行拼接
	str := bytes.Join([][]byte{txsBytes, block.PrevHash, time, nonce}, []byte{})
	//拼接后进行获取hash，然后返回
	return tools.GetSha256Hash(str), nil
}
