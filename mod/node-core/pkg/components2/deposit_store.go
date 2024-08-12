// SPDX-License-Identifier: BUSL-1.1
//
// Copyright (C) 2024, Berachain Foundation. All rights reserved.
// Use of this software is governed by the Business Source License included
// in the LICENSE file of this repository and at www.mariadb.com/bsl11.
//
// ANY USE OF THE LICENSED WORK IN VIOLATION OF THIS LICENSE WILL AUTOMATICALLY
// TERMINATE YOUR RIGHTS UNDER THIS LICENSE FOR THE CURRENT AND ALL OTHER
// VERSIONS OF THE LICENSED WORK.
//
// THIS LICENSE DOES NOT GRANT YOU ANY RIGHT IN ANY TRADEMARK OR LOGO OF
// LICENSOR OR ITS AFFILIATES (PROVIDED THAT YOU MAY USE A TRADEMARK OR LOGO OF
// LICENSOR AS EXPRESSLY REQUIRED BY THIS LICENSE).
//
// TO THE EXTENT PERMITTED BY APPLICABLE LAW, THE LICENSED WORK IS PROVIDED ON
// AN “AS IS” BASIS. LICENSOR HEREBY DISCLAIMS ALL WARRANTIES AND CONDITIONS,
// EXPRESS OR IMPLIED, INCLUDING (WITHOUT LIMITATION) WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, NON-INFRINGEMENT, AND
// TITLE.

package components

import (
	"cosmossdk.io/depinject"
	"cosmossdk.io/log"
	storev2 "cosmossdk.io/store/v2/db"
	"github.com/berachain/beacon-kit/mod/async/pkg/broker"
	"github.com/berachain/beacon-kit/mod/execution/pkg/deposit"
	"github.com/berachain/beacon-kit/mod/node-core/pkg/components/storage"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/common"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/math"
	depositstore "github.com/berachain/beacon-kit/mod/storage/pkg/deposit"
	"github.com/berachain/beacon-kit/mod/storage/pkg/manager"
	"github.com/berachain/beacon-kit/mod/storage/pkg/pruner"
	"github.com/cosmos/cosmos-sdk/client/flags"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/spf13/cast"
)

// DepositStoreInput is the input for the dep inject framework.
type DepositStoreInput struct {
	depinject.In
	AppOpts servertypes.AppOptions
}

// ProvideDepositStore is a function that provides the module to the
// application.
func ProvideDepositStore[
	DepositT Deposit[DepositT, ForkDataT, WithdrawalCredentialsT],
	ForkDataT any,
	WithdrawalCredentialsT WithdrawalCredentials,
](
	in DepositStoreInput,
) (*depositstore.KVStore[DepositT], error) {
	name := "deposits"
	dir := cast.ToString(in.AppOpts.Get(flags.FlagHome)) + "/data"
	kvp, err := storev2.NewDB(storev2.DBTypePebbleDB, name, dir, nil)
	if err != nil {
		return nil, err
	}

	return depositstore.NewStore[DepositT](storage.NewKVStoreProvider(kvp)), nil
}

// DepositPrunerInput is the input for the deposit pruner.
type DepositPrunerInput[
	BeaconBlockT any,
	BeaconBlockEventT Event[BeaconBlockT],
	DepositT any,
	DepositStoreT DepositStore[DepositT],
] struct {
	depinject.In
	BlockBroker  *broker.Broker[BeaconBlockEventT]
	ChainSpec    common.ChainSpec
	DepositStore DepositStoreT
	Logger       log.Logger
}

// ProvideDepositPruner provides a deposit pruner for the depinject framework.
func ProvideDepositPruner[
	BeaconBlockT interface {
		GetSlot() math.U64
		GetBody() BeaconBlockBodyT
	},
	BeaconBlockBodyT interface {
		GetDeposits() []DepositT
		GetExecutionPayload() ExecutionPayloadT
	},
	BeaconBlockEventT Event[BeaconBlockT],
	DepositT Deposit[DepositT, ForkDataT, WithdrawalCredentialsT],
	DepositStoreT DepositStore[DepositT],
	ExecutionPayloadT ExecutionPayload[
		ExecutionPayloadT,
		ExecutionPayloadHeaderT,
		WithdrawalsT,
	],
	ExecutionPayloadHeaderT any,
	ForkDataT any,
	WithdrawalsT any,
	WithdrawalCredentialsT WithdrawalCredentials,
](
	in DepositPrunerInput[BeaconBlockT, BeaconBlockEventT, DepositT, DepositStoreT],
) (pruner.Pruner[DepositStoreT], error) {
	subCh, err := in.BlockBroker.Subscribe()
	if err != nil {
		in.Logger.Error("failed to subscribe to block feed", "err", err)
		return nil, err
	}

	return pruner.NewPruner[
		BeaconBlockT,
		BeaconBlockEventT,
		DepositStoreT,
	](
		in.Logger.With("service", manager.DepositPrunerName),
		in.DepositStore,
		manager.DepositPrunerName,
		subCh,
		deposit.BuildPruneRangeFn[
			BeaconBlockT,
			BeaconBlockBodyT,
			BeaconBlockEventT,
			DepositT,
			ExecutionPayloadT,
			WithdrawalCredentialsT,
		](in.ChainSpec),
	), nil
}