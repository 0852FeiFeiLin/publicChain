package wallet
//注意：这个address里面的所有全部挪移到了wallet里面,所以把address的所有函数小写，在其他地方不可见
//
//
//import (
//	"bytes"
//	"crypto/ecdsa"
//	"crypto/elliptic"
//	"crypto/rand"
//	"errors"
//	"math/big"
//	"publicChain/tools"
//)
///*
///**
//* @author: linfeifei
//* @email: 2778368047@qq.com
//* @phone: 18170618733
//* @DateTime: 2022/5/21 9:10
//* @Description:
//	生成比特币地址和校验:
//		1、生成随机私钥
//		2、根据私钥得到公钥，
//		3、对公钥进行双重hash计算得到pubHash
//			a.第一重是sha256计算
//			b.第二重是ripemd160计算
//		4、拼接版本号0x00到开头，得到ver_pubHash
//		5、对ver_pubHash进行双重sha256计算，取前四个字节作为校验码checkCode
//		6、拼接校验码到末尾，得到ver_pubHash_CheckCode
//		7、对上一步的结果进行base58编码，得到比特币地址，btcAddress
//		8、对btcAddress反解码,得到ver_pubHash_checkCode查看是否有效
//			a.截取后四个字符check1，(假设是正确答案)
//			b.截取除去检验位[:len(address)-4]的数据,也就是ver_pubHash
//			c.再进行两次sha256计算，因为hash是不可逆的，得到hash值，然后再取前四个字符check2(待验证的答案)
//		9、比较check1和check2是否一致，一致说明比特币地址有效，验证通过，反之无效
//**/
//
////比特币地址版本号
////const version = 0x00
//
///*
//	生成比特币地址
//*/
//func createAddress() (string, error) {
//	//得到公私钥
//	_, pubKeyByte, err := CreateKeys()
//	//持久化存储私钥hash   key ---> address  value ---> priHash  显然是使用钱包结构
//	if err != nil {
//		return "", err
//	}
//	//对公钥进行sha256计算 + ripemd160计算
//	hash := tools.GetSha256Hash(pubKeyByte)
//	pubHash := tools.GetRipemd160(hash) //得到公钥Hahs
//	//version + pubHash
//	ver_pubHash := append(pubHash, []byte{version}...)
//	//对上一步的结果进行双重sha256计算，截取前四个字节，得到校验位checkCode
//	hash1 := tools.GetSha256Hash(ver_pubHash)
//	hash2 := tools.GetSha256Hash(hash1)
//	checkCode := hash2[:4]
//	//version + pubHash + checkCode
//	ver_pubHash_checkCode := append(ver_pubHash, checkCode...)
//	//进行base58编码
//	btcAddress := tools.Encode(ver_pubHash_checkCode)
//	//返回地址
//	return btcAddress, nil
//}
//
////生成公钥和私钥   (没进行拼接，也没进行压缩.....)
//func createKeys() (*ecdsa.PrivateKey, []byte, error) {
//	//曲线方程
//	curve := elliptic.P256()
//	//生成私钥
//	privateKey, err := ecdsa.GenerateKey(curve, rand.Reader)
//	if err != nil {
//		return nil, nil, err
//	}
//	//获取公钥
//	publicKey := privateKey.PublicKey
//	pubByte := elliptic.Marshal(curve, publicKey.X, publicKey.Y) //04  x   y (非压缩公钥)
//
//	//压缩公钥
//	//catPubKey := CatDownPubKey(publicKey)
//
//	//返回私钥和非压缩公钥字节
//	return privateKey, pubByte, nil
//}
//
///*
//	检查地址是否有效
//*/
//func addressVerify(btcAddress string) bool {
//	if btcAddress == "" {
//		return false
//	}
//	//base58反解码
//	ver_pubHash_checkCode := tools.Decode(btcAddress)
//	//check1
//	check1 := ver_pubHash_checkCode[len(ver_pubHash_checkCode)-4:]
//	//截取得到ver_pubHash
//	ver_pubHash := ver_pubHash_checkCode[:len(ver_pubHash_checkCode)-4]
//	//双重hash后，取前四字符作为check2
//	hash1 := tools.GetSha256Hash(ver_pubHash)
//	hash2 := tools.GetSha256Hash(hash1)
//	check2 := hash2[:4]
//	//判断check1 和 check2 是否相等
//	return bytes.Compare(check1, check2) == 0
//}
//
///*
//	返回压缩公钥的字节切片
//*/
//func catDownPubKey(publicKey ecdsa.PublicKey) []byte {
//	var pub []byte
//	b := big.NewInt(0)
//	//判断y坐标，
//	if publicKey.Y.Cmp(b) == -1 { //y < 0 ==> 02
//		xByte := publicKey.X.Bytes()
//		pub = append([]byte{02}, xByte...) //压缩公钥
//	} else { //y > 0 ==> 03
//		xByte := publicKey.X.Bytes()
//		pub = append([]byte{03}, xByte...) //压缩公钥
//	}
//	return pub
//}
//
///*
//	获取到address中的公钥pubHash部分，并返回
//*/
//func  getPubHash(addr string)([]byte,error){
//	//1、验证地址是否正确
//	verify := AddressVerify(addr)
//	if !verify {
//		return nil,errors.New("地址不合法")
//	}
//	//2、解码地址
//	ver_pubHash_check := tools.Decode(addr)
//	ver_pubHash := ver_pubHash_check[:len(ver_pubHash_check)-4]
//	//3、得到公钥hash
//	pubHash := ver_pubHash[1:]//0x00 版本号只占据一个字节，所以从下标1开始
//	return pubHash,nil
//}
//
///*
//	获取到普通公钥的hash值
//*/
//func hashPubKey(pub []byte)([]byte){
//	//sha256
//	hash1 := tools.GetSha256Hash(pub)
//	//ripemd160计算
//	pubHash := tools.GetRipemd160(hash1)
//	//得到公钥hash
//	return pubHash
//}