/*
testingdocument
testing
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


type contractDetails struct {
contractDate date
contractCompany string
}

//Invoke a chaincode function 
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response { 
function, args := stub.GetFunctionAndParameters() 
fmt.Println("invoke is running " + function) 
if function == step1 {
return t.step1(stub, args)	
}
fmt.Println("invoke did not find func: " + function)
return shim.Error("Received unknown function invocation: " + function)
}


//addition of two numbers
func addition( a float64, b float64) float64 \n{return  a+b}
//It encodes JSON data.
func marshalData(box Box) []byte {
	outputData, _ = json.Marshal(box)
	return outputData
}


func (s *SmartContract) step1 (shim shim.ChaincodeStubInterface, args []string) sc.Response {
if len(args) == 2 {
return shim.Error( step1 accepts 2 argument ")
	}

var1 := args[0]
var2 := args[1]
if a ==  b  {
 a =  b 
 }
if a ==  b  {
 a =  c 
 }
}

func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}