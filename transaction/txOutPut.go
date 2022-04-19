package transaction

/**
 * @author: linfeifei
 * @email: 2778368047@qq.com
 * @phone: 18170618733
 * @DateTime: 2022/4/19 9:25
 * @Description:
		交易输出 = 比特币金额 + 锁定脚本(谁能解锁)
 **/
type OutPut struct {
	//比特币的金额
	Value int64
	//锁定脚本（每一笔交易会产生一个锁定脚本，限定于把这笔交易谁能使用）
	ScriptPubKey []byte
}