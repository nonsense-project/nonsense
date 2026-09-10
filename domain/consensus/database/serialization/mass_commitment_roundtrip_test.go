package serialization

import (
	"testing"

	"github.com/nonsense-project/nonsense/v2/domain/consensus/model/externalapi"
	"github.com/nonsense-project/nonsense/v2/domain/consensus/utils/constants"
)

func TestDatabaseTransactionMassCommitmentRoundTrip(t *testing.T) {
	const massCommitment = uint64(2036)
	tx := &externalapi.DomainTransaction{MassCommitment: massCommitment}

	dbTx := DomainTransactionToDbTransaction(tx)
	if dbTx.Mass != massCommitment {
		t.Fatalf("database mass is %d, expected %d", dbTx.Mass, massCommitment)
	}

	roundTripped, err := DbTransactionToDomainTransaction(dbTx)
	if err != nil {
		t.Fatalf("DbTransactionToDomainTransaction failed: %s", err)
	}
	if roundTripped.MassCommitment != massCommitment {
		t.Fatalf("round-tripped mass commitment is %d, expected %d", roundTripped.MassCommitment, massCommitment)
	}
}

func TestDatabaseTransactionDoesNotGuessMissingMassCommitment(t *testing.T) {
	tx := &externalapi.DomainTransaction{
		Inputs: []*externalapi.DomainTransactionInput{{SigOpCount: 1}},
		Outputs: []*externalapi.DomainTransactionOutput{{
			Value:           constants.SompiPerNonsense,
			ScriptPublicKey: &externalapi.ScriptPublicKey{Script: []byte{0x51}},
		}},
	}
	legacyDatabaseTransaction := DomainTransactionToDbTransaction(tx)
	legacyDatabaseTransaction.Mass = 0
	recovered, err := DbTransactionToDomainTransaction(legacyDatabaseTransaction)
	if err != nil {
		t.Fatalf("DbTransactionToDomainTransaction failed: %s", err)
	}
	if recovered.MassCommitment != 0 {
		t.Fatalf("unknown missing mass commitment was guessed as %d", recovered.MassCommitment)
	}
}
