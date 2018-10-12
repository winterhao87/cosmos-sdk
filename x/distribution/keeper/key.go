package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// keys/key-prefixes
var (
	FeePoolKey               = []byte{0x00} // key for global distribution state
	ValidatorDistInfoKey     = []byte{0x01} // prefix for each key to a validator distribution
	DelegationDistInfoKey    = []byte{0x02} // prefix for each key to a delegation distribution
	DelegatorWithdrawInfoKey = []byte{0x03} // prefix for each key to a delegator withdraw info

	// transient
	ProposerKey          = []byte{0x00} // key for storing the proposer operator address
	SumPrecommitPowerKey = []byte{0x01} // key for storing the power of the precommit validators
)

// nolint
const (
	ParamStoreKeyCommunityTax = "distr/community-tax"
)

// gets the key for the validator distribution info from address
// VALUE: distribution/types.ValidatorDistInfo
func GetValidatorDistInfoKey(operatorAddr sdk.ValAddress) []byte {
	return append(ValidatorDistInfoKey, operatorAddr.Bytes()...)
}

// gets the key for delegator distribution for a validator
// VALUE: distribution/types.DelegationDistInfo
func GetDelegationDistInfoKey(delAddr sdk.AccAddress, valAddr sdk.ValAddress) []byte {
	return append(GetDelegationDistInfosKey(delAddr), valAddr.Bytes()...)
}

// gets the prefix for a delegator's distributions across all validators
func GetDelegationDistInfosKey(delAddr sdk.AccAddress) []byte {
	return append(DelegationDistInfoKey, delAddr.Bytes()...)
}

// gets the prefix for a delegator's withdraw info
func GetDelegatorWithdrawAddrKey(delAddr sdk.AccAddress) []byte {
	return append(DelegatorWithdrawInfoKey, delAddr.Bytes()...)
}