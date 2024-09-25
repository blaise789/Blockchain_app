package main
type Blockchain struct{
	blocks []*Block
}
func (bc *Blockchain) AddBlock(data string){
	prevBlock:=bc.blocks[len(bc.blocks)-1]
	// last Block
	newBlock:=NewBlock(data,prevBlock.Hash)
	bc.blocks=append(bc.blocks, newBlock)

}
func NewBlockchain()*Blockchain{
	return &Blockchain{[]*Block{NewGensisBlock()}}
}
func NewGensisBlock()*Block{
	return NewBlock("Genesis Block",[]byte{})
}