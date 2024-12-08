package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/scalarorg/scalar-core/utils/clog"

	nexus "github.com/axelarnetwork/axelar-core/x/nexus/exported"
)

// NewAddressValidator returns the callback for validating hex-encoded EVM addresses
func NewAddressValidator() nexus.AddressValidator {
	return func(ctx sdk.Context, address nexus.CrossChainAddress) error {
		// TODO: validate btc address
		clog.Red("TODO: validate btc address", address)

		return nil
	}
}
