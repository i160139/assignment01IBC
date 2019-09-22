package assignment01IBC

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type Block struct {
	transaction string
	prevPointer *Block
	prevHash    string
}

func InsertBlock(transaction string, chainHead *Block) *Block {
	return &Block{
		transaction: transaction,
		prevPointer: chainHead,
		prevHash:    calculateHash(chainHead),
	}
}
func ListBlocks(chainHead *Block) {
	for ; chainHead != nil; {
		fmt.Println("Transaction: " + chainHead.transaction + "\n Previous Hash: " + chainHead.prevHash + "\n")
		chainHead = chainHead.prevPointer
	}
	fmt.Println()
}
func ChangeBlock(oldTrans string, newTrans string, chainHead *Block) {
	for ; chainHead.transaction != oldTrans && chainHead != nil; {
		chainHead = chainHead.prevPointer
	}
	if chainHead != nil {
		chainHead.transaction = newTrans

	}
}
func VerifyChain(chainHead *Block) {
	if chainHead != nil {
		for ; chainHead.prevPointer != nil; {
			if chainHead.prevHash != calculateHash(chainHead.prevPointer) {
				fmt.Println(chainHead.prevPointer.transaction + " has current hash as " + calculateHash(chainHead.prevPointer) + " but it should be " + chainHead.prevHash)
			}
			chainHead = chainHead.prevPointer
		}
	}
}

func calculateHash(block *Block) string {
	if block == nil {
		return "nil"
	} else {

		record := block.transaction + block.prevHash
		h := sha256.New()
		h.Write([]byte(record))
		hashed := h.Sum(nil)
		return hex.EncodeToString(hashed)
	}
}
