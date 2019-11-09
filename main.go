package main

import "fmt"




//定义结构
type Block struct {
	//前区块哈希
	PreHash []byte
	//当前区块哈希
	Hash  []byte
	//交易数据
	Data []byte
}

//创建区块
func NewBlock(data string ,preBlockHash []byte) *Block {
	block:=Block{
		PreHash:preBlockHash,
		Hash:[]byte{}, //先填空，后面再计算 //todo
		Data:[]byte(data),

	}
	return &block
}

func main() {

	block:=NewBlock("擦浪嘿呦，小饼干",[]byte{})

	fmt.Printf("前区块哈希： %x\n",block.PreHash)
	fmt.Printf("当前区块哈希： %x\n",block.Hash)
	fmt.Printf("交易数据： %s\n",block.Data)


	fmt.Println("擦浪嘿呦，小饼干")
}
