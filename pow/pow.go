package pow

import (
	"bytes"
	"math/big"
	"publicChain/tools"
	"publicChain/transaction"
	"strconv"
)

/**
 * @author: linfeifei
 * @email: 2778368047@qq.com
 * @phone: 18170618733
 * @DateTime: 2022/3/1 9:35
 **/

/*
	难度系数：也就是说我的那个二进制前面有20个0 (前面的0越多，难度就系数越高)
*/
const BITS = 10

/*
	工作量证明：找到随机数
*/
type ProofOfWork struct {
	//给哪一个区块工作
	/*Block *entity.Block*/

	//系统给定的hash目标值
	Target *big.Int
	/*
	PrevHash []byte
	Data []byte
	TimeStamp int64*/
	//使用接口解决包循环引用问题
	Block     BlockInterface
}

/*
	接口
*/
type BlockInterface interface {
	GetTimeStamp() int64
	GetPrevHash() []byte
	GetTxs() []transaction.Transaction
}

/*
	创建一个Pow结构体，并返回，并给target赋值
	1、实例化一个结构体
	2、算出target，并赋值
	3、返回pow结构体
*/

/*func NewPow(data []byte,prh []byte,timeStamp int64) *ProofOfWork {   //把字段传过来，上一个区块的信息 + 交易信息 + 时间戳
	//声明一个大整数类型的变量，
	target := big.NewInt(1) //值是1，因为是转为二进制，也就是0 1的组合
	//找到二进制的hash值
	//左移    目标hash：  256 - 20 -1   结果： 20个0 + 1 + 256-21个0
	target = target.Lsh(target, 256-BITS-1) //移动的数字，移动的位数
	pow := ProofOfWork{
		PrevHash: prh,
		TimeStamp: timeStamp,
		Data: data,
		Target: target,  //系统目标hash
	}
	return &pow
}*/
/*
	使用接口
*/
func NewPow(block BlockInterface) *ProofOfWork { //把字段传过来，上一个区块的信息 + 交易信息 + 时间戳
	//声明一个大整数类型的变量，
	target := big.NewInt(1) //值是1，因为是转为二进制，也就是0 1的组合
	//找到二进制的hash值
	//左移    目标hash：  256 - 20 -1   结果： 20个0 + 1 + 256-21个0
	target = target.Lsh(target, 256-BITS-1) //移动的数字，移动的位数
	pow := ProofOfWork{
		Block:  block,
		Target: target, //系统目标hash
	}
	return &pow
}

/*
	通过工作量证明pow结构体，找到随机数，并把区块hash值 和 随机数 返回
*/
func (pow *ProofOfWork) Run() ([]byte, int64) {
	//随机数
	var nonce int64
	nonce = 0 //从0开始找

	block := pow.Block
	//时间戳转为[]byte类型
	time := []byte(strconv.FormatInt(block.GetTimeStamp(), 10))
	//循环比对
	/*
		循环中做的事情：
			1.  准备数据
			2.  用 SHA-256 对数据进行哈希
			3.  将哈希转换成一个大整数
		 	4.将这个大整数与目标进行比较
	*/
	for {
		//fmt.Println("随机数是：",nonce)
		//把随机数转为[]byte类型
		nonceByte := []byte(strconv.FormatInt(nonce, 10))

		//交易切片序列化，[]byte类型
		//循环遍历交易集合，然后进行序列化，然后加入交易集合里面,类型是[]byte
		txsBytes := []byte{}
		for _, value := range block.GetTxs() {
			txsByet, _ := value.Serialize()
			txsBytes = append(txsBytes, txsByet...)
		}
		//拼接：时间戳 + 上一个区块Hash + 交易信息 + 随机数   （交易信息改变成交易集合）
		byteHash := bytes.Join([][]byte{time, block.GetPrevHash(), txsBytes, nonceByte}, []byte{})

		//获取区块新的hash值
		hash := tools.GetSha256Hash(byteHash)

		//把[]byte类型的hash值 转为 大整数类型bigInt
		num := big.NewInt(0)
		//传入字节切片，返回大整数
		num = num.SetBytes(hash)

		//进行比较
		/*
			规则：
				if(a < target){
				//a：区块的hash值   target：系统给定的hahs值
				}
		*/
		if num.Cmp(pow.Target) == -1 { //如果 a<target，就是找到了
			/*
					-1 if x <  y
				     0 if x == y
				    +1 if x >  y
			*/

			//找到了，返回Hash值和随机数
			return hash, nonce
		}
		nonce++ //算错一次，随机数++
	}
	return nil, 0
}
