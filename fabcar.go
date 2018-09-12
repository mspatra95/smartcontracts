/*
Letter Of Creditdocument
A letter of credit is a document from a bank that guarantees payment.
*/
 package main
/* Imports 
* 2 specific Hyperledger Fabric specific libraries for Smart Contracts
*/
import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	 "github.com/hyperledger/fabric/core/chaincode/shim"
	 pb "github.com/hyperledger/fabric/protos/peer" 
 )

// Define the Smart Contract structure
type SmartContract struct {
}



//Invoke a chaincode function 
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response { 
function, args := stub.GetFunctionAndParameters() 
fmt.Println("invoke is running " + function) 
if function == locstep1 {
return t.locstep1(stub, args)	
}
fmt.Println("invoke did not find func: " + function)
return shim.Error("Received unknown function invocation: " + function)
}


func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
