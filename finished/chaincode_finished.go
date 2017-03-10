/*
Copyright IBM Corp 2016 All Rights Reserved.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
		 http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"bytes"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type SimpleChaincode struct {
}

var key string ;
var value string ;


   
var counter int = 0;
var stringvalues []string;



func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

// Init resets all the things
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1")
	}

	err := stub.PutState("abc", []byte(args[0]))
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// Invoke isur entry point to invoke a chaincode function
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("invoke is running " + function)

	// Handle different functions
	if function == "init" {
		return t.Init(stub, "init", args)
	} else if function == "write" {
		return t.write(stub, args)
	} else if function == "addLoc" {
		fmt.Println("**** First argument in addLoc:****" + args[0])
		return t.addLoc(stub, args)
	} else if function == "updateLocStatus" {
		fmt.Println("**** First argument in updateLocStatus:****" + args[0])
		return t.updateLocStatus(stub, args)
	} else if function == "uploadBol" {
		fmt.Println("**** First argument in uploadBol:****" + args[0])
		return t.uploadBol(stub, args)
	} else if function == "uploadContract" {
		fmt.Println("**** First argument in uploadContract:****" + args[0])
		return t.uploadContract(stub, args)
	}
	
	fmt.Println("invoke did not find func: " + function)

	return nil, errors.New("Received unknown function invocation")
}

// Query is our entry point for queries
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("query is running " + function)

	// Handle different functions
	if function == "read" { //read a variable
		return t.read(stub, args)
	} else if function == "getLoc" {
	//	i,err := strconv.Atoi(args[0])
	//	fmt.Println(err); 
		return t.getLoc(stub, args);
		 
	} else if function == "getList" {
	
		return t.getLocList(stub, args);
	} else if function == "getNumberOfLocs" {
	
		return t.getNumberOfLocs(stub, args);
	} 
	
	fmt.Println("query did not find func: " + function)

	return nil, errors.New("Received unknown function query")
}

// write - invoke function to write key/value pair
func (t *SimpleChaincode) write(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	
	var err error
	fmt.Println("running write()")

	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2. name of the key and value to set")
	}

	key = args[0] //rename for funsies
	value = args[1]
	
	fmt.Println("saving state for key: " + key);
	
	err = stub.PutState(key, []byte(value)) //write the variable into the chaincode state
	if err != nil {
		return nil, err
	}
	
	
	return nil, nil
}

