package protowire

import (
	"testing"

	"github.com/nonsense-project/nonsense/v2/app/appmessage"
)

func TestP2PTransactionMassCommitmentRoundTrip(t *testing.T) {
	const massCommitment = uint64(2036)
	msgTx := &appmessage.MsgTx{MassCommitment: massCommitment}

	wireTx := new(TransactionMessage)
	wireTx.fromAppMessage(msgTx)
	if wireTx.Mass != massCommitment {
		t.Fatalf("wire mass is %d, expected %d", wireTx.Mass, massCommitment)
	}

	roundTrippedMessage, err := wireTx.toAppMessage()
	if err != nil {
		t.Fatalf("toAppMessage failed: %s", err)
	}
	roundTripped := roundTrippedMessage.(*appmessage.MsgTx)
	if roundTripped.MassCommitment != massCommitment {
		t.Fatalf("round-tripped mass commitment is %d, expected %d", roundTripped.MassCommitment, massCommitment)
	}
}
