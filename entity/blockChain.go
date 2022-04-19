package entity

import (
	"errors"
	"github.com/boltdb/bolt"
	"publicChain/transaction"
)

/**
 * @author: linfeifei
 * @email: 2778368047@qq.com
 * @phone: 18170618733
 * @DateTime: 2022/2/22 10:45
 **/
const BLOCKCHAIN_DB_PATH = "./blockchain.db" //存数据库的区块链文件
const BUCKET_BLOCK = "chain_blocks"          //存区块的桶名字
const BUCKET_STATUS = "chain_status"         //获取最后一个hash值的桶名字
const LAST_HASH = "last_hash"                //第二个桶的key值，存的是最后一个区块hash
/*
	用于创建区块链
*/
type BlockChain struct {
	/*	Blocks []*Block //多个区块组成区块链，区块类型的切片*/
	DB       *bolt.DB //将区块存入bolt数据库里面,数据库连接对象
	LastHash []byte   //最后一个hash值
}

/*
	创建区块链 ---> 改成blot区块链数据库
		1、打开数据库
		2、Update存入数据
		3、先直接使用桶，判断桶是否存在，不存在创建，(避免出现桶存在报错问题)
		4、桶不存在: 1.创建桶。 2.获取到创世区块，存入桶中。  3.创建桶2,存入最后一个区块的hash。
		5、桶存在: 直接使用那个桶2，获取到最后一个区块的hash
		6、给区块链赋值: db对象 + 最后一个区块hash
*/
func NewBlockChain(address string) (*BlockChain, error) {
	var lastHash []byte //用于接收lastHash
	//打开数据库
	db, err := bolt.Open(BLOCKCHAIN_DB_PATH, 0600, nil)
	if err != nil {
		return nil, err
	}
	//存入数据到区块链数据库里面
	err = db.Update(func(tx *bolt.Tx) error {
		//先直接使用桶，如果没有桶再创建
		bucket := tx.Bucket([]byte(BUCKET_BLOCK))
		if bucket == nil { //如果桶为空，说明还没有区块链，就要创建区块链  桶1 = 区块链
			//获取到创世区块,(1.调用方法。2.传入coinbase交易)
			coinbase, _ := transaction.NewCoinBase(address)//放入交易包里面。功能单一
			genesic := NewGenesisBlock(*coinbase)
			//创建第一个桶1，存储区块
			bk, err := tx.CreateBucket([]byte(BUCKET_BLOCK))
			if err != nil {
				return err
			}
			//把区块先转为[]byte
			byteGenesic, err := genesic.Serialize()
			if err != nil {
				return err
			}
			//把区块添加进去  key:区块hash  value:区块
			bk.Put(genesic.NowHash, byteGenesic)

			//第二个桶2，存储最后一个区块的hash值
			bk2, err := tx.CreateBucket([]byte(BUCKET_STATUS))
			if err != nil {
				return err
			}
			//存入lastHash
			bk2.Put([]byte(LAST_HASH), genesic.NowHash)
			//更新区块链的lastHash
			lastHash = genesic.NowHash
		} else { //如果有桶，有区块链了，那么我们从他那个2获取到最后一个区块hash，赋值给变量
			bk2 := tx.Bucket([]byte(BUCKET_STATUS))
			//获取最后一个hash (取值)
			lastHash = bk2.Get([]byte(LAST_HASH))
		}
		return nil
	})
	//以上都是准备工作，这里是给区块链结构赋值，也就是创建区块链
	bc := BlockChain{
		DB:       db,
		LastHash: lastHash,
	}
	//返回区块链和错误信息
	return &bc, err
}

/*
	创建创世区块
*/
func NewGenesisBlock(tx transaction.Transaction) *Block {
	//创世区块 (交易信息data,上一个区块hash:32个0特殊化,)
	return NewBlock([]transaction.Transaction{tx}, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
	/*
		解释：你传一个切片给我，但是切片里面只有一个txs，一个交易信息，
	*/
}

/*
	添加区块到区块链数据库中：
		1、创建区块，上一个区块的hash从区块链的LastHash属性中得到，因为我们之前都给区块链重新赋值了
		2、直接使用桶，然后判断存区块的桶1是否存在
		3、不存在返回错误，存在直接使用，将行区块序列化成byte数据，然后添加到桶中
		4、更新桶2，也就是把最后一个hash值变为当前新区块的hash值
		5、给区块链也重新赋值
*/
func (bc *BlockChain) AddBlockToChain(txs []transaction.Transaction) error {
	//1、创建区块
	newBlock := NewBlock(txs, bc.LastHash) //上一个区块的hash值，直接从区块链获取
	//2、添加至区块链数据库
	err := bc.DB.Update(func(tx *bolt.Tx) error {
		//有桶了，直接用
		bucket := tx.Bucket([]byte(BUCKET_BLOCK))
		if bucket == nil {
			return errors.New("还没有创建桶呢") //自定义错误信息返回
		}
		//将新区块序列化
		serialize, err := newBlock.Serialize()
		if err != nil {
			return err
		}
		//添加：区块hash,区块
		bucket.Put(newBlock.NowHash, serialize)

		//将桶2更新，将本区块的hash值存入桶2，成为lastHash
		bk2 := tx.Bucket([]byte(BUCKET_STATUS))
		if bk2 == nil {
			return errors.New("还没有创建桶2呢")
		}
		//赋值
		bk2.Put([]byte(LAST_HASH), newBlock.NowHash)

		//给区块链的LastHash赋值
		bc.LastHash = newBlock.NowHash
		return nil
	})
	return err
}

/*
	创建迭代器，返回迭代器（属于区块链）
*/
func (bc *BlockChain) Iterator() *ChainIterator {
	//创建迭代器
	chainIterator := ChainIterator{
		DB:          bc.DB,
		CurrentHash: bc.LastHash,
	}
	//使用数据库
	return &chainIterator
}

/*
	获取到最后一个区块信息
*/
func (bc *BlockChain) GetBlockInfo(hash []byte) ([]byte, error) {

	return nil, nil
}

/*
	查看区块个数  没有用到
*/
func (bc *BlockChain) GetBlockCount() (count int, err error) {
	//遍历区块链，然后conut++。计算总数
	iterator := bc.Iterator()
	for {
		if iterator.HahNext() {
			_, err := iterator.Next()
			if err != nil {
				break
			}
			count++
		} else {
			break
		}
	}

	//返回总数
	return count, err
}

/*
	获取到最后一个区块信息  没有用到
*/func (bc *BlockChain) GetLastBlock() (block *Block, err error) {
	bc.DB.View(func(tx *bolt.Tx) error {
		//有桶1，直接用
		bucket := tx.Bucket([]byte(BUCKET_BLOCK))
		if bucket == nil {
			return errors.New("还没有创建桶呢") //自定义错误信息返回
		}
		//桶2
		bk2 := tx.Bucket([]byte(BUCKET_STATUS))
		last := bk2.Get([]byte(LAST_HASH))
		lastBlock := bucket.Get(last)
		//反序列化
		block, err = block.DeSerialize(lastBlock)
		if err != nil {
			return err
		}
		return nil
	})
	return block, err
}

/*
	获取所有区块对象，返回切片对象
*/
func (bc *BlockChain) GetAllBlock() (blocks []*Block, err error) {
	iterator := bc.Iterator()
	for {
		if iterator.HahNext() {
			//还有区块，获取block,赋值给Block切片
			block, err := iterator.Next()
			if err != nil {
				return nil, err
			}
			//追加
			blocks = append(blocks, block)
		} else { //迭代完了，没有对象了
			break
		}
	}
	return blocks, nil
}
