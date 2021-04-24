package main

import (
	"encoding/json"
	"fmt"
	"github.com/geometry-labs/protoserdes/mocker"
	"google.golang.org/protobuf/types/known/structpb"
)

func main() {
	fmt.Println("starting...")
	// 1. create jsonstr
	jsonstr := mocker.GetSampleNewContractTxJsonString()
	fmt.Printf("1. jsonstr:%s\n", jsonstr)
	// 2. convert json str to object, note: needs struct definition
	var newContractTx mocker.NewContractTx //Can handle nested json objects
	_ = json.Unmarshal([]byte(jsonstr), &newContractTx)
	fmt.Printf("2. %s\n", newContractTx.Data.Params.Txhash)
	// 3. convert json str to map // note: need no struct definition, dynamically through map
	var newContractTxDynamic map[string]interface{}
	_ = json.Unmarshal([]byte(jsonstr), &newContractTxDynamic)
	fmt.Printf("3. %v\n", newContractTxDynamic["data"]) // but accessing fields is a headache
	// 4. convert map to proto
	var newContractTxDynamic1 interface{}
	_ = json.Unmarshal([]byte(jsonstr), &newContractTxDynamic1)
	newContractTxDynamic1Typecasted := newContractTxDynamic1.(map[string]interface{})
	fmt.Printf("4. %v\n", newContractTxDynamic1Typecasted["data"]) // but accessing fields is a headache
	newContractTxProto, _ := structpb.NewStruct(newContractTxDynamic)
	fmt.Printf("5. %v\n", newContractTxProto.GetFields()["data"])

}
