package codec

import (
	codectypes "github.com/reapchain/cosmos-sdk/codec/types"
	cryptotypes "github.com/reapchain/cosmos-sdk/crypto/types"

	"github.com/reapchain/ethermint/crypto/ethsecp256k1"
)

// RegisterInterfaces register the Ethermint key concrete types.
func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterImplementations((*cryptotypes.PubKey)(nil), &ethsecp256k1.PubKey{})
}
