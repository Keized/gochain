package blockchain

type BlockChain struct {
	Blocks []*Block
}

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nonce    int
}

func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, new)
}

func CreateBlock(data string, prevHash []byte) *Block {
	// Create a block from the provided data
	block := &Block{[]byte{}, []byte(data), prevHash, 0}
	// Create a proof of work that contains the block and the target
	// the miner must find a "nonce" which, once appended to the data and "hashed" must have a lower value than the target
	pow := NewProofOfWork(block)
	// Guess the nonce
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}
