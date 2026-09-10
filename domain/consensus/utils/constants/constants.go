package constants

import "math"

const (
	// BlockVersionKHashV1 represents the block version before the HF
	BlockVersionKHashV1 uint16 = 1

	// BlockVersionKHashV2 represents the block version after the HF
	BlockVersionKHashV2 uint16 = 2

	// MaxTransactionVersion is the current latest supported transaction version.
	MaxTransactionVersion uint16 = 0

	// MaxScriptPublicKeyVersion is the current latest supported public key script version.
	MaxScriptPublicKeyVersion uint16 = 0

	// SompiPerNonsense is the number of sompi in one nonsense (1 NNN).
	SompiPerNonsense = 100_000_000

	// MaxSompi is the maximum transaction amount allowed in sompi.
	MaxSompi = uint64(2_000_000_000 * SompiPerNonsense)

	// Transaction mass parameters are consensus constants on every Nonsense
	// network. They are exported so legacy V2 transactions can be repaired
	// while being decoded from the first daemon's incomplete wire format.
	MassPerTxByte           = uint64(1)
	MassPerScriptPubKeyByte = uint64(10)
	MassPerSigOp            = uint64(1000)

	// MaxTxInSequenceNum is the maximum sequence number the sequence field
	// of a transaction input can be.
	MaxTxInSequenceNum uint64 = math.MaxUint64

	// SequenceLockTimeDisabled is a flag that if set on a transaction
	// input's sequence number, the sequence number will not be interpreted
	// as a relative locktime.
	SequenceLockTimeDisabled uint64 = 1 << 63

	// SequenceLockTimeMask is a mask that extracts the relative locktime
	// when masked against the transaction input sequence number.
	SequenceLockTimeMask uint64 = 0x00000000ffffffff

	// LockTimeThreshold is the number below which a lock time is
	// interpreted to be a DAA score.
	LockTimeThreshold = 5e11 // Tue Nov 5 00:53:20 1985 UTC

	// UnacceptedDAAScore is used to for UTXOEntries that were created by transactions in the mempool, or otherwise
	// not-yet-accepted transactions.
	UnacceptedDAAScore = math.MaxUint64
)
