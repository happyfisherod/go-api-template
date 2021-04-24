package serdes

import (
	"encoding/json"
	"fmt"
	"github.com/geometry-labs/protoserdes/mocker"
	"google.golang.org/protobuf/types/known/structpb"
	"reflect"
)

func JsonToProto() {
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
	data := newContractTxDynamic["data"]
	fmt.Printf("3. %v\n", data) // but accessing fields is a headache
	typeOfData := reflect.TypeOf(data)
	fmt.Println(typeOfData)
	fmt.Println("printLeafValues golang struct::")
	printLeafValues(newContractTxDynamic)
	//// 4. convert json to map
	//var newContractTxDynamic1 interface{}
	//_ = json.Unmarshal([]byte(jsonstr), &newContractTxDynamic1)
	//newContractTxDynamic1Typecasted := newContractTxDynamic1.(map[string]interface{})
	//fmt.Printf("4. %v\n", newContractTxDynamic1Typecasted["data"]) // but accessing fields is a headache
	// 5. convert map to proto
	newContractTxProto, _ := structpb.NewStruct(newContractTxDynamic)
	fmt.Printf("5. %v\n", newContractTxProto.GetFields()["data"]) // but accessing fields is a headache
	fmt.Println("printLeafValues proto ::")
	printLeafValues(newContractTxDynamic)
	// so finally: we get a json, 1st we convert it to map[string]interface{}, 2nd we convert it to proto to communicate

}

func printLeafValues(data map[string]interface{}) {
	for k, v := range data {
		if reflect.ValueOf(v).Kind().String() == "map" {
			printLeafValues(v.(map[string]interface{}))
		} else {
			fmt.Println("k:", k, ", v:", v)
		}
	}
}