package loader

import (
	"fmt"
	"github.com/geometry-labs/go-service-template/core"
	"github.com/geometry-labs/go-service-template/models"
	"go.uber.org/zap"
)

func StartBlockRawsLoader() {
	go BlockRawsLoader()
}

func BlockRawsLoader() {
	var block *models.BlockRaw
	postgresLoaderChan := core.GetGlobal().Blocks.GetWriteChan()
	for {
		block = <-postgresLoaderChan
		core.GetGlobal().Blocks.Create(block)
		zap.S().Debug(fmt.Sprintf(
			"Loader BlockRaws: Loaded in postgres table BlockRaws, Block Number %d", block.Number),
		)
	}
}
