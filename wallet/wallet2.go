package wallet

import "crypto/ecdsa"

/**
 * @author: linfeifei
 * @email: 2778368047@qq.com
 * @phone: 18170618733
 * @DateTime: 2022/5/31 9:12
 * @Description:
 **/


/*
	钱包：一个钱包只有一个密钥对而已
		私钥
		公钥
*/
type Wallet2 struct {
	PrivateKey ecdsa.PrivateKey
	PublicKey ecdsa.PublicKey

}
/*
	钱包组合：
		我们需要 `Wallets` 类型来保存多个钱包的组合，将它们保存到文件中
*/
type Wallets struct {
	Wallets map[string]*Wallet
}

