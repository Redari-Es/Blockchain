package main

func main() {
	// bc := NewBlockchain()
	// bc.AddBlock("Send 1 BTC to Shon")
	// bc.AddBlock("Send 2 more BTC to Shon")
	// for _, block := range bc.blocks {
	// 	fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
	// 	fmt.Printf("Data: %s\n", block.Data)
	// 	fmt.Printf("Hash: %x\n", block.Hash)
	// 	pow := NewProofOfWork(block)
	// 	fmt.Printf("Pow: %s\n", strconv.FormatBool(pow.Validate()))
	// 	fmt.Println()
	//
	// }

	bc := NewBlockchain()
	defer bc.db.Close()
	cli := CLI{bc}
	cli.Run()
}
