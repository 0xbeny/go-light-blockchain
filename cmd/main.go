package main

import (
	"fmt"
	go_blockchain "github.com/behnammohammadkhani/go-blockchain"
)

func main() {

	bc := go_blockchain.NewBlockchain(2)
	bc.Add("2nd")
	bc.Add("3rd")

	//bc.Blocks[1].Data = []byte("2nd copyyyyy")

	//bc.Timestamp = time.Now() -> Change the timestamp of the block
	//if err := bc.Validate(); err != nil {
	//	log.Fatalln("Hash of Block is not Valid")
	//}
	fmt.Println(bc)
}
