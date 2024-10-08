package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
)
 var (maxNonce=math.MaxInt64)
const targetBits=24

type ProofOfWork struct{
	block *Block
	target *big.Int
}
func (pow *ProofOfWork) prepareData(nonce int) []byte{
	data:=bytes.Join(
		[][]byte{
			pow.block.Data,
			pow.block.PreviousBlockHash,
			IntToHex(pow.block.Timestamp),
			IntToHex(int64(targetBits)),
			IntToHex(int64(nonce)),
		},[]byte{})
	return data

}
// Run Performs proof of work 


func (pow *ProofOfWork) Run()(int,[]byte){
	var hashInt big.Int
	var hash [32]byte
	nonce:=0
	fmt.Printf("Mining the block containing \"%s\"\n",pow.block.Data)
	for nonce <maxNonce{
		data:=pow.prepareData(nonce)
		hash=sha256.Sum256(data)
		fmt.Printf("r%x",hash)
		hashInt.SetBytes(hash[:])
		if hashInt.Cmp(pow.target)==-1{
			break
		}else{
			nonce++
		}
		
	}
	fmt.Print("\n\n")
	return  nonce,hash[:]
}

func NewProofOfWork(b *Block) *ProofOfWork{
	target:=big.NewInt(1)
	target.Lsh(target,uint(256-targetBits))
    pow:=&ProofOfWork{b,target}
	return pow
}