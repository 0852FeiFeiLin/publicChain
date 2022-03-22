package tools

import "crypto/sha256"

/**
 * @author: linfeifei
 * @email: 2778368047@qq.com
 * @phone: 18170618733
 * @DateTime: 2022/2/22 9:14
 **/
/*
	计算hash值，并返回
*/
func GetSha256Hash(data []byte)[]byte{
	hash  := sha256.New()
	hash.Write(data)
	return hash.Sum(nil)
}