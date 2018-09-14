/*
Letter of Creditdocument
Use case for LoC between importer and exporter
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


type Lc struct {
ShipmentId string
ExporterCompany string
ExporterBank string
}

//Invoke a chaincode function 
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response { 
function, args := stub.GetFunctionAndParameters() 
fmt.Println("invoke is running " + function) 
if function == checkContarct {
return t.checkContarct(stub, args)	
}
else if function == getAllLCs {
return t.getAllLCs(stub, args)
}
else if function == write {
return t.write(stub, args)
}
else if function == read {
return t.read(stub, args)
}
else if function == createLC {
return t.createLC(stub, args)
}
fmt.Println("invoke did not find func: " + function)
return shim.Error("Received unknown function invocation: " + function)
}


//It encodes JSON data.
func marshalData(box Box) []byte {
	outputData, _ = json.Marshal(box)
	return outputData
}
//Use this function to read data from BlockChain
func (t *SimpleChaincode) read(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var key, jsonResp string
	var err error

    if len(args) != 1 {
        return shim.Error("Incorrect number of arguments. Expecting name of the key to query")
    }

    key = args[0]
    valAsbytes, err := stub.GetState(key)
    if err != nil {
        jsonResp = "{\"Error\":\"Failed to get state for " + key + "\"}"
        return shim.Error(jsonResp)
    }
    return shim.Success(valAsbytes)
}
//Use this function to write data into the BlockChain
func (t *SimpleChaincode) write(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var key, value string
	var err error
	fmt.Println("running write()")

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2. name of the key and value to set")
	}

	key = args[0]                            //rename for fun
	value = args[1]
	err = stub.PutState(key, []byte(value))  //write the variable into the chaincode state
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}
//retrieve all contracts
func custom_getalllcs() {

/*
	@description:retrieve all contracts
	@input variables:
	@input variables type:
	@output:
	@output type: 
*/
var allLCs []LC

	// Get list of all the keys
	keysBytes, err := stub.GetState("LCKeys")
	if err != nil {
		fmt.Println("Error retrieving LC keys")
		return shim.Error("Error retrieving LC keys")
	}
	var keys []string
	err = json.Unmarshal(keysBytes, &keys)
	if err != nil {
		fmt.Println("Error unmarshalling LC keys" + err.Error())
		return shim.Error("Error unmarshalling LC keys")
	}

	// Get all the lcs
	for _, value := range keys {
		lcBytes, err := stub.GetState(value)

		var lc LC
		err = json.Unmarshal(lcBytes, &lc)
		if err != nil {
			fmt.Println("Error retrieving lc " + value)
			return shim.Error("Error retrieving lc " + value)
		}

		fmt.Println("Appending LC" + value)
		allLCs = append(allLCs, lc)
	}
	
	allLCsAsBytes, _ := json.Marshal(allLCs);
	return shim.Success(allLCsAsBytes)
}
//create LC contract
func custom_createLC() {

/*
	@description:create LC contract
	@input variables:
	@input variables type:
	@output:
	@output type: 
*/

if len(args) != 1 {
	fmt.Println("Error obtaining LC JSON string")
	return shim.Error("createLC accepts 1 argument (LCJSONString)")
}

var lc LC
var err error
fmt.Println("Unmarshalling LC");

err = json.Unmarshal([]byte(args[0]), &lc)
if err != nil {
	fmt.Println("error invalid LC string")
	return shim.Error("Invalid LC string")
}

var shipmentId string
shipmentId = lc.ShipmentId

existingBytes, err := stub.GetState(shipmentId);
 
	lc.CurrentStatus = "Created"
	lc.ExporterBankApproved = false
	lc.ExporterDocsUploaded = false
	lc.CustomsApproved = false
	lc.PaymentComplete = false
	
	lcBytes, _ := json.Marshal(lc)
	if err != nil {
		fmt.Println("Error marshalling lc")
		return shim.Error("Error marshalling LC")
	}
	
	err = stub.PutState(shipmentId, lcBytes)
	if err != nil {
		fmt.Println("Error creating LC")
		return shim.Error("Error creating LC")
	}
return shim.Success(nil)

}
//this is a sample custom LC function
func customLCFunction() {

/*
	@description:this is a sample custom LC function
	@input variables:
	@input variables type:
	@output:
	@output type: 
*/
    return True
}
//power of n numbers
func power(x float64, y float64) float64{
    return math.Pow(x,y)
}
//truncate of n numbers
func truncate(x float64) float64 {
    return math.Trunc(x)
}


func (s *SmartContract) checkContarct (shim shim.ChaincodeStubInterface, args []string) sc.Response {
if len(args) == 4 {
return shim.Error( checkContarct accepts 4 argument ")
	}

ShipmentId := args[0]
ExporterCompany := args[1]
ExporterBank := args[2]
ImporterCompany := args[3]
if obj1.pr1 !=  obj2.pr2 &&  5 ==  6  {
 a =  b 
 }
if Status ==  approve  {
 existingorder_state =  2 
 }
if Status ==  rejected  {
 existingorder_state =  3 
 }
}

func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
