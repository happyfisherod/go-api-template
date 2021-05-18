package core

import (
	"github.com/geometry-labs/app/crud/postgres_crud"
	"sync"
)

type Global struct {
	Blocks *postgres_crud.BlockRawModel
}

var globalInstance *Global
var globalOnce sync.Once

func GetGlobal() *Global {
	globalOnce.Do(func() {
		globalInstance = &Global{
			Blocks: postgres_crud.GetBlockRawModel(),
		}
	})
	return globalInstance
}
