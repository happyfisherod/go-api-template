package crud

import (
	"github.com/geometry-labs/app/models"
	"gorm.io/gorm"
	"sync"
)

type BlockRawModel struct {
	db    *gorm.DB
	model *models.BlockRaw
}

var blockRawModelInstance *BlockRawModel
var blockRawModelOnce sync.Once

func GetBlockRawModel() *BlockRawModel {
	blockRawModelOnce.Do(func() {
		blockRawModelInstance = &BlockRawModel{
			db:    GetPostgresConn().conn,
			model: &models.BlockRaw{},
		}
	})
	return blockRawModelInstance
}

func NewBlockRawModel(conn *gorm.DB) *BlockRawModel {
	blockRawModelInstance = &BlockRawModel{
		db:    conn,
		model: &models.BlockRaw{},
	}
	return blockRawModelInstance
}

func (m *BlockRawModel) Migrate() error {
	err := m.db.AutoMigrate(m.model)
	return err
}

func (m *BlockRawModel) Create(block *models.BlockRaw) *gorm.DB {
	tx := m.db.Create(block)
	return tx
}

func (m *BlockRawModel) Update(oldBlock *models.BlockRaw, newBlock *models.BlockRaw, whereClause ...interface{}) *gorm.DB {
	tx := m.db.Model(oldBlock).Where(whereClause[0], whereClause[1:]).Updates(newBlock)
	return tx
}

func (m *BlockRawModel) Delete(conds ...interface{}) *gorm.DB {
	tx := m.db.Delete(m.model, conds...)
	return tx
}

func (m *BlockRawModel) FindOne(conds ...interface{}) (*models.BlockRaw, *gorm.DB) {
	block := &models.BlockRaw{}
	tx := m.db.Find(block, conds...)
	return block, tx
}

func (m *BlockRawModel) FindAll(conds ...interface{}) (*[]models.BlockRaw, *gorm.DB) {
	blocks := &[]models.BlockRaw{}
	tx := m.db.Find(blocks, conds...)
	return blocks, tx
}