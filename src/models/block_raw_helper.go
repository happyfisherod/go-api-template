package models

import (
	"go.uber.org/zap"
	"google.golang.org/protobuf/encoding/protojson"
)

func ConvertToBlockRaw(value []byte) (*BlockRaw, error) {
	block := BlockRaw{}
	err := protojson.Unmarshal(value, &block)
	if err != nil {
		zap.S().Error("Block_raw_helper: Error in ConvertToBlockRaw: %v", err)
	}
	return &block, err
}
