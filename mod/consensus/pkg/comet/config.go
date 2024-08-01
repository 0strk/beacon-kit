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

package cometbft

import (
	"time"

	"github.com/cometbft/cometbft/config"
)

type Config struct {
	CometConfig *config.Config `mapstructure:",squash"`

	// Filepaths
	NodeKeyFile            string `mapstructure:"node_key_file"`
	PrivValidatorKeyFile   string `mapstructure:"priv_validator_key_file"`
	PrivValidatorStateFile string `mapstructure:"priv_validator_state_file"`
}

func DefaultConfig() *Config {
	// the SDK is very opinionated about these values, so we override them
	// if they aren't already set
	//nolint:mnd // 5 seconds
	cfg := config.DefaultConfig()
	cfg.Consensus.TimeoutCommit = 5 * time.Second
	cfg.RPC.PprofListenAddress = "localhost:6060"
	return &Config{
		CometConfig:            cfg,
		NodeKeyFile:            "config/node_key.json",
		PrivValidatorKeyFile:   "config/priv_validator_key.json",
		PrivValidatorStateFile: "config/priv_validator_state.json",
	}
}
