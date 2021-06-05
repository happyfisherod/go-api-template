// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: block_raw.proto

/*
Package models is a generated protocol buffer package.

It is generated from these files:
	block_raw.proto
	blocks_raw.proto

It has these top-level messages:
	BlockRaw
	BlocksRaw
*/
package models

import context "context"

import errors1 "github.com/infobloxopen/protoc-gen-gorm/errors"
import field_mask1 "google.golang.org/genproto/protobuf/field_mask"
import gorm1 "github.com/jinzhu/gorm"
import gorm2 "github.com/infobloxopen/atlas-app-toolkit/gorm"

import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = fmt.Errorf
var _ = math.Inf

type BlockRawORM struct {
	Hash             string `gorm:"primary_key"`
	ItemId           string
	ItemTimestamp    string
	MerkleRootHash   string
	NextLeader       string
	Number           uint32
	ParentHash       string
	PeerId           string
	Signature        string
	Timestamp        uint64
	TransactionCount uint32
	Type             string
	Version          string
}

// TableName overrides the default tablename generated by GORM
func (BlockRawORM) TableName() string {
	return "block_raws"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *BlockRaw) ToORM(ctx context.Context) (BlockRawORM, error) {
	to := BlockRawORM{}
	var err error
	if prehook, ok := interface{}(m).(BlockRawWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Signature = m.Signature
	to.ItemId = m.ItemId
	to.NextLeader = m.NextLeader
	to.TransactionCount = m.TransactionCount
	to.Type = m.Type
	to.Version = m.Version
	to.PeerId = m.PeerId
	to.Number = m.Number
	to.MerkleRootHash = m.MerkleRootHash
	to.ItemTimestamp = m.ItemTimestamp
	to.Hash = m.Hash
	to.ParentHash = m.ParentHash
	to.Timestamp = m.Timestamp
	if posthook, ok := interface{}(m).(BlockRawWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *BlockRawORM) ToPB(ctx context.Context) (BlockRaw, error) {
	to := BlockRaw{}
	var err error
	if prehook, ok := interface{}(m).(BlockRawWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Signature = m.Signature
	to.ItemId = m.ItemId
	to.NextLeader = m.NextLeader
	to.TransactionCount = m.TransactionCount
	to.Type = m.Type
	to.Version = m.Version
	to.PeerId = m.PeerId
	to.Number = m.Number
	to.MerkleRootHash = m.MerkleRootHash
	to.ItemTimestamp = m.ItemTimestamp
	to.Hash = m.Hash
	to.ParentHash = m.ParentHash
	to.Timestamp = m.Timestamp
	if posthook, ok := interface{}(m).(BlockRawWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type BlockRaw the arg will be the target, the caller the one being converted from

// BlockRawBeforeToORM called before default ToORM code
type BlockRawWithBeforeToORM interface {
	BeforeToORM(context.Context, *BlockRawORM) error
}

// BlockRawAfterToORM called after default ToORM code
type BlockRawWithAfterToORM interface {
	AfterToORM(context.Context, *BlockRawORM) error
}

// BlockRawBeforeToPB called before default ToPB code
type BlockRawWithBeforeToPB interface {
	BeforeToPB(context.Context, *BlockRaw) error
}

// BlockRawAfterToPB called after default ToPB code
type BlockRawWithAfterToPB interface {
	AfterToPB(context.Context, *BlockRaw) error
}

// DefaultCreateBlockRaw executes a basic gorm create call
func DefaultCreateBlockRaw(ctx context.Context, in *BlockRaw, db *gorm1.DB) (*BlockRaw, error) {
	if in == nil {
		return nil, errors1.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(BlockRawORMWithBeforeCreate_); ok {
		if db, err = hook.BeforeCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(BlockRawORMWithAfterCreate_); ok {
		if err = hook.AfterCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

type BlockRawORMWithBeforeCreate_ interface {
	BeforeCreate_(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type BlockRawORMWithAfterCreate_ interface {
	AfterCreate_(context.Context, *gorm1.DB) error
}

// DefaultApplyFieldMaskBlockRaw patches an pbObject with patcher according to a field mask.
func DefaultApplyFieldMaskBlockRaw(ctx context.Context, patchee *BlockRaw, patcher *BlockRaw, updateMask *field_mask1.FieldMask, prefix string, db *gorm1.DB) (*BlockRaw, error) {
	if patcher == nil {
		return nil, nil
	} else if patchee == nil {
		return nil, errors1.NilArgumentError
	}
	var err error
	for _, f := range updateMask.Paths {
		if f == prefix+"Signature" {
			patchee.Signature = patcher.Signature
			continue
		}
		if f == prefix+"ItemId" {
			patchee.ItemId = patcher.ItemId
			continue
		}
		if f == prefix+"NextLeader" {
			patchee.NextLeader = patcher.NextLeader
			continue
		}
		if f == prefix+"TransactionCount" {
			patchee.TransactionCount = patcher.TransactionCount
			continue
		}
		if f == prefix+"Type" {
			patchee.Type = patcher.Type
			continue
		}
		if f == prefix+"Version" {
			patchee.Version = patcher.Version
			continue
		}
		if f == prefix+"PeerId" {
			patchee.PeerId = patcher.PeerId
			continue
		}
		if f == prefix+"Number" {
			patchee.Number = patcher.Number
			continue
		}
		if f == prefix+"MerkleRootHash" {
			patchee.MerkleRootHash = patcher.MerkleRootHash
			continue
		}
		if f == prefix+"ItemTimestamp" {
			patchee.ItemTimestamp = patcher.ItemTimestamp
			continue
		}
		if f == prefix+"Hash" {
			patchee.Hash = patcher.Hash
			continue
		}
		if f == prefix+"ParentHash" {
			patchee.ParentHash = patcher.ParentHash
			continue
		}
		if f == prefix+"Timestamp" {
			patchee.Timestamp = patcher.Timestamp
			continue
		}
	}
	if err != nil {
		return nil, err
	}
	return patchee, nil
}

// DefaultListBlockRaw executes a gorm list call
func DefaultListBlockRaw(ctx context.Context, db *gorm1.DB) ([]*BlockRaw, error) {
	in := BlockRaw{}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(BlockRawORMWithBeforeListApplyQuery); ok {
		if db, err = hook.BeforeListApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	db, err = gorm2.ApplyCollectionOperators(ctx, db, &BlockRawORM{}, &BlockRaw{}, nil, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(BlockRawORMWithBeforeListFind); ok {
		if db, err = hook.BeforeListFind(ctx, db); err != nil {
			return nil, err
		}
	}
	db = db.Where(&ormObj)
	db = db.Order("hash")
	ormResponse := []BlockRawORM{}
	if err := db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(BlockRawORMWithAfterListFind); ok {
		if err = hook.AfterListFind(ctx, db, &ormResponse); err != nil {
			return nil, err
		}
	}
	pbResponse := []*BlockRaw{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

type BlockRawORMWithBeforeListApplyQuery interface {
	BeforeListApplyQuery(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type BlockRawORMWithBeforeListFind interface {
	BeforeListFind(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type BlockRawORMWithAfterListFind interface {
	AfterListFind(context.Context, *gorm1.DB, *[]BlockRawORM) error
}
