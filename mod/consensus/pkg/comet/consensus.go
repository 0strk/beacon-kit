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

package comet

import (
	"context"

	appmodulev2 "cosmossdk.io/core/appmodule/v2"
	"cosmossdk.io/core/transaction"
	"cosmossdk.io/server/v2/cometbft/handlers"
	"github.com/berachain/beacon-kit/mod/consensus"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/math"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/transition"
	cmtabci "github.com/cometbft/cometbft/abci/types"
	"github.com/cosmos/gogoproto/proto"
	"github.com/sourcegraph/conc/iter"
)

var _ consensus.Consensus[
	transaction.Tx,
	appmodulev2.ValidatorUpdate,
] = (*Consensus[
	transaction.Tx,
	appmodulev2.ValidatorUpdate,
])(nil)

// NewConsensus returns a new consensus middleware.
func NewConsensus[
	T transaction.Tx, ValidatorUpdateT any,
](
	txCodec transaction.Codec[T],
	m Middleware,
) *Consensus[T, ValidatorUpdateT] {
	return &Consensus[T, ValidatorUpdateT]{
		txCodec:    txCodec,
		Middleware: m,
	}
}

// Consensus is used to decouple the Comet consensus engine from the Cosmos SDK.
// Right now, it is very coupled to the sdk base app and we will
// eventually fully decouple this.
type Consensus[T transaction.Tx, ValidatorUpdateT any] struct {
	txCodec transaction.Codec[T]
	Middleware
}

func (c *Consensus[T, ValidatorUpdateT]) InitGenesis(
	ctx context.Context,
	bz []byte,
) ([]ValidatorUpdateT, error) {
	updates, err := c.Middleware.InitGenesis(ctx, bz)
	if err != nil {
		return nil, err
	}
	// Convert updates into the Cosmos SDK format.
	return iter.MapErr[
		*transition.ValidatorUpdate, ValidatorUpdateT,
	](updates, convertValidatorUpdate)
}

// TODO: Decouple Comet Types
func (c *Consensus[T, ValidatorUpdateT]) Prepare(
	ctx context.Context,
	am handlers.AppManager[T],
	txs []T,
	req proto.Message,
) ([]T, error) {
	abciReq, ok := req.(*cmtabci.PrepareProposalRequest)
	if !ok {
		return nil, ErrInvalidRequestType
	}
	slot := math.Slot(abciReq.Height)
	blkBz, sidecarsBz, err := c.Middleware.PrepareProposal(ctx, slot)
	if err != nil {
		return nil, err
	}
	blkTx, err := c.txCodec.Decode(blkBz)
	if err != nil {
		return nil, err
	}
	sidecarsTx, err := c.txCodec.Decode(sidecarsBz)
	if err != nil {
		return nil, err
	}

	return []T{blkTx, sidecarsTx}, nil
}

// TODO: Decouple Comet Types
func (c *Consensus[T, ValidatorUpdateT]) Process(
	ctx context.Context,
	_ handlers.AppManager[T],
	txs []T,
	req proto.Message,
) error {
	abciReq, ok := req.(*cmtabci.ProcessProposalRequest)
	if !ok {
		return ErrInvalidRequestType
	}
	return c.Middleware.ProcessProposal(ctx, abciReq)
}

// TODO: Decouple Comet Types
func (c *Consensus[T, ValidatorUpdateT]) PreBlock(
	ctx context.Context, msg proto.Message,
) error {
	req, ok := msg.(*cmtabci.FinalizeBlockRequest)
	if !ok {
		return ErrInvalidRequestType
	}
	return c.Middleware.PreBlock(ctx, req)
}

func (c *Consensus[T, ValidatorUpdateT]) EndBlock(
	ctx context.Context,
) ([]ValidatorUpdateT, error) {
	updates, err := c.Middleware.EndBlock(ctx)
	if err != nil {
		return nil, err
	}
	return iter.MapErr[
		*transition.ValidatorUpdate, ValidatorUpdateT,
	](updates, convertValidatorUpdate)
}