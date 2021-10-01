package blockchain

import (
	"bytes"
	"encoding/gob"
	"log"
)

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

func Genesis() *Block {
	// Bloco inicial da rede
	return CreateBlock("Genesis", []byte{})
}

func (b *Block) Serialize() []byte {
	var res bytes.Buffer
	encoder := gob.NewEncoder(&res)

	err := encoder.Encode(b)

	Handle(err)

	return res.Bytes()
}

func Deserialize(data []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(data))

	err := decoder.Decode(&block)

	Handle(err)

	return &block
}

func Handle(err error) {
	if err != nil {
		log.Panic(err)
	}
}