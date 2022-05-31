package tools

import "golang.org/x/crypto/ripemd160"

/**
 * @author: linfeifei
 * @email: 2778368047@qq.com
 * @phone: 18170618733
 * @DateTime: 2022/5/21 9:40
 * @Description: 对数据进行ripemd160计算
 **/

func GetRipemd160(data []byte)([]byte){
	ripemd := ripemd160.New()
	ripemd.Write(data)
	return ripemd.Sum(nil)
}
