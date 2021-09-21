package blockchain

type BlockChain struct {
	Blocks []*Block
}

type Block struct {
	Hash 	 []byte // Hash é um algoritmo ou fórmula muito complicada que converte qualquer sequência de caracteres em uma sequência de 64 caracteres ou números
	Data	 []byte
	PrevHash []byte
	Nonce    int // Nonce é um número aleatório usado apenas uma vez
}

func CreateBlock(data string, prevHash []byte) *Block {
	// Criar a estrutura do novo bloco
	block := &Block{[]byte{}, []byte(data), prevHash, 0}
	/*
	Prova de trabalho - Proof of Work
	----------------------------------
	O Proof of Work é um algoritmo de consenso no qual é caro 
	e demorado produzir uma parte dos dados, mas é fácil para 
	outras pessoas verificarem se os dados estão corretos.
	*/
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
