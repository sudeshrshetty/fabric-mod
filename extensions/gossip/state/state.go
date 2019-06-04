/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package state

import (
	"github.com/hyperledger/fabric/extensions/gossip/api"
	common2 "github.com/hyperledger/fabric/gossip/common"
	"github.com/hyperledger/fabric/gossip/discovery"
	"github.com/hyperledger/fabric/gossip/protoext"
	"github.com/hyperledger/fabric/gossip/util"
	"github.com/hyperledger/fabric/protos/common"
	proto "github.com/hyperledger/fabric/protos/gossip"
)

//GossipStateProviderExtension extends GossipStateProvider features
type GossipStateProviderExtension interface {

	//HandleStateRequest can used to extend given request handle
	HandleStateRequest(func(msg protoext.ReceivedMessage)) func(msg protoext.ReceivedMessage)

	//AntiEntropy can be used to extend Anti Entropy features
	AntiEntropy(func()) func()

	//Predicate can used to override existing predicate to filter peers to be asked for blocks
	Predicate(func(peer discovery.NetworkMember) bool) func(peer discovery.NetworkMember) bool

	//AddPayload can used to extend given add payload handle
	AddPayload(func(payload *proto.Payload, blockingMode bool) error) func(payload *proto.Payload, blockingMode bool) error

	//StoreBlock  can used to extend given store block handle
	StoreBlock(func(block *common.Block, pvtData util.PvtDataCollections) error) func(block *common.Block, pvtData util.PvtDataCollections) error
}

// GossipServiceMediator aggregated adapter interface to compound basic mediator services
// required by state transfer into single struct
type GossipServiceMediator interface {
	// VerifyBlock returns nil if the block is properly signed, and the claimed seqNum is the
	// sequence number that the block's header contains.
	// else returns error
	VerifyBlock(chainID common2.ChainID, seqNum uint64, signedBlock []byte) error

	// PeersOfChannel returns the NetworkMembers considered alive
	// and also subscribed to the channel given
	PeersOfChannel(common2.ChainID) []discovery.NetworkMember

	// Gossip sends a message to other peers to the network
	Gossip(msg *proto.GossipMessage)
}

//AddBlockHandler handles state update in gossip
func AddBlockHandler(cid string, publisher api.BlockPublisher) {
	//do nothing
}

//NewGossipStateProviderExtension returns new GossipStateProvider Extension implementation
func NewGossipStateProviderExtension(chainID string, mediator GossipServiceMediator) GossipStateProviderExtension {
	return &gossipStateProviderExtension{}
}

type gossipStateProviderExtension struct {
}

func (s *gossipStateProviderExtension) HandleStateRequest(handle func(msg protoext.ReceivedMessage)) func(msg protoext.ReceivedMessage) {
	return handle
}

func (s *gossipStateProviderExtension) AntiEntropy(handle func()) func() {
	return handle
}

func (s *gossipStateProviderExtension) Predicate(handle func(peer discovery.NetworkMember) bool) func(peer discovery.NetworkMember) bool {
	return handle
}

func (s *gossipStateProviderExtension) AddPayload(handle func(payload *proto.Payload, blockingMode bool) error) func(payload *proto.Payload, blockingMode bool) error {
	return handle
}

func (s *gossipStateProviderExtension) StoreBlock(handle func(block *common.Block, pvtData util.PvtDataCollections) error) func(block *common.Block, pvtData util.PvtDataCollections) error {
	return handle
}
