/*
qwedqwddocument
wdqwdwqd
*/
 package main
/* Imports 
* 2 specific Hyperledger Fabric specific libraries for Smart Contracts
*/
import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// Define the Smart Contract structure
type SmartContract struct {
}



//Invoke a chaincode function 
func (t *SmartContract) Invoke(stub shim.ChaincodeStubInterface) pb.Response { 
	function, args := stub.GetFunctionAndParameters() 
	fmt.Println("invoke is running " + function) 
	if function == "www" {
		return t.www(stub, args)	
	}
	fmt.Println("invoke did not find func: " + function)
	return shim.Error("Received unknown function invocation: " + function)
}


//It encodes JSON data.
func marshalData(box Box) []byte {
	outputData, _ = json.Marshal(box)
	return outputData
}
//addition of two numbers
func addition( a float64, b float64) float64 \n{return  a+b}
//It generates JSON/Structure.
func unMarshalData(bytes []byte) Data {
	var data Data
	json.Unmarshal(data, &box1)
	return box1
}


func (t *SmartContract) www(shim shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error( "www accepts 1 argument ")
	}

	asda := args[0]
}

func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}