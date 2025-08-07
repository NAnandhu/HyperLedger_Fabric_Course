package main
import (
    "log"
    "github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func main() { 
carContract := new(CarContract)
chaincode, err := contractapi.NewChaincode(carContract)
if err != nil { 

log.Panicf("Could not create chaincode." + err.Error()) } 
err = chaincode.Start() 

if err != nil { 
log.Panicf("Failed to start chaincode. " + err.Error()) }
}