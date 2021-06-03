package loader

import (
	"fmt"
	"github.com/geometry-labs/go-service-template/core"
	log "github.com/sirupsen/logrus"
)

func StartBlockRawsLoader() {
	go BlockRawsLoader()
}

func BlockRawsLoader() {
	for {
		block := <-core.GetGlobal().Blocks.GetWriteChan()
		core.GetGlobal().Blocks.Create(block)
		log.Debug(fmt.Sprintf("Loader BlockRaws: Loaded in postgres table BlockRaws, Block Number %d", block.Number))
	}
}
