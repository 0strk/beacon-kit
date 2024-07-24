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

package types

import (
	gethprimitives "github.com/berachain/beacon-kit/mod/geth-primitives"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/common"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/math"
	"github.com/karalabe/ssz"
)

type Eth1Data struct {
	// DepositRoot is the root of the deposit tree.
	DepositRoot common.Root `json:"depositRoot"`
	// DepositCount is the number of deposits in the deposit tree.
	DepositCount math.U64 `json:"depositCount"`
	// BlockHash is the hash of the block corresponding to the Eth1Data.
	BlockHash gethprimitives.ExecutionHash `json:"blockHash"`
}

// New creates a new Eth1Data.
func (e *Eth1Data) New(
	depositRoot common.Root,
	depositCount math.U64,
	blockHash gethprimitives.ExecutionHash,
) *Eth1Data {
	e = &Eth1Data{
		DepositRoot:  depositRoot,
		DepositCount: depositCount,
		BlockHash:    blockHash,
	}
	return e
}

// SizeSSZ returns the size of the Eth1Data object in SSZ encoding.
func (*Eth1Data) SizeSSZ() uint32 {
	//nolint:mnd // 32 + 8 + 32
	return 72
}

// DefineSSZ defines the SSZ encoding for the Eth1Data object.
func (e *Eth1Data) DefineSSZ(codec *ssz.Codec) {
	ssz.DefineStaticBytes(codec, &e.DepositRoot)
	ssz.DefineUint64(codec, &e.DepositCount)
	ssz.DefineStaticBytes(codec, &e.BlockHash)
}

// HashTreeRoot computes the SSZ hash tree root of the Eth1Data object.
func (e *Eth1Data) HashTreeRoot() ([32]byte, error) {
	return ssz.HashSequential(e), nil
}

// MarshalSSZ marshals the Eth1Data object to SSZ format.
func (e *Eth1Data) MarshalSSZ() ([]byte, error) {
	buf := make([]byte, e.SizeSSZ())
	return buf, ssz.EncodeToBytes(buf, e)
}

// MarshalSSZTo marshals the Eth1Data object into a pre-allocated byte slice.
func (e *Eth1Data) MarshalSSZTo(dst []byte) ([]byte, error) {
	return dst, ssz.EncodeToBytes(dst, e)
}

// UnmarshalSSZ unmarshals the Eth1Data object from SSZ format.
func (e *Eth1Data) UnmarshalSSZ(buf []byte) error {
	return ssz.DecodeFromBytes(buf, e)
}

// GetDepositCount returns the deposit count.
func (e *Eth1Data) GetDepositCount() math.U64 {
	return e.DepositCount
}