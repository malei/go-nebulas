// Copyright (C) 2017 go-nebulas authors
//
// This file is part of the go-nebulas library.
//
// the go-nebulas library is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// the go-nebulas library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with the go-nebulas library.  If not, see <http://www.gnu.org/licenses/>.
//

package pow

import (
	"errors"

	"github.com/nebulasio/go-nebulas/components/net"
	"github.com/nebulasio/go-nebulas/consensus"
	"github.com/nebulasio/go-nebulas/core"

	log "github.com/sirupsen/logrus"
)

// Errors in PoW Consensus
var (
	ErrInvalidDataType   = errors.New("invalid data type, should be *core.Block")
	ErrInvalidBlockNonce = errors.New("invalid block nonce")
)

/*
Pow implementation of Proof-of-Work consensus, designed to be a state machine.
The following is the state diagram:

@startuml
[*] --> Prepare
Prepare --> Mining : start mining
Mining --> Prepare : new block received
Mining --> Minted : found the nonce/block
Minted --> Prepare : broadcast the block, and start over
Prepare --> [*] : stop
Mining --> [*] : stop
Minted --> [*] : stop
@enduml
*/
type Pow struct {
	quitCh chan bool

	chain *core.BlockChain
	nm    net.Manager

	states            consensus.States
	currentState      consensus.State
	stateTransitionCh chan *stateTransitionArgs

	miningBlock      *core.Block
	newBlockReceived bool
}

type stateTransitionArgs struct {
	nextState consensus.State
	data      interface{}
}

// NewPow create Pow instance.
func NewPow(bc *core.BlockChain, nm net.Manager) *Pow {
	p := &Pow{
		chain:             bc,
		nm:                nm,
		quitCh:            make(chan bool, 5),
		stateTransitionCh: make(chan *stateTransitionArgs, 10),
	}

	p.states = consensus.States{
		Mining:  NewMiningState(p),
		Minted:  NewMintedState(p),
		Prepare: NewPrepareState(p),
		Stopped: NewStoppedState(p),
	}
	p.currentState = p.states[Prepare]

	return p
}

// Start start pow service.
func (p *Pow) Start() {
	// start state machine.
	go p.stateLoop()

	// start goroutine to process received message.
	go p.blockLoop()
}

// Stop stop pow service.
func (p *Pow) Stop() {
	// cleanup.
	p.quitCh <- true
	p.quitCh <- true
}

/*
Event handle events from Network or State.
The whole event process should be as the following:
1. dispatch to currentState to process.
2. if currentState does not captured it, consensus process it by default.
*/
func (p *Pow) Event(e consensus.Event) {
	captured, nextState := p.currentState.Event(e)
	if captured {
		if nextState != nil && p.currentState != nextState {
			p.Transit(nextState, nil)
		}
		return
	}

	// default procedure.
	et := e.EventType()
	switch et {
	case consensus.NewBlockEvent:
		block := e.Data().(*core.Block)
		log.WithFields(log.Fields{
			"block": block,
		}).Info("Pow.Event: handle BlockMessage.")

	default:
		log.WithFields(log.Fields{
			"eventType": e,
		}).Info("Pow.Event: handle this event.")
	}
}

// TransitByKey transit state by stateKey.
func (p *Pow) TransitByKey(stateKey string, data interface{}) {
	p.Transit(p.states[stateKey], data)
}

// Transit transit state.
func (p *Pow) Transit(nextState consensus.State, data interface{}) {
	if p.currentState == nextState {
		return
	}

	p.stateTransitionCh <- &stateTransitionArgs{nextState: nextState, data: data}
}

// VerifyBlock return nil if nonce is right, otherwise return error.
func (p *Pow) VerifyBlock(block *core.Block) error {
	if block == nil {
		log.WithFields(log.Fields{
			"func": "Pow.VerifyBlock",
			"err":  ErrInvalidDataType,
		}).Error("data is not valid block")
		return ErrInvalidDataType
	}

	ret := HashAndVerify(block)
	if ret == nil {
		log.WithFields(log.Fields{
			"func":  "Pow.VerifyBlock",
			"err":   ErrInvalidBlockNonce,
			"block": block,
		}).Error("invalid block nonce.")
		return ErrInvalidBlockNonce
	}

	return nil
}

func (p *Pow) stateLoop() {
	p.currentState.Enter(nil)

	for {
		select {
		case args := <-p.stateTransitionCh:
			nextState := args.nextState
			data := args.data

			if p.currentState == nextState {
				continue
			}

			p.currentState.Leave(data)
			p.currentState = nextState
			p.currentState.Enter(data)

		case <-p.quitCh:
			log.Info("quit Pow.loop.")
			return
		}
	}
}

func (p *Pow) blockLoop() {
	count := 0
	for {
		select {
		case block := <-p.chain.BlockPool().ReceivedBlockCh():
			count++
			log.Debugf("Pow.blockLoop: new block message received. Count=%d", count)
			p.newBlockReceived = true
			p.Event(consensus.NewBaseEvent(consensus.NewBlockEvent, block))
		case <-p.quitCh:
			// TODO: should provide base goroutine start/stop func to graceful stop them.
			/*
				for example,

				type Stopper struct {
					quitCh chan int // maybe int is better than bool, less confuss.
					count int q		// should use thread-safe int, eg. AtomicInt.
				}
				func NewStopper() *Stopper {
					s := &Stopper{quitCh: make(chan int，16), count : 0}
					return s
				}
				func (s *Stopper) CountMe() {
					s.count++
				}
				func (s *Stopper) QuitMe() {
					for i :=0 ; i<s.count; i++ {
						s.quitCh <- 0
					}
				}
			*/
			log.Info("Pow.blockLoop: quit.")
			return
		}
	}
}

func (p *Pow) resetMiningStatus() {
	p.miningBlock = nil
	p.newBlockReceived = false
}
