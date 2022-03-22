package test

import "crypto/sha256"

/**
 * @author: linfeifei
 * @email: 2778368047@qq.com
 * @phone: 18170618733
 * @DateTime: 2022/3/1 8:47
 **/
func GetData(data []byte) []byte{
	hash := sha256.New()
	hash.Write(data)
	return hash.Sum(nil)
}