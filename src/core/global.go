package core

import (
	"sync"

	"github.com/geometry-labs/go-service-template/crud"
)

type Global struct {
	Blocks *crud.BlockRawModel
}

var globalInstance *Global
var globalOnce sync.Once

func GetGlobal() *Global {
	globalOnce.Do(func() {
		globalInstance = &Global{
			Blocks: crud.GetBlockRawModel(),
		}
	})
	return globalInstance
}
