package blockchain

import (
	"bytes"
	"crypto/sha256"
)

type BlockChain struct {
	blocks []*Block
}

type Block struct {
	Hash 	 []byte
	Data	 []byte
	PrevHash []byte
}

func (b *Block) DeriveHash() {
	// Gerar um novo hash baseado nas infos da Data e PrevHash
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

func CreateBlock(data string, prevHash []byte) *Block {
	// Criar a estrutura do novo bloco
	block := &Block{[]byte{}, []byte(data), prevHash}
	block.DeriveHash()
	return block
}

func (chain *BlockChain) AddBlock(data string) {
	// Adicionar o block na corrente
	prevBlock := chain.blocks[len(chain.blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, new)
}

func Genesis() *Block {
	// Bloco inicial da rede
	return CreateBlock("Genesis", []byte{})
}

func InitBlockChain() *BlockChain {
	// Inicialização da BlockChain
	return &BlockChain{[]*Block{Genesis()}}
}
