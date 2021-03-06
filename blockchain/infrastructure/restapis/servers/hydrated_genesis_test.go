package servers

import (
	"testing"
	"time"

	"github.com/deepvalue-network/software/blockchain/domain/genesis"
	"github.com/deepvalue-network/software/libs/hydro"
)

func TestHydrate_genesis_Success(t *testing.T) {
	// init:
	Init(time.Duration(time.Second), nil, "2006-01-02T15:04:05.000Z")

	// build a genesis:
	gen := genesis.CreateGenesisForTests()

	// execute:
	hydro.VerifyAdapterUsingJSForTests(internalHydroAdapter, gen, t)
}
