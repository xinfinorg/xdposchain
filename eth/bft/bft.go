package bft

import (
	"github.com/XinFinOrg/XDPoSChain/consensus/XDPoS"
	"github.com/XinFinOrg/XDPoSChain/consensus/XDPoS/utils"
)

type collectVoteFn func(utils.VoteType) error

type collectTimeoutFn func(utils.TimeoutType) error

type updateRoundFn func(utils.SyncInfoType) error

type broadcastVoteFn func(utils.VoteType)

type broadcastTimeoutFn func(utils.TimeoutType)

type broadcastSyncInfoFn func(utils.SyncInfoType)

type broadcastTCFn func(utils.TCType)

type BFT struct {
	messageBus chan interface{}
	quit       chan struct{}
	engine     ConsensusFns
	broadcast  BroadcastFns
}

type ConsensusFns struct {
	collectVote    collectVoteFn
	collectTimeout collectTimeoutFn
	updateRound    updateRoundFn
}

type BroadcastFns struct {
	Vote     broadcastVoteFn
	Timeout  broadcastTimeoutFn
	SyncInfo broadcastSyncInfoFn
	TC       broadcastTCFn
}

func New(engine *XDPoS.XDPoS, broadcasts BroadcastFns) *BFT {
	consensus := ConsensusFns{
		collectVote:    engine.CollectVote,
		collectTimeout: engine.CollectTimeout,
		updateRound:    engine.UpdateRound,
	}
	return &BFT{
		messageBus: make(chan interface{}),
		engine:     consensus,
		broadcast:  broadcasts,
	}
}

func (b *BFT) Vote(vote interface{}) {
	b.engine.collectVote(vote)
}

func (b *BFT) Timeout(timeout interface{}) {
	b.engine.collectTimeout(timeout)
}

func (b *BFT) SyncInfo(syncInfo interface{}) {
	b.engine.updateRound(syncInfo)
}
func (b *BFT) Start() {
	go b.loop()
}

func (b *BFT) loop() {

	for {
		select {
		case <-b.quit:
			return
		case obj := <-b.messageBus:
			switch v := obj.(type) {
			case utils.VoteType:
				b.broadcast.Vote(v)
			case utils.TimeoutType:
				b.broadcast.Timeout(v)
			case utils.SyncInfoType:
				b.broadcast.SyncInfo(v)
			case utils.TCType:
				b.broadcast.TC(v)
			default:

			}
		}
		//TODO: stop routine
	}
}
