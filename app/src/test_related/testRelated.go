package test_related

import (
	"encoding/json"
	"github.com/geometry-labs/app/models"
	"io/ioutil"
)

type TestCrudData struct {
	Blocks *[]models.BlockRaw
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func LoadTestBlocks(filepath string) *[]models.BlockRaw {
	var blocks []models.BlockRaw
	dat, err := ioutil.ReadFile(filepath)
	check(err)
	json.Unmarshal(dat, &blocks)
	return &blocks
}
