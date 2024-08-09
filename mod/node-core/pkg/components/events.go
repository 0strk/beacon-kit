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
	"github.com/berachain/beacon-kit/mod/async/pkg/messaging"
	"github.com/berachain/beacon-kit/mod/async/pkg/server"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/messages"
)

// EventServerInput is the input for the event server.
type EventServerInput struct {
	depinject.In
	BeaconBlockFinalizedPublisher *BeaconBlockFinalizedPublisher
}

// ProvideEventServer provides an event server.
func ProvideEventServer(in EventServerInput) *EventServer {
	es := server.NewEventServer()
	es.RegisterPublisher(
		in.BeaconBlockFinalizedPublisher.EventID(),
		in.BeaconBlockFinalizedPublisher,
	)
	return es
}

// ProvideBeaconBlockFinalizedPublisher provides a publisher for beacon block
// finalized events.
func ProvideBeaconBlockFinalizedPublisher() *BeaconBlockFinalizedPublisher {
	return messaging.NewPublisher[*FinalizedBlockEvent](
		messages.BeaconBlockFinalizedEvent,
	)
}

// EventServerComponents returns all the depinject providers for the event
// server.
func EventServerComponents() []any {
	return []any{
		ProvideEventServer,
		ProvideBeaconBlockFinalizedPublisher,
	}
}
