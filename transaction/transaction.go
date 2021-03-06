package transaction

import (
	"bytes"
	"encoding/gob"
	"publicChain/tools"
	"publicChain/wallet"
	"time"
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
	//时间戳
	TimeStrap int64
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
	coinbase交易  移动到了blockChain中
*/
func NewCoinBase(address string) (*Transaction, error) { //address 是矿工的账户
	//获取到公钥hash
	pubHash, err2 := wallet.GetPubHash(address)
	if err2 != nil {
		return nil, err2
	}
	//实例化交易对象
	cb := Transaction{
		Input: nil,
		OutPut: []OutPut{
			//coinbase只有一个交易输出，所以就写一个大括号
			{
				Value: 50,
				//锁定脚本里面的是一个账户，公钥  --->  修改成公钥hash
				//ScriptPubKey: []byte(address),   //[]byte("zhang")
				ScriptPubKey: pubHash,
			},
			//正常的交易，是有两个交易输出，也就是两个大括号{}，{}代表两个交易输出
		},
	}
	cb.TimeStrap = time.Now().Unix()
	//先把交易对象序列化，
	txsByte, err := cb.Serialize()
	//计算出hash值，然后当作txid
	hash := tools.GetSha256Hash(txsByte)
	if err != nil {
		return nil, err
	}
	cb.TXid = hash
	//返回的是coinBase交易对象
	return &cb, nil
}

/*
	系统奖励的挖矿奖励费  50 --> 25 --> 12.5 --> 6.25（现今）  四年减半 (假设我们模拟的系统奖励是交易金额的20%)
 */
func GetCoinBase(address string)(*Transaction ,error){
	return nil,nil
}
/*
	创建交易,返回交易
		参数:(交易发送者，接受者，金额，需要花费的交易输出)
*/
func NewTransaction(from, to string, amount uint,spendOutputs []UTXO) (*Transaction, error) {
	/*
		1、创建Input
			a、在已经有的交易中，去寻找可用的交易输出，
				怎么找？ （全部在blockChain的NewTranaction中，找到spedAmount）
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
	//1、创建Input
	//c、构建input （因为一笔交易可能会有多个input，[10,10,20,30]）
	allInput := make([]Input, 0) //这次交易要用到的所有交易输入 [10,10,20,30]
	for _, output := range spendOutputs {
		//遍历form要给的钱

		//公钥hash：需要把对应的私钥存到数据库中，之后才能获取对应的公钥hash，
		input := NewInput(output.Txid, output.Index, nil,nil) //参数三是from，是因为from要用这笔钱，所以是from的script
		allInput = append(allInput, input)                         //这个就是outPut里面的Input
	}

	//2、创建OutPut  。。。。 遍历的是to收到钱
	allOutPut := make([]OutPut, 0)
	var totalNums uint //纪录每一次循环的 钱累加
	for _, out := range spendOutputs {
		totalNums += out.Value //累加 10 + 10 + 20 + 30
		/*
			[10,10,20,30] 70 (余额刚好够的情况) --> 不需要进行找零和构建最后一张面额
		*/
		pubHash, err := wallet.GetPubHash(to)
		if err != nil {
			return nil, err
		}
		if totalNums <= amount {
			//构建output,属于to的
			output := NewOutPut(out.Value,pubHash)
			allOutPut = append(allOutPut, output)
			/*
				[10,10,20,30] 50 (余额有多的情况) --> 需要找零20，和构建最后一张面额，
				最终的交易输出 ---> [10,10,20, 10(50-40的还需要凑的钱), 20(这个是30-10的找零)])
			*/
		} else { //进入这里面就是余额大于交易金额了，也就是70了，所以我们要减去上一次累加的金额，70-30 = 40
			//需求1:最后一张面额 需求2:构建找零
			//最后一张面额，寻找还需要凑多少钱  70-30 = 40
			totalNums -= out.Value
			//这就是还需要给的钱， 50 - 40 = 10 ，还需要构建Ouput
			needAmount := amount - totalNums
			to_pubHash, err := wallet.GetPubHash(to)
			if err != nil {
				return nil, err
			}

			output := NewOutPut(needAmount, to_pubHash)
			//[10,10,20,10]
			allOutPut = append(allOutPut, output)
			//找零
			from_pubHash, err := wallet.GetPubHash(from)
			if err != nil {
				return nil, err
			}
			//锁定到的是公钥hash上面
			backChange := NewOutPut(out.Value-needAmount, from_pubHash)
			allOutPut = append(allOutPut, backChange) //找零也添加到本次交易的交易输出中
		}

	}
	//3、给txid赋值
	tx := Transaction{ //实例化交易对象
		OutPut: allOutPut,
		Input:  allInput,
	}
	tx.TimeStrap = time.Now().Unix()
	//序列化
	byteTx, err := tx.Serialize()
	if err != nil {
		return nil, err
	}
	hahs := tools.GetSha256Hash(byteTx)
	tx.TXid = hahs
	//4、返回交易对象
	return &tx, nil
}
