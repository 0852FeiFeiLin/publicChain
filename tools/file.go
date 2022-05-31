package tools

import (
	"bytes"
	"encoding/gob"
	"os"
)

/**
 * @author: linfeifei
 * @email: 2778368047@qq.com
 * @phone: 18170618733
 * @DateTime: 2022/3/29 11:23
 **/

/*
	判断文件是否存在:
			存在  --> 返回true
			不存在 --->  返回false
*/
func FileExits(path string) bool {
	/*
		返回值1：文件描述
		返回值2：文件存在err为空，不存在err为错误
	*/
	_, err := os.Lstat(path)
	/*
		os.IsNotExist
			错误存在 --->返回true，代表文件不存在
			错误不存在 --->  返回false，代表文件存在
	*/
	return !os.IsNotExist(err) //假设存错误：false  返回值!false   -->true

}

/*
	序列化数据
 */
func Serialize(data interface{}) ([]byte,error) {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(data)
	if err != nil {
		return nil,err
	}
	return result.Bytes(),nil
}

//反序列化
 func DeSerialize([]byte)(interface{}){
 	return nil
 }
