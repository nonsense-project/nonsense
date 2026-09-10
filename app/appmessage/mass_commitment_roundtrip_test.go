package appmessage

import (
	"testing"

	"github.com/nonsense-project/nonsense/v2/domain/consensus/model/externalapi"
	"github.com/nonsense-project/nonsense/v2/domain/consensus/utils/constants"
)

func TestDomainTransactionMassCommitmentRoundTrip(t *testing.T) {
	const massCommitment = uint64(2036)
	tx := &externalapi.DomainTransaction{MassCommitment: massCommitment}

	msgTx := DomainTransactionToMsgTx(tx)
	if msgTx.MassCommitment != massCommitment {
		t.Fatalf("message mass commitment is %d, expected %d", msgTx.MassCommitment, massCommitment)
	}

	roundTripped := MsgTxToDomainTransaction(msgTx)
	if roundTripped.MassCommitment != massCommitment {
		t.Fatalf("round-tripped mass commitment is %d, expected %d", roundTripped.MassCommitment, massCommitment)
	}
}

func TestMsgTransactionDoesNotGuessMissingMassCommitment(t *testing.T) {
	tx := &externalapi.DomainTransaction{
		Inputs: []*externalapi.DomainTransactionInput{{SigOpCount: 1}},
		Outputs: []*externalapi.DomainTransactionOutput{{
			Value:           constants.SompiPerNonsense,
			ScriptPublicKey: &externalapi.ScriptPublicKey{Script: []byte{0x51}},
		}},
	}
	legacyMessage := DomainTransactionToMsgTx(tx)
	legacyMessage.MassCommitment = 0
	recovered := MsgTxToDomainTransaction(legacyMessage)
	if recovered.MassCommitment != 0 {
		t.Fatalf("unknown missing mass commitment was guessed as %d", recovered.MassCommitment)
	}
}
