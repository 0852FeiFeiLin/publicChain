package block

import (
	"bytes"
	"errors"
	"github.com/boltdb/bolt"
)

/**
 * @author: linfeifei
 * @email: 2778368047@qq.com
 * @phone: 18170618733
 * @DateTime: 2022/3/15 10:36
 **/
/*
	区块链迭代器结构体，但是创建迭代器需要先有区块链，也就是属于区块链，所以(创建迭代器方法)写到blockChain里面
*/
type ChainIterator struct {
	//数据库对象
	DB *bolt.DB
	//标志位，是会变的的值  == 当前区块的位置
	CurrentHash []byte
}

/*
	迭代方法：
		使用迭代器依次获取当前区块的信息,
*/
func (iterator *ChainIterator) Next() (*Block, error) {
	var block *Block
	var err error
	//获取数据
	err = iterator.DB.View(func(tx *bolt.Tx) error {
		//1、通过currentHash获取到最后一个区块
		bucket := tx.Bucket([]byte(BUCKET_BLOCK)) //桶1 == 区块链
		if bucket == nil {
			return errors.New("还没创建桶")
		}
		//通过标志位hash从桶1找到对应区块（标志位在blockchain里面赋值了 = lastHash）
		hashBytes := bucket.Get(iterator.CurrentHash)
		//反序列化返回区块
		block, err = block.DeSerialize(hashBytes)
		//更新标志位  = 当前区块的prevHash  这样下一次出来的就是上一个区块
		iterator.CurrentHash = block.PrevHash
		return nil
	})
	return block, err
}

/*
	判断是否还有下一个区块，因为避免到创世区块的时候又继续往前迭代,导致报错
*/
func (iterator *ChainIterator) HahNext() bool {
	//返回int类型的标志数字
	cph := bytes.Compare(iterator.CurrentHash,[]byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0})
	//如果currentHash == 创世区块prevHash ，也就是空那就停止
	return cph != 0
	/*
		如果int不等于0,说明还有下一个，那么就是返回true,
		如返回值是0，那就返回false
	*/
}
