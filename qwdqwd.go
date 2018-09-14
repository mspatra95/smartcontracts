/*
qwdrqwddocument
qdqwdwqd
*/
 package main
/* Imports 
* 2 specific Hyperledger Fabric specific libraries for Smart Contracts
*/
import ( 
 github.com/hyperledger/fabric/core/chaincode/shim
 sc github.com/hyperledger/fabric/protos/peer 
 )

// Define the Smart Contract structure
	type SmartContract struct {
	}

type RentalStruct struct {qwqwdsqwd integer}

//Invoke a chaincode function 
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response { 
function, args := stub.GetFunctionAndParameters() 
fmt.Println("invoke is running " + function) 
if function == weqe {
return t.weqe(stub, args)	
}
fmt.Println("invoke did not find func: " + function)
return shim.Error("Received unknown function invocation: " + function)
}


func (s *SmartContract) weqe (shim shim.ChaincodeStubInterface, args []string) sc.Response {
if len(args) == 1 {
return shim.Error( weqe accepts 1 argument ")
	}

qwdq := args[0]
}func (s *SmartContract) blockchainUpdate (shim shim.ChaincodeStubInterface, args []string) sc.Response {

qwqwdsqwd := args[0]
 shim.PutState(key, RentalStruct)
 }


func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}