package state

import (
	"bytes"
	"math/big"

	"github.com/XinFinOrg/XDPoSChain/common"
	lru "github.com/hashicorp/golang-lru"

)

var (
	SlotXDC21Issuer = map[string]uint64{
		"minCap":      0,
		"tokens":      1,
		"tokensState": 2,
	}
	SlotXDC21Token = map[string]uint64{
		"balances": 0,
		"minFee":   1,
		"issuer":   2,
	}
	transferFuncHex     = common.Hex2Bytes("0xa9059cbb")
	transferFromFuncHex = common.Hex2Bytes("0x23b872dd")
	cache, _            = lru.NewARC(128)
)

func GetXDC21FeeCapacityFromStateWithCache(trieRoot common.Hash, statedb *StateDB) map[common.Address]*big.Int {
	if statedb == nil {
		return map[common.Address]*big.Int{}
	}
	data, _ := cache.Get(trieRoot)
	var info map[common.Address]*big.Int
	if data != nil {
		info = data.(map[common.Address]*big.Int)
	} else {
		info = GetXDC21FeeCapacityFromState(statedb)
	}
	cache.Add(trieRoot, info)
	tokensFee := map[common.Address]*big.Int{}
	for key, value := range info {
		tokensFee[key] = big.NewInt(0).SetBytes(value.Bytes())
	}
	return tokensFee
}
func GetXDC21FeeCapacityFromState(statedb *StateDB) map[common.Address]*big.Int {
	if statedb == nil {
		return map[common.Address]*big.Int{}
	}
	tokensCapacity := map[common.Address]*big.Int{}
	slotTokens := SlotXDC21Issuer["tokens"]
	slotTokensHash := common.BigToHash(new(big.Int).SetUint64(slotTokens))
	slotTokensState := SlotXDC21Issuer["tokensState"]
	tokenCount := statedb.GetState(common.XDC21IssuerSMC, slotTokensHash).Big().Uint64()
	for i := uint64(0); i < tokenCount; i++ {
		key := GetLocDynamicArrAtElement(slotTokensHash, i, 1)
		value := statedb.GetState(common.XDC21IssuerSMC, key)
		if !common.EmptyHash(value) {
			token := common.BytesToAddress(value.Bytes())
			balanceKey := GetLocMappingAtKey(token.Hash(), slotTokensState)
			balanceHash := statedb.GetState(common.XDC21IssuerSMC, common.BigToHash(balanceKey))
			tokensCapacity[common.BytesToAddress(token.Bytes())] = balanceHash.Big()
		}
	}
	return tokensCapacity
}

func PayFeeWithXDC21TxFail(statedb *StateDB, from common.Address, token common.Address) {
	if statedb == nil {
		return
	}
	slotBalanceXdc21 := SlotXDC21Token["balances"]
	balanceKey := GetLocMappingAtKey(from.Hash(), slotBalanceXdc21)
	balanceHash := statedb.GetState(token, common.BigToHash(balanceKey))
	if !common.EmptyHash(balanceHash) {
		balance := balanceHash.Big()
		feeUsed := big.NewInt(0)
		if balance.Cmp(feeUsed) <= 0 {
			return
		}
		issuerTokenKey := GetLocSimpleVariable(SlotXDC21Token["issuer"])
		if common.EmptyHash(issuerTokenKey) {
			return
		}
		issuerAddr := common.BytesToAddress(statedb.GetState(token, issuerTokenKey).Bytes())
		feeTokenKey := GetLocSimpleVariable(SlotXDC21Token["minFee"])
		feeHash := statedb.GetState(token, feeTokenKey)
		fee := feeHash.Big()
		if balance.Cmp(fee) < 0 {
			feeUsed = balance
		} else {
			feeUsed = fee
		}
		balance = balance.Sub(balance, feeUsed)
		statedb.SetState(token, common.BigToHash(balanceKey), common.BigToHash(balance))

		issuerBalanceKey := GetLocMappingAtKey(issuerAddr.Hash(), slotBalanceXdc21)
		issuerBalanceHash := statedb.GetState(token, common.BigToHash(issuerBalanceKey))
		issuerBalance := issuerBalanceHash.Big()
		issuerBalance = issuerBalance.Add(issuerBalance, feeUsed)
		statedb.SetState(token, common.BigToHash(issuerBalanceKey), common.BigToHash(issuerBalance))
	}
}

func ValidateXDC21Tx(statedb *StateDB, from common.Address, token common.Address, data []byte) bool {
	if data == nil || statedb == nil {
		return false
	}
	slotBalanceXdc21 := SlotXDC21Token["balances"]
	balanceKey := GetLocMappingAtKey(from.Hash(), slotBalanceXdc21)
	balanceHash := statedb.GetState(token, common.BigToHash(balanceKey))

	if !common.EmptyHash(balanceHash) {
		balance := balanceHash.Big()
		minFeeTokenKey := GetLocSimpleVariable(SlotXDC21Token["minFee"])
		minFeeHash := statedb.GetState(token, minFeeTokenKey)
		requiredMinBalance := minFeeHash.Big()
		funcHex := data[:4]
		value := big.NewInt(0)
		if bytes.Equal(funcHex, transferFuncHex) && len(data) == 68 {
			value = common.BytesToHash(data[36:]).Big()
		} else {
			if bytes.Equal(funcHex, transferFromFuncHex) && len(data) == 80 {
				value = common.BytesToHash(data[68:]).Big()
			}
		}
		requiredMinBalance = requiredMinBalance.Add(requiredMinBalance, value)
		if balance.Cmp(requiredMinBalance) < 0 {
			return false
		} else {
			return true
		}
	} else {
		// we both accept tx with balance = 0 and fee = 0
		minFeeTokenKey := GetLocSimpleVariable(SlotXDC21Token["minFee"])
		if !common.EmptyHash(minFeeTokenKey) {
			return true
		}
	}

	return false
}

func UpdateXDC21Fee(statedb *StateDB, newBalance map[common.Address]*big.Int, totalFeeUsed *big.Int) {
	if statedb == nil || len(newBalance) == 0 {
		return
	}
	slotTokensState := SlotXDC21Issuer["tokensState"]
	for token, value := range newBalance {
		balanceKey := GetLocMappingAtKey(token.Hash(), slotTokensState)
		statedb.SetState(common.XDC21IssuerSMC, common.BigToHash(balanceKey), common.BigToHash(value))
	}
	statedb.SubBalance(common.XDC21IssuerSMC, totalFeeUsed)
}
