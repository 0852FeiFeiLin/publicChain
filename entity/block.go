package entity

import (
	"bytes"
	"encoding/gob"
	"publicChain/pow"
	"publicChain/tools"
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
	Data      []byte //交易信息
	Nonce int64  //随机数
}

/*
	创建区块:
		1、传递交易信息 + 上一个区块hash
		2、调用pow工人、找到区块hash值和随机数
		3、返回区块
*/
func NewBlock(data []byte, prevHash []byte) *Block { //交易信息，。上一个区块hash
	//实例化结构体，创建区块
	block := Block{
		TimeStamp: time.Now().Unix(),
		PrevHash:  prevHash,
		Data:      data,
	}

	//找到pow结构，也就是工人
	pow := pow.NewPow(block.PrevHash,block.TimeStamp,block.Data)
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
func (block *Block) Serialize()([]byte,error){
	var result bytes.Buffer
	//创建序列化对象
	en := gob.NewEncoder(&result)
	//进行序列化
	err := en.Encode(block)
	if err != nil {
		return nil,err
	}
	//返回序列化结果
	return result.Bytes(),nil
}

/*
	反序列化：将[]byte数据 ---> 结构体数据，
*/
func DeSerialize(data []byte)(*Block,error){
	//将字节切片转为io.Reader
	reader := bytes.NewReader(data)
	de := gob.NewDecoder(reader)
	//转为结构数据
	var block *Block
	err := de.Decode(&block) //(要转成什么类型)
	if err != nil {
		return nil,err
	}
	return block,nil
}



/*
	计算hash值 :不用了
		组成部分：时间戳 + 上一个区块hash + 交易信息 拼接组成字符串进行hash计算。
*/
func (block *Block) SetHash() []byte {
	time := []byte(strconv.FormatInt(block.TimeStamp, 10)) //要转的数字,是什么进制
	//把随机数加入
	nonce := []byte(strconv.FormatInt(block.Nonce, 10))

	//bytes.Join拼接字符串：拼接内容,以什么方式进行拼接
	str := bytes.Join([][]byte{time, block.PrevHash, block.Data,nonce}, []byte{})
	//拼接后进行获取hash，然后返回
	return tools.GetSha256Hash(str)
}
