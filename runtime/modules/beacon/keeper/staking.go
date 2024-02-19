// SPDX-License-Identifier: MIT
//
// Copyright (c) 2024 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person
// obtaining a copy of this software and associated documentation
// files (the "Software"), to deal in the Software without
// restriction, including without limitation the rights to use,
// copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following
// conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// OTHER DEALINGS IN THE SOFTWARE.

package keeper

import (
	"context"
	"errors"

	sdkmath "cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	sdkcrypto "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkstaking "github.com/cosmos/cosmos-sdk/x/staking/types"

	consensusv1 "github.com/itsdevbear/bolaris/types/consensus/v1"
	enginev1 "github.com/itsdevbear/bolaris/types/engine/v1"
)

// delegate delegates the deposit to the validator.
func (k *Keeper) delegate(ctx context.Context, deposit *consensusv1.Deposit) (uint64, error) {
	validatorPK := &ed25519.PubKey{}
	err := validatorPK.Unmarshal(deposit.GetPubkey())
	if err != nil {
		return 0, err
	}
	amount := deposit.GetAmount()
	valConsAddr := sdk.GetConsAddress(validatorPK)
	validator, err := k.stakingKeeper.GetValidator(ctx, sdk.ValAddress(valConsAddr))
	if err != nil {
		if errors.Is(err, sdkstaking.ErrNoValidatorFound) {
			validator, err = k.createValidator(validatorPK, amount)
			return validator.DelegatorShares.BigInt().Uint64(), err
		}
		return 0, err
	}
	newShares, err := k.stakingKeeper.Delegate(
		ctx, sdk.AccAddress(valConsAddr),
		sdkmath.NewIntFromUint64(amount),
		sdkstaking.Unbonded, validator, true)
	return newShares.BigInt().Uint64(), err
}

// undelegate undelegates the validator.
func (k *Keeper) undelegate(_ context.Context, _ *enginev1.Withdrawal) (uint64, error) {
	// TODO: implement undelegate
	return 0, nil
}

// createValidator creates a new validator with the given public key and amount of tokens.
func (k *Keeper) createValidator(
	validatorPK sdkcrypto.PubKey,
	amount uint64) (sdkstaking.Validator, error) {
	stake := sdkmath.NewIntFromUint64(amount)
	valConsAddr := sdk.GetConsAddress(validatorPK)
	operator := sdk.ValAddress(valConsAddr).String()
	val, err := sdkstaking.NewValidator(
		operator, validatorPK,
		sdkstaking.Description{Moniker: validatorPK.String()})
	val.Tokens = stake
	val.DelegatorShares = sdkmath.LegacyNewDecFromInt(val.Tokens)
	return val, err
}

// ApplyValsetChanges applies the deposits and withdrawals
// to the validator set in the underlying staking module.
func (k *Keeper) ApplyValsetChanges(
	ctx context.Context,
	deposits []*consensusv1.Deposit,
	withdrawals []*enginev1.Withdrawal,
) error {
	for _, deposit := range deposits {
		_, err := k.delegate(ctx, deposit)
		if err != nil {
			return err
		}
	}
	for _, withdrawal := range withdrawals {
		_, err := k.undelegate(ctx, withdrawal)
		if err != nil {
			return err
		}
	}
	return nil
}
