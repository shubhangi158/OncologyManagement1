package main

import(
	//"fmt"
	"errors"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"encoding/json"
	//"strconv"
)

var logger = shim.NewLogger("PatientRecord")

type Chaincode struct{
}

type Contact struct{
	ContactId string
	Name string
	Age int
	Gender string
	Race string
}

type Case struct{
	N_Stage string
	T_Stage string
	Grade string
	Condition string
	Cancer_Diagnosis string
	Cancer_Stage string
	Metastasis_Location string
	NCCN_Distress_Score	int
	Survival_Time int									//Survival Time is in percentage.
} 

type BackgroundInformation struct{
	Affected_Breast string
	ER_Status_LB string
	ER_Status_RB string
	HER2_Status_LB string
	HER2_Status_RB string
	PR_Status_LB string
	PR_Status_RB string
}

type MedicalRecord struct{
	Contact_Rec Contact
	Case_Rec Case
	BackgroundInfo BackgroundInformation
}

func main(){
	err := shim.Start(new(Chaincode))
	if err != nil{
		errors.New("MAIN: Error in starting chaincode.")
	}
}

func (c *Chaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error){

	err := stub.PutState(args[0], []byte(args[0]))
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (c *Chaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error){
	
	if len(args) != 21 {
		errors.New("INVOKE: Incorrect number of arguments passed.")
	}
	
	/*age, err := strconv.Atoi(args[2])
	if err != nil{
		errors.New("INVOKE: String to int conversion failed.")
	}*/
	
	if function == "writeMedicalRecord" {
		return c.writeMedicalRecord(stub, args)
	} 
	/*else if function == "updateMedicalRecord" {
		return c.updateMedicalRecord(stub, args)
	}*/
	
	return nil, nil
}

func (c *Chaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error){

	/*if len(args) != 1{
		errors.New("QUERY: Incorrect number of arguments passed.")
	}*/

	if function == "readMedicalRecord" {
		return c.readMedicalRecord(stub, args[0])
	}

	return nil, nil
}

/*func (c *Chaincode) writeMedicalRecord(stub shim.ChaincodeStubInterface, args []string) ([]byte, error){

	if args[0] == ""{
		return nil, errors.New("WRITE_MEDICAL_RECORD: Invalid Contact ID provided.")
	}
	
	var cont Contact
	contactJson  := []byte(`{"ContactId":"` + args[0] + `","Name":"` + args[1] + `","Age":` + args[2] + `,"Gender":"` + args[3] + `","Race":"` + args[4] + `"}`)
	err := json.Unmarshal(contactJson, &cont)
	if err != nil{
		errors.New("WRITE_MEDICAL_RECORD: Invalid JSON object.")
	}

	var record MedicalRecord
	recordJson  := []byte(`{"ContactId":"` + contactId + `","Name":"` + name + `","Age":` + strconv.Itoa(age) + `,"Gender":"` + gender + `","Race":"` + race + `"}`)
	
	err := json.Unmarshal(recordJson, &record)
	if err != nil{
		errors.New("WRITE_MEDICAL_RECORD: Invalid JSON object.")
	}
	
	c.saveRecord(stub, cont)
	
	return nil, nil
}*/

func (c *Chaincode) writeMedicalRecord(stub shim.ChaincodeStubInterface, args []string) ([]byte, error){

	if args[0] == ""{
		return nil, errors.New("WRITE_MEDICAL_RECORD: Invalid Contact ID provided.")
	}
	
	err := stub.PutState(args[0], []byte(args[1]))
	if err != nil{
		errors.New("SAVE_RECORD: Error saving Medical Record.")
	}
	
	return nil, nil
}

func (c *Chaincode) saveRecord(stub shim.ChaincodeStubInterface, cont Contact) (bool, error){
	
	contact, err := json.Marshal(cont)
	if err != nil{
		errors.New("SAVE_RECORD: Error encoding Medical Record.")
	}
	
	err = stub.PutState(cont.ContactId, contact)
	if err != nil{
		errors.New("SAVE_RECORD: Error saving Medical Record.")
	}
	
	return true, nil
}

func (c *Chaincode) readMedicalRecord(stub shim.ChaincodeStubInterface, contactId string) ([]byte, error){

	//var cont Contact

	recordAsBytes, err := stub.GetState(contactId);
	if err != nil{
		errors.New("READ_MEDICAL_RECORD: Failed to get Medical Record")
	}
	
	/*err = json.Unmarshal(recordAsBytes, &cont)
	if err != nil{
		errors.New("READ_MEDICAL_RECORD: Corrupt Medical Record")
	}*/
	
	return recordAsBytes, nil
}

/*func (c *Chaincode) updateMedicalRecord(stub shim.ChaincodeStubInterface, args []string) ([]byte, error){
	return nil, nil
}*/
