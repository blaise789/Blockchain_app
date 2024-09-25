package main

import (
	"fmt"

)
// constants 




func main(){
	bc:=NewBlockchain()
	bc.AddBlock("Sending 1 BTC TO Kevin")
	bc.AddBlock("Sending 2 BTC TO  Kevin")
	
   for _,block:=range bc.blocks{
	fmt.Printf("Data:%s\n",block.Data)
	fmt.Printf("Hash: %x\n",block.Hash)
   }
}