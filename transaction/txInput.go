package transaction

import (
	"bytes"
	"publicChain/wallet"
)

/**
 * @author: linfeifei
 * @email: 2778368047@qq.com
 * @phone: 18170618733
 * @DateTime: 2022/4/19 9:28
 * @Description:
		交易输入 = tXid + 交易输出的索引 + 解锁脚本(使用者提供)
 **/
type Input struct {
	//交易hash ，表示存储的是之前交易的 ID，代表钱是从哪里来的
	TXid []byte
	//交易输出的索引，标志是引用的这笔交易里面的第几个交易输出
	VOut int
	//解锁脚本（后期使用者提供解锁脚本和公钥去验证这笔钱能不能被使用）
	//ScriptSing []byte   签名  + 公钥
	Sig    []byte
	PubKey []byte //需要持久化存储私钥，因为私钥可以直接.PubLicKey得到公钥
}

/*
	判断能不能锁定  有问题：判断嫩不能解锁不就是代表判断能不能使用这笔钱，但是解锁脚本的标准不是要判断 "使用者提供解锁脚本和公钥"
	需要获取到两个公钥Hash进行比较：
			from的公钥Hash
			input.PubKey的公钥Hash
*/
func (in *Input) IsLocked(from string) bool {
	if from == "" {
		return false
	}
	//判断这个人的公钥hash是够一致

	//先把in.PubKey公钥获得这个消费的拥有者的pubHash
	lockPubHash := wallet.HashPubKey(in.PubKey)

	//获取到from的公钥hash
	pubHash, err := wallet.GetPubHash(from)
	if err != nil {
		return false
	}
	return bytes.Compare(lockPubHash, pubHash) == 0
	//lockPubHash 是当前的消费属于这个人，假设属于张三
	//pubHah   要验证的人的公钥hash，验证是不是是不是张三
}

/*
	构建Input
*/
func NewInput(txid []byte, vout int, pubKey []byte, sig []byte) Input {
	return Input{
		TXid:   txid,
		VOut:   vout,
		PubKey: pubKey,
		Sig:    sig,
	}
}
