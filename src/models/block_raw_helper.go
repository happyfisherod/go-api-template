package models

import (
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/encoding/protojson"
)

func ConvertToBlockRaw(value []byte) (*BlockRaw, error) {
	block := BlockRaw{}
	err := protojson.Unmarshal(value, &block)
	if err != nil {
		log.Error("Block_raw_helper: Error in ConvertToBlockRaw: %v", err)
	}
	return &block, err
}
