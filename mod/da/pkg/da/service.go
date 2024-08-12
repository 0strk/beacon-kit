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

package da

import (
	"context"

	asynctypes "github.com/berachain/beacon-kit/mod/async/pkg/types"
	"github.com/berachain/beacon-kit/mod/log"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/messages"
)

type Service[
	AvailabilityStoreT AvailabilityStore[BeaconBlockBodyT, BlobSidecarsT],
	BeaconBlockBodyT any,
	BlobSidecarsT BlobSidecar,

	ExecutionPayloadT any,
] struct {
	avs AvailabilityStoreT
	bp  BlobProcessor[
		AvailabilityStoreT, BeaconBlockBodyT,
		BlobSidecarsT, ExecutionPayloadT,
	]
	dispatcher             asynctypes.MessageDispatcher
	logger                 log.Logger[any]
	processSidecarRequests chan asynctypes.Message[BlobSidecarsT]
	verifySidecarRequests  chan asynctypes.Message[BlobSidecarsT]
}

// NewService returns a new DA service.
func NewService[
	AvailabilityStoreT AvailabilityStore[
		BeaconBlockBodyT, BlobSidecarsT,
	],
	BeaconBlockBodyT any,
	BlobSidecarsT BlobSidecar,

	ExecutionPayloadT any,
](
	avs AvailabilityStoreT,
	bp BlobProcessor[
		AvailabilityStoreT, BeaconBlockBodyT,
		BlobSidecarsT, ExecutionPayloadT,
	],
	dispatcher asynctypes.MessageDispatcher,
	logger log.Logger[any],
) *Service[
	AvailabilityStoreT, BeaconBlockBodyT,
	BlobSidecarsT, ExecutionPayloadT,
] {
	return &Service[
		AvailabilityStoreT, BeaconBlockBodyT,
		BlobSidecarsT, ExecutionPayloadT,
	]{
		avs:                    avs,
		bp:                     bp,
		dispatcher:             dispatcher,
		logger:                 logger,
		processSidecarRequests: make(chan asynctypes.Message[BlobSidecarsT]),
		verifySidecarRequests:  make(chan asynctypes.Message[BlobSidecarsT]),
	}
}

// Name returns the name of the service.
func (s *Service[_, _, _, _]) Name() string {
	return "da"
}

// Start registers this service as the recipient of ProcessSidecars and
// VerifySidecars messages, and begins listening for these requests.
func (s *Service[_, _, BlobSidecarsT, _]) Start(ctx context.Context) error {
	var err error
	// register as recipient of ProcessSidecars messages.
	if err = s.dispatcher.RegisterMsgReceiver(
		messages.ProcessSidecars, s.processSidecarRequests,
	); err != nil {
		return err
	}

	// register as recipient of VerifySidecars messages.
	if err = s.dispatcher.RegisterMsgReceiver(
		messages.VerifySidecars, s.verifySidecarRequests,
	); err != nil {
		return err
	}

	// start a goroutine to listen for requests and handle accordingly
	go s.start(ctx)
	return nil
}

// start starts listens for ProcessSidecars and VerifySidecars messages and
// handles them accordingly.
func (s *Service[_, _, BlobSidecarsT, _]) start(
	ctx context.Context,
) {
	for {
		select {
		case <-ctx.Done():
			return
		case msg := <-s.processSidecarRequests:
			s.handleBlobSidecarsProcessRequest(msg)
		case msg := <-s.verifySidecarRequests:
			s.handleSidecarsVerifyRequest(msg)
		}
	}
}

/* -------------------------------------------------------------------------- */
/*                               Message Handlers                             */
/* -------------------------------------------------------------------------- */

// handleBlobSidecarsProcessRequest handles the BlobSidecarsProcessRequest
// event.
// It processes the sidecars and publishes a BlobSidecarsProcessed event.
func (s *Service[_, _, BlobSidecarsT, _]) handleBlobSidecarsProcessRequest(
	msg asynctypes.Message[BlobSidecarsT],
) {
	var err error
	err = s.processSidecars(msg.Context(), msg.Data())
	if err != nil {
		s.logger.Error(
			"Failed to process blob sidecars",
			"error",
			err,
		)
	}

	// dispatch a response to acknowledge the request.
	if err = s.dispatcher.SendResponse(
		asynctypes.NewMessage(
			msg.Context(),
			messages.ProcessSidecars,
			msg.Data(),
			nil,
		),
	); err != nil {
		s.logger.Error("failed to respond", "err", err)
	}
}

// handleSidecarsVerifyRequest handles the SidecarsVerifyRequest event.
// It verifies the sidecars and publishes a SidecarsVerified event.
func (s *Service[_, _, BlobSidecarsT, _]) handleSidecarsVerifyRequest(
	msg asynctypes.Message[BlobSidecarsT],
) {
	var err error
	// verify the sidecars.
	if err = s.verifySidecars(msg.Data()); err != nil {
		s.logger.Error(
			"Failed to receive blob sidecars",
			"error",
			err,
		)
	}

	// dispatch a response to acknowledge the request.
	if err = s.dispatcher.SendResponse(
		asynctypes.NewMessage(
			msg.Context(),
			messages.VerifySidecars,
			msg.Data(),
			nil,
		),
	); err != nil {
		s.logger.Error("failed to respond", "err", err)
	}
}

/* -------------------------------------------------------------------------- */
/*                                   helpers                                  */
/* -------------------------------------------------------------------------- */

// ProcessSidecars processes the blob sidecars.
func (s *Service[_, _, BlobSidecarsT, _]) processSidecars(
	_ context.Context,
	sidecars BlobSidecarsT,
) error {
	// startTime := time.Now()
	// defer s.metrics.measureBlobProcessingDuration(startTime)
	return s.bp.ProcessSidecars(
		s.avs,
		sidecars,
	)
}

// VerifyIncomingBlobs receives blobs from the network and processes them.
func (s *Service[_, _, BlobSidecarsT, _]) verifySidecars(
	sidecars BlobSidecarsT,
) error {
	// If there are no blobs to verify, return early.
	if sidecars.IsNil() || sidecars.Len() == 0 {
		return nil
	}

	s.logger.Info(
		"Received incoming blob sidecars",
	)

	// Verify the blobs and ensure they match the local state.
	if err := s.bp.VerifySidecars(sidecars); err != nil {
		s.logger.Error(
			"rejecting incoming blob sidecars",
			"reason", err,
		)
		return err
	}

	s.logger.Info(
		"Blob sidecars verification succeeded - accepting incoming blob sidecars",
		"num_blobs",
		sidecars.Len(),
	)

	return nil
}
