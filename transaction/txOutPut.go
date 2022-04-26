package transaction

import "bytes"

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

/*
	判断某笔交易是不是某人的(被解锁),返回布尔值
*/
func (o *OutPut) IsUnlock(from string) bool { //判断在这笔交易能不能被某人解锁
	/*
		判断from名字(也就是地址)是否和锁定脚本一致
	 */
	if from == "" {
		 return false
	}
	return bytes.Compare(o.ScriptPubKey,[]byte(from)) == 0
	//相等返回值就是0.返回true、如果不是0代表不相等，那么就返回false
}
