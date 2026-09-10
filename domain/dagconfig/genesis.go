// Copyright (c) 2014-2016 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package dagconfig

import (
	"math/big"

	"github.com/kaspanet/go-muhash"
	"github.com/nonsense-project/nonsense/v2/domain/consensus/model/externalapi"
	"github.com/nonsense-project/nonsense/v2/domain/consensus/utils/blockheader"
	"github.com/nonsense-project/nonsense/v2/domain/consensus/utils/subnetworks"
	"github.com/nonsense-project/nonsense/v2/domain/consensus/utils/transactionhelper"
)

var genesisTxOuts = []*externalapi.DomainTransactionOutput{}

var genesisTxPayload = append([]byte{
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Blue score
	0x00, 0xE1, 0xF5, 0x05, 0x00, 0x00, 0x00, 0x00, // Subsidy
	0x00, 0x00, // Script version
	0x01, // Varint
	0x00, // OP-FALSE
}, []byte("Nonsense mainnet genesis 2026-08-30 | supply 2000000000 NNN | no premine")...)

// genesisCoinbaseTx is the coinbase transaction for the genesis blocks for
// the main network.
var genesisCoinbaseTx = transactionhelper.NewSubnetworkTransaction(0, []*externalapi.DomainTransactionInput{}, genesisTxOuts,
	&subnetworks.SubnetworkIDCoinbase, 0, genesisTxPayload)

// genesisHash is the hash of the first block in the block DAG for the main
// network (genesis block).
var genesisHash = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	0xb6, 0xa9, 0xf4, 0x54, 0x3f, 0xe8, 0x03, 0x8b,
	0x61, 0x22, 0xc3, 0xac, 0x7d, 0xb3, 0x8b, 0x97,
	0x60, 0x0f, 0x1e, 0x24, 0x35, 0x7a, 0x72, 0xca,
	0xb5, 0xf0, 0x31, 0xb1, 0x6d, 0x40, 0xbc, 0x4f,
})

// genesisMerkleRoot is the hash of the first transaction in the genesis block
// for the main network.
var genesisMerkleRoot = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	0x0b, 0xbc, 0x4d, 0x1f, 0xec, 0xae, 0x80, 0xa1,
	0xf3, 0xdd, 0xd2, 0x92, 0x44, 0x06, 0x7b, 0x64,
	0xe1, 0x40, 0x30, 0x31, 0x20, 0xc2, 0xa4, 0x5e,
	0xf9, 0x3a, 0x50, 0xce, 0xfe, 0x32, 0x23, 0x03,
})

// genesisBlock defines the genesis block of the block DAG which serves as the
// public transaction ledger for the main network.
var genesisBlock = externalapi.DomainBlock{
	Header: blockheader.NewImmutableBlockHeader(
		0,
		[]externalapi.BlockLevelParents{},
		genesisMerkleRoot,
		&externalapi.DomainHash{},
		externalapi.NewDomainHashFromByteArray(muhash.EmptyMuHashHash.AsArray()),
		1788048000000,
		0x207fffff,
		0,
		0,
		0,
		big.NewInt(0),
		&externalapi.DomainHash{},
	),
	Transactions: []*externalapi.DomainTransaction{genesisCoinbaseTx},
}

var devnetGenesisTxOuts = []*externalapi.DomainTransactionOutput{}

var devnetGenesisTxPayload = append([]byte{
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Blue score
	0x00, 0xE1, 0xF5, 0x05, 0x00, 0x00, 0x00, 0x00, // Subsidy
	0x00, 0x00, // Script version
	0x01, // Varint
	0x00, // OP-FALSE
}, []byte("Nonsense devnet genesis 2026-08-30")...)

// devnetGenesisCoinbaseTx is the coinbase transaction for the genesis blocks for
// the development network.
var devnetGenesisCoinbaseTx = transactionhelper.NewSubnetworkTransaction(0,
	[]*externalapi.DomainTransactionInput{}, devnetGenesisTxOuts,
	&subnetworks.SubnetworkIDCoinbase, 0, devnetGenesisTxPayload)

// devGenesisHash is the hash of the first block in the block DAG for the development
// network (genesis block).
var devnetGenesisHash = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	0x95, 0xa2, 0xa3, 0x69, 0x40, 0xa4, 0xf7, 0xdb,
	0xe2, 0x53, 0xc6, 0x06, 0x93, 0xc5, 0x6d, 0xad,
	0x5c, 0x8c, 0x5a, 0xa9, 0x0e, 0x33, 0xbf, 0x75,
	0x0d, 0xc4, 0x6f, 0x5b, 0xcb, 0xd4, 0x2d, 0x65,
})

// devnetGenesisMerkleRoot is the hash of the first transaction in the genesis block
// for the devopment network.
var devnetGenesisMerkleRoot = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	0xa0, 0x7b, 0x46, 0x52, 0x0b, 0xae, 0x07, 0x9c,
	0xce, 0x4d, 0x1a, 0xc4, 0xbc, 0x40, 0x01, 0xb9,
	0x62, 0x9c, 0x5e, 0x9e, 0x80, 0xc7, 0xd4, 0x9c,
	0x45, 0x4b, 0x38, 0x9a, 0x3e, 0xbd, 0x4c, 0x2d,
})

// devnetGenesisBlock defines the genesis block of the block DAG which serves as the
// public transaction ledger for the development network.
var devnetGenesisBlock = externalapi.DomainBlock{
	Header: blockheader.NewImmutableBlockHeader(
		0,
		[]externalapi.BlockLevelParents{},
		devnetGenesisMerkleRoot,
		&externalapi.DomainHash{},
		externalapi.NewDomainHashFromByteArray(muhash.EmptyMuHashHash.AsArray()),
		1788048000003,
		0x207fffff,
		2,
		0,
		0,
		big.NewInt(0),
		&externalapi.DomainHash{},
	),
	Transactions: []*externalapi.DomainTransaction{devnetGenesisCoinbaseTx},
}

var simnetGenesisTxOuts = []*externalapi.DomainTransactionOutput{}

var simnetGenesisTxPayload = append([]byte{
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Blue score
	0x00, 0xE1, 0xF5, 0x05, 0x00, 0x00, 0x00, 0x00, // Subsidy
	0x00, 0x00, // Script version
	0x01, // Varint
	0x00, // OP-FALSE
}, []byte("Nonsense simnet genesis 2026-08-30")...)

// simnetGenesisCoinbaseTx is the coinbase transaction for the simnet genesis block.
var simnetGenesisCoinbaseTx = transactionhelper.NewSubnetworkTransaction(0,
	[]*externalapi.DomainTransactionInput{}, simnetGenesisTxOuts,
	&subnetworks.SubnetworkIDCoinbase, 0, simnetGenesisTxPayload)

// simnetGenesisHash is the hash of the first block in the block DAG for
// the simnet (genesis block).
var simnetGenesisHash = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	0xc1, 0xb2, 0xae, 0xca, 0x96, 0xd3, 0x0c, 0x21,
	0x87, 0x9c, 0x7d, 0x16, 0x7b, 0x8c, 0x2a, 0x64,
	0x00, 0x2a, 0x2d, 0xfd, 0xed, 0x09, 0xb5, 0x07,
	0xa2, 0x8b, 0x2b, 0x17, 0x7c, 0x98, 0xf3, 0xb2,
})

// simnetGenesisMerkleRoot is the hash of the first transaction in the genesis block
// for the development network.
var simnetGenesisMerkleRoot = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	0x03, 0xda, 0x32, 0x49, 0x1c, 0x63, 0x88, 0x16,
	0x9e, 0xae, 0x6d, 0x43, 0x9c, 0xa7, 0x48, 0x1d,
	0x5b, 0x12, 0x68, 0x7c, 0x55, 0xe9, 0x98, 0x9c,
	0x24, 0x7b, 0x59, 0xc2, 0x5d, 0x78, 0x2b, 0x2e,
})

// simnetGenesisBlock defines the genesis block of the block DAG which serves as the
// public transaction ledger for the development network.
var simnetGenesisBlock = externalapi.DomainBlock{
	Header: blockheader.NewImmutableBlockHeader(
		0,
		[]externalapi.BlockLevelParents{},
		simnetGenesisMerkleRoot,
		&externalapi.DomainHash{},
		externalapi.NewDomainHashFromByteArray(muhash.EmptyMuHashHash.AsArray()),
		1788048000002,
		0x207fffff,
		4,
		0,
		0,
		big.NewInt(0),
		&externalapi.DomainHash{},
	),
	Transactions: []*externalapi.DomainTransaction{simnetGenesisCoinbaseTx},
}

var testnetGenesisTxOuts = []*externalapi.DomainTransactionOutput{}

var testnetGenesisTxPayload = append([]byte{
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Blue score
	0x00, 0xE1, 0xF5, 0x05, 0x00, 0x00, 0x00, 0x00, // Subsidy
	0x00, 0x00, // Script version
	0x01, // Varint
	0x00, // OP-FALSE
}, []byte("Nonsense testnet genesis 2026-08-30")...)

// testnetGenesisCoinbaseTx is the coinbase transaction for the testnet genesis block.
var testnetGenesisCoinbaseTx = transactionhelper.NewSubnetworkTransaction(0,
	[]*externalapi.DomainTransactionInput{}, testnetGenesisTxOuts,
	&subnetworks.SubnetworkIDCoinbase, 0, testnetGenesisTxPayload)

// testnetGenesisHash is the hash of the first block in the block DAG for the test
// network (genesis block).
var testnetGenesisHash = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	0x67, 0xff, 0xfe, 0x4c, 0x60, 0x4a, 0x4a, 0xf0,
	0x27, 0x39, 0xd2, 0x9b, 0xca, 0xc8, 0x94, 0x81,
	0x82, 0x46, 0xda, 0x01, 0x85, 0xaa, 0x4a, 0xd3,
	0x39, 0xda, 0x30, 0x53, 0x32, 0x74, 0x77, 0xf4,
})

// testnetGenesisMerkleRoot is the hash of the first transaction in the genesis block
// for testnet.
var testnetGenesisMerkleRoot = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	0x8b, 0xf4, 0x5c, 0x30, 0xa7, 0x70, 0xa2, 0x2d,
	0xf2, 0x73, 0x40, 0x6b, 0x21, 0x46, 0xd8, 0x49,
	0x88, 0x03, 0x44, 0x28, 0x6a, 0x38, 0x67, 0x46,
	0xdf, 0xed, 0xb5, 0x44, 0xb7, 0xee, 0xd0, 0x85,
})

// testnetGenesisBlock defines the genesis block of the block DAG which serves as the
// public transaction ledger for testnet.
var testnetGenesisBlock = externalapi.DomainBlock{
	Header: blockheader.NewImmutableBlockHeader(
		0,
		[]externalapi.BlockLevelParents{},
		testnetGenesisMerkleRoot,
		&externalapi.DomainHash{},
		externalapi.NewDomainHashFromByteArray(muhash.EmptyMuHashHash.AsArray()),
		1788048000001,
		0x207fffff,
		0,
		0,
		0,
		big.NewInt(0),
		&externalapi.DomainHash{},
	),
	Transactions: []*externalapi.DomainTransaction{testnetGenesisCoinbaseTx},
}
