package main

import (
  "fmt"
  "crypto/sha256"
  "time"
  "bytes"
  "strconv"
)

type Block struct {
  Timestamp int64
  Data []byte
  PrevBlockHash []byte
  Hash []byte
}

func (b *Block) SetHash() {
  timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
  headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
  hash := sha256.Sum256(headers)

  b.Hash = hash[:]
  
}

func NewBlock(data string, prevBlockHash []byte) *Block {
  block := &Block{
    Timestamp: time.Now().Unix(),
    Data: []byte(data),
    PrevBlockHash: prevBlockHash,
    Hash: []byte{},
  }

  block.SetHash()
  return block
}

type Blockchain struct {
  blocks []*Block
}

func (bc *Blockchain) AddBlock(data string) {
  prevBlock := bc.blocks[len(bc.blocks)-1]
  newBlock := NewBlock(data, prevBlock.Hash)
  bc.blocks = append(bc.blocks, newBlock)

}

func NewGenesisBlock() *Block {
  return NewBlock("Genesis Block", []byte{})
  
}

func NewBlockchain() *Blockchain {
  return &Blockchain{[]*Block{NewGenesisBlock()}}

}

func main() {
  bc := NewBlockchain()

  bc.AddBlock("Send 1 BTC to Ian")
  bc.AddBlock("Send 2 more BTC to Ian")

  for i := 0; i < len(bc.blocks); i++ {

    fmt.Printf("PrevHash %x\n", bc.blocks[i].PrevBlockHash)
    fmt.Printf("Hash %s\n", bc.blocks[i].Data)
    fmt.Printf("Hash %x\n", bc.blocks[i].Hash)
  }
}


