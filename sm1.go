/*
Mortgagedocument
Mortgage Desc.
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



//Invoke a chaincode function 
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response { 
function, args := stub.GetFunctionAndParameters() 
fmt.Println("invoke is running " + function) 
if function == step1desc {
return t.step1desc(stub, args)	
}
fmt.Println("invoke did not find func: " + function)
return shim.Error("Received unknown function invocation: " + function)
}


func (s *SmartContract) step1desc (shim shim.ChaincodeStubInterface, args []string) sc.Response {
if len(args) == 3 {
return shim.Error( step1desc accepts 3 argument ")
	}
var1 := args[0]var2 := args[1]var3 := args[2]

func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}