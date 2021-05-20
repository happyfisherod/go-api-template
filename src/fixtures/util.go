package fixtures

import (
	"encoding/json"
	"fmt"
	"github.com/geometry-labs/app/models"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
)

func getWorkingDir() string {
	// Todo: generalize
	dir := "/Users/anuragjha/go/src/go-api-template/src/fixtures/"
	return dir
}

var fixtureAbsDir = getWorkingDir()

const (
	Block_raws_fixture = "block_raws.json"
)

type Fixtures []Fixture
type Fixture struct {
	Input    map[string]interface{}
	Expected map[string]interface{}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func LoadTestFixtures(file string) Fixtures {
	var fs Fixtures
	dat, err := ioutil.ReadFile(fixtureAbsDir + file)
	check(err)
	json.Unmarshal(dat, &fs)
	return fs
}

func ReadCurrentDir() {
	file, err := os.Open(".")
	if err != nil {
		log.Fatalf("failed opening directory: %s", err)
	}
	defer file.Close()

	list, _ := file.Readdirnames(0) // 0 to read all files and folders
	for _, name := range list {
		fmt.Println(name)
	}
}

func (f *Fixture) GetBlock(data map[string]interface{}) *models.BlockRaw {
	block := models.BlockRaw{
		Signature:        data["signature"].(string),
		ItemId:           data["item_id"].(string),
		NextLeader:       data["next_leader"].(string),
		TransactionCount: uint32(data["transaction_count"].(float64)),
		Type:             data["type"].(string),
		Version:          data["version"].(string),
		PeerId:           data["peer_id"].(string),
		Number:           uint32(data["number"].(float64)),
		MerkleRootHash:   data["merkle_root_hash"].(string),
		ItemTimestamp:    data["item_timestamp"].(string),
		Hash:             data["hash"].(string),
		ParentHash:       data["parent_hash"].(string),
		Timestamp:        uint64(data["timestamp"].(float64)),
	}
	return &block
}

// todo
///0. fixture file path map/enum => generalize to get correct fixture dir
////1. fixtures  - load and return []{input, expected} / expected can be an error => DONE
//2. create setup and teradon func here
////3. make init of test suite eg. encapsulate postgress connection - so that it can be called for test in other packages => DONE
///?4. param a test and loop over a array of fixtures
//5. unit vs integration test convention
//
//
//test:
//1. unit
//2. integration
//3. production

////////
//to get up and running
// 1. sarama
// 2. test
