package main

import (
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// NFT represents a unique non-fungible token
type NFT struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// SmartContract represents the logic for the contract
type SmartContract struct {
	contractapi.Contract
}
./network.sh createChannel
// Init is called when the contract is deployed
func (s *SmartContract) Init(ctx contractapi.TransactionContextInterface) error {
	fmt.Println("Smart contract initialized!")
	return nil
}

// MintNFT creates a new NFT and adds it to the ledger
func (s *SmartContract) MintNFT(ctx contractapi.TransactionContextInterface, id string, name string) error {
	nft := NFT{
		ID:   id,
		Name: name,
	}

	// Add the NFT to the ledger
	if err := ctx.GetStub().PutState(id, []byte(name)); err != nil {
		return err
	}

	fmt.Printf("Minted NFT with ID %s and name %s\n", id, name)
	return nil
}

func main() {
	chaincode, err := contractapi.NewChaincode(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating smart contract: %s", err)
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting smart contract: %s", err)
	}
}
