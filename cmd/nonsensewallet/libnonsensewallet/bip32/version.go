package bip32

import "github.com/pkg/errors"

// BitcoinMainnetPrivate is the version that is used for
// bitcoin mainnet bip32 private extended keys.
// Ecnodes to xprv in base58.
var BitcoinMainnetPrivate = [4]byte{
	0x04,
	0x88,
	0xad,
	0xe4,
}

// BitcoinMainnetPublic is the version that is used for
// bitcoin mainnet bip32 public extended keys.
// Ecnodes to xpub in base58.
var BitcoinMainnetPublic = [4]byte{
	0x04,
	0x88,
	0xb2,
	0x1e,
}

// NonsenseMainnetPrivate is the version that is used for
// nonsense mainnet bip32 private extended keys.
// Ecnodes to xprv in base58.
var NonsenseMainnetPrivate = [4]byte{
	0x03,
	0x8f,
	0x2e,
	0xf4,
}

// NonsenseMainnetPublic is the version that is used for
// nonsense mainnet bip32 public extended keys.
// Ecnodes to kpub in base58.
var NonsenseMainnetPublic = [4]byte{
	0x03,
	0x8f,
	0x33,
	0x2e,
}

// NonsenseTestnetPrivate is the version that is used for
// nonsense testnet bip32 public extended keys.
// Ecnodes to ktrv in base58.
var NonsenseTestnetPrivate = [4]byte{
	0x03,
	0x90,
	0x9e,
	0x07,
}

// NonsenseTestnetPublic is the version that is used for
// nonsense testnet bip32 public extended keys.
// Ecnodes to ktub in base58.
var NonsenseTestnetPublic = [4]byte{
	0x03,
	0x90,
	0xa2,
	0x41,
}

// NonsenseDevnetPrivate is the version that is used for
// nonsense devnet bip32 public extended keys.
// Ecnodes to kdrv in base58.
var NonsenseDevnetPrivate = [4]byte{
	0x03,
	0x8b,
	0x3d,
	0x80,
}

// NonsenseDevnetPublic is the version that is used for
// nonsense devnet bip32 public extended keys.
// Ecnodes to xdub in base58.
var NonsenseDevnetPublic = [4]byte{
	0x03,
	0x8b,
	0x41,
	0xba,
}

// NonsenseSimnetPrivate is the version that is used for
// nonsense simnet bip32 public extended keys.
// Ecnodes to ksrv in base58.
var NonsenseSimnetPrivate = [4]byte{
	0x03,
	0x90,
	0x42,
	0x42,
}

// NonsenseSimnetPublic is the version that is used for
// nonsense simnet bip32 public extended keys.
// Ecnodes to xsub in base58.
var NonsenseSimnetPublic = [4]byte{
	0x03,
	0x90,
	0x46,
	0x7d,
}

func toPublicVersion(version [4]byte) ([4]byte, error) {
	switch version {
	case BitcoinMainnetPrivate:
		return BitcoinMainnetPublic, nil
	case NonsenseMainnetPrivate:
		return NonsenseMainnetPublic, nil
	case NonsenseTestnetPrivate:
		return NonsenseTestnetPublic, nil
	case NonsenseDevnetPrivate:
		return NonsenseDevnetPublic, nil
	case NonsenseSimnetPrivate:
		return NonsenseSimnetPublic, nil
	}

	return [4]byte{}, errors.Errorf("unknown version %x", version)
}

func isPrivateVersion(version [4]byte) bool {
	switch version {
	case BitcoinMainnetPrivate:
		return true
	case NonsenseMainnetPrivate:
		return true
	case NonsenseTestnetPrivate:
		return true
	case NonsenseDevnetPrivate:
		return true
	case NonsenseSimnetPrivate:
		return true
	}

	return false
}
