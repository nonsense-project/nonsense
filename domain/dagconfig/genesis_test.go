// Copyright (c) 2014-2016 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package dagconfig

import (
	"testing"

	"github.com/nonsense-project/nonsense/v2/domain/consensus/utils/consensushashing"
	"github.com/nonsense-project/nonsense/v2/domain/consensus/utils/merkle"
)

func assertGenesisCommitments(t *testing.T, params *Params) {
	t.Helper()
	hash := consensushashing.BlockHash(params.GenesisBlock)
	if !params.GenesisHash.Equal(hash) {
		t.Fatalf("genesis block hash is %v, expected %v", hash, params.GenesisHash)
	}
	root := merkle.CalculateHashMerkleRoot(params.GenesisBlock.Transactions)
	if !params.GenesisBlock.Header.HashMerkleRoot().Equal(root) {
		t.Fatalf("genesis transaction merkle root is %v, expected %v", root, params.GenesisBlock.Header.HashMerkleRoot())
	}
	if params.GenesisBlock.Header.Bits() != 0x207fffff {
		t.Fatalf("genesis difficulty bits are %08x, expected 207fffff", params.GenesisBlock.Header.Bits())
	}
	if len(params.GenesisBlock.Transactions) != 1 || len(params.GenesisBlock.Transactions[0].Outputs) != 0 {
		t.Fatal("genesis must contain one coinbase transaction and no spendable outputs")
	}
}

// TestGenesisBlock tests the genesis block of the main network for validity by
// checking the encoded hash.
func TestGenesisBlock(t *testing.T) {
	assertGenesisCommitments(t, &MainnetParams)
}

// TestTestnetGenesisBlock tests the genesis block of the test network for
// validity by checking the hash.
func TestTestnetGenesisBlock(t *testing.T) {
	assertGenesisCommitments(t, &TestnetParams)
}

// TestSimnetGenesisBlock tests the genesis block of the simulation test network
// for validity by checking the hash.
func TestSimnetGenesisBlock(t *testing.T) {
	assertGenesisCommitments(t, &SimnetParams)
}

// TestDevnetGenesisBlock tests the genesis block of the development network
// for validity by checking the encoded hash.
func TestDevnetGenesisBlock(t *testing.T) {
	assertGenesisCommitments(t, &DevnetParams)
}
