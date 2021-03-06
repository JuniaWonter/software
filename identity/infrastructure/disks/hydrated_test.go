package disks

import (
	"testing"

	"github.com/deepvalue-network/software/identity/domain/accesses"
	"github.com/deepvalue-network/software/identity/domain/accesses/access"
	"github.com/deepvalue-network/software/identity/domain/users"
	"github.com/deepvalue-network/software/libs/hydro"
)

func TestHydrate_access_Success(t *testing.T) {
	// build an access:
	access := access.CreateAccessForTests(4096)

	// execute:
	hydro.VerifyAdapterUsingJSForTests(hydroAdapter, access, t)
}

func TestHydrate_accesses_Success(t *testing.T) {
	// build an accesses:
	accesses := accesses.CreateAccessesForTests(4096)

	// execute:
	hydro.VerifyAdapterUsingJSForTests(hydroAdapter, accesses, t)
}

func TestHydrate_user_Success(t *testing.T) {
	// build an accesses:
	user := users.CreateUserForTests(4096)

	// execute:
	hydro.VerifyAdapterUsingJSForTests(hydroAdapter, user, t)
}
