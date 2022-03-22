package test

import (
	"bytes"
	"strconv"
	"time"
)

/**
 * @author: linfeifei
 * @email: 2778368047@qq.com
 * @phone: 18170618733
 * @DateTime: 2022/3/1 8:40
 **/
type BLock struct {
	TimeStamp int64
	PrevHash []byte
	Data []byte
	Hash []byte
}
/*
	创建区块，返回区块
*/
func CreateBlock(prevHash []byte,data []byte)*BLock{
	block := BLock{
		TimeStamp: time.Now().Unix(),
		PrevHash: prevHash,
		Data: data,
	}
	block.Hash = block.SetHash();
	return &block
}
/*
	hash值的计算方法：data + 时间戳 +
*/
func (block *BLock)SetHash()[]byte{
	time := []byte(strconv.FormatInt(block.TimeStamp,10))
	hashData := bytes.Join([][]byte{block.Data,time,block.PrevHash},[]byte{})
	return GetData(hashData)
}