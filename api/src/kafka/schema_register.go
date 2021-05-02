package kafka

import (
	"encoding/binary"
	"fmt"
	"github.com/riferrei/srclient"
	"io/ioutil"
)

func RegisterSchema(topic string, isKey bool, srcSchemaFile string, forceUpdate bool) (int, error) {
	schemaRegistryClient := srclient.CreateSchemaRegistryClient("http://schemaregistry:8081")
	schema, err := schemaRegistryClient.GetLatestSchema(topic, false)
	if schema == nil {
		schema, err = registerSchema(schemaRegistryClient, topic, isKey, srcSchemaFile)
	} else if forceUpdate {
		schema, err = registerSchema(schemaRegistryClient, topic, isKey, srcSchemaFile) //TODO: Resolve update not happening
	}

	if err != nil {
		return 0, err
	}
	schemaIDBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(schemaIDBytes, uint32(schema.ID()))
	return schema.ID(), nil
}

func registerSchema(schemaRegistryClient *srclient.SchemaRegistryClient, topic string, isKey bool, srcSchemaFile string) (*srclient.Schema, error) {
	filePath := "models/" + srcSchemaFile + ".proto"
	fmt.Printf("Adding/Updating Schema from: %s\n", filePath)
	schemaBytes, _ := ioutil.ReadFile(filePath)
	schema, err := schemaRegistryClient.CreateSchema(topic, string(schemaBytes), srclient.Protobuf, isKey)
	if err != nil {
		//panic(fmt.Sprintf("Error creating the schema %s", err))
		fmt.Printf("Error creating the schema %s\n", err)
		return nil, err
	}
	return schema, nil
}
