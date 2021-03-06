package servers

import (
	"testing"
	"time"

	"github.com/deepvalue-network/software/blockchain/domain/chains/peers"
	"github.com/deepvalue-network/software/libs/hydro"
)

func TestHydrate_peer_Success(t *testing.T) {
	// init:
	Init(time.Duration(time.Second), nil, "2006-01-02T15:04:05.000Z")

	// build a peer:
	peer := peers.CreatePeerForTests()

	// execute:
	hydro.VerifyAdapterUsingJSForTests(internalHydroAdapter, peer, t)
}
