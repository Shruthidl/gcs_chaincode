/*
Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at
  http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied.  See the License for the
specific language governing permissions and limitations
under the License.
*/

package main

import (
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

// ============================================================================================================================
// Main
// ============================================================================================================================
func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode - %s", err)
	}
}

// ============================================================================================================================
// Init - initialize the chaincode - marbles don’t need anything initlization, so let's run a dead simple test instead
// ============================================================================================================================
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	_, args := stub.GetFunctionAndParameters()
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	err = stub.PutState("marbles_ui", []byte(args[0]))
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}


// ============================================================================================================================
// Invoke - Our entry point for Invocations
// ============================================================================================================================
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	fmt.Println("starting invoke, for - " + function)

  if function != "invoke" {
        return shim.Error("Unknown function call")
	}
	if function == "invoke" {
        function = args[0]
	} 
  
	// Handle different functions
	if function == "init" {                     //initialize the chaincode state, used as reset
		return t.Init(stub)
	} else if function == "read" {              //generic read ledger
		return read(stub, args)
	} else if function == "write" {             //generic writes to ledger
		return write(stub, args)
	} else if function == "addOutClearFile" {   //add outclear file
		return addOutClearFile(stub, args)
	} else if function == "addInClearFile" {    //add inclear file
		return addInClearFile(stub, args)
	} else if function == "markTxnCleared" {    //mark txn cleared
		return markTxnCleared(stub, args)
	} else if function == "markFilesCleared"{   //mark files cleared
		return markFilesCleared(stub, args)
	} else if function == "getFiles"{           //get files
		return getFiles(stub)
	} else if function == "getAlltxns"{         //get txns
		return getAlltxns(stub, args)
	} else if function == "getCurrentFileId"{   // get current id
		return getCurrentFileId(stub, args)
	} else if function == "getCounts"{          //get counts
		return getCounts(stub, args)
	}

	// error out
	fmt.Println("Received unknown invoke function name - " + function)
	return shim.Error("Received unknown invoke function name - '" + function + "'")
}

// ============================================================================================================================
// Query - legacy function
// ============================================================================================================================
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Error("Unknown supported call - Query()")
}
