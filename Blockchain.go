package go_blockchain

import (
	"bytes"
	"fmt"
	"time"
)

type Block struct {
	Timestamp time.Time
	Nonce     int
	Hash      []byte
	PrevHash  []byte
	Data      []byte
}

func (b *Block) String() string {
	return fmt.Sprintf(
		"Time: %s \nData: %s \nPrevHash: %x \nHash: %x \nNonce: %d\n--------------\n",
		b.Timestamp, b.Data, b.PrevHash, b.Hash, b.Nonce,
	)
}
func (b *Block) Validate(mask []byte) error {
	hash := EasyHash(b.Timestamp.UnixNano(), b.Data, b.PrevHash, b.Nonce)
	if !bytes.Equal(hash, b.Hash) {
		return fmt.Errorf("the hash is invalid it is %x should be %x\n----------------\n", hash, b.Hash)
	}
	if isGoodEnough(mask, b.Hash) {
		return fmt.Errorf("invalid mask")
	}
	return nil
}

func NewBlock(data string, mask []byte, prevHash []byte) *Block {
	b := Block{
		Timestamp: time.Now(),
		Data:      []byte(data),
		PrevHash:  prevHash,
	}
	b.Hash, b.Nonce = DifficultHash(mask, b.Timestamp.UnixNano(), b.Data, b.PrevHash)

	return &b
}

type Blockchain struct {
	Difficulty int
	Mask       []byte
	Blocks     []*Block
}

func (b *Blockchain) Add(data string) {
	ln := len(b.Blocks)
	if ln == 0 {
		panic("Why?")
	}
	b.Blocks = append(b.Blocks, NewBlock(data, b.Mask, b.Blocks[ln-1].Hash))

}
func (b *Blockchain) String() string {
	var a string
	for _, i := range b.Blocks {
		a += i.String()
	}
	return a

}
func (bc *Blockchain) Validate() error {
	for index, i := range bc.Blocks {
		if err := i.Validate(bc.Mask); err != nil {
			return fmt.Errorf("blockchain is not valid %w", err)
		}

		if index == 0 {
			continue
		}
		if !bytes.Equal(i.PrevHash, bc.Blocks[index-1].Hash) {
			return fmt.Errorf("invalid order")
		}
	}

	return nil
}
func NewBlockchain(difficulty int) *Blockchain {
	mask := GenerateMask(difficulty)
	bc := Blockchain{
		Difficulty: difficulty,
		Mask:       mask,
	}
	bc.Blocks = []*Block{NewBlock("Genesis", mask, []byte{})}
	return &bc
}
