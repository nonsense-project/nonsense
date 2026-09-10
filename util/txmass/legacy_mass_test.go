package txmass

import (
	"testing"

	"github.com/nonsense-project/nonsense/v2/domain/consensus/model/externalapi"
)

func TestPopulateMissingMassCommitment(t *testing.T) {
	tests := []struct {
		transactionID string
		mass          uint64
	}{
		{"e98ba22770a21b2b4e3d359a4e726d08f5c1ebb91e4e3622a91d34832ceca2da", 10026},
		{"9096f2d07ac92d66fe7fa0cfca71d5e4e071773057020da43079eddf9a4af8c4", 30012},
		{"1f093f04b7f507fc18948f1da0f2fd4b044a25b3b157e4ea871d439cb3f4b011", 60229},
	}

	for _, test := range tests {
		transactionID, err := externalapi.NewDomainTransactionIDFromString(test.transactionID)
		if err != nil {
			t.Fatalf("invalid test transaction ID: %s", err)
		}
		transaction := &externalapi.DomainTransaction{ID: transactionID}
		PopulateMissingMassCommitment(transaction)
		if transaction.MassCommitment != test.mass {
			t.Fatalf("transaction %s got mass %d, expected %d", test.transactionID, transaction.MassCommitment, test.mass)
		}
	}

	unknownID, err := externalapi.NewDomainTransactionIDFromString("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	if err != nil {
		t.Fatal(err)
	}
	unknown := &externalapi.DomainTransaction{ID: unknownID}
	PopulateMissingMassCommitment(unknown)
	if unknown.MassCommitment != 0 {
		t.Fatalf("unknown transaction got guessed mass %d", unknown.MassCommitment)
	}
}
