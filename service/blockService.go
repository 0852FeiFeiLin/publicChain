package service

/**
 * @author: linfeifei
 * @email: 2778368047@qq.com
 * @phone: 18170618733
 * @DateTime: 2022/2/22 9:57
 **/

/*
<<<<<<< HEAD
	创建区块
=======
	对应client里面的所有功能
>>>>>>> 1d33b38 (1.0)
*//*
func NewBlock(data []byte,prevHash []byte) *block.Block { //交易信息，。上一个区块hash
	//实例化结构体，创建区块
	block := block.Block{
		TimeStamp: time.Now().Unix(),
		PrevHash:  prevHash,
		Data:      data,
	}
	//调用结构体方法，计算当前区块hash值
	block.NowHash = block.SetHash()
	return &block
}
*/