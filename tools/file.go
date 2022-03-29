package tools

import "os"

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
	_, err := os.Lstat(path) //参数1：文件的描述信息  参数2：不存在返回错误，存在为空

	return !os.IsNotExist(err)  //假设存在：false  返回值!false   -->true
}