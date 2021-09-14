package blockchain

type BlockChain struct {
	Blocks []*Block
}

type Block struct {
	Hash 	 []byte
	Data	 []byte
	PrevHash []byte
	Nonce    int
}

func CreateBlock(data string, prevHash []byte) *Block {
	// Criar a estrutura do novo bloco
	block := &Block{[]byte{}, []byte(data), prevHash, 0}
	pow := NewProof(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

func (chain *BlockChain) AddBlock(data string) {
	// Adicionar o block na corrente
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, new)
}

func Genesis() *Block {
	// Bloco inicial da rede
	return CreateBlock("Genesis", []byte{})
}

func InitBlockChain() *BlockChain {
	// Inicialização da BlockChain
	return &BlockChain{[]*Block{Genesis()}}
}
