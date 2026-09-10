package consensusstatemanager

import (
	"github.com/nonsense-project/nonsense/v2/domain/consensus/model"
	"github.com/nonsense-project/nonsense/v2/domain/consensus/model/externalapi"
	"github.com/nonsense-project/nonsense/v2/domain/consensus/model/testapi"
)

type testConsensusStateManager struct {
	*consensusStateManager
}

// NewTestConsensusStateManager creates an instance of a TestConsensusStateManager
func NewTestConsensusStateManager(baseConsensusStateManager model.ConsensusStateManager) testapi.TestConsensusStateManager {
	return &testConsensusStateManager{consensusStateManager: baseConsensusStateManager.(*consensusStateManager)}
}

func (csm *testConsensusStateManager) AddUTXOToMultiset(
	multiset model.Multiset, entry externalapi.UTXOEntry, outpoint *externalapi.DomainOutpoint) error {

	return addUTXOToMultiset(multiset, entry, outpoint)
}

func (csm *testConsensusStateManager) ResolveBlockStatus(stagingArea *model.StagingArea, blockHash *externalapi.DomainHash,
	useSeparateStagingAreaPerBlock bool) (externalapi.BlockStatus, error) {

	status, _, err := csm.resolveBlockStatus(stagingArea, blockHash, useSeparateStagingAreaPerBlock)
	return status, err
}
