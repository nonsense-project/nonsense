package model

import "github.com/nonsense-project/nonsense/v2/domain/consensus/model/externalapi"

// GHOSTDAGDataStore represents a store of BlockGHOSTDAGData
type GHOSTDAGDataStore interface {
	Store
	Stage(stagingArea *StagingArea, blockHash *externalapi.DomainHash, blockGHOSTDAGData *externalapi.BlockGHOSTDAGData, isTrustedData bool)
	IsStaged(stagingArea *StagingArea) bool
	Get(dbContext DBReader, stagingArea *StagingArea, blockHash *externalapi.DomainHash, isTrustedData bool) (*externalapi.BlockGHOSTDAGData, error)
	UnstageAll(stagingArea *StagingArea)
}
