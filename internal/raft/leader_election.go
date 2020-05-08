package raft

import "github.com/tomarrell/lbadd/internal/raft/message"

// StartElection enables a node in the cluster to start the election.
func StartElection(Server Node) {
	Server.State = CandidateState
	Server.PersistentState.CurrentTerm++

	var votes int

	for i := range Server.PersistentState.PeerIPs {
		// parallely request votes from all the other peers.
		go func(i int) {
			if Server.PersistentState.PeerIPs[i] != Server.PersistentState.SelfIP {
				// send a requestVotesRPC
				req := message.NewRequestVoteRequest(
					int32(Server.PersistentState.CurrentTerm),
					Server.PersistentState.SelfID,
					int32(len(Server.PersistentState.Log)),
					int32(Server.PersistentState.Log[len(Server.PersistentState.Log)-1].Term),
				)
				res := RequestVote(req)
				if res.VoteGranted {
					votes++
				}
			}
		}(i)
	}
}
