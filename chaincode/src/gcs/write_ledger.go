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
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// ============================================================================================================================
// write() - genric write variable into ledger
// 
// Shows Off PutState() - writting a key/value into the ledger
//
// Inputs - Array of strings
//    0   ,    1
//   key  ,  value
//  "abc" , "test"
// ============================================================================================================================

func write(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var key, value string
	var err error
	fmt.Println("starting write")

	if len(args) < 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2. key of the variable and value to set")
	}

  key = args[0]                                   //rename for funsies
	value = args[1]
	err = stub.PutState(key, []byte(value))         //write the variable into the ledger
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- end write")
	return shim.Success(nil)
}


// Add outclear files

func addOutClearFile(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    var err error;
    var counter1 int;
    var stringslice []string;

    //prepareData
    err = stub.PutState("364924",[]byte("City Bank - 130"))
    if err != nil {
       return shim.Error(err.Error())
    }
    
    err = stub.PutState("364914",[]byte("I Bank - 120"))
    if err != nil {
       return shim.Error(err.Error())
    }
    err = stub.PutState("364927",[]byte("My Bank - 140"))
    if err != nil {
       return shim.Error(err.Error())
    }
    err = stub.PutState("4321432100",[]byte("DCB Bank - 25"))
    if err != nil {
       return shim.Error(err.Error())
    }
    err = stub.PutState("1234123400",[]byte("Src Bank - 29"))
    if err != nil {
        return shim.Error(err.Error())
    }

	// add out clear files
    valAsbytes,err :=stub.GetState(strconv.Itoa(counter))
    s:=string(valAsbytes);

     if len(s) != 0 {
	     lastByByte := s[len(s)-1:]
       counter1, err =  strconv.Atoi(lastByByte)
       if err != nil {
          return shim.Error(err.Error())
        }
    } else {
          counter1 = 0
    }

       counter = counter1+1;

      counter_s := strconv.Itoa(counter)
      acq_name,err :=stub.GetState(args[0])
      stringslice = append(stringslice,args[2],args[1],args[3],args[4],string(acq_name),"In Process");
      stringvalues = append(stringslice,counter_s,counter_s);//string array (value)
      s_requester := counter_s //counter value(key)

	    var mCount int = len(stringvalues);
                parts  := make([]string, mCount );
                parts = stringvalues;

		if(!strings.HasPrefix(args[5] , "H-")){

			 parts[5] = "Rejected";
       stringBytes1 := strings.Join(parts, "|")

			 err = stub.PutState(s_requester, []byte(stringBytes1));
			 if err != nil {
					return shim.Error(err.Error())
			 }
			    return shim.Success(nil);
		}

		 if(!strings.HasPrefix(args[6] , "T-")){

			 parts[5] = "Rejected";
       stringBytes2 := strings.Join(parts, "|")
      
       err = stub.PutState(s_requester, []byte(stringBytes2));
			 if err != nil {
					return shim.Error(err.Error())
			 }
			
			    return shim.Success(nil)
		 }


	  parts[5] = "Validated";
		stringBytes := strings.Join(parts, "|")

		err = stub.PutState(s_requester, []byte(stringBytes));
			if err != nil {
			 return shim.Error(err.Error())
			}



	//enrich data
	 content := strings.Split(args[7], ",");
	 var m int = len(content);
   cont  := make([]string, m );
   cont = content;

   for i := 0; i < len(cont); i++ {
	 txncounter = txncounter + 1;
	 fmt.Println(cont[i]);
	 content1:=strings.Split(cont[i],"|")
   var mCount1 int = len(content1);
   parts1  := make([]string, mCount1 );
   parts1 = content1;
    
   var buffer bytes.Buffer
   buffer.WriteString(strconv.Itoa(txncounter));
   fmt.Println(buffer);
   buffer.WriteString("|");
   buffer.WriteString(strconv.Itoa(counter));
   buffer.WriteString("|");
   buffer.WriteString(cont[i]);
	 buffer.WriteString("|");
	 status := "Validated|20-01-2017 07:20AM";
	 if(strings.HasPrefix(parts1[0], "1240")){
         status = "Validated|20-01-2017 07:20AM";
	 }else{
          status = "Invalid|20-01-2017 07:20AM";
	 }
   fmt.Println(status);
	 card := "364924";
   
	 if(strings.HasPrefix(parts1[0], "364924")){
           card = "364924";
   } else if(strings.HasPrefix(parts1[0], "364914")){
          card = "364914";
   } else if(strings.HasPrefix(parts1[0], "364927")){
          card = "364927";
   }

     fmt.Println(card);

		 cardname,err :=stub.GetState(card)
		 cardname1,err :=stub.GetState(parts1[4])
		 buffer.WriteString(string(cardname));
	   buffer.WriteString("|");
	   buffer.WriteString(string(cardname1));
		 buffer.WriteString("|0.25|");
		 buffer.WriteString(status);

		err = stub.PutState("t"+strconv.Itoa(txncounter), []byte(buffer.String()));
      if err != nil {
					return shim.Error(err.Error())
			}
	}
      return shim.Success(nil)
}


// Add inclear files

func addInClearFile(stub shim.ChaincodeStubInterface, args []string) pb.Response {
  var err error;
  var counter2 int;
  var stringslice1 []string;

	//prepareData
	err = stub.PutState("364924",[]byte("City Bank - 130"))
      if err != nil {
          return shim.Error(err.Error())
      }
	err = stub.PutState("364914",[]byte("I Bank - 120"))
      if err != nil {
        return shim.Error(err.Error())
      }
	err = stub.PutState("364927",[]byte("My Bank - 140"))
      if err != nil {
          return shim.Error(err.Error())
      }
	err = stub.PutState("4321432100",[]byte("DCB Bank - 25"))
      if err != nil {
        return shim.Error(err.Error())
      }
	err = stub.PutState("1234123400",[]byte("Src Bank - 29"))
      if err != nil {
        return shim.Error(err.Error())
      }

	valAsbytes,err :=stub.GetState(strconv.Itoa(counter))
  s:=string(valAsbytes);

   if len(s) != 0 {
   lastByByte := s[len(s)-1:]
   counter2, err =  strconv.Atoi(lastByByte)
      if err != nil {
        return shim.Error(err.Error())
      }
   } else {
      counter2 = 0
   }

	counter = counter2+1;

	counter_s := strconv.Itoa(counter)
	acq_name,err :=stub.GetState(args[0])
	stringslice1 = append(stringslice1,args[2],args[1],args[3],args[4],string(acq_name),"In Process");

	stringvalues = append(stringslice1,counter_s,counter_s);//string array (value)
  s_requester := counter_s //counter value(key)
  stringBytes := strings.Join(stringvalues, "|")

	err = stub.PutState(s_requester, []byte(stringBytes));
		 if err != nil {
				return shim.Error(err.Error())
			}

	  return shim.Success(nil)
}


//Mark Txns Cleared

func markTxnCleared(stub shim.ChaincodeStubInterface, args []string) pb.Response {
  var s = args;
	var str1 bytes.Buffer;
  var mCount int = len(s);
  parts  := make([]string, mCount );
  parts = s;
        
  for j := 0; j < len(parts); j++ {

        valueAsBytes , err := stub.GetState("t"+parts[j]);
          if err != nil {
            return shim.Error(err.Error())
          }
        var str bytes.Buffer;
        str.WriteString(string(valueAsBytes));
        str.WriteString("|Cleared");
        err = stub.PutState("t"+parts[j], []byte(str.String()));
        str1 = str;
        if err != nil {
            return shim.Error(err.Error())
          }
  }
	
	    return shim.Success([]byte(str1.String()))
}


// Mark Files Cleared

func markFilesCleared(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	  var s = args;
    var mCount int = len(s);
    parts  := make([]string, mCount );
   	parts = s;

for j := 0; j < len(parts); j++ {
		valueAsBytes , err := stub.GetState(parts[j]);
      if err != nil {
       return shim.Error(err.Error())
      }
	  s1:=string(valueAsBytes);
	  var s2 = strings.Split(s1, "|");;
    var mCount int = len(s2);
    parts1  := make([]string, mCount );
   	parts1 = s2;
	  parts1[5] = "Cleared";
    stringbyte := strings.Join(parts1, "|")
		err = stub.PutState(parts[j],[]byte(stringbyte));

      if err != nil {
        return shim.Error(err.Error())
      }
 }
          return shim.Success(nil)
}

