package core

import (
	"github.com/geometry-labs/app/crud"
	"sync"
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
