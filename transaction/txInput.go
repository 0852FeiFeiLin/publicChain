package transaction

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
	ScriptSing []byte
}

