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

package api_test

import (
	"github.com/berachain/beacon-kit/testing/e2e/config"
	"github.com/berachain/beacon-kit/testing/e2e/suite"
)

// BeaconAPISuite is a suite of beacon node-api tests with full simulation of a
// beacon-kit network.
type BeaconAPISuite struct {
	suite.KurtosisE2ESuite
}

// TestBeaconAPISuite tests that the api test suite is setup correctly with a
// working beacon node-api client.
func (s *BeaconAPISuite) TestBeaconAPISuite() {
	executionBlockNum := uint64(5)

	// Wait for execution block 5.
	err := s.WaitForFinalizedBlockNumber(executionBlockNum)
	s.Require().NoError(err)

	// Get the consensus client.
	client := s.ConsensusClients()[config.DefaultClient]
	s.Require().NotNil(client)

	// Get the latest beacon block.
	slot, htr, err := client.GetLatestBeaconBlock(s.Ctx())
	s.Require().NoError(err)
	s.Require().GreaterOrEqual(slot, executionBlockNum)
	s.Require().NotEmpty(htr)
}