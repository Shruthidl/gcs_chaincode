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
	"bytes"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// ============================================================================================================================
// Read - read a generic variable from ledger
//
// Shows Off GetState() - reading a key/value from the ledger
//
// Inputs - Array of strings
//  0
//  key
//  "abc"
// 
// Returns - string
// ============================================================================================================================

func read(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var key, jsonResp string
	var err error
	fmt.Println("starting read")

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting key of the var to query")
	}

	// input sanitation
	err = sanitize_arguments(args)
	if err != nil {
		return shim.Error(err.Error())
	}

	key = args[0]
	valAsbytes, err := stub.GetState(key)           //get the var from ledger
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + key + "\"}"
		return shim.Error(jsonResp)
	}

	fmt.Println("- end read")
	return shim.Success(valAsbytes)                  //send it onward
}


//Return All Files

func getFiles(stub shim.ChaincodeStubInterface) pb.Response {

    	var list []string;

      for i := 1; i <=counter; i++ {
           valueAsBytes , err := stub.GetState(strconv.Itoa(i));
          if err != nil {
           return shim.Error(err.Error())
          }
          s:=string(valueAsBytes);
          var s2 = strings.Split(s, "|");

          parts1  := make([]string, len(s2)+2);
	        parts1[0] = strconv.Itoa(i)
	        copy(parts1[1:], s2)
	        parts1[1] = strconv.Itoa(i)
	        copy(parts1[2:], s2)
   	      s2 = parts1;
          s2= append(s2[:9], s2[9+1:]...)
          s2= append(s2[:8], s2[8+1:]...)
          s = strings.Join(s2, "|")
          list =append(list,s);
	   }

	  listByte := strings.Join(list, ",");
    return shim.Success([]byte(listByte))
}

//Return All transactions

func getAlltxns(stub shim.ChaincodeStubInterface, args []string) pb.Response {

    	var list []string;

	for i := 1; i <=txncounter; i++ {
	  valueAsBytes , err := stub.GetState("t"+strconv.Itoa(i));
	    if err != nil {
	     return shim.Error(err.Error())
      }
      s:=string(valueAsBytes);
      list =append(list,s);
  }

	  txnsByte := strings.Join(list, ",");
	  return shim.Success([]byte(txnsByte))

}

// Return current file ID

func getCurrentFileId(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	   return shim.Success([]byte(strconv.Itoa(counter)))
}

// Return Count

func getCounts(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	        var str bytes.Buffer;
		      str.WriteString("Files:");
          str.WriteString(strconv.Itoa(counter));
	        str.WriteString(",");
	        str.WriteString("Txns:");
          str.WriteString(strconv.Itoa(txncounter));
	      
					return shim.Success([]byte(str.String()))
    }

