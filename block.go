package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)
type Block struct{
	Timestamp int64
	Data []byte
	PreviousBlockHash []byte
	Hash []byte
	Nonce int



}

func(b *Block) SetHash(){
	timestamp:=[]byte(strconv.FormatInt(b.Timestamp,10))
	headers:=bytes.Join([][]byte{b.PreviousBlockHash,b.Data,timestamp},[]byte{})
	hash:=sha256.Sum256(headers)
	b.Hash=hash[:]
}
func NewBlock(data string,prevBlockHash []byte) *Block{
   block:=&Block{time.Now().Unix(),[]byte(data),prevBlockHash,[]byte{},0}
   block.SetHash()
   return block
}