package transaction

/**
 * @author: linfeifei
 * @email: 2778368047@qq.com
 * @phone: 18170618733
 * @DateTime: 2022/5/3 9:42
 * @Description:  未消费的交易输出UTXO  可消费的余额
 **/

/*
	该交易输出在哪个交易中：txid
	该交友输出在交易中的下标vout:  index
	该交易输出属于谁 （Output:scriptPub）
	该交易输出的面额 （OutPut:value）
*/
type UTXO struct {
	Txid    []byte
	Index   int
	OutPut //匿名字段：作用就会让UTXo结构体默认包含OutPut中的两个字段（scriptPub属于谁、value面额）
	//不要使用这个引用传递值，不然就会是引用原来的值 所以我们把*号去掉就行了
}

//实例化UTXO结构体方法,并返回Utxo
func NewUTXO(txid []byte, index int, output OutPut) UTXO {
	return UTXO{txid, index, output}
}
